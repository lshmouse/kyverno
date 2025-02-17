/*
Copyright The Kubernetes Authors.

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

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "github.com/kyverno/kyverno/pkg/api/kyverno/v1alpha2"
	scheme "github.com/kyverno/kyverno/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ReportChangeRequestsGetter has a method to return a ReportChangeRequestInterface.
// A group's client should implement this interface.
type ReportChangeRequestsGetter interface {
	ReportChangeRequests(namespace string) ReportChangeRequestInterface
}

// ReportChangeRequestInterface has methods to work with ReportChangeRequest resources.
type ReportChangeRequestInterface interface {
	Create(ctx context.Context, reportChangeRequest *v1alpha2.ReportChangeRequest, opts v1.CreateOptions) (*v1alpha2.ReportChangeRequest, error)
	Update(ctx context.Context, reportChangeRequest *v1alpha2.ReportChangeRequest, opts v1.UpdateOptions) (*v1alpha2.ReportChangeRequest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.ReportChangeRequest, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.ReportChangeRequestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ReportChangeRequest, err error)
	ReportChangeRequestExpansion
}

// reportChangeRequests implements ReportChangeRequestInterface
type reportChangeRequests struct {
	client rest.Interface
	ns     string
}

// newReportChangeRequests returns a ReportChangeRequests
func newReportChangeRequests(c *KyvernoV1alpha2Client, namespace string) *reportChangeRequests {
	return &reportChangeRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the reportChangeRequest, and returns the corresponding reportChangeRequest object, and an error if there is any.
func (c *reportChangeRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ReportChangeRequest, err error) {
	result = &v1alpha2.ReportChangeRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reportchangerequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ReportChangeRequests that match those selectors.
func (c *reportChangeRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ReportChangeRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.ReportChangeRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reportchangerequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested reportChangeRequests.
func (c *reportChangeRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("reportchangerequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a reportChangeRequest and creates it.  Returns the server's representation of the reportChangeRequest, and an error, if there is any.
func (c *reportChangeRequests) Create(ctx context.Context, reportChangeRequest *v1alpha2.ReportChangeRequest, opts v1.CreateOptions) (result *v1alpha2.ReportChangeRequest, err error) {
	result = &v1alpha2.ReportChangeRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("reportchangerequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(reportChangeRequest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a reportChangeRequest and updates it. Returns the server's representation of the reportChangeRequest, and an error, if there is any.
func (c *reportChangeRequests) Update(ctx context.Context, reportChangeRequest *v1alpha2.ReportChangeRequest, opts v1.UpdateOptions) (result *v1alpha2.ReportChangeRequest, err error) {
	result = &v1alpha2.ReportChangeRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("reportchangerequests").
		Name(reportChangeRequest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(reportChangeRequest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the reportChangeRequest and deletes it. Returns an error if one occurs.
func (c *reportChangeRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reportchangerequests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *reportChangeRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reportchangerequests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched reportChangeRequest.
func (c *reportChangeRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ReportChangeRequest, err error) {
	result = &v1alpha2.ReportChangeRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("reportchangerequests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
