package rules

import (
    "strings"
    "github.com/terraform-linters/tflint-plugin-sdk/hclext"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AzurePublicIPCompatibility struct {
    tflint.DefaultRule
}

func (r *AzurePublicIPCompatibility) Name() string {
    return "azure_public_ip_compatibility"
}

func (r *AzurePublicIPCompatibility) Enabled() bool {
    return true
}

func (r *AzurePublicIPCompatibility) Severity() tflint.Severity {
    return tflint.ERROR
}

func (r *AzurePublicIPCompatibility) Link() string {
    return "https://learn.microsoft.com/azure/virtual-network/ip-services/public-ip-addresses"
}

func (r *AzurePublicIPCompatibility) Check(runner tflint.Runner) error {
    resources, err := runner.GetResourceContent("azurerm_virtual_network_gateway", &hclext.BodySchema{
        Attributes: []hclext.AttributeSchema{{Name: "sku"}, {Name: "ip_configuration"}},
    }, nil)
    if err != nil {
        return err
    }

    for _, resource := range resources.Blocks {
        skuAttr, exists := resource.Body.Attributes["sku"]
        if !exists {
            continue
        }

        var sku string
        if err := runner.EvaluateExpr(skuAttr.Expr, &sku, nil); err != nil {
            return err
        }

        if strings.HasSuffix(sku, "AZ") {
            ipConfigsAttr, exists := resource.Body.Attributes["ip_configuration"]
            if !exists {
                continue
            }

            var ipConfigs []map[string]interface{}
            if err := runner.EvaluateExpr(ipConfigsAttr.Expr, &ipConfigs, nil); err != nil {
                return err
            }

            for _, config := range ipConfigs {
                if _, ok := config["public_ip_address_id"].(string); ok {
                    pipResources, _ := runner.GetResourceContent("azurerm_public_ip", &hclext.BodySchema{
                        Attributes: []hclext.AttributeSchema{{Name: "sku"}},
                    }, nil)

                    for _, pip := range pipResources.Blocks {
                        skuAttr, exists := pip.Body.Attributes["sku"]
                        if !exists {
                            continue
                        }

                        var pipSKU string
                        if err := runner.EvaluateExpr(skuAttr.Expr, &pipSKU, nil); err != nil {
                            return err
                        }

                        if pipSKU == "Basic" {
                            runner.EmitIssue(
                                r,
                                "SKU de IP pública debe ser 'Standard' con gateways AZ",
                                skuAttr.Expr.Range(),
                            )
                        }
                    }
                }
            }
        }
    }
    return nil
}
