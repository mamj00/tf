package rules

import (
    "strings"
    "github.com/terraform-linters/tflint-plugin-sdk/hclext"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint")

type AzurePublicIPCompatibility struct {
    tflint.DefaultRule
} // <-- Añadir llave faltante

func (r *AzurePublicIPCompatibility) Name() string {
    return "azure_public_ip_compatibility"
} // <-- Añadir llave

func (r *AzurePublicIPCompatibility) Enabled() bool {
    return true
} // <-- Añadir llave

func (r *AzurePublicIPCompatibility) Severity() tflint.Severity {
    return tflint.ERROR
} // <-- Añadir llave

func (r *AzurePublicIPCompatibility) Link() string {
    return "https://learn.microsoft.com/azure/virtual-network/ip-services/public-ip-addresses"
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
            // Eliminar pipID si no se usa
            _, exists := config["public_ip_address_id"]
            if !exists {
                continue
            }
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
