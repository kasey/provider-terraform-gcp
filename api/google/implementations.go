package google

import (
	iam "github.com/crossplane/provider-terraform-gcp/api/google/iam/v1alpha1"
	gcp "github.com/crossplane/provider-terraform-gcp/api/google/v1alpha1"
	"github.com/crossplane/terraform-provider-runtime/pkg/plugin"
)

// TODO: output this list somewhere in the codegen pipeline
var ResourceImplementations = []*plugin.Implementation{
	iam.ServiceAccountImplementation(),
}

var providerEntry = gcp.GetProviderEntry()