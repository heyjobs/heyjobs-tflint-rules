package rules

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_StackTmHclRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing stack.tm.hcl in non-module directory",
			Content: `
resource "aws_s3_bucket" "test" {
  bucket = "test-bucket"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewStackTmHclRule(),
					Message: "stack.tm.hcl file is required in non-module directory: /test",
				},
			},
		},
		{
			Name: "in module directory - should not trigger",
			Content: `
resource "aws_s3_bucket" "test" {
  bucket = "test-bucket"
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewStackTmHclRule()

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{
				"main.tf": tc.Content,
			})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
