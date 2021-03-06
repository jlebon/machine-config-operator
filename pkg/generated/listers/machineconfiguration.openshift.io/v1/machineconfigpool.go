// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineConfigPoolLister helps list MachineConfigPools.
type MachineConfigPoolLister interface {
	// List lists all MachineConfigPools in the indexer.
	List(selector labels.Selector) (ret []*v1.MachineConfigPool, err error)
	// MachineConfigPools returns an object that can list and get MachineConfigPools.
	MachineConfigPools(namespace string) MachineConfigPoolNamespaceLister
	MachineConfigPoolListerExpansion
}

// machineConfigPoolLister implements the MachineConfigPoolLister interface.
type machineConfigPoolLister struct {
	indexer cache.Indexer
}

// NewMachineConfigPoolLister returns a new MachineConfigPoolLister.
func NewMachineConfigPoolLister(indexer cache.Indexer) MachineConfigPoolLister {
	return &machineConfigPoolLister{indexer: indexer}
}

// List lists all MachineConfigPools in the indexer.
func (s *machineConfigPoolLister) List(selector labels.Selector) (ret []*v1.MachineConfigPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MachineConfigPool))
	})
	return ret, err
}

// MachineConfigPools returns an object that can list and get MachineConfigPools.
func (s *machineConfigPoolLister) MachineConfigPools(namespace string) MachineConfigPoolNamespaceLister {
	return machineConfigPoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineConfigPoolNamespaceLister helps list and get MachineConfigPools.
type MachineConfigPoolNamespaceLister interface {
	// List lists all MachineConfigPools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MachineConfigPool, err error)
	// Get retrieves the MachineConfigPool from the indexer for a given namespace and name.
	Get(name string) (*v1.MachineConfigPool, error)
	MachineConfigPoolNamespaceListerExpansion
}

// machineConfigPoolNamespaceLister implements the MachineConfigPoolNamespaceLister
// interface.
type machineConfigPoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineConfigPools in the indexer for a given namespace.
func (s machineConfigPoolNamespaceLister) List(selector labels.Selector) (ret []*v1.MachineConfigPool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MachineConfigPool))
	})
	return ret, err
}

// Get retrieves the MachineConfigPool from the indexer for a given namespace and name.
func (s machineConfigPoolNamespaceLister) Get(name string) (*v1.MachineConfigPool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("machineconfigpool"), name)
	}
	return obj.(*v1.MachineConfigPool), nil
}
