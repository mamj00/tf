// tflint-ruleset-azurerm-custom/apis/azure/client.go
package azure

import (
    "context"
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

type Client struct {
    SKUClient *armnetwork.VirtualNetworkGatewaysClient
}

func NewClient(subscriptionID string) (*Client, error) {
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        return nil, err
    }

    skuClient, err := armnetwork.NewVirtualNetworkGatewaysClient(subscriptionID, cred, nil)
    return &Client{
        SKUClient: skuClient,
    }, nil
}

func (c *Client) GetValidSKUs(ctx context.Context, location string) (map[string]bool, error) {
    skus := make(map[string]bool)
    
    pager := c.SKUClient.NewListPager(nil)
    for pager.More() {
        page, err := pager.NextPage(ctx)
        if err != nil {
            return nil, err
        }
        
        for _, sku := range page.Value {
            if *sku.Location == location {
                skus[*sku.Name] = true
            }
        }
    }
    
    return skus, nil
}
