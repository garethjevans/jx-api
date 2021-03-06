// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/jx-api/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CommitStatusesGetter has a method to return a CommitStatusInterface.
// A group's client should implement this interface.
type CommitStatusesGetter interface {
	CommitStatuses(namespace string) CommitStatusInterface
}

// CommitStatusInterface has methods to work with CommitStatus resources.
type CommitStatusInterface interface {
	Create(*v1.CommitStatus) (*v1.CommitStatus, error)
	Update(*v1.CommitStatus) (*v1.CommitStatus, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.CommitStatus, error)
	List(opts metav1.ListOptions) (*v1.CommitStatusList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CommitStatus, err error)
	CommitStatusExpansion
}

// commitStatuses implements CommitStatusInterface
type commitStatuses struct {
	client rest.Interface
	ns     string
}

// newCommitStatuses returns a CommitStatuses
func newCommitStatuses(c *JenkinsV1Client, namespace string) *commitStatuses {
	return &commitStatuses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the commitStatus, and returns the corresponding commitStatus object, and an error if there is any.
func (c *commitStatuses) Get(name string, options metav1.GetOptions) (result *v1.CommitStatus, err error) {
	result = &v1.CommitStatus{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("commitstatuses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CommitStatuses that match those selectors.
func (c *commitStatuses) List(opts metav1.ListOptions) (result *v1.CommitStatusList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CommitStatusList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("commitstatuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested commitStatuses.
func (c *commitStatuses) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("commitstatuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a commitStatus and creates it.  Returns the server's representation of the commitStatus, and an error, if there is any.
func (c *commitStatuses) Create(commitStatus *v1.CommitStatus) (result *v1.CommitStatus, err error) {
	result = &v1.CommitStatus{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("commitstatuses").
		Body(commitStatus).
		Do().
		Into(result)
	return
}

// Update takes the representation of a commitStatus and updates it. Returns the server's representation of the commitStatus, and an error, if there is any.
func (c *commitStatuses) Update(commitStatus *v1.CommitStatus) (result *v1.CommitStatus, err error) {
	result = &v1.CommitStatus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("commitstatuses").
		Name(commitStatus.Name).
		Body(commitStatus).
		Do().
		Into(result)
	return
}

// Delete takes name of the commitStatus and deletes it. Returns an error if one occurs.
func (c *commitStatuses) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("commitstatuses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *commitStatuses) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("commitstatuses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched commitStatus.
func (c *commitStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CommitStatus, err error) {
	result = &v1.CommitStatus{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("commitstatuses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
