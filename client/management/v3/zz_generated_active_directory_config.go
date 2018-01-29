package client

import "github.com/rancher/norman/types"

const (
	ActiveDirectoryConfigType                             = "activeDirectoryConfig"
	ActiveDirectoryConfigFieldAnnotations                 = "annotations"
	ActiveDirectoryConfigFieldConnectionTimeout           = "connectionTimeout"
	ActiveDirectoryConfigFieldCreated                     = "created"
	ActiveDirectoryConfigFieldCreatorID                   = "creatorId"
	ActiveDirectoryConfigFieldDomain                      = "domain"
	ActiveDirectoryConfigFieldEnabled                     = "enabled"
	ActiveDirectoryConfigFieldGroupDNField                = "groupDNField"
	ActiveDirectoryConfigFieldGroupMemberMappingAttribute = "groupMemberMappingAttribute"
	ActiveDirectoryConfigFieldGroupMemberUserAttribute    = "groupMemberUserAttribute"
	ActiveDirectoryConfigFieldGroupNameField              = "groupNameField"
	ActiveDirectoryConfigFieldGroupObjectClass            = "groupObjectClass"
	ActiveDirectoryConfigFieldGroupSearchDomain           = "groupSearchDomain"
	ActiveDirectoryConfigFieldGroupSearchField            = "groupSearchField"
	ActiveDirectoryConfigFieldLabels                      = "labels"
	ActiveDirectoryConfigFieldLoginDomain                 = "loginDomain"
	ActiveDirectoryConfigFieldName                        = "name"
	ActiveDirectoryConfigFieldOwnerReferences             = "ownerReferences"
	ActiveDirectoryConfigFieldPort                        = "port"
	ActiveDirectoryConfigFieldRemoved                     = "removed"
	ActiveDirectoryConfigFieldServer                      = "server"
	ActiveDirectoryConfigFieldServiceAccountPassword      = "serviceAccountPassword"
	ActiveDirectoryConfigFieldServiceAccountUsername      = "serviceAccountUsername"
	ActiveDirectoryConfigFieldTLS                         = "tls"
	ActiveDirectoryConfigFieldUserDisabledBitMask         = "userDisabledBitMask"
	ActiveDirectoryConfigFieldUserEnabledAttribute        = "userEnabledAttribute"
	ActiveDirectoryConfigFieldUserLoginField              = "userLoginField"
	ActiveDirectoryConfigFieldUserNameField               = "userNameField"
	ActiveDirectoryConfigFieldUserObjectClass             = "userObjectClass"
	ActiveDirectoryConfigFieldUserSearchField             = "userSearchField"
	ActiveDirectoryConfigFieldUuid                        = "uuid"
)

type ActiveDirectoryConfig struct {
	types.Resource
	Annotations                 map[string]string `json:"annotations,omitempty"`
	ConnectionTimeout           *int64            `json:"connectionTimeout,omitempty"`
	Created                     string            `json:"created,omitempty"`
	CreatorID                   string            `json:"creatorId,omitempty"`
	Domain                      string            `json:"domain,omitempty"`
	Enabled                     *bool             `json:"enabled,omitempty"`
	GroupDNField                string            `json:"groupDNField,omitempty"`
	GroupMemberMappingAttribute string            `json:"groupMemberMappingAttribute,omitempty"`
	GroupMemberUserAttribute    string            `json:"groupMemberUserAttribute,omitempty"`
	GroupNameField              string            `json:"groupNameField,omitempty"`
	GroupObjectClass            string            `json:"groupObjectClass,omitempty"`
	GroupSearchDomain           string            `json:"groupSearchDomain,omitempty"`
	GroupSearchField            string            `json:"groupSearchField,omitempty"`
	Labels                      map[string]string `json:"labels,omitempty"`
	LoginDomain                 string            `json:"loginDomain,omitempty"`
	Name                        string            `json:"name,omitempty"`
	OwnerReferences             []OwnerReference  `json:"ownerReferences,omitempty"`
	Port                        *int64            `json:"port,omitempty"`
	Removed                     string            `json:"removed,omitempty"`
	Server                      string            `json:"server,omitempty"`
	ServiceAccountPassword      string            `json:"serviceAccountPassword,omitempty"`
	ServiceAccountUsername      string            `json:"serviceAccountUsername,omitempty"`
	TLS                         *bool             `json:"tls,omitempty"`
	UserDisabledBitMask         *int64            `json:"userDisabledBitMask,omitempty"`
	UserEnabledAttribute        string            `json:"userEnabledAttribute,omitempty"`
	UserLoginField              string            `json:"userLoginField,omitempty"`
	UserNameField               string            `json:"userNameField,omitempty"`
	UserObjectClass             string            `json:"userObjectClass,omitempty"`
	UserSearchField             string            `json:"userSearchField,omitempty"`
	Uuid                        string            `json:"uuid,omitempty"`
}
type ActiveDirectoryConfigCollection struct {
	types.Collection
	Data   []ActiveDirectoryConfig `json:"data,omitempty"`
	client *ActiveDirectoryConfigClient
}

type ActiveDirectoryConfigClient struct {
	apiClient *Client
}

type ActiveDirectoryConfigOperations interface {
	List(opts *types.ListOpts) (*ActiveDirectoryConfigCollection, error)
	Create(opts *ActiveDirectoryConfig) (*ActiveDirectoryConfig, error)
	Update(existing *ActiveDirectoryConfig, updates interface{}) (*ActiveDirectoryConfig, error)
	ByID(id string) (*ActiveDirectoryConfig, error)
	Delete(container *ActiveDirectoryConfig) error

	ActionTestAndApply(*ActiveDirectoryConfig, *ActiveDirectoryConfigApplyInput) (*ActiveDirectoryConfig, error)
}

func newActiveDirectoryConfigClient(apiClient *Client) *ActiveDirectoryConfigClient {
	return &ActiveDirectoryConfigClient{
		apiClient: apiClient,
	}
}

func (c *ActiveDirectoryConfigClient) Create(container *ActiveDirectoryConfig) (*ActiveDirectoryConfig, error) {
	resp := &ActiveDirectoryConfig{}
	err := c.apiClient.Ops.DoCreate(ActiveDirectoryConfigType, container, resp)
	return resp, err
}

func (c *ActiveDirectoryConfigClient) Update(existing *ActiveDirectoryConfig, updates interface{}) (*ActiveDirectoryConfig, error) {
	resp := &ActiveDirectoryConfig{}
	err := c.apiClient.Ops.DoUpdate(ActiveDirectoryConfigType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ActiveDirectoryConfigClient) List(opts *types.ListOpts) (*ActiveDirectoryConfigCollection, error) {
	resp := &ActiveDirectoryConfigCollection{}
	err := c.apiClient.Ops.DoList(ActiveDirectoryConfigType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ActiveDirectoryConfigCollection) Next() (*ActiveDirectoryConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ActiveDirectoryConfigCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ActiveDirectoryConfigClient) ByID(id string) (*ActiveDirectoryConfig, error) {
	resp := &ActiveDirectoryConfig{}
	err := c.apiClient.Ops.DoByID(ActiveDirectoryConfigType, id, resp)
	return resp, err
}

func (c *ActiveDirectoryConfigClient) Delete(container *ActiveDirectoryConfig) error {
	return c.apiClient.Ops.DoResourceDelete(ActiveDirectoryConfigType, &container.Resource)
}

func (c *ActiveDirectoryConfigClient) ActionTestAndApply(resource *ActiveDirectoryConfig, input *ActiveDirectoryConfigApplyInput) (*ActiveDirectoryConfig, error) {

	resp := &ActiveDirectoryConfig{}

	err := c.apiClient.Ops.DoAction(ActiveDirectoryConfigType, "testAndApply", &resource.Resource, input, resp)

	return resp, err
}
