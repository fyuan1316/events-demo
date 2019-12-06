/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mlv1alpha1 "github.com/fyuan1316/events/api/v1alpha1"
)

// SourceWatcherReconciler reconciles a SourceWatcher object
type SourceWatcherReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=ml.fyuan.io,resources=sourcewatchers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ml.fyuan.io,resources=sourcewatchers/status,verbs=get;update;patch

func (r *SourceWatcherReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("sourcewatcher", req.NamespacedName)
	// your logic here
	watcher := mlv1alpha1.SourceWatcher{}
	//if err := r.Client.Get(ctx, req.NamespacedName, &watcher); err != nil {
	//	log.Error(err, "failed to get MyKind resource")
	//	// Ignore NotFound errors as they will be retried automatically if the
	//	// resource is created in future.
	//	return ctrl.Result{}, client.IgnoreNotFound(err)
	//}
	r.Recorder.Eventf(&watcher, core.EventTypeNormal, "created", "detail", "res-name")

	return ctrl.Result{}, nil
}

func (r *SourceWatcherReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mlv1alpha1.SourceWatcher{}).
		Complete(r)
}
