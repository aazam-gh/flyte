// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFlyteWorkflows implements FlyteWorkflowInterface
type FakeFlyteWorkflows struct {
	Fake *FakeFlyteworkflowV1alpha1
	ns   string
}

var flyteworkflowsResource = schema.GroupVersionResource{Group: "flyteworkflow.flyte.net", Version: "v1alpha1", Resource: "flyteworkflows"}

var flyteworkflowsKind = schema.GroupVersionKind{Group: "flyteworkflow.flyte.net", Version: "v1alpha1", Kind: "FlyteWorkflow"}

// Get takes name of the flyteWorkflow, and returns the corresponding flyteWorkflow object, and an error if there is any.
func (c *FakeFlyteWorkflows) Get(name string, options v1.GetOptions) (result *v1alpha1.FlyteWorkflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(flyteworkflowsResource, c.ns, name), &v1alpha1.FlyteWorkflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlyteWorkflow), err
}

// List takes label and field selectors, and returns the list of FlyteWorkflows that match those selectors.
func (c *FakeFlyteWorkflows) List(opts v1.ListOptions) (result *v1alpha1.FlyteWorkflowList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(flyteworkflowsResource, flyteworkflowsKind, c.ns, opts), &v1alpha1.FlyteWorkflowList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FlyteWorkflowList{ListMeta: obj.(*v1alpha1.FlyteWorkflowList).ListMeta}
	for _, item := range obj.(*v1alpha1.FlyteWorkflowList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested flyteWorkflows.
func (c *FakeFlyteWorkflows) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(flyteworkflowsResource, c.ns, opts))

}

// Create takes the representation of a flyteWorkflow and creates it.  Returns the server's representation of the flyteWorkflow, and an error, if there is any.
func (c *FakeFlyteWorkflows) Create(flyteWorkflow *v1alpha1.FlyteWorkflow) (result *v1alpha1.FlyteWorkflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(flyteworkflowsResource, c.ns, flyteWorkflow), &v1alpha1.FlyteWorkflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlyteWorkflow), err
}

// Update takes the representation of a flyteWorkflow and updates it. Returns the server's representation of the flyteWorkflow, and an error, if there is any.
func (c *FakeFlyteWorkflows) Update(flyteWorkflow *v1alpha1.FlyteWorkflow) (result *v1alpha1.FlyteWorkflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(flyteworkflowsResource, c.ns, flyteWorkflow), &v1alpha1.FlyteWorkflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlyteWorkflow), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFlyteWorkflows) UpdateStatus(flyteWorkflow *v1alpha1.FlyteWorkflow) (*v1alpha1.FlyteWorkflow, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(flyteworkflowsResource, "status", c.ns, flyteWorkflow), &v1alpha1.FlyteWorkflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlyteWorkflow), err
}

// Delete takes name of the flyteWorkflow and deletes it. Returns an error if one occurs.
func (c *FakeFlyteWorkflows) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(flyteworkflowsResource, c.ns, name), &v1alpha1.FlyteWorkflow{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFlyteWorkflows) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(flyteworkflowsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FlyteWorkflowList{})
	return err
}

// Patch applies the patch and returns the patched flyteWorkflow.
func (c *FakeFlyteWorkflows) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FlyteWorkflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(flyteworkflowsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FlyteWorkflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlyteWorkflow), err
}
