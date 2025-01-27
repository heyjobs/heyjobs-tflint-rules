# HeyJobs TFLint Rules

Custom TFLint ruleset for HeyJobs' Terraform configurations. This ruleset helps enforce best practices and standards across our infrastructure code.

## Requirements

- TFLint v0.42+
- Go v1.23

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
```

2. Build the plugin:
```bash
make
```

3. Install locally:
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
