// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": Application User Types
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
)

// chartPostBody user type.
type chartPostBody struct {
	// Chart identifier
	Application *string `form:"application,omitempty" json:"application,omitempty" xml:"application,omitempty"`
	// Chart version string
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// Validate validates the chartPostBody type instance.
func (ut *chartPostBody) Validate() (err error) {
	if ut.Application == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "application"))
	}
	if ut.Version == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}

// Publicize creates ChartPostBody from chartPostBody
func (ut *chartPostBody) Publicize() *ChartPostBody {
	var pub ChartPostBody
	if ut.Application != nil {
		pub.Application = *ut.Application
	}
	if ut.Version != nil {
		pub.Version = *ut.Version
	}
	return &pub
}

// ChartPostBody user type.
type ChartPostBody struct {
	// Chart identifier
	Application string `form:"application" json:"application" xml:"application"`
	// Chart version string
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the ChartPostBody type instance.
func (ut *ChartPostBody) Validate() (err error) {
	if ut.Application == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "application"))
	}
	if ut.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}

// mongoPostBody user type.
type mongoPostBody struct {
	// Appplication Registry Identifier
	Application *string `form:"application,omitempty" json:"application,omitempty" xml:"application,omitempty"`
	// Appplication Version
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// Finalize sets the default values for mongoPostBody type instance.
func (ut *mongoPostBody) Finalize() {
	var defaultApplication = "quay.io/samsung_cnct/mongodb-replicaset"
	if ut.Application == nil {
		ut.Application = &defaultApplication
	}
	var defaultVersion = "v1.2.0"
	if ut.Version == nil {
		ut.Version = &defaultVersion
	}
}

// Validate validates the mongoPostBody type instance.
func (ut *mongoPostBody) Validate() (err error) {
	if ut.Application == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "application"))
	}
	if ut.Version == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}

// Publicize creates MongoPostBody from mongoPostBody
func (ut *mongoPostBody) Publicize() *MongoPostBody {
	var pub MongoPostBody
	if ut.Application != nil {
		pub.Application = *ut.Application
	}
	if ut.Version != nil {
		pub.Version = *ut.Version
	}
	return &pub
}

// MongoPostBody user type.
type MongoPostBody struct {
	// Appplication Registry Identifier
	Application string `form:"application" json:"application" xml:"application"`
	// Appplication Version
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the MongoPostBody type instance.
func (ut *MongoPostBody) Validate() (err error) {
	if ut.Application == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "application"))
	}
	if ut.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}
