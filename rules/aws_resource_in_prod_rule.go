package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AWSResourceInProdRule checks if specific AWS resource types are used in production environments
type AWSResourceInProdRule struct {
	tflint.DefaultRule
}

// NewAWSResourceInProdRule returns a new rule
func NewAWSResourceInProdRule() *AWSResourceInProdRule {
	return &AWSResourceInProdRule{}
}

// Name returns the rule name
func (r *AWSResourceInProdRule) Name() string {
	return "aws_resource_in_prod"
}

// Enabled returns whether the rule is enabled by default
func (r *AWSResourceInProdRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AWSResourceInProdRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Check runs the rule
func (r *AWSResourceInProdRule) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	// List of AWS resource types to check (can be customized)
	restrictedTypes := []string{
		"aws_t_", // Any T family instance
	}

	for _, block := range content.Blocks {
		// Check if we're in a production environment by looking at the file path
		filePath := strings.ToLower(block.DefRange.Filename)
		isProd := strings.Contains(filePath, "/prod/") || strings.Contains(filePath, "/production/")

		if !isProd {
			continue
		}

		if len(block.Labels) > 1 {
			resourceType := block.Labels[0]

			// Check if the resource type matches any restricted type
			for _, restrictedType := range restrictedTypes {
				if strings.HasPrefix(resourceType, restrictedType) {
					runner.EmitIssue(
						r,
						fmt.Sprintf("Resource type '%s' is not allowed in production environment", resourceType),
						block.DefRange,
					)
					break
				}
			}
		}
	}

	return nil
}
