package connections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewConnectionsClientWithBaseURI(api environments.Api) (*ConnectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "connections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectionsClient: %+v", err)
	}

	return &ConnectionsClient{
		Client: client,
	}, nil
}
