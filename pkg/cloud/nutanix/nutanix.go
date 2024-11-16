// Package nutanix contains the functions for creating Nutanix resources.
package nutanix

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	nutanixschema "homelab/pkg/cloud/nutanix/schema"

	utils "homelab/pkg/utils"
)

// CreateResources creates the Nutanix resources.
func CreateResources(ctx *pulumi.Context, config *nutanixschema.Config) error {

	utils.LogInfo(ctx, "✨ Creating Nutanix resources...")

	utils.LogInfo(ctx, "Nutanix configuration: %+v", config)

	utils.LogInfo(ctx, "🚀 Successfully created Nutanix resources!")

	return nil
}
