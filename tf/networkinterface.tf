resource "azurerm_network_interface" "nicexample" {
  name = "nicexample"
  location = azurerm_resource_group.rgexample2.location

  ip_configuration {
    name = "ipconfig"
    private_ip_address_allocation = "Dynamic"
    subnet_id = azurerm_subnet.snet2.id
  }
  resource_group_name = azurerm_resource_group.rgexample2.name
}

