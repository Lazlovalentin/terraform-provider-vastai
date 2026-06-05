// Package provider is a thin, non-internal re-export of the Terraform
// Plugin Framework provider constructor so that external Go modules
// (e.g. upjet-based Crossplane providers) can embed the Vast.ai
// provider directly without spawning the `terraform` binary.
package provider

import (
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	internalprovider "github.com/realnedsanders/terraform-provider-vastai/internal/provider"
)

// New returns a constructor for the Vast.ai Terraform Plugin Framework
// provider. The semantics match the upstream `internal/provider.New`.
func New(version string) func() fwprovider.Provider {
	return internalprovider.New(version)
}
