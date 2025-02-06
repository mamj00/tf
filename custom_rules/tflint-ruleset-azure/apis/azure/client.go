package azure

import (
    "context"
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5" // Usar armcompute
)

type Client struct {
    SKUClient *armcompute.ResourceSKUsClient // Cambiar tipo de cliente
}

func NewClient(subscriptionID string) (*Client, error) {
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        return nil, err
    }

    skuClient, err := armcompute.NewResourceSKUsClient(subscriptionID, cred, nil)
    return &Client{
        SKUClient: skuClient,
    }, nil
}


func (c *Client) GetValidSKUs(ctx context.Context, location string) (map[string]bool, error) {
    skus := make(map[string]bool)
    
    pager := c.SKUClient.NewListPager(&armcompute.ResourceSKUsClientListOptions{
        Filter: to.Ptr("location eq '" + location + "'"), // Filtro por ubicación
    })
    
    for pager.More() {
        page, err := pager.NextPage(ctx)
        if err != nil {
            return nil, err
        }
        
        for _, sku := range page.Value {
            if *sku.ResourceType == "virtualNetworkGateways" { // Filtrar por tipo de recurso
                skus[*sku.Name] = true
            }
        }
    }
    
    return skus, nil
}

