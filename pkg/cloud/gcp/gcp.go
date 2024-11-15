package gcp

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	gcpconfig "homelab/pkg/cloud/gcp/config"
)

func CreateResources(ctx *pulumi.Context, config *gcpconfig.Config) error {

	var message string

	ctx.Log.Info("TASK: Creating GCP resources...", nil)

	message = "TODO: CREATE GCP RESOURCES HERE..."
	ctx.Log.Info(message, nil)

	message = fmt.Sprintf("GCP configuration: %+v", config)
	ctx.Log.Info(message, nil)

	return nil
}
