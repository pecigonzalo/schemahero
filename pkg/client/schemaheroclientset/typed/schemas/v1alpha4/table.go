/*
Copyright 2021 The SchemaHero Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha4

import (
	"context"
	"time"

	v1alpha4 "github.com/schemahero/schemahero/pkg/apis/schemas/v1alpha4"
	scheme "github.com/schemahero/schemahero/pkg/client/schemaheroclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TablesGetter has a method to return a TableInterface.
// A group's client should implement this interface.
type TablesGetter interface {
	Tables(namespace string) TableInterface
}

// TableInterface has methods to work with Table resources.
type TableInterface interface {
	Create(ctx context.Context, table *v1alpha4.Table, opts v1.CreateOptions) (*v1alpha4.Table, error)
	Update(ctx context.Context, table *v1alpha4.Table, opts v1.UpdateOptions) (*v1alpha4.Table, error)
	UpdateStatus(ctx context.Context, table *v1alpha4.Table, opts v1.UpdateOptions) (*v1alpha4.Table, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha4.Table, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha4.TableList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha4.Table, err error)
	TableExpansion
}

// tables implements TableInterface
type tables struct {
	client rest.Interface
	ns     string
}

// newTables returns a Tables
func newTables(c *SchemasV1alpha4Client, namespace string) *tables {
	return &tables{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the table, and returns the corresponding table object, and an error if there is any.
func (c *tables) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha4.Table, err error) {
	result = &v1alpha4.Table{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Tables that match those selectors.
func (c *tables) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha4.TableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha4.TableList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tables.
func (c *tables) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a table and creates it.  Returns the server's representation of the table, and an error, if there is any.
func (c *tables) Create(ctx context.Context, table *v1alpha4.Table, opts v1.CreateOptions) (result *v1alpha4.Table, err error) {
	result = &v1alpha4.Table{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(table).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a table and updates it. Returns the server's representation of the table, and an error, if there is any.
func (c *tables) Update(ctx context.Context, table *v1alpha4.Table, opts v1.UpdateOptions) (result *v1alpha4.Table, err error) {
	result = &v1alpha4.Table{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tables").
		Name(table.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(table).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tables) UpdateStatus(ctx context.Context, table *v1alpha4.Table, opts v1.UpdateOptions) (result *v1alpha4.Table, err error) {
	result = &v1alpha4.Table{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tables").
		Name(table.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(table).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the table and deletes it. Returns an error if one occurs.
func (c *tables) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tables").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tables) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tables").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched table.
func (c *tables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha4.Table, err error) {
	result = &v1alpha4.Table{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tables").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
