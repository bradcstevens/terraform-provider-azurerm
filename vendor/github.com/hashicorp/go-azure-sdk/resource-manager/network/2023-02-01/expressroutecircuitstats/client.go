package expressroutecircuitstats

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCircuitStatsClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteCircuitStatsClientWithBaseURI(api environments.Api) (*ExpressRouteCircuitStatsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "expressroutecircuitstats", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteCircuitStatsClient: %+v", err)
	}

	return &ExpressRouteCircuitStatsClient{
		Client: client,
	}, nil
}
