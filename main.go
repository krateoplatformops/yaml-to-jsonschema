package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/krateoplatformops/yaml-to-jsonschema/internal/config"
	"github.com/krateoplatformops/yaml-to-jsonschema/internal/schema"
	"gopkg.in/yaml.v3"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if cfg.YAMLFile == "" {
		fmt.Fprintln(os.Stderr, "error: missing source YAML file")
		os.Exit(1)
	}

	content, err := os.ReadFile(cfg.YAMLFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var values yaml.Node
	err = yaml.Unmarshal(content, &values)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	dir := filepath.Dir(cfg.YAMLFile)
	base := filepath.Base(cfg.YAMLFile)
	ext := filepath.Ext(cfg.YAMLFile)

	res := schema.FromYAML(dir, &values, nil)

	sch, err := res.ToJson()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	err = os.MkdirAll(cfg.DestinationDir, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: unable to create destination dir: %v\n", err)
		os.Exit(1)
	}

	schemaFilePath := filepath.Join(cfg.DestinationDir,
		fmt.Sprintf("%s.schema.json", strings.TrimSuffix(base, ext)))

	err = os.WriteFile(schemaFilePath, sch, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
