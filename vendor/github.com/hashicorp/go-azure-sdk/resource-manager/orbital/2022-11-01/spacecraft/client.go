package spacecraft

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SpacecraftClient struct {
	Client *resourcemanager.Client
}

func NewSpacecraftClientWithBaseURI(api environments.Api) (*SpacecraftClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "spacecraft", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SpacecraftClient: %+v", err)
	}

	return &SpacecraftClient{
		Client: client,
	}, nil
}
