// tflint-ruleset-azurerm-custom/rules/azure_gateway_valid_sku.go
package rules

import (
    "context"
    "strings"
    
    "github.com/terraform-linters/tflint-plugin-sdk/hclext"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "./apis/azure"
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
    ctx := context.Background()
    subscriptionID := runner.AzureCredentials().SubscriptionID
    
    client, err := azure.NewClient(subscriptionID)
    if err != nil {
        return err
    }

    resources, err := runner.GetResourceContent("azurerm_virtual_network_gateway", &hclext.BodySchema{
        Attributes: []hclext.AttributeSchema{
            {Name: "sku"}, 
            {Name: "location"},
            {Name: "type"},
        },
    }, nil)
    if err != nil {
        return err
    }

    for _, resource := range resources.Blocks {
        skuAttr, exists := resource.Body.Attributes["sku"]
        if !exists {
            continue
        }

        locationAttr, exists := resource.Body.Attributes["location"]
        if !exists {
            continue
        }

        var sku, location string
        if err := runner.EvaluateExpr(skuAttr.Expr, &sku, nil); err != nil {
            return err
        }
        if err := runner.EvaluateExpr(locationAttr.Expr, &location, nil); err != nil {
            return err
        }

        validSKUs, err := client.GetValidSKUs(ctx, location)
        if err != nil {
            return err
        }

        if !validSKUs[sku] {
            runner.EmitIssue(
                r,
                "SKU inválido para "+location+". Válidos: "+strings.Join(getKeys(validSKUs), ", "),
                skuAttr.Expr.Range(),
            )
        }
    }
    return nil
}

func getKeys(m map[string]bool) []string {
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
