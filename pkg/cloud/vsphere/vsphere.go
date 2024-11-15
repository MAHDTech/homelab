package vsphere

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	vsphereconfig "homelab/pkg/cloud/vsphere/config"
)

func CreateResources(ctx *pulumi.Context, config *vsphereconfig.Config) error {

	var message string

	ctx.Log.Info("TASK: Creating VSphere resources...", nil)

	message = fmt.Sprintf("VSphere configuration: %+v", config)
	ctx.Log.Info(message, nil)

	return nil
}
