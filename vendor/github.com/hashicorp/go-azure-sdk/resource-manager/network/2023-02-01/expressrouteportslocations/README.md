
## `github.com/hashicorp/go-azure-sdk/resource-manager/network/2023-02-01/expressrouteportslocations` Documentation

The `expressrouteportslocations` SDK allows for interaction with the Azure Resource Manager Service `network` (API Version `2023-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/network/2023-02-01/expressrouteportslocations"
```


### Client Initialization

```go
client := expressrouteportslocations.NewExpressRoutePortsLocationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExpressRoutePortsLocationsClient.Get`

```go
ctx := context.TODO()
id := expressrouteportslocations.NewExpressRoutePortsLocationID("12345678-1234-9876-4563-123456789012", "expressRoutePortsLocationValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExpressRoutePortsLocationsClient.List`

```go
ctx := context.TODO()
id := expressrouteportslocations.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
