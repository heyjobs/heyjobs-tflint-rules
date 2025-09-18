package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_BackendProviderNamingConventionRule(t *testing.T) {
	cases := []struct {
		Name     string
		Files    map[string]string
		Expected helper.Issues
	}{
		{
			Name: "provider in wrong file",
			Files: map[string]string{
				"main.tf": `
provider "aws" {
  region = "us-west-2"
}`,
			},
			Expected: helper.Issues{
				{
					Rule:    NewBackendProviderNamingConventionRule(),
					Message: "Provider configuration found in main.tf. Provider configurations should be in provider.tf",
					Range:   hcl.Range{Filename: "main.tf", Start: hcl.Pos{Line: 2, Column: 1}, End: hcl.Pos{Line: 2, Column: 15}},
				},
			},
		},
		{
			Name: "provider in correct file",
			Files: map[string]string{
				"provider.tf": `
provider "aws" {
  region = "us-west-2"
}`,
			},
			Expected: helper.Issues{},
		},
		{
			Name: "backend in wrong file",
			Files: map[string]string{
				"main.tf": `
terraform {
  backend "s3" {
    bucket = "my-bucket"
  }
}`,
			},
			Expected: helper.Issues{
				{
					Rule:    NewBackendProviderNamingConventionRule(),
					Message: "Backend configuration found in main.tf. Backend configurations should be in backend.tf",
					Range:   hcl.Range{Filename: "main.tf", Start: hcl.Pos{Line: 2, Column: 1}, End: hcl.Pos{Line: 2, Column: 10}},
				},
			},
		},
		{
			Name: "backend in correct file",
			Files: map[string]string{
				"backend.tf": `
terraform {
  backend "s3" {
    bucket = "my-bucket"
  }
}`,
			},
			Expected: helper.Issues{},
		},
		{
			Name: "multiple files with wrong configurations",
			Files: map[string]string{
				"infrastructure.tf": `
provider "aws" {
  region = "us-west-2"
}`,
				"setup.tf": `
terraform {
  backend "s3" {
    bucket = "my-bucket"
  }
}`,
			},
			Expected: helper.Issues{
				{
					Rule:    NewBackendProviderNamingConventionRule(),
					Message: "Provider configuration found in infrastructure.tf. Provider configurations should be in provider.tf",
					Range:   hcl.Range{Filename: "infrastructure.tf", Start: hcl.Pos{Line: 2, Column: 1}, End: hcl.Pos{Line: 2, Column: 15}},
				},
				{
					Rule:    NewBackendProviderNamingConventionRule(),
					Message: "Backend configuration found in setup.tf. Backend configurations should be in backend.tf",
					Range:   hcl.Range{Filename: "setup.tf", Start: hcl.Pos{Line: 2, Column: 1}, End: hcl.Pos{Line: 2, Column: 10}},
				},
			},
		},
		{
			Name: "terraform block without backend",
			Files: map[string]string{
				"main.tf": `
terraform {
  required_version = "~> 1.0.0"
}`,
			},
			Expected: helper.Issues{},
		},
	}

	rule := NewBackendProviderNamingConventionRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, tc.Files)
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
} 