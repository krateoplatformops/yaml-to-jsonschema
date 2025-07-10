# `YAML to JSON Schema`

This GitHub Action generate JSON Schema from YAML values.

## Features

- Converts a specified YAML file into a JSON Schema.
- Outputs the generated schema into a configurable directory.
- Runs inside a Docker container for consistent environments.

## Inputs 

| Input            | Description                                             | Required | Default       |
| ---------------- | ------------------------------------------------------- | -------- | ------------- |
| `destinationDir` | Directory where the generated JSON Schema will be saved | No       | `schema`      |
| `yamlFile`       | The source YAML file used for JSON Schema generation    | No       | `values.yaml` |


## Usage example

```yaml
name: Generate JSON Schema from Helm Values

on:
  push:
    tags:
      - 'v*'

jobs:
  run-action:
    runs-on: ubuntu-latest
    permissions:
      contents: write
      packages: read
    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Generate JSON Schema from YAML
        uses: krateoplatformops/yaml-to-jsonschema@v0.1.0
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          yamlFile: values.yaml            # optional, defaults to values.yaml
          destinationDir: schema           # optional, defaults to schema

      - name: Commit changes
        run: |
          git config --global user.name "github-actions[bot]"
          git config --global user.email "github-actions[bot]@users.noreply.github.com"
          git add .
          git commit -m "chore: update JSON Schema via GitHub Actions"
          git push origin HEAD:main
```