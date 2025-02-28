package iotdpsresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotDpsResourceClient struct {
	Client *resourcemanager.Client
}

func NewIotDpsResourceClientWithBaseURI(api environments.Api) (*IotDpsResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "iotdpsresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IotDpsResourceClient: %+v", err)
	}

	return &IotDpsResourceClient{
		Client: client,
	}, nil
}
