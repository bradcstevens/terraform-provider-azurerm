package networkmanageractiveconnectivityconfigurations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkManagerActiveConnectivityConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewNetworkManagerActiveConnectivityConfigurationsClientWithBaseURI(api environments.Api) (*NetworkManagerActiveConnectivityConfigurationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "networkmanageractiveconnectivityconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkManagerActiveConnectivityConfigurationsClient: %+v", err)
	}

	return &NetworkManagerActiveConnectivityConfigurationsClient{
		Client: client,
	}, nil
}
