package rules

import (
	"fmt"
	"path/filepath"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// BackendProviderNamingConventionRule checks if provider and backend configurations are in correctly named files
type BackendProviderNamingConventionRule struct {
	tflint.DefaultRule
}

// NewBackendProviderNamingConventionRule returns a new rule
func NewBackendProviderNamingConventionRule() *BackendProviderNamingConventionRule {
	return &BackendProviderNamingConventionRule{}
}

// Name returns the rule name
func (r *BackendProviderNamingConventionRule) Name() string {
	return "backend_provider_naming_convention"
}

// Enabled returns whether the rule is enabled by default
func (r *BackendProviderNamingConventionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *BackendProviderNamingConventionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Check runs the rule
func (r *BackendProviderNamingConventionRule) Check(runner tflint.Runner) error {
	// Check for provider blocks
	providers, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "provider",
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, provider := range providers.Blocks {
		filename := filepath.Base(provider.DefRange.Filename)
		if filename != "provider.tf" {
			runner.EmitIssue(
				r,
				fmt.Sprintf("Provider configuration found in %s. Provider configurations should be in provider.tf", filename),
				provider.DefRange,
			)
		}
	}

	// Check for terraform blocks with backend configuration
	terraformBlocks, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "terraform",
				Body: &hclext.BodySchema{
					Blocks: []hclext.BlockSchema{
						{
							Type: "backend",
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range terraformBlocks.Blocks {
		if len(block.Body.Blocks) > 0 { // Has backend configuration
			filename := filepath.Base(block.DefRange.Filename)
			if filename != "backend.tf" {
				runner.EmitIssue(
					r,
					fmt.Sprintf("Backend configuration found in %s. Backend configurations should be in backend.tf", filename),
					block.DefRange,
				)
			}
		}
	}

	return nil
} 