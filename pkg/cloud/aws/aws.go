// Package aws contains the functions for creating AWS resources.
package aws

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	awsschema "homelab/pkg/cloud/aws/schema"

	utils "homelab/pkg/utils"
)

// CreateResources creates the AWS resources.
func CreateResources(ctx *pulumi.Context, config *awsschema.Config) error {

	utils.LogInfo(ctx, "✨ Creating AWS resources...")

	utils.LogInfo(ctx, "AWS configuration: %+v", config)

	utils.LogInfo(ctx, "🚀 Successfully created AWS resources!")

	return nil
}
