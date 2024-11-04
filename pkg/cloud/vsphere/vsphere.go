package vsphere

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	config "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	omniStackTalos "omnistack/talos"
)

func Init(ctx *pulumi.Context, cfg *config.Config) (*omniStackTalos.TalosDeployment, error) {

	cloudProvider := cfg.Get("cloud")

	message := fmt.Sprintf("TASK: Initializing OmniStack deployment on %s...", cloudProvider)
	ctx.Log.Info(message, nil)

	return &omniStackTalos.TalosDeployment{}, nil

}