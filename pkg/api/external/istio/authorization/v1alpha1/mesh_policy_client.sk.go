// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type MeshPolicyClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*MeshPolicy, error)
	Write(resource *MeshPolicy, opts clients.WriteOpts) (*MeshPolicy, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (MeshPolicyList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan MeshPolicyList, <-chan error, error)
}

type meshPolicyClient struct {
	rc clients.ResourceClient
}

func NewMeshPolicyClient(rcFactory factory.ResourceClientFactory) (MeshPolicyClient, error) {
	return NewMeshPolicyClientWithToken(rcFactory, "")
}

func NewMeshPolicyClientWithToken(rcFactory factory.ResourceClientFactory, token string) (MeshPolicyClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &MeshPolicy{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base MeshPolicy resource client")
	}
	return NewMeshPolicyClientWithBase(rc), nil
}

func NewMeshPolicyClientWithBase(rc clients.ResourceClient) MeshPolicyClient {
	return &meshPolicyClient{
		rc: rc,
	}
}

func (client *meshPolicyClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *meshPolicyClient) Register() error {
	return client.rc.Register()
}

func (client *meshPolicyClient) Read(namespace, name string, opts clients.ReadOpts) (*MeshPolicy, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*MeshPolicy), nil
}

func (client *meshPolicyClient) Write(meshPolicy *MeshPolicy, opts clients.WriteOpts) (*MeshPolicy, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(meshPolicy, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*MeshPolicy), nil
}

func (client *meshPolicyClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *meshPolicyClient) List(namespace string, opts clients.ListOpts) (MeshPolicyList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToMeshPolicy(resourceList), nil
}

func (client *meshPolicyClient) Watch(namespace string, opts clients.WatchOpts) (<-chan MeshPolicyList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	meshpoliciesChan := make(chan MeshPolicyList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				meshpoliciesChan <- convertToMeshPolicy(resourceList)
			case <-opts.Ctx.Done():
				close(meshpoliciesChan)
				return
			}
		}
	}()
	return meshpoliciesChan, errs, nil
}

func convertToMeshPolicy(resources resources.ResourceList) MeshPolicyList {
	var meshPolicyList MeshPolicyList
	for _, resource := range resources {
		meshPolicyList = append(meshPolicyList, resource.(*MeshPolicy))
	}
	return meshPolicyList
}
