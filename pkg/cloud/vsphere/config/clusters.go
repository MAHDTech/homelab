// Package vsphereconfig contains the functions for verifying the VMware configuration.
package vsphereconfig

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	vsphereschema "homelab/pkg/cloud/vsphere/schema"

	utils "homelab/pkg/utils"
)

// verifyHostsAndClustersConfig ensures the Hosts and Clusters configuration is valid.
// If the validation fails, an error is returned.
func verifyHostsAndClustersConfig(
	ctx *pulumi.Context,
	configRaw interface{},
) (vsphereschema.HostsAndClusters, error) {

	var configFinal vsphereschema.HostsAndClusters

	// Take the raw configuration and unmarshal it into the schema.
	err := utils.TryObject(configRaw, "hosts_and_clusters", &configFinal)
	if err != nil {
		return configFinal, err
	}

	// nolint:errcheck
	ctx.Log.Info("TODO: Validate the configuration for 'HostsAndClusters' here.", nil)

	return configFinal, nil
}
