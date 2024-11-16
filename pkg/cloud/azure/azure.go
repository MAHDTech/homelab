// Package azure contains the functions for creating Azure resources.
package azure

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	azureschema "homelab/pkg/cloud/azure/schema"
	utils "homelab/pkg/utils"
)

// CreateResources creates the Azure resources.
func CreateResources(ctx *pulumi.Context, config *azureschema.Config) error {

	utils.LogInfo(ctx, "✨ Creating Azure resources...")

	utils.LogInfo(ctx, "Azure configuration: %+v", config)

	utils.LogInfo(ctx, "🚀 Successfully created Azure resources!")

	return nil
}
