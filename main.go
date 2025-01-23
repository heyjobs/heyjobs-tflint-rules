package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/heyjobs/heyjobs-tflint-rules/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "heyjobs",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewTerraformNamingConventionRule(),
				rules.NewBackendProviderNamingConventionRule(),
			},
		},
	})
}
