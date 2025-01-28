# HeyJobs TFLint Rules

Custom TFLint ruleset for HeyJobs' Terraform configurations. This ruleset helps enforce best practices and standards across our infrastructure code.

## Requirements

- TFLint v0.42+
- Go v1.23

## Development Setup

### Installing Dependencies

We use Homebrew for managing development dependencies. A `Brewfile` is provided with all necessary tools:

```bash
# Install Homebrew if you haven't already
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install all dependencies
brew bundle
```

The Brewfile includes:
- `direnv` - Environment variable manager
- `go` - Go programming language
- `tflint` - Terraform linter
- `git` - Version control
- `goreleaser` - Release automation tool
- `make` - Build automation
- `golangci-lint` - Go linters aggregator
- `pre-commit` - Git hooks framework

### Environment Setup

The repository includes a `.envrc` file for direnv that sets up the Go environment:

```bash
# Allow the direnv configuration
direnv allow
```

This will automatically set up your Go environment variables when you enter the directory.

## Installation

Add the following configuration to your `.tflint.hcl`:

```hcl
plugin "heyjobs" {
    enabled = true
    version = "0.1.0"
    source  = "github.com/heyjobs/heyjobs-tflint-rules"
}
```

## Available Rules

| Name | Description | Severity | Enabled |
|------|-------------|----------|---------|
| aws_instance_example_type | Validates instance types against allowed values | ERROR | ✔ |
| aws_s3_bucket_example_lifecycle_rule | Ensures S3 buckets have proper lifecycle rules | ERROR | ✔ |
| terraform_backend_type | Validates backend configuration | ERROR | ✔ |

## Development

### Building the Plugin

1. Clone the repository:
```bash
git clone git@github.com:heyjobs/heyjobs-tflint-rules.git
cd heyjobs-tflint-rules
```

2. Install dependencies:
```bash
brew bundle
```

3. Build the plugin:
```bash
make
```

4. Install locally:
```bash
make install
```

### Testing

Run the test suite:
```bash
make test
```

### Local Testing

To test the plugin locally:

1. Create a `.tflint.hcl` file:
```hcl
plugin "heyjobs" {
    enabled = true
}
```

2. Run TFLint:
```bash
tflint
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

Copyright © 2024 HeyJobs GmbH 