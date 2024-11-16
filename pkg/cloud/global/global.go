// Package global contains the functions for creating global resources.
package global

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	globalschema "homelab/pkg/cloud/global/schema"

	utils "homelab/pkg/utils"
)

// CreateResources creates the global resources.
func CreateResources(ctx *pulumi.Context, config *globalschema.Config) error {

	utils.LogInfo(ctx, "✨ Creating global resources...")

	utils.LogInfo(ctx, "Global configuration: %+v", config)

	utils.LogInfo(ctx, "🚀 Successfully created global resources!")

	return nil
}
