package cloud

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	aws "homelab/pkg/cloud/aws"
	azure "homelab/pkg/cloud/azure"
	gcp "homelab/pkg/cloud/gcp"
	global "homelab/pkg/cloud/global"
	nutanix "homelab/pkg/cloud/nutanix"
	vsphere "homelab/pkg/cloud/vsphere"
	config "homelab/pkg/config"
)

func CreateResources(ctx *pulumi.Context, configVerified *config.Config) error {

	var err error

	ctx.Log.Info("TASK: Creating cloud resources...", nil)

	// Create the global resources if enabled.
	if configVerified.Global.Enabled {
		err = global.CreateResources(ctx, &configVerified.Global)
		if err != nil {
			message := "Failed to create global resources!"
			ctx.Log.Error(message, nil)
			return err
		}
	}

	// Create the AWS resources if enabled.
	if configVerified.AWS.Enabled {
		err = aws.CreateResources(ctx, &configVerified.AWS)
		if err != nil {
			message := "Failed to create AWS resources!"
			ctx.Log.Error(message, nil)
			return err
		}
	}

	// Create the Azure resources if enabled.
	if configVerified.Azure.Enabled {
		err = azure.CreateResources(ctx, &configVerified.Azure)
		if err != nil {
			message := "Failed to create Azure resources!"
			ctx.Log.Error(message, nil)
			return err
		}
	}

	// Create the GCP resources if enabled.
	if configVerified.GCP.Enabled {
		err = gcp.CreateResources(ctx, &configVerified.GCP)
		if err != nil {
			message := "Failed to create GCP resources!"
			ctx.Log.Error(message, nil)
			return err
		}
	}

	// Create the Nutanix resources if enabled.
	if configVerified.Nutanix.Enabled {
		err = nutanix.CreateResources(ctx, &configVerified.Nutanix)
		if err != nil {
			message := "Failed to create Nutanix resources!"
			ctx.Log.Error(message, nil)
			return err
		}
	}

	// Create the VSphere resources if enabled.
	if configVerified.VSphere.Enabled {
		err = vsphere.CreateResources(ctx, &configVerified.VSphere)
		if err != nil {
			message := "Failed to create VSphere resources!"
			ctx.Log.Error(message, nil)
			return err
		}
	}

	return nil
}
