package shim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/vantage-sh/terraform-provider-vantage/vantage"
)

func NewProvider() provider.Provider {
	prov := vantage.New()

	return prov
}