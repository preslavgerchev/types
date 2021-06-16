# Podkrepi.bg types
This repository contains the protobuf contracts and the generated Go types used by podkrepi.bg.

- [Linting](#linting)
- [Usage](#usage)
    + [Go](#go)
    + [C#](#c-)
- [Updating the protos](#updating-the-protos)

## Linting
Before pushing the repo make sure you lint your proto files. You can do this by running:
```bash
task lint_proto
```

## Usage
The section explains how to use the types in different languages.

### Go
To use the types in a Go project run:
```bash
go get github.com/daritelska-platforma/types/go-types
```

### C#
TODO

## Updating the protos
The easiest way to update the protos is:
- Open the repo in Visual Studio code
- Click Ctrl/Cmd + Shift + P and select `Remote-Containers: Reopen in Container`
- The workspace will be automatically opened in a Docker container with all required dependencies to update and build the proto contracts

After the proto files are updated you can run the following command to update the generated Go types:
```bash
task build_go_types
```