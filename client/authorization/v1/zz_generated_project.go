package client

import (
	"github.com/rancher/norman/types"
)

const (
	ProjectType             = "project"
	ProjectFieldAPIVersion  = "apiVersion"
	ProjectFieldClusterName = "clusterName"
	ProjectFieldKind        = "kind"
	ProjectFieldObjectMeta  = "objectMeta"
)

type Project struct {
	types.Resource
	APIVersion  string     `json:"apiVersion,omitempty"`
	ClusterName string     `json:"clusterName,omitempty"`
	Kind        string     `json:"kind,omitempty"`
	ObjectMeta  ObjectMeta `json:"objectMeta,omitempty"`
}
type ProjectCollection struct {
	types.Collection
	Data   []Project `json:"data,omitempty"`
	client *ProjectClient
}

type ProjectClient struct {
	apiClient *Client
}

type ProjectOperations interface {
	List(opts *types.ListOpts) (*ProjectCollection, error)
	Create(opts *Project) (*Project, error)
	Update(existing *Project, updates interface{}) (*Project, error)
	ByID(id string) (*Project, error)
	Delete(container *Project) error
}

func newProjectClient(apiClient *Client) *ProjectClient {
	return &ProjectClient{
		apiClient: apiClient,
	}
}

func (c *ProjectClient) Create(container *Project) (*Project, error) {
	resp := &Project{}
	err := c.apiClient.Ops.DoCreate(ProjectType, container, resp)
	return resp, err
}

func (c *ProjectClient) Update(existing *Project, updates interface{}) (*Project, error) {
	resp := &Project{}
	err := c.apiClient.Ops.DoUpdate(ProjectType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProjectClient) List(opts *types.ListOpts) (*ProjectCollection, error) {
	resp := &ProjectCollection{}
	err := c.apiClient.Ops.DoList(ProjectType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProjectCollection) Next() (*ProjectCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProjectCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProjectClient) ByID(id string) (*Project, error) {
	resp := &Project{}
	err := c.apiClient.Ops.DoByID(ProjectType, id, resp)
	return resp, err
}

func (c *ProjectClient) Delete(container *Project) error {
	return c.apiClient.Ops.DoResourceDelete(ProjectType, &container.Resource)
}
