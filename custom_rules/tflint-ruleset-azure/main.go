package main

import (
    "github.com/terraform-linters/tflint-plugin-sdk/plugin"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/mamj00/tf/custom_rules/tflint-ruleset-azure/rules"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        RuleSet: &tflint.BuiltinRuleSet{
            Name:    "cloudima",
            Version: "0.1.0",
            Rules: []tflint.Rule{
                &rules.AzureGatewayValidSKU{},
				&rules.AzurePublicIPCompatibility{},
            },
        },
    })
}
