package config

import (
	"flag"
	"os"
	"path/filepath"
)

const (
	appName = "yaml-to-jsonschema"
)

func Load() (cfg Config, err error) {
	flag.StringVar(&cfg.GithubToken, "github-token", os.Getenv("GITHUB_TOKEN"), "GitHub token")
	flag.StringVar(&cfg.YAMLFile, "yaml-file", os.Getenv("INPUT_YAMLFILE"), "Path to YAML file")
	flag.StringVar(&cfg.DestinationDir, "destination-dir", os.Getenv("INPUT_DESTINATIONDIR"), "Destination directory")

	flag.CommandLine.SetOutput(os.Stderr)

	flag.Parse()

	if cfg.DestinationDir == "" {
		cfg.DestinationDir = filepath.Dir(cfg.YAMLFile)
	}

	return
}

type Config struct {
	GithubToken    string
	YAMLFile       string
	DestinationDir string
}
