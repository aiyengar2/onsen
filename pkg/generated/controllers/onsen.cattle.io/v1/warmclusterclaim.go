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
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
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

type WarmClusterClaimHandler func(string, *v1.WarmClusterClaim) (*v1.WarmClusterClaim, error)

type WarmClusterClaimController interface {
	generic.ControllerMeta
	WarmClusterClaimClient

	OnChange(ctx context.Context, name string, sync WarmClusterClaimHandler)
	OnRemove(ctx context.Context, name string, sync WarmClusterClaimHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() WarmClusterClaimCache
}

type WarmClusterClaimClient interface {
	Create(*v1.WarmClusterClaim) (*v1.WarmClusterClaim, error)
	Update(*v1.WarmClusterClaim) (*v1.WarmClusterClaim, error)
	UpdateStatus(*v1.WarmClusterClaim) (*v1.WarmClusterClaim, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.WarmClusterClaim, error)
	List(namespace string, opts metav1.ListOptions) (*v1.WarmClusterClaimList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.WarmClusterClaim, err error)
}

type WarmClusterClaimCache interface {
	Get(namespace, name string) (*v1.WarmClusterClaim, error)
	List(namespace string, selector labels.Selector) ([]*v1.WarmClusterClaim, error)

	AddIndexer(indexName string, indexer WarmClusterClaimIndexer)
	GetByIndex(indexName, key string) ([]*v1.WarmClusterClaim, error)
}

type WarmClusterClaimIndexer func(obj *v1.WarmClusterClaim) ([]string, error)

type warmClusterClaimController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewWarmClusterClaimController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) WarmClusterClaimController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &warmClusterClaimController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromWarmClusterClaimHandlerToHandler(sync WarmClusterClaimHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.WarmClusterClaim
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.WarmClusterClaim))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *warmClusterClaimController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.WarmClusterClaim))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateWarmClusterClaimDeepCopyOnChange(client WarmClusterClaimClient, obj *v1.WarmClusterClaim, handler func(obj *v1.WarmClusterClaim) (*v1.WarmClusterClaim, error)) (*v1.WarmClusterClaim, error) {
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

func (c *warmClusterClaimController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *warmClusterClaimController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *warmClusterClaimController) OnChange(ctx context.Context, name string, sync WarmClusterClaimHandler) {
	c.AddGenericHandler(ctx, name, FromWarmClusterClaimHandlerToHandler(sync))
}

func (c *warmClusterClaimController) OnRemove(ctx context.Context, name string, sync WarmClusterClaimHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromWarmClusterClaimHandlerToHandler(sync)))
}

func (c *warmClusterClaimController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *warmClusterClaimController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *warmClusterClaimController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *warmClusterClaimController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *warmClusterClaimController) Cache() WarmClusterClaimCache {
	return &warmClusterClaimCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *warmClusterClaimController) Create(obj *v1.WarmClusterClaim) (*v1.WarmClusterClaim, error) {
	result := &v1.WarmClusterClaim{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *warmClusterClaimController) Update(obj *v1.WarmClusterClaim) (*v1.WarmClusterClaim, error) {
	result := &v1.WarmClusterClaim{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *warmClusterClaimController) UpdateStatus(obj *v1.WarmClusterClaim) (*v1.WarmClusterClaim, error) {
	result := &v1.WarmClusterClaim{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *warmClusterClaimController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *warmClusterClaimController) Get(namespace, name string, options metav1.GetOptions) (*v1.WarmClusterClaim, error) {
	result := &v1.WarmClusterClaim{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *warmClusterClaimController) List(namespace string, opts metav1.ListOptions) (*v1.WarmClusterClaimList, error) {
	result := &v1.WarmClusterClaimList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *warmClusterClaimController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *warmClusterClaimController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.WarmClusterClaim, error) {
	result := &v1.WarmClusterClaim{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type warmClusterClaimCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *warmClusterClaimCache) Get(namespace, name string) (*v1.WarmClusterClaim, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.WarmClusterClaim), nil
}

func (c *warmClusterClaimCache) List(namespace string, selector labels.Selector) (ret []*v1.WarmClusterClaim, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.WarmClusterClaim))
	})

	return ret, err
}

func (c *warmClusterClaimCache) AddIndexer(indexName string, indexer WarmClusterClaimIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.WarmClusterClaim))
		},
	}))
}

func (c *warmClusterClaimCache) GetByIndex(indexName, key string) (result []*v1.WarmClusterClaim, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.WarmClusterClaim, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.WarmClusterClaim))
	}
	return result, nil
}

type WarmClusterClaimStatusHandler func(obj *v1.WarmClusterClaim, status v1.WarmClusterClaimStatus) (v1.WarmClusterClaimStatus, error)

type WarmClusterClaimGeneratingHandler func(obj *v1.WarmClusterClaim, status v1.WarmClusterClaimStatus) ([]runtime.Object, v1.WarmClusterClaimStatus, error)

func RegisterWarmClusterClaimStatusHandler(ctx context.Context, controller WarmClusterClaimController, condition condition.Cond, name string, handler WarmClusterClaimStatusHandler) {
	statusHandler := &warmClusterClaimStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromWarmClusterClaimHandlerToHandler(statusHandler.sync))
}

func RegisterWarmClusterClaimGeneratingHandler(ctx context.Context, controller WarmClusterClaimController, apply apply.Apply,
	condition condition.Cond, name string, handler WarmClusterClaimGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &warmClusterClaimGeneratingHandler{
		WarmClusterClaimGeneratingHandler: handler,
		apply:                             apply,
		name:                              name,
		gvk:                               controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterWarmClusterClaimStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type warmClusterClaimStatusHandler struct {
	client    WarmClusterClaimClient
	condition condition.Cond
	handler   WarmClusterClaimStatusHandler
}

func (a *warmClusterClaimStatusHandler) sync(key string, obj *v1.WarmClusterClaim) (*v1.WarmClusterClaim, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type warmClusterClaimGeneratingHandler struct {
	WarmClusterClaimGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *warmClusterClaimGeneratingHandler) Remove(key string, obj *v1.WarmClusterClaim) (*v1.WarmClusterClaim, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.WarmClusterClaim{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *warmClusterClaimGeneratingHandler) Handle(obj *v1.WarmClusterClaim, status v1.WarmClusterClaimStatus) (v1.WarmClusterClaimStatus, error) {
	objs, newStatus, err := a.WarmClusterClaimGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
