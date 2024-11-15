package global

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	globalconfig "homelab/pkg/cloud/global/config"
)

func CreateResources(ctx *pulumi.Context, config *globalconfig.Config) error {

	ctx.Log.Info("TASK: Creating global resources...", nil)

	return nil
}
