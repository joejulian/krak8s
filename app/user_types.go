// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": Application User Types
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// applicationPostBody user type.
type applicationPostBody struct {
	// Application chart's channel
	Channel *string `form:"channel,omitempty" json:"channel,omitempty" xml:"channel,omitempty"`
	// Cluster application deployment name
	DeploymentName *string `form:"deployment_name,omitempty" json:"deployment_name,omitempty" xml:"deployment_name,omitempty"`
	// Application chart's json values string
	JSONValues *string `form:"json_values,omitempty" json:"json_values,omitempty" xml:"json_values,omitempty"`
	// Application chart name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The related namespace's generated unique id, not the namespace's name
	NamespaceID *string `form:"namespace_id,omitempty" json:"namespace_id,omitempty" xml:"namespace_id,omitempty"`
	// Application chart's registry
	Registry *string `form:"registry,omitempty" json:"registry,omitempty" xml:"registry,omitempty"`
	// Application chart registry host server
	Server *string `form:"server,omitempty" json:"server,omitempty" xml:"server,omitempty"`
	// Application chart config --set argument string
	Set *string `form:"set,omitempty" json:"set,omitempty" xml:"set,omitempty"`
	// Application chart version string
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// Finalize sets the default values for applicationPostBody type instance.
func (ut *applicationPostBody) Finalize() {
	var defaultDeploymentName = "samsung-mongodb-replicaset"
	if ut.DeploymentName == nil {
		ut.DeploymentName = &defaultDeploymentName
	}
	var defaultRegistry = "application/samsung_cnct"
	if ut.Registry == nil {
		ut.Registry = &defaultRegistry
	}
	var defaultServer = "quay.io"
	if ut.Server == nil {
		ut.Server = &defaultServer
	}
	var defaultVersion = "latest"
	if ut.Version == nil {
		ut.Version = &defaultVersion
	}
}

// Validate validates the applicationPostBody type instance.
func (ut *applicationPostBody) Validate() (err error) {
	if ut.DeploymentName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "deployment_name"))
	}
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Version == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	if ut.NamespaceID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "namespace_id"))
	}
	return
}

// Publicize creates ApplicationPostBody from applicationPostBody
func (ut *applicationPostBody) Publicize() *ApplicationPostBody {
	var pub ApplicationPostBody
	if ut.Channel != nil {
		pub.Channel = ut.Channel
	}
	if ut.DeploymentName != nil {
		pub.DeploymentName = *ut.DeploymentName
	}
	if ut.JSONValues != nil {
		pub.JSONValues = ut.JSONValues
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.NamespaceID != nil {
		pub.NamespaceID = *ut.NamespaceID
	}
	if ut.Registry != nil {
		pub.Registry = *ut.Registry
	}
	if ut.Server != nil {
		pub.Server = *ut.Server
	}
	if ut.Set != nil {
		pub.Set = ut.Set
	}
	if ut.Version != nil {
		pub.Version = *ut.Version
	}
	return &pub
}

// ApplicationPostBody user type.
type ApplicationPostBody struct {
	// Application chart's channel
	Channel *string `form:"channel,omitempty" json:"channel,omitempty" xml:"channel,omitempty"`
	// Cluster application deployment name
	DeploymentName string `form:"deployment_name" json:"deployment_name" xml:"deployment_name"`
	// Application chart's json values string
	JSONValues *string `form:"json_values,omitempty" json:"json_values,omitempty" xml:"json_values,omitempty"`
	// Application chart name
	Name string `form:"name" json:"name" xml:"name"`
	// The related namespace's generated unique id, not the namespace's name
	NamespaceID string `form:"namespace_id" json:"namespace_id" xml:"namespace_id"`
	// Application chart's registry
	Registry string `form:"registry" json:"registry" xml:"registry"`
	// Application chart registry host server
	Server string `form:"server" json:"server" xml:"server"`
	// Application chart config --set argument string
	Set *string `form:"set,omitempty" json:"set,omitempty" xml:"set,omitempty"`
	// Application chart version string
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the ApplicationPostBody type instance.
func (ut *ApplicationPostBody) Validate() (err error) {
	if ut.DeploymentName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "deployment_name"))
	}
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	if ut.NamespaceID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "namespace_id"))
	}
	return
}

// clusterPostBody user type.
type clusterPostBody struct {
	// The related namespace's generated unique id, not the namespace's name
	NamespaceID *string `form:"namespace_id,omitempty" json:"namespace_id,omitempty" xml:"namespace_id,omitempty"`
	// The number of worker nodes in the projects resource pool
	NodePoolSize *int `form:"nodePoolSize,omitempty" json:"nodePoolSize,omitempty" xml:"nodePoolSize,omitempty"`
}

// Finalize sets the default values for clusterPostBody type instance.
func (ut *clusterPostBody) Finalize() {
	var defaultNodePoolSize = 3
	if ut.NodePoolSize == nil {
		ut.NodePoolSize = &defaultNodePoolSize
	}
}

// Validate validates the clusterPostBody type instance.
func (ut *clusterPostBody) Validate() (err error) {
	if ut.NodePoolSize == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "nodePoolSize"))
	}
	if ut.NamespaceID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "namespace_id"))
	}
	if ut.NodePoolSize != nil {
		if *ut.NodePoolSize < 3 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, *ut.NodePoolSize, 3, true))
		}
	}
	if ut.NodePoolSize != nil {
		if *ut.NodePoolSize > 11 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, *ut.NodePoolSize, 11, false))
		}
	}
	return
}

// Publicize creates ClusterPostBody from clusterPostBody
func (ut *clusterPostBody) Publicize() *ClusterPostBody {
	var pub ClusterPostBody
	if ut.NamespaceID != nil {
		pub.NamespaceID = *ut.NamespaceID
	}
	if ut.NodePoolSize != nil {
		pub.NodePoolSize = *ut.NodePoolSize
	}
	return &pub
}

// ClusterPostBody user type.
type ClusterPostBody struct {
	// The related namespace's generated unique id, not the namespace's name
	NamespaceID string `form:"namespace_id" json:"namespace_id" xml:"namespace_id"`
	// The number of worker nodes in the projects resource pool
	NodePoolSize int `form:"nodePoolSize" json:"nodePoolSize" xml:"nodePoolSize"`
}

// Validate validates the ClusterPostBody type instance.
func (ut *ClusterPostBody) Validate() (err error) {

	if ut.NamespaceID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "namespace_id"))
	}
	if ut.NodePoolSize < 3 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, ut.NodePoolSize, 3, true))
	}
	if ut.NodePoolSize > 11 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, ut.NodePoolSize, 11, false))
	}
	return
}
