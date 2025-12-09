package config

type AndroidConfig struct {
	BuildGradlePath     string
	ManifestPath        string
	ManifestDebugPath   string
	ManifestProfilePath string
	ActivityPath        string
}

type Config struct {
	AndroidConfig AndroidConfig
}
