// Package cloud is the main entry point for creating the resources
// for each of the supported cloud providers.
package cloud

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"homelab/pkg/cloud/aws"
	"homelab/pkg/cloud/azure"
	"homelab/pkg/cloud/gcp"
	"homelab/pkg/cloud/nutanix"
	"homelab/pkg/cloud/vsphere"
	schema "homelab/pkg/schema"
	utils "homelab/pkg/utils"
)

// CreateResources is responsible for creating the resources for each
// of the supported cloud providers.
func CreateResources(ctx *pulumi.Context, configVerified *schema.Config) error {

	var err error

	// If the global configuration is disabled, we can skip the rest of the function.
	if !configVerified.Global.Enabled {
		utils.LogInfo(
			ctx,
			"🛈 The 'global' configuration section is disabled. Skipping creation of all cloud resources.",
		)
		return nil
	}

	utils.LogInfo(ctx, "✨ Creating cloud resources...")

	// Create the AWS resources if enabled.
	if configVerified.AWS.Enabled {
		err = aws.CreateResources(ctx, &configVerified.AWS)
		if err != nil {
			utils.LogError(ctx, "💥 Failed to create AWS resources!")
			return err
		}
	}

	// Create the Azure resources if enabled.
	if configVerified.Azure.Enabled {
		err = azure.CreateResources(ctx, &configVerified.Azure)
		if err != nil {
			utils.LogError(ctx, "💥 Failed to create Azure resources!")
			return err
		}
	}

	// Create the GCP resources if enabled.
	if configVerified.GCP.Enabled {
		err = gcp.CreateResources(ctx, &configVerified.GCP)
		if err != nil {
			utils.LogError(ctx, "💥 Failed to create GCP resources!")
			return err
		}
	}

	// Create the Nutanix resources if enabled.
	if configVerified.Nutanix.Enabled {
		err = nutanix.CreateResources(ctx, &configVerified.Nutanix)
		if err != nil {
			utils.LogError(ctx, "💥 Failed to create Nutanix resources!")
			return err
		}
	}

	// Create the vSphere resources if enabled.
	if configVerified.VSphere.Enabled {
		err = vsphere.CreateResources(ctx, &configVerified.VSphere)
		if err != nil {
			utils.LogError(ctx, "💥 Failed to create VSphere resources!")
			return err
		}
	}

	return nil
}
