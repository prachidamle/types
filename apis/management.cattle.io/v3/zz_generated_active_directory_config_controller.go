package v3

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ActiveDirectoryConfigGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ActiveDirectoryConfig",
	}
	ActiveDirectoryConfigResource = metav1.APIResource{
		Name:         "activedirectoryconfigs",
		SingularName: "activedirectoryconfig",
		Namespaced:   false,
		Kind:         ActiveDirectoryConfigGroupVersionKind.Kind,
	}
)

type ActiveDirectoryConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiveDirectoryConfig
}

type ActiveDirectoryConfigHandlerFunc func(key string, obj *ActiveDirectoryConfig) error

type ActiveDirectoryConfigLister interface {
	List(namespace string, selector labels.Selector) (ret []*ActiveDirectoryConfig, err error)
	Get(namespace, name string) (*ActiveDirectoryConfig, error)
}

type ActiveDirectoryConfigController interface {
	Informer() cache.SharedIndexInformer
	Lister() ActiveDirectoryConfigLister
	AddHandler(name string, handler ActiveDirectoryConfigHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler ActiveDirectoryConfigHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ActiveDirectoryConfigInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*ActiveDirectoryConfig) (*ActiveDirectoryConfig, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ActiveDirectoryConfig, error)
	Get(name string, opts metav1.GetOptions) (*ActiveDirectoryConfig, error)
	Update(*ActiveDirectoryConfig) (*ActiveDirectoryConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ActiveDirectoryConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ActiveDirectoryConfigController
	AddHandler(name string, sync ActiveDirectoryConfigHandlerFunc)
	AddLifecycle(name string, lifecycle ActiveDirectoryConfigLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync ActiveDirectoryConfigHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle ActiveDirectoryConfigLifecycle)
}

type activeDirectoryConfigLister struct {
	controller *activeDirectoryConfigController
}

func (l *activeDirectoryConfigLister) List(namespace string, selector labels.Selector) (ret []*ActiveDirectoryConfig, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ActiveDirectoryConfig))
	})
	return
}

func (l *activeDirectoryConfigLister) Get(namespace, name string) (*ActiveDirectoryConfig, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    ActiveDirectoryConfigGroupVersionKind.Group,
			Resource: "activeDirectoryConfig",
		}, name)
	}
	return obj.(*ActiveDirectoryConfig), nil
}

type activeDirectoryConfigController struct {
	controller.GenericController
}

func (c *activeDirectoryConfigController) Lister() ActiveDirectoryConfigLister {
	return &activeDirectoryConfigLister{
		controller: c,
	}
}

func (c *activeDirectoryConfigController) AddHandler(name string, handler ActiveDirectoryConfigHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ActiveDirectoryConfig))
	})
}

func (c *activeDirectoryConfigController) AddClusterScopedHandler(name, cluster string, handler ActiveDirectoryConfigHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}

		if !controller.ObjectInCluster(cluster, obj) {
			return nil
		}

		return handler(key, obj.(*ActiveDirectoryConfig))
	})
}

type activeDirectoryConfigFactory struct {
}

func (c activeDirectoryConfigFactory) Object() runtime.Object {
	return &ActiveDirectoryConfig{}
}

func (c activeDirectoryConfigFactory) List() runtime.Object {
	return &ActiveDirectoryConfigList{}
}

func (s *activeDirectoryConfigClient) Controller() ActiveDirectoryConfigController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.activeDirectoryConfigControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ActiveDirectoryConfigGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &activeDirectoryConfigController{
		GenericController: genericController,
	}

	s.client.activeDirectoryConfigControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type activeDirectoryConfigClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   ActiveDirectoryConfigController
}

func (s *activeDirectoryConfigClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *activeDirectoryConfigClient) Create(o *ActiveDirectoryConfig) (*ActiveDirectoryConfig, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ActiveDirectoryConfig), err
}

func (s *activeDirectoryConfigClient) Get(name string, opts metav1.GetOptions) (*ActiveDirectoryConfig, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ActiveDirectoryConfig), err
}

func (s *activeDirectoryConfigClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ActiveDirectoryConfig, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ActiveDirectoryConfig), err
}

func (s *activeDirectoryConfigClient) Update(o *ActiveDirectoryConfig) (*ActiveDirectoryConfig, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ActiveDirectoryConfig), err
}

func (s *activeDirectoryConfigClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *activeDirectoryConfigClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *activeDirectoryConfigClient) List(opts metav1.ListOptions) (*ActiveDirectoryConfigList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ActiveDirectoryConfigList), err
}

func (s *activeDirectoryConfigClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *activeDirectoryConfigClient) Patch(o *ActiveDirectoryConfig, data []byte, subresources ...string) (*ActiveDirectoryConfig, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*ActiveDirectoryConfig), err
}

func (s *activeDirectoryConfigClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *activeDirectoryConfigClient) AddHandler(name string, sync ActiveDirectoryConfigHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *activeDirectoryConfigClient) AddLifecycle(name string, lifecycle ActiveDirectoryConfigLifecycle) {
	sync := NewActiveDirectoryConfigLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *activeDirectoryConfigClient) AddClusterScopedHandler(name, clusterName string, sync ActiveDirectoryConfigHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *activeDirectoryConfigClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle ActiveDirectoryConfigLifecycle) {
	sync := NewActiveDirectoryConfigLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
