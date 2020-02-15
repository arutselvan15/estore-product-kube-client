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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/arutselvan15/estore-product-kube-client/pkg/apis/estore/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProductLister helps list Products.
type ProductLister interface {
	// List lists all Products in the indexer.
	List(selector labels.Selector) (ret []*v1.Product, err error)
	// Products returns an object that can list and get Products.
	Products(namespace string) ProductNamespaceLister
	ProductListerExpansion
}

// productLister implements the ProductLister interface.
type productLister struct {
	indexer cache.Indexer
}

// NewProductLister returns a new ProductLister.
func NewProductLister(indexer cache.Indexer) ProductLister {
	return &productLister{indexer: indexer}
}

// List lists all Products in the indexer.
func (s *productLister) List(selector labels.Selector) (ret []*v1.Product, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Product))
	})
	return ret, err
}

// Products returns an object that can list and get Products.
func (s *productLister) Products(namespace string) ProductNamespaceLister {
	return productNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProductNamespaceLister helps list and get Products.
type ProductNamespaceLister interface {
	// List lists all Products in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Product, err error)
	// Get retrieves the Product from the indexer for a given namespace and name.
	Get(name string) (*v1.Product, error)
	ProductNamespaceListerExpansion
}

// productNamespaceLister implements the ProductNamespaceLister
// interface.
type productNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Products in the indexer for a given namespace.
func (s productNamespaceLister) List(selector labels.Selector) (ret []*v1.Product, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Product))
	})
	return ret, err
}

// Get retrieves the Product from the indexer for a given namespace and name.
func (s productNamespaceLister) Get(name string) (*v1.Product, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("product"), name)
	}
	return obj.(*v1.Product), nil
}
