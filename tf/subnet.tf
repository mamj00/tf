resource "azurerm_subnet" "subnetgw" {
  name = "GatewaySubnet"
  address_prefixes = ["10.0.0.0/24"]
  resource_group_name = azurerm_resource_group.rgexample.name
  virtual_network_name = azurerm_virtual_network.vnetexample.name
}

resource "azurerm_subnet" "subnetb" {
  name = "subnetb"
  address_prefixes = ["10.0.1.0/24"]
  resource_group_name = azurerm_resource_group.rgexample.name
  virtual_network_name = azurerm_virtual_network.vnetexample.name
}

resource "azurerm_subnet" "snet2" {
  name = "snetexample"
  address_prefixes = ["10.1.0.0/24"]
  resource_group_name = azurerm_resource_group.rgexample2.name
  virtual_network_name = azurerm_virtual_network.vnet2.name
}

