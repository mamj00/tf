package rules

import (
    "github.com/terraform-linters/tflint-plugin-sdk/hclext"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AzureGatewayValidSKU struct {
    tflint.DefaultRule
}

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

func (r *AzureGatewayValidSKU) Check(runner tflint.Runner) error {
    allowedSKUs := map[string]bool{
        "VpnGw1AZ": true, "VpnGw2AZ": true, "VpnGw3AZ": true,
        "VpnGw4AZ": true, "VpnGw5AZ": true,
        "VpnGw1": true, "VpnGw2": true, "VpnGw3": true, "VpnGw4": true, "VpnGw5": true,
    }

    resources, err := runner.GetResourceContent("azurerm_virtual_network_gateway", &hclext.BodySchema{
        Attributes: []hclext.AttributeSchema{{Name: "sku"}},
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

        if !allowedSKUs[sku] {
            runner.EmitIssue(
                r,
                "SKU inválido para Azure Gateway. Válidos: VpnGw1-5 o VpnGw1AZ-5AZ",
                skuAttr.Expr.Range(),
            )
        }
    }
    return nil
}
