package main

import (
    "github.com/terraform-linters/tflint-plugin-sdk/plugin"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/mamj00/tf/rules"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        RuleSet: &tflint.BuiltinRuleSet{
            Name:    "cloudima",
            Version: "0.1.1",
            Rules: []tflint.Rule{
                &rules.AzureGatewayValidSKU{},
				&rules.AzurePublicIPCompatibility{},
            },
        },
    })
}
