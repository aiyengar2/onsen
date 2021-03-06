/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/aiyengar2/onsen/pkg/apis/onsen.cattle.io/v1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type ClusterBootstrapConfigHandler func(string, *v1.ClusterBootstrapConfig) (*v1.ClusterBootstrapConfig, error)

type ClusterBootstrapConfigController interface {
	generic.ControllerMeta
	ClusterBootstrapConfigClient

	OnChange(ctx context.Context, name string, sync ClusterBootstrapConfigHandler)
	OnRemove(ctx context.Context, name string, sync ClusterBootstrapConfigHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ClusterBootstrapConfigCache
}

type ClusterBootstrapConfigClient interface {
	Create(*v1.ClusterBootstrapConfig) (*v1.ClusterBootstrapConfig, error)
	Update(*v1.ClusterBootstrapConfig) (*v1.ClusterBootstrapConfig, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.ClusterBootstrapConfig, error)
	List(namespace string, opts metav1.ListOptions) (*v1.ClusterBootstrapConfigList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterBootstrapConfig, err error)
}

type ClusterBootstrapConfigCache interface {
	Get(namespace, name string) (*v1.ClusterBootstrapConfig, error)
	List(namespace string, selector labels.Selector) ([]*v1.ClusterBootstrapConfig, error)

	AddIndexer(indexName string, indexer ClusterBootstrapConfigIndexer)
	GetByIndex(indexName, key string) ([]*v1.ClusterBootstrapConfig, error)
}

type ClusterBootstrapConfigIndexer func(obj *v1.ClusterBootstrapConfig) ([]string, error)

type clusterBootstrapConfigController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewClusterBootstrapConfigController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ClusterBootstrapConfigController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &clusterBootstrapConfigController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromClusterBootstrapConfigHandlerToHandler(sync ClusterBootstrapConfigHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.ClusterBootstrapConfig
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.ClusterBootstrapConfig))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *clusterBootstrapConfigController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.ClusterBootstrapConfig))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateClusterBootstrapConfigDeepCopyOnChange(client ClusterBootstrapConfigClient, obj *v1.ClusterBootstrapConfig, handler func(obj *v1.ClusterBootstrapConfig) (*v1.ClusterBootstrapConfig, error)) (*v1.ClusterBootstrapConfig, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *clusterBootstrapConfigController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *clusterBootstrapConfigController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *clusterBootstrapConfigController) OnChange(ctx context.Context, name string, sync ClusterBootstrapConfigHandler) {
	c.AddGenericHandler(ctx, name, FromClusterBootstrapConfigHandlerToHandler(sync))
}

func (c *clusterBootstrapConfigController) OnRemove(ctx context.Context, name string, sync ClusterBootstrapConfigHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromClusterBootstrapConfigHandlerToHandler(sync)))
}

func (c *clusterBootstrapConfigController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *clusterBootstrapConfigController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *clusterBootstrapConfigController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *clusterBootstrapConfigController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *clusterBootstrapConfigController) Cache() ClusterBootstrapConfigCache {
	return &clusterBootstrapConfigCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *clusterBootstrapConfigController) Create(obj *v1.ClusterBootstrapConfig) (*v1.ClusterBootstrapConfig, error) {
	result := &v1.ClusterBootstrapConfig{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *clusterBootstrapConfigController) Update(obj *v1.ClusterBootstrapConfig) (*v1.ClusterBootstrapConfig, error) {
	result := &v1.ClusterBootstrapConfig{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *clusterBootstrapConfigController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *clusterBootstrapConfigController) Get(namespace, name string, options metav1.GetOptions) (*v1.ClusterBootstrapConfig, error) {
	result := &v1.ClusterBootstrapConfig{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *clusterBootstrapConfigController) List(namespace string, opts metav1.ListOptions) (*v1.ClusterBootstrapConfigList, error) {
	result := &v1.ClusterBootstrapConfigList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *clusterBootstrapConfigController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *clusterBootstrapConfigController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.ClusterBootstrapConfig, error) {
	result := &v1.ClusterBootstrapConfig{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type clusterBootstrapConfigCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *clusterBootstrapConfigCache) Get(namespace, name string) (*v1.ClusterBootstrapConfig, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.ClusterBootstrapConfig), nil
}

func (c *clusterBootstrapConfigCache) List(namespace string, selector labels.Selector) (ret []*v1.ClusterBootstrapConfig, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterBootstrapConfig))
	})

	return ret, err
}

func (c *clusterBootstrapConfigCache) AddIndexer(indexName string, indexer ClusterBootstrapConfigIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.ClusterBootstrapConfig))
		},
	}))
}

func (c *clusterBootstrapConfigCache) GetByIndex(indexName, key string) (result []*v1.ClusterBootstrapConfig, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.ClusterBootstrapConfig, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.ClusterBootstrapConfig))
	}
	return result, nil
}
