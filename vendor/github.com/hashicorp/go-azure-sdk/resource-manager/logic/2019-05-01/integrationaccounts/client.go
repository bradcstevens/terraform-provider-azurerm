package integrationaccounts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountsClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationAccountsClientWithBaseURI(api environments.Api) (*IntegrationAccountsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "integrationaccounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationAccountsClient: %+v", err)
	}

	return &IntegrationAccountsClient{
		Client: client,
	}, nil
}
