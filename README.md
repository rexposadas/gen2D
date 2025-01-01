[![Go Report Card](https://goreportcard.com/badge/github.com/rexposadas/gen2D)](https://goreportcard.com/report/github.com/rexposadas/gen2D)

# Gen2D
Gen2D is a command-line interface (CLI) tool that simplifies the generation of 2D art in Go.

This tool is built on top of the [generativeart](https://github.com/jdxyw/generativeart) library.

# Quick Start

Build and run the application:

```bash
go build -o ./gen2D && ./gen2D circles
```

This command creates a directory called "output" containing your generated artwork. 

On macOS, you can view the generated files using:
```bash
open output/*
```

# Available Commands

To see all available options, run:
```bash
./gen2D
```

Some example commands to try:
```bash
./gen2D random-shapes
./gen2D perls
```

# Using the Sample Makefile

The repository includes a sample Makefile to help you generate art. To create artwork using all the commands in the Makefile:

```bash
make -f sample-Makefile all
```

# Configuration Files

You can customize your artwork using configuration files. Here's an example configuration file (in the input folder):

```json
{
  "out": {
    "dir": "output"
  },
  "canvas": {
    "width": 300,
    "height": 300
  }
}
```

The configuration file allows you to set:
- Output directory location
- Canvas dimensions

To use a configuration file:

```bash
./gen2D circles grid -c 10 -f input/config.json
```

This command generates 10 images using the specified circle configuration.

# Contributing

To contribute:
1. Create an issue describing the proposed change
2. Submit a pull request with your implementation

# Planned Features

- [ ] Generate a configuration file for each image, documenting the parameters used for future recreation
- [ ] Add a flag to specify the output directory
- [ ] Create subdirectories with timestamps for each generation run
- [ ] Switch from JSON to TOML configuration files for improved readability and conciseness

# License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.