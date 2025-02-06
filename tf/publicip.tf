resource "azurerm_public_ip" "pip" {
  sku = "Standard"
  name = "pipgw"
  location = azurerm_resource_group.rgexample.location
  allocation_method = "Static"
  resource_group_name = azurerm_resource_group.rgexample.name
}

