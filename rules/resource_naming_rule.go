package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TerraformNamingConventionRule checks if resource names use underscores instead of dashes
type TerraformNamingConventionRule struct {
	tflint.DefaultRule
}

// NewTerraformNamingConventionRule returns a new rule
func NewTerraformNamingConventionRule() *TerraformNamingConventionRule {
	return &TerraformNamingConventionRule{}
}

// Name returns the rule name
func (r *TerraformNamingConventionRule) Name() string {
	return "terraform_naming_convention"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformNamingConventionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TerraformNamingConventionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Check runs the rule
func (r *TerraformNamingConventionRule) Check(runner tflint.Runner) error {
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

	for _, block := range content.Blocks {
		if len(block.Labels) > 1 {
			resourceName := block.Labels[1]
			if strings.Contains(resourceName, "-") {
				runner.EmitIssue(
					r,
					fmt.Sprintf("Resource name '%s' contains dashes. Use underscores instead.", resourceName),
					block.DefRange,
				)
			}
		}
	}

	return nil
}