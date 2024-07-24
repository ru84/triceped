@description('The name of the resource group')
param resourceGroupName string = 'Test-Terraform-RG'

@description('Azure region to deploy resources')
param location string = 'East US'

@description('SSH public key for virtual machine')
param adminPublicKey string

@description('Azure tenant ID')
param tenantId string

// Resource Group
resource resourceGroup 'Microsoft.Resources/resourceGroups@2021-04-01' = {
  name: resourceGroupName
  location: location
}

// Virtual Network
resource virtualNetwork 'Microsoft.Network/virtualNetworks@2021-02-01' = {
  name: 'test-vnet'
  location: resourceGroup.location
  properties: {
    addressSpace: {
      addressPrefixes: [
        '10.0.0.0/16'
      ]
    }
  }
}

// Subnet
resource subnet 'Microsoft.Network/virtualNetworks/subnets@2021-02-01' = {
  name: 'test-subnet'
  parent: virtualNetwork
  properties: {
    addressPrefix: '10.0.1.0/24'
  }
}

// Network Security Group
resource networkSecurityGroup 'Microsoft.Network/networkSecurityGroups@2021-02-01' = {
  name: 'test-nsg'
  location: resourceGroup.location
}

// Public IP
resource publicIP 'Microsoft.Network/publicIPAddresses@2021-02-01' = {
  name: 'test-public-ip'
  location: resourceGroup.location
  properties: {
    publicIPAllocationMethod: 'Dynamic'
  }
}

// Network Interface
resource networkInterface 'Microsoft.Network/networkInterfaces@2021-02-01' = {
  name: 'test-nic'
  location: resourceGroup.location
  properties: {
    ipConfigurations: [
      {
        name: 'internal'
        properties: {
          subnet: {
            id: subnet.id
          }
          privateIPAllocationMethod: 'Dynamic'
          publicIPAddress: {
            id: publicIP.id
          }
        }
      }
    ]
  }
}

// Key Vault
resource keyVault 'Microsoft.KeyVault/vaults@2021-04-01' = {
  name: 'test-keyvault'
  location: resourceGroup.location
  properties: {
    tenantId: tenantId
    sku: {
      name: 'standard'
    }
    enableRbacAuthorization: true
    enablePurgeProtection: true
    softDeleteRetentionInDays: 90
  }
}

// Virtual Machine
resource virtualMachine 'Microsoft.Compute/virtualMachines@2021-07-01' = {
  name: 'test-vm'
  location: resourceGroup.location
  properties: {
    hardwareProfile: {
      vmSize: 'Standard_DS1_v2'
    }
    osProfile: {
      computerName: 'test-vm'
      adminUsername: 'azureuser'
      linuxConfiguration: {
        disablePasswordAuthentication: true
        ssh: {
          publicKeys: [
            {
              path: '/home/azureuser/.ssh/authorized_keys'
              keyData: adminPublicKey
            }
          ]
        }
      }
    }
    storageProfile: {
      osDisk: {
        caching: 'ReadWrite'
        managedDisk: {
          storageAccountType: 'Standard_LRS'
        }
        createOption: 'FromImage'
      }
      imageReference: {
        publisher: 'Canonical'
        offer: 'UbuntuServer'
        sku: '18.04-LTS'
        version: 'latest'
      }
    }
    networkProfile: {
      networkInterfaces: [
        {
          id: networkInterface.id
        }
      ]
    }
  }
}

// Storage Account
resource storageAccount 'Microsoft.Storage/storageAccounts@2021-09-01' = {
  name: 'testsa'
  location: resourceGroup.location
  properties: {
    accessTier: 'Hot'
  }
  sku: {
    name: 'Standard_LRS'
  }
}

// Outputs
output resourceGroupName string = resourceGroup.name
output vmIPAddress string = publicIP.properties.ipAddress
output keyVaultId string = keyVault.id
