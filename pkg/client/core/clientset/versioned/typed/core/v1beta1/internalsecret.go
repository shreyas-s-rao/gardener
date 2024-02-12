// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	scheme "github.com/gardener/gardener/pkg/client/core/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InternalSecretsGetter has a method to return a InternalSecretInterface.
// A group's client should implement this interface.
type InternalSecretsGetter interface {
	InternalSecrets(namespace string) InternalSecretInterface
}

// InternalSecretInterface has methods to work with InternalSecret resources.
type InternalSecretInterface interface {
	Create(ctx context.Context, internalSecret *v1beta1.InternalSecret, opts v1.CreateOptions) (*v1beta1.InternalSecret, error)
	Update(ctx context.Context, internalSecret *v1beta1.InternalSecret, opts v1.UpdateOptions) (*v1beta1.InternalSecret, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.InternalSecret, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.InternalSecretList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.InternalSecret, err error)
	InternalSecretExpansion
}

// internalSecrets implements InternalSecretInterface
type internalSecrets struct {
	client rest.Interface
	ns     string
}

// newInternalSecrets returns a InternalSecrets
func newInternalSecrets(c *CoreV1beta1Client, namespace string) *internalSecrets {
	return &internalSecrets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the internalSecret, and returns the corresponding internalSecret object, and an error if there is any.
func (c *internalSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.InternalSecret, err error) {
	result = &v1beta1.InternalSecret{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("internalsecrets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InternalSecrets that match those selectors.
func (c *internalSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.InternalSecretList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.InternalSecretList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("internalsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested internalSecrets.
func (c *internalSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("internalsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a internalSecret and creates it.  Returns the server's representation of the internalSecret, and an error, if there is any.
func (c *internalSecrets) Create(ctx context.Context, internalSecret *v1beta1.InternalSecret, opts v1.CreateOptions) (result *v1beta1.InternalSecret, err error) {
	result = &v1beta1.InternalSecret{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("internalsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(internalSecret).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a internalSecret and updates it. Returns the server's representation of the internalSecret, and an error, if there is any.
func (c *internalSecrets) Update(ctx context.Context, internalSecret *v1beta1.InternalSecret, opts v1.UpdateOptions) (result *v1beta1.InternalSecret, err error) {
	result = &v1beta1.InternalSecret{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("internalsecrets").
		Name(internalSecret.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(internalSecret).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the internalSecret and deletes it. Returns an error if one occurs.
func (c *internalSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("internalsecrets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *internalSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("internalsecrets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched internalSecret.
func (c *internalSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.InternalSecret, err error) {
	result = &v1beta1.InternalSecret{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("internalsecrets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
