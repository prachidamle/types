package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ActiveDirectoryConfigLifecycle interface {
	Create(obj *ActiveDirectoryConfig) (*ActiveDirectoryConfig, error)
	Remove(obj *ActiveDirectoryConfig) (*ActiveDirectoryConfig, error)
	Updated(obj *ActiveDirectoryConfig) (*ActiveDirectoryConfig, error)
}

type activeDirectoryConfigLifecycleAdapter struct {
	lifecycle ActiveDirectoryConfigLifecycle
}

func (w *activeDirectoryConfigLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ActiveDirectoryConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *activeDirectoryConfigLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ActiveDirectoryConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *activeDirectoryConfigLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ActiveDirectoryConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewActiveDirectoryConfigLifecycleAdapter(name string, clusterScoped bool, client ActiveDirectoryConfigInterface, l ActiveDirectoryConfigLifecycle) ActiveDirectoryConfigHandlerFunc {
	adapter := &activeDirectoryConfigLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ActiveDirectoryConfig) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
