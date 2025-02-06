resource "azurerm_firewall" "firewallexample" {
  name = "firewall"
  location = azurerm_resource_group.rgexample.location
  sku_name = "AZFW_Hub"
  sku_tier = "Basic"
  resource_group_name = azurerm_resource_group.rgexample.name
}

