package azure

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	azureconfig "homelab/pkg/cloud/azure/config"
)

func CreateResources(ctx *pulumi.Context, config *azureconfig.Config) error {

	var message string

	ctx.Log.Info("TASK: Creating Azure resources...", nil)

	message = fmt.Sprintf("Azure configuration: %+v", config)
	ctx.Log.Info(message, nil)

	return nil
}
