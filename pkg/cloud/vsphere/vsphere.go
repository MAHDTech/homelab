// Package vsphere contains the functions for creating VMware resources.
package vsphere

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	vsphereschema "homelab/pkg/cloud/vsphere/schema"

	utils "homelab/pkg/utils"
)

// CreateResources creates the VMware resources.
func CreateResources(ctx *pulumi.Context, config *vsphereschema.Config) error {

	utils.LogInfo(ctx, "✨ Creating vSphere resources...")

	utils.LogInfo(ctx, "vSphere configuration: %+v", config)

	utils.LogInfo(ctx, "🚀 Successfully created vSphere resources!")

	return nil
}
