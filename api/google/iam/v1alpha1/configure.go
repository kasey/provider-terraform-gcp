package v1alpha1

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/terraform-provider-runtime/pkg/client"
	"github.com/crossplane/terraform-provider-runtime/pkg/controller/terraform"
	"github.com/crossplane/terraform-provider-runtime/pkg/registry"
	ctrl "sigs.k8s.io/controller-runtime"
)

// ConfigureReconciler adds a controller that reconciles the autogenerated managed.Resources in this package
func ConfigureReconciler(mgr ctrl.Manager, l logging.Logger, reg *registry.Registry, pool *client.ProviderPool) error {
	name := managed.ControllerName(ServiceAccountGroupKind)
	r := managed.NewReconciler(mgr,
		resource.ManagedKind(ServiceAccountGroupVersionKind),
		managed.WithTimeout(time.Duration(3600*time.Second)),
		managed.WithExternalConnecter(&terraform.Connector{KubeClient: mgr.GetClient(), Registry: reg, Logger: l, Pool: pool}),
		managed.WithInitializers(managed.NewNameAsExternalName(mgr.GetClient())),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&ServiceAccount{}).
		Complete(r)
}
