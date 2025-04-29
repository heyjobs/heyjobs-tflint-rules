package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AWSResourceInProdRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		FilePath string
		Expected helper.Issues
	}{
		{
			Name: "AWS T instance in production",
			Content: `
resource "aws_t_instance" "example" {
  instance_type = "t3.micro"
}`,
			FilePath: "environments/prod/main.tf",
			Expected: helper.Issues{
				{
					Rule:    NewAWSResourceInProdRule(),
					Message: "Resource type 'aws_t_instance' is not allowed in production environment",
					Range:   hcl.Range{Filename: "environments/prod/main.tf", Start: hcl.Pos{Line: 2, Column: 1}, End: hcl.Pos{Line: 2, Column: 36}},
				},
			},
		},
		{
			Name: "AWS T instance in non-production",
			Content: `
resource "aws_t_instance" "example" {
  instance_type = "t3.micro"
}`,
			FilePath: "environments/staging/main.tf",
			Expected: helper.Issues{},
		},
		{
			Name: "Non-T instance in production",
			Content: `
resource "aws_m5_instance" "example" {
  instance_type = "m5.large"
}`,
			FilePath: "environments/production/main.tf",
			Expected: helper.Issues{},
		},
		{
			Name: "Multiple T instances in production",
			Content: `
resource "aws_t2_instance" "example1" {
  instance_type = "t2.micro"
}

resource "aws_t3_instance" "example2" {
  instance_type = "t3.micro"
}`,
			FilePath: "environments/prod/main.tf",
			Expected: helper.Issues{
				{
					Rule:    NewAWSResourceInProdRule(),
					Message: "Resource type 'aws_t2_instance' is not allowed in production environment",
					Range:   hcl.Range{Filename: "environments/prod/main.tf", Start: hcl.Pos{Line: 2, Column: 1}, End: hcl.Pos{Line: 2, Column: 38}},
				},
				{
					Rule:    NewAWSResourceInProdRule(),
					Message: "Resource type 'aws_t3_instance' is not allowed in production environment",
					Range:   hcl.Range{Filename: "environments/prod/main.tf", Start: hcl.Pos{Line: 6, Column: 1}, End: hcl.Pos{Line: 6, Column: 38}},
				},
			},
		},
	}

	rule := NewAWSResourceInProdRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{tc.FilePath: tc.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
