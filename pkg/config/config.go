package config

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// VerifyConfig checks for required configuration items
// in the Global section.
// If a required item is missing, an err is returned.
// Unlike the built-in pulumi "Require" this doesn't panic
// and display a stack trace, rather just the specific err
// message which is more user friendly.
func VerifyConfig(ctx *pulumi.Context, cfg config.Config) error {

	ctx.Log.Info("TASK: Verifying user provided configuration...", nil)

	// REQUIRED: environment.domain
	environmentDomain, err := cfg.Try("environment.domain")
	if err != nil {
		message := "ERROR: Missing configuration item: environment.domain"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: environment.domain: " + environmentDomain
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: talos.cluster.name
	talosClusterName, err := cfg.Try("talos.cluster.name")
	if err != nil {
		message := "ERROR: Missing configuration item: talos.cluster.name"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: talos.cluster.name: " + talosClusterName
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: talos.nodes.count
	talosNodesCount, err := cfg.TryInt("talos.nodes.count")
	if err != nil {
		message := "ERROR: Missing configuration item: talos.nodes.count"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: talos.nodes.count: " + fmt.Sprint(talosNodesCount)
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: talos.infra.user
	talosInfraUser, err := cfg.Try("talos.infra.user")
	if err != nil {
		message := "ERROR: Missing configuration item: talos.infra.user"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: talos.infra.user: " + talosInfraUser
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: talos.infra.sshPrivateKey
	talosInfraSSHPrivateKey, err := cfg.Try("talos.infra.sshPrivateKey")
	if err != nil {
		message := "ERROR: Missing configuration item: talos.infra.sshPrivateKey"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: talos.infra.privateKey: " + talosInfraSSHPrivateKey
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: talos.infra.sshPublicKey
	talosInfraSSHPublicKey, err := cfg.Try("talos.infra.sshPublicKey")
	if err != nil {
		message := "ERROR: Missing configuration item: talos.infra.sshPublicKey"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: talos.infra.privateKey: " + talosInfraSSHPublicKey
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: vsphere.datacenter.name
	vsphereDatacenterName, err := cfg.Try("vsphere.datacenter.name")
	if err != nil {
		message := "ERROR: Missing configuration item: vsphere.datacenter.name"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: vsphere.datacenter.name: " + vsphereDatacenterName
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: vsphere.cluster.name
	vsphereClusterName, err := cfg.Try("vsphere.cluster.name")
	if err != nil {
		message := "ERROR: Missing configuration item: vsphere.cluster.name"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: vsphere.cluster.name: " + vsphereClusterName
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: vsphere.host.name
	vsphereHostName, err := cfg.Try("vsphere.host.name")
	if err != nil {
		message := "ERROR: Missing configuration item: vsphere.host.name"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: vsphere.host.name: " + vsphereHostName
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: vsphere.datastore.name
	vsphereDatastoreName, err := cfg.Try("vsphere.datastore.name")
	if err != nil {
		message := "ERROR: Missing configuration item: vsphere.datastore.name"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: vsphere.datastore.name: " + vsphereDatastoreName
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: vsphere.network.name
	vsphereNetworkName, err := cfg.Try("vsphere.network.name")
	if err != nil {
		message := "ERROR: Missing configuration item: vsphere.network.name"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: vsphere.network.name: " + vsphereNetworkName
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: vsphere.resourcepool.name
	vsphereResourcePoolName, err := cfg.Try("vsphere.resourcepool.name")
	if err != nil {
		message := "ERROR: Missing configuration item: vsphere.resourcepool.name"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: vsphere.resourcepool.name: " + vsphereResourcePoolName
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: vsphere.folder.name
	vSphereFolderRoot, err := cfg.Try("vsphere.folder.name")
	if err != nil {
		message := "ERROR: Missing configuration item: vsphere.folder.name"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: vsphere.folder.name: " + vSphereFolderRoot
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: vsphere.templates.infra
	vSphereTemplateInfra, err := cfg.Try("vsphere.templates.infra")
	if err != nil {
		message := "ERROR: Missing configuration item: vsphere.templates.infra"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: vsphere.templates.infra: " + vSphereTemplateInfra
		ctx.Log.Info(message, nil)
	}

	// REQUIRED: vsphere.templates.node
	vSphereTemplateNode, err := cfg.Try("vsphere.templates.node")
	if err != nil {
		message := "ERROR: Missing configuration item: vsphere.templates.node"
		ctx.Log.Error(message, nil)
		return err
	} else {
		message := "OK: vsphere.templates.node: " + vSphereTemplateNode
		ctx.Log.Info(message, nil)
	}

	return nil

}

// VerifyConfigCloud checks for required configuration items
// in the Cloud section.
// If a required item is missing, an err is returned.
// Unlike the built-in pulumi "Require" this doesn't panic
// and display a stack trace, rather just the specific err
// message which is more user friendly.
func VerifyConfigCloud(ctx *pulumi.Context, cfg config.Config, cloud string) error {

	switch cloud {
	
		case "aws":
			return VerifyConfigAWS(ctx, cfg)
		
		case "azure":
			return VerifyConfigAzure(ctx, cfg)
		
		case "gcp":
			return VerifyConfigGCP(ctx, cfg)
		
		case "vsphere":
			return VerifyConfigVSphere(ctx, cfg)

		case "nutanix":
			return VerifyConfigNutanix(ctx, cfg)

		default:
			return fmt.Errorf("unsupported cloud provider: %s", cloud)
	}

	return nil
}
