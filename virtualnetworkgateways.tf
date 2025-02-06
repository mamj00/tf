resource "azurerm_virtual_network_gateway" "virtualgw" {
  sku = "Standard"
  name = "virtualgw"
  type = "Vpn"
  location = azurerm_resource_group.rgexample.location

  ip_configuration {
    public_ip_address_id = azurerm_public_ip.pip.id
    subnet_id = azurerm_subnet.subnetgw.id
  }
  resource_group_name = azurerm_resource_group.rgexample.name
}

