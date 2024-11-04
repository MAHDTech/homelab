package main

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	config "homelab/pkg/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Read the pulumi stack variables into a config object.
		stackConfig := pulumiConfig.New(ctx, "homelab")

		// Verify the configuration.
		configVerified, err := config.VerifyConfig(ctx, *stackConfig)
		if err != nil {
			message := "Failed to verify configuration!"
			ctx.Log.Error(message, nil)
			return err
		}

		// TODO: Remove debug message displaying the verified config.
		message := fmt.Sprintf("Verified configuration: %+v", configVerified)
		ctx.Log.Info(message, nil)

		// TODO: Create the resources for enabled cloud providers.

		return nil

	})

}
