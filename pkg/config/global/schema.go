package config

const CloudProviders = []string{
	"aws",
	"azure",
	"gcp",
	"vsphere",
	"nutanix",
}

type ConfigGCP struct {
	Enabled bool
}

type ConfigVSphere struct {
	Enabled bool
}

type ConfigNutanix struct {
	Enabled bool
}

type Config struct {

	ConfigGlobal struct {
		Debug bool
		Clouds []string
	}

	AWS ConfigAWS

	GCP ConfigGCP

	VSphere ConfigVSphere

	Nutanix ConfigNutanix

}
