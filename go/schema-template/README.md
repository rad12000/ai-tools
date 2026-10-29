# schema-template

Renders Go template files, validating their inputs against JSON Schema files (written in YAML or JSON).

## Install

### Prebuilt binary

Download the binary for your platform from the [releases page](https://github.com/rad12000/ai-tools/releases?q=schema-template) and put it on your `PATH`.

macOS / Linux:

```sh
VERSION=0.1.0                                   # pick a release
OS=$(uname -s | tr '[:upper:]' '[:lower:]')     # darwin or linux
ARCH=$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/')
curl -fL -o schema-template \
  "https://github.com/rad12000/ai-tools/releases/download/schema-template/v${VERSION}/schema-template_${OS}_${ARCH}"
chmod +x schema-template
sudo mv schema-template /usr/local/bin/
```

On macOS, if Gatekeeper blocks the binary, run `xattr -d com.apple.quarantine /usr/local/bin/schema-template`.

Windows (PowerShell):

```powershell
$Version = "0.1.0"
Invoke-WebRequest -OutFile schema-template.exe `
  "https://github.com/rad12000/ai-tools/releases/download/schema-template/v$Version/schema-template_windows_amd64.exe"
```

Move `schema-template.exe` to a directory on your `PATH`. On ARM machines, use `schema-template_windows_arm64.exe`.

### From source

Requires Go 1.25+:

```sh
go install github.com/rad12000/ai-tools/go/schema-template@latest
```

## Usage

```sh
schema-template get {template_name}                               # print the template's JSON Schema
echo '{"key": "value"}' | schema-template execute {template_name} # render the template
```
