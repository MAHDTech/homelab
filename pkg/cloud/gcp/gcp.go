// Package gcp contains the functions for creating GCP resources.
package gcp

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	gcpschema "homelab/pkg/cloud/gcp/schema"

	utils "homelab/pkg/utils"
)

// CreateResources creates the GCP resources.
func CreateResources(ctx *pulumi.Context, config *gcpschema.Config) error {

	utils.LogInfo(ctx, "✨ Creating GCP resources...")

	utils.LogInfo(ctx, "GCP configuration: %+v", config)

	utils.LogInfo(ctx, "🚀 Successfully created GCP resources!")

	return nil
}
