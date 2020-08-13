package google

import (
	gcp "github.com/crossplane/provider-terraform-gcp/api/google/v1alpha1"
	"github.com/crossplane/terraform-provider-runtime/pkg/plugin"
)

const ProviderReferenceName string = "google"

func Index(idxr *plugin.Indexer) {
	for _, impl := range ResourceImplementations {
		idxr.Overlay(impl)
	}
}

func ProviderInit() *plugin.ProviderInit {
	return gcp.GetProviderInit()
}
