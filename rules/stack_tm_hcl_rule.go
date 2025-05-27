package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// StackTmHclRule checks if a stack.tm.hcl file exists in non-module directories
type StackTmHclRule struct {
	tflint.DefaultRule
}

// NewStackTmHclRule returns a new rule
func NewStackTmHclRule() *StackTmHclRule {
	return &StackTmHclRule{}
}

// Name returns the rule name
func (r *StackTmHclRule) Name() string {
	return "stack_tm_hcl"
}

// Enabled returns whether the rule is enabled by default
func (r *StackTmHclRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *StackTmHclRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Check runs the rule
func (r *StackTmHclRule) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{}, nil)
	if err != nil {
		return err
	}

	// Get the current file's directory
	currentFile := content.Blocks[0].DefRange.Filename
	dir := filepath.Dir(currentFile)

	// Skip if the path contains "modules"
	if strings.Contains(dir, "modules") {
		return nil
	}

	// Check if stack.tm.hcl exists in the directory
	stackFile := filepath.Join(dir, "stack.tm.hcl")
	if _, err := os.Stat(stackFile); os.IsNotExist(err) {
		runner.EmitIssue(
			r,
			fmt.Sprintf("stack.tm.hcl file is required in non-module directory: %s", dir),
			content.Blocks[0].DefRange,
		)
	}

	return nil
}
