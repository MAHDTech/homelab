package nutanix

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	nutanixconfig "homelab/pkg/cloud/nutanix/config"
)

func CreateResources(ctx *pulumi.Context, config *nutanixconfig.Config) error {

	var message string

	ctx.Log.Info("TASK: Creating Nutanix resources...", nil)

	message = fmt.Sprintf("Nutanix configuration: %+v", config)
	ctx.Log.Info(message, nil)

	return nil
}
