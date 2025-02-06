plugin "azurerm" {
  enabled = true
  version = "0.27.0" # Sustituye por la versión más reciente
  source  = "github.com/terraform-linters/tflint-ruleset-azurerm"
}

plugin "template" {
  enabled = true
}

rule "azure_gateway_valid_sku" {
  enabled = true
}

rule "azure_public_ip_compatibility" {
  enabled = true
}