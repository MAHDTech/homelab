package main

import (
	godotenv "github.com/joho/godotenv"
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	cloud "homelab/pkg/cloud"
	config "homelab/pkg/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Load the dotenv file.
		err := godotenv.Load()
		if err != nil {
			message := "Failed to load the file '.env'. Please check that the file exists and is readable."
			ctx.Log.Error(message, nil)
			return err
		}

		// Read the pulumi stack variables into a config object.
		stackConfig := pulumiConfig.New(ctx, "homelab")

		// Verify the configuration.
		configVerified, err := config.VerifyConfig(ctx, *stackConfig)
		if err != nil {
			message := "Failed to verify configuration!"
			ctx.Log.Error(message, nil)
			return err
		}

		// Create the cloud resources.
		err = cloud.CreateResources(ctx, &configVerified)
		if err != nil {
			message := "Failed to create cloud resources!"
			ctx.Log.Error(message, nil)
			return err
		}

		return nil

	})
}
