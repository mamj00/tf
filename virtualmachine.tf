resource "azurerm_windows_virtual_machine" "vm315" {
  name = "vmexample"
  size = "Standard_B2s"

  os_disk {
    caching = "None"
    name = "disk"
    storage_account_type = "Standard_LRS"
  }
  location = azurerm_resource_group.rgexample2.location
  admin_password = "Password?.2025"
  admin_username = "Useradmin"
  resource_group_name = azurerm_resource_group.rgexample2.name
  network_interface_ids = [azurerm_network_interface.nicexample.id]

  source_image_reference {
    offer = "WindowsServer"
    publisher = "MicrosoftWindowsServer"
    sku = "2016-Datacenter"
    version = "latest"
  }
}

