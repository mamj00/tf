package rules

import (
    "context"
    "strings"
    
    "github.com/terraform-linters/tflint-plugin-sdk/hclext"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "./apis/azure" // Ruta local corregida
)

type AzureGatewayValidSKU struct {
    tflint.DefaultRule
}

// Métodos corregidos con llaves de cierre
func (r *AzureGatewayValidSKU) Name() string {
    return "azure_gateway_valid_sku"
}

func (r *AzureGatewayValidSKU) Enabled() bool {
    return true
}

func (r *AzureGatewayValidSKU) Severity() tflint.Severity {
    return tflint.ERROR
}

func (r *AzureGatewayValidSKU) Link() string {
    return "https://learn.microsoft.com/azure/vpn-gateway/vpn-gateway-about-vpngateways"
}

func (r *AzurePublicIPCompatibility) Check(runner tflint.Runner) error {
    gateways, err := runner.GetResourceContent("azurerm_virtual_network_gateway", &hclext.BodySchema{
        Attributes: []hclext.AttributeSchema{
            {Name: "sku"},
            {Name: "ip_configuration"},
        },
    }, nil)
    if err != nil {
        return err
    }

    for _, gateway := range gateways.Blocks {
        skuAttr, exists := gateway.Body.Attributes["sku"]
        if !exists {
            continue
        }

        var gatewaySKU string
        if err := runner.EvaluateExpr(skuAttr.Expr, &gatewaySKU, nil); err != nil {
            return err
        }

        if !strings.HasSuffix(gatewaySKU, "AZ") {
            continue
        }

        ipConfigsAttr, exists := gateway.Body.Attributes["ip_configuration"]
        if !exists {
            continue
        }

        var ipConfigs []map[string]interface{}
        if err := runner.EvaluateExpr(ipConfigsAttr.Expr, &ipConfigs, nil); err != nil {
            return err
        }

        for _, config := range ipConfigs {
            pipID, ok := config["public_ip_address_id"].(string)
            if !ok {
                continue
            }

            pipResource, err := runner.GetResourceContent("azurerm_public_ip", &hclext.BodySchema{
                Attributes: []hclext.AttributeSchema{{Name: "sku"}},
            }, nil)
            if err != nil {
                return err
            }

            for _, pip := range pipResource.Blocks {
                skuAttr, exists := pip.Body.Attributes["sku"]
                if !exists {
                    continue
                }

                var sku string
                if err := runner.EvaluateExpr(skuAttr.Expr, &sku, nil); err != nil {
                    return err
                }

                if sku == "Basic" {
                    runner.EmitIssue(
                        r,
                        "Las IP públicas deben ser Standard con gateways AZ",
                        skuAttr.Expr.Range(),
                    )
                }
            }
        }
    }
    return nil
}
