// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/sriov-network-operator/pkg/client/clientset/versioned/typed/sriovnetwork/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSriovnetworkV1 struct {
	*testing.Fake
}

func (c *FakeSriovnetworkV1) SriovNetworks(namespace string) v1.SriovNetworkInterface {
	return &FakeSriovNetworks{c, namespace}
}

func (c *FakeSriovnetworkV1) SriovNetworkNodePolicies(namespace string) v1.SriovNetworkNodePolicyInterface {
	return &FakeSriovNetworkNodePolicies{c, namespace}
}

func (c *FakeSriovnetworkV1) SriovNetworkNodeStates(namespace string) v1.SriovNetworkNodeStateInterface {
	return &FakeSriovNetworkNodeStates{c, namespace}
}

func (c *FakeSriovnetworkV1) SriovOperatorConfigs(namespace string) v1.SriovOperatorConfigInterface {
	return &FakeSriovOperatorConfigs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSriovnetworkV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
