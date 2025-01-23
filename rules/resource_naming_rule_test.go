package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_TerraformNamingConventionRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "resource with dash in name",
			Content: `
resource "aws_instance" "my-instance" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTerraformNamingConventionRule(),
					Message: "Resource name 'my-instance' contains dashes. Use underscores instead.",
					Range:   hcl.Range{Filename: "main.tf", Start: hcl.Pos{Line: 2, Column: 1}, End: hcl.Pos{Line: 2, Column: 38}},
				},
			},
		},
		{
			Name: "resource with underscore in name",
			Content: `
resource "aws_instance" "my_instance" {
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "multiple resources with mixed naming",
			Content: `
resource "aws_instance" "my-bad-instance" {
}

resource "aws_s3_bucket" "my_good_bucket" {
}

resource "aws_lambda" "another-bad-name" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTerraformNamingConventionRule(),
					Message: "Resource name 'my-bad-instance' contains dashes. Use underscores instead.",
					Range:   hcl.Range{Filename: "main.tf", Start: hcl.Pos{Line: 2, Column: 1}, End: hcl.Pos{Line: 2, Column: 42}},
				},
				{
					Rule:    NewTerraformNamingConventionRule(),
					Message: "Resource name 'another-bad-name' contains dashes. Use underscores instead.",
					Range:   hcl.Range{Filename: "main.tf", Start: hcl.Pos{Line: 8, Column: 1}, End: hcl.Pos{Line: 8, Column: 41}},
				},
			},
		},
	}

	rule := NewTerraformNamingConventionRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tc.Content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}