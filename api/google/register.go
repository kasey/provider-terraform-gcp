package google

import (
	"github.com/crossplane/terraform-provider-runtime/pkg/plugin"
)

const ProviderReferenceName string = "google"

func Index(idxr *plugin.Indexer) {
	for _, ft := range ResourceFuncTables {
		idxr.IndexFuncTable(ft)
	}
}

func ProviderInit() *plugin.ProviderInit {
	return providerEntry
}
