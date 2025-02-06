resource "azurerm_virtual_network" "vnetexample" {
  name = "vnetexample"
  location = azurerm_resource_group.rgexample.location
  address_space = ["10.0.0.0/16"]
  resource_group_name = azurerm_resource_group.rgexample.name
}

resource "azurerm_virtual_network" "vnet2" {
  name = "vnetexample2"
  location = azurerm_resource_group.rgexample2.location
  address_space = ["10.1.0.0/16"]
  resource_group_name = azurerm_resource_group.rgexample2.name
}

