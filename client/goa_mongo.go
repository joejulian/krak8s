// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": goa_mongo Resource Client
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateGoaMongoPath computes a request path to the create action of goa_mongo.
func CreateGoaMongoPath() string {

	return fmt.Sprintf("/v1/mongo/")
}

// Create a MongoDB for a user in a namespace
func (c *Client) CreateGoaMongo(ctx context.Context, path string, payload *MongoPostBody) (*http.Response, error) {
	req, err := c.NewCreateGoaMongoRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateGoaMongoRequest create the request corresponding to the create action endpoint of the goa_mongo resource.
func (c *Client) NewCreateGoaMongoRequest(ctx context.Context, path string, payload *MongoPostBody) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// DeleteGoaMongoPath computes a request path to the delete action of goa_mongo.
func DeleteGoaMongoPath(user string, ns string) string {
	param0 := user
	param1 := ns

	return fmt.Sprintf("/v1/mongo/%s/%s", param0, param1)
}

// Delete the user/namespace specified MongoDB Deloyment
func (c *Client) DeleteGoaMongo(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteGoaMongoRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteGoaMongoRequest create the request corresponding to the delete action endpoint of the goa_mongo resource.
func (c *Client) NewDeleteGoaMongoRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ReadGoaMongoPath computes a request path to the read action of goa_mongo.
func ReadGoaMongoPath(user string, ns string) string {
	param0 := user
	param1 := ns

	return fmt.Sprintf("/v1/mongo/%s/%s", param0, param1)
}

// Get the status of the specified user/namespace(ns) MongoDB Deloyment
func (c *Client) ReadGoaMongo(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewReadGoaMongoRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReadGoaMongoRequest create the request corresponding to the read action endpoint of the goa_mongo resource.
func (c *Client) NewReadGoaMongoRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}