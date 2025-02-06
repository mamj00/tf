plugin "azurerm" {
  enabled = true
  version = "0.27.0" # Sustituye por la versión más reciente
  source  = "github.com/terraform-linters/tflint-ruleset-azurerm"
}

plugin "cloudima" {
  enabled = true
  version = "0.1.1"
  source = "github.com/mamj00/tf"
}

rule "azure_gateway_valid_sku" {
  enabled = true
}

rule "azure_public_ip_compatibility" {
  enabled = true
}