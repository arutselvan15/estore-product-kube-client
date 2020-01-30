/*
Copyright (c) 2017 SAP SE or an SAP affiliate company. All rights reserved.

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

package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	ProductLabelClusterName = "product.k8s.io/cluster-name"
	ProductLabelNamespace   = "product.k8s.io/namespace"
	ProductLabelBrand       = "product.k8s.io/brand"
	ProductLabelCategory    = "product.k8s.io/category"

	ProductAnnotationOutOfStock  = "product.k8s.io/out-of-stock"
	ProductAnnotationApplyOffer  = "product.k8s.io/apply-offer"
	ProductAnnotationForceDelete = "product.k8s.io/force-delete"
)

type ProductConditionType string

const (
	ConditionTypeWarehouse    ProductConditionType = "WarehouseCheck"
	ConditionTypeInventory    ProductConditionType = "InventoryCheck"
	ConditionTypeOffer        ProductConditionType = "ApplyOffers"
	ConditionTypeOutOfService ProductConditionType = "StockCheck"
)

// ProductPhase is a label for the condition of a product at the current time.
type ProductPhase string

// These are the valid statuses of products.
const (
	// ProductPending means that the product is waiting for inventory
	ProductPending ProductPhase = "Pending"
	// ProductAvailable means that product is present on provider but hasn't joined cluster yet
	ProductAvailable ProductPhase = "Available"
	// ProductOutOfStock means that the product is out of stock
	ProductOutOfStock ProductPhase = "OutOfStock"
	// ProductDeleting means the deletion timestamp is set
	ProductDeleting ProductPhase = "Deleting"
	// ProductUnknown indicates that the product is not ready at the moment
	ProductUnknown ProductPhase = "Unknown"
	// ProductFailed means operation failed leading to product status failure
	ProductFailed ProductPhase = "Failed"
)

// ProductState is a high level state on the progress of a product
type ProductState string

// These are the valid statuses of products.
const (
	// ProductStatePending means there are operations pending on this product state
	ProductStateProcessing ProductState = "Processing"
	// ProductStateFailed means operation failed leading to product status failure
	ProductStateFailed ProductState = "Failed"
	// ProductStateSuccessful indicates that the node operation succeeded
	ProductStateSuccessful ProductState = "Successful"
)

// ProductOperationType is a label for the operation performed on a product object.
type ProductOperationType string

// These are the valid statuses of products.
const (
	// ProductOperationCreate indicates that the operation was a create
	ProductOperationCreate ProductOperationType = "Create"
	// ProductOperationUpdate indicates that the operation was an update
	ProductOperationUpdate ProductOperationType = "Update"
	// ProductOperationHealthCheck indicates that the operation was a health check
	ProductOperationHealthCheck ProductOperationType = "HealthCheck"
	// ProductOperationDelete indicates that the operation was a delete
	ProductOperationDelete ProductOperationType = "Delete"
)

// The below types are used by kube_client and api_server.

type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition;
// "ConditionFalse" means a resource is not in the condition; "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Product is a top level type
type Product struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProductSpec `json:"spec"`
	// +optional
	Status ProductStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProductList product lister
type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// +optional
	Status ProductStatus `json:"status,omitempty"`
	Items  []Product     `json:"items"`
}

// ProductSpec custom spec
type ProductSpec struct {
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	// +optional
	Brand string `json:"brand,omitempty"`
	// +optional
	Price float64 `json:"price,omitempty"`
	// +optional
	Categories []string `json:"categories,omitempty"`
}

// ProductStatus product status
type ProductStatus struct {
	// Conditions of this product
	Conditions []ProductCondition `json:"conditions,omitempty"`

	// LastOperation details the last operation performed on the product
	LastOperation LastOperation `json:"lastOperation,omitempty"`

	// CurrentStatus is the current status of the product
	CurrentStatus CurrentStatus `json:"currentStatus,omitempty"`
}

// ProductCondition product condition
type ProductCondition struct {
	Type   ProductConditionType `json:"type"`
	Status ConditionStatus      `json:"status"`
	// +optional
	Reason string `json:"reason,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
}

// LastOperation suggests the last operation performed on the object
type LastOperation struct {
	// Description of the current operation
	Description string `json:"description,omitempty"`

	// Last update time of current operation
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// State of operation
	State ProductState `json:"state,omitempty"`

	// Type of operation
	Type ProductOperationType `json:"type,omitempty"`
}

// CurrentStatus is the current status of the product
type CurrentStatus struct {
	// Phase is the current phase of the product
	Phase ProductPhase `json:"phase,omitempty"`

	// TimeoutActive states whether the timeout has been triggered or not
	TimeoutActive bool `json:"timeoutActive,omitempty"`

	// LastUpdateTime is the last update time of current status
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
}
