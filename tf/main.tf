terraform {
  backend "azurerm" {
    resource_group_name  = "RG-Cloudima"
    storage_account_name = "cloudimamycd"
    container_name       = "tfstate"
    key                  = "tfstates/9d0b82cf-3dcb-43fc-b1ec-a24abac86b62/terraform.tfstate"
  }
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.96.0"
    }
  }
}

provider "azurerm" {
  features {}
}
