// package rules

// import (
// 	"testing"

// 	"github.com/terraform-linters/tflint-plugin-sdk/helper"
// )

// func Test_AWSResourceInProdRule(t *testing.T) {
// 	cases := []struct {
// 		Name     string
// 		Content  string
// 		FilePath string
// 		Expected helper.Issues
// 	}{
// 		{
// 			Name: "AWS T instance in production",
// 			Content: `
// resource "aws_t_instance" "example" {
//   instance_type = "t3.micro"
// }`,
// 			FilePath: "environments/prod/main.tf",
// 			Expected: helper.Issues{
// 				{
// 					Rule:    NewAWSResourceInProdRule(),
// 					Message: "Resource type 'aws_t_instance' is not allowed in production environment",
// 				},
// 			},
// 		},
// 		{
// 			Name: "AWS T instance in non-production",
// 			Content: `
// resource "aws_t_instance" "example" {
//   instance_type = "t3.micro"
// }`,
// 			FilePath: "environments/staging/main.tf",
// 			Expected: helper.Issues{},
// 		},
// 		{
// 			Name: "Non-T instance in production",
// 			Content: `
// resource "aws_m5_instance" "example" {
//   instance_type = "m5.large"
// }`,
// 			FilePath: "environments/production/main.tf",
// 			Expected: helper.Issues{},
// 		},
// 	}

// 	rule := NewAWSResourceInProdRule()

// 	for _, tc := range cases {
// 		runner := helper.TestRunner(t, map[string]string{tc.FilePath: tc.Content})

// 		if err := rule.Check(runner); err != nil {
// 			t.Fatalf("Unexpected error occurred: %s", err)
// 		}

// 		helper.AssertIssues(t, tc.Expected, runner.Issues)
// 	}
// }
