package aws

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	awsconfig "homelab/pkg/cloud/aws/config"
)

func CreateResources(ctx *pulumi.Context, config *awsconfig.Config) error {

	var message string

	ctx.Log.Info("TASK: Creating AWS resources...", nil)

	message = fmt.Sprintf("AWS configuration: %+v", config)
	ctx.Log.Info(message, nil)

	return nil
}
