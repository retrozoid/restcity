package trac

import (
	"fmt"
	"net/http"
)

type UserAPI struct {
	HTTPClient *http.Client
}

// GetUserProperties Get all properties of the matching user.
func (t UserAPI) GetUserProperties(userLocator UserLocator, fields string) (value Properties, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s/properties", EncodeLocator(userLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetUserGroup Get a user group of the matching user.
func (t UserAPI) GetUserGroup(userLocator UserLocator, groupLocator UserGroupLocator, fields string) (value Group, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s/groups/%s", EncodeLocator(userLocator), EncodeLocator(groupLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// RemoveUserFromGroup Remove the matching user from the specific group.
func (t UserAPI) RemoveUserFromGroup(userLocator UserLocator, groupLocator UserGroupLocator, fields string) error {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s/groups/%s", EncodeLocator(userLocator), EncodeLocator(groupLocator)),
		Values: Values{"fields": fields},
	}
	return req.Delete(t.HTTPClient)
}

// GetAllUsers Get all users.
func (t UserAPI) GetAllUsers(userLocator UserLocator, fields string) (value Users, err error) {
	req := Request{
		Path:   "/app/rest/users",
		Values: Values{"fields": fields, "locator": EncodeLocator(userLocator)},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddUser Create a new user.
func (t UserAPI) AddUser(data User, fields string) (value User, err error) {
	req := Request{
		Path:   "/app/rest/users",
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// GetUser Get user matching the locator.
func (t UserAPI) GetUser(userLocator UserLocator, fields string) (value User, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s", EncodeLocator(userLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// ReplaceUser Update user matching the locator.
func (t UserAPI) ReplaceUser(userLocator UserLocator, data User, fields string) (value User, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s", EncodeLocator(userLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteUser Delete user matching the locator.
func (t UserAPI) DeleteUser(userLocator UserLocator) error {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s", EncodeLocator(userLocator)),
	}
	return req.Delete(t.HTTPClient)
}

// GetUserPermissions Get all permissions effective for the matching user.
func (t UserAPI) GetUserPermissions(userLocator UserLocator, locator string, fields string) (value PermissionAssignments, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s/permissions", EncodeLocator(userLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// GetUserRolesAtScope Get a user role with the specific scope from the matching user.
func (t UserAPI) GetUserRolesAtScope(userLocator UserLocator, roleId string, scope string) (value Role, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/roles/%s/%s", EncodeLocator(userLocator), roleId, scope),
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddRoleToUserAtScope Add a role with the specific scope to the matching user.
func (t UserAPI) AddRoleToUserAtScope(userLocator UserLocator, roleId string, scope string) (value Role, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/roles/%s/%s", EncodeLocator(userLocator), roleId, scope),
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// RemoveUserRoleAtScope Remove a role with the specific scope from the matching user.
func (t UserAPI) RemoveUserRoleAtScope(userLocator UserLocator, roleId string, scope string) error {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/roles/%s/%s", EncodeLocator(userLocator), roleId, scope),
	}
	return req.Delete(t.HTTPClient)
}

// GetUserField Get a field of the matching user.
func (t UserAPI) GetUserField(userLocator UserLocator, field string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/users/%s/%s", EncodeLocator(userLocator), field),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetUserField Update a field of the matching user.
func (t UserAPI) SetUserField(userLocator UserLocator, field string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/users/%s/%s", EncodeLocator(userLocator), field),
		Consumer: Coders.String,
		Data:     data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// DeleteUserField Remove a property of the matching user.
func (t UserAPI) DeleteUserField(userLocator UserLocator, field string) error {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/%s", EncodeLocator(userLocator), field),
	}
	return req.Delete(t.HTTPClient)
}

// DeleteUserToken Remove an authentication token from the matching user.
func (t UserAPI) DeleteUserToken(userLocator UserLocator, name string) error {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/tokens/%s", EncodeLocator(userLocator), name),
	}
	return req.Delete(t.HTTPClient)
}

// GetAllUserRoles Get all user roles of the matching user.
func (t UserAPI) GetAllUserRoles(userLocator UserLocator) (value Roles, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/roles", EncodeLocator(userLocator)),
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddRoleToUser Add a role to the matching user.
func (t UserAPI) AddRoleToUser(userLocator UserLocator, data Role) (value Role, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/roles", EncodeLocator(userLocator)),
		Data: data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}

// SetUserRoles Update user roles of the matching user.
func (t UserAPI) SetUserRoles(userLocator UserLocator, data Roles) (value Roles, err error) {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/roles", EncodeLocator(userLocator)),
		Data: data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// GetAllUserGroups Get all groups of the matching user.
func (t UserAPI) GetAllUserGroups(userLocator UserLocator, fields string) (value Groups, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s/groups", EncodeLocator(userLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetUserGroups Update groups of the matching user.
func (t UserAPI) SetUserGroups(userLocator UserLocator, data Groups, fields string) (value Groups, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s/groups", EncodeLocator(userLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// RemoveUserRememberMe Remove the RememberMe data of the matching user.
func (t UserAPI) RemoveUserRememberMe(userLocator UserLocator) error {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/debug/rememberMe", EncodeLocator(userLocator)),
	}
	return req.Delete(t.HTTPClient)
}

// GetUserProperty Get a property of the matching user.
func (t UserAPI) GetUserProperty(userLocator UserLocator, name string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/users/%s/properties/%s", EncodeLocator(userLocator), name),
		Consumer: Coders.String,
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// SetUserProperty Update a property of the matching user.
func (t UserAPI) SetUserProperty(userLocator UserLocator, name string, data string) (value string, err error) {
	req := Request{
		Path:     fmt.Sprintf("/app/rest/users/%s/properties/%s", EncodeLocator(userLocator), name),
		Data:     data,
		Consumer: Coders.String,
	}
	err = req.Put(t.HTTPClient, &value)
	return
}

// RemoveUserProperty Remove a property of the matching user.
func (t UserAPI) RemoveUserProperty(userLocator UserLocator, name string) error {
	req := Request{
		Path: fmt.Sprintf("/app/rest/users/%s/properties/%s", EncodeLocator(userLocator), name),
	}
	return req.Delete(t.HTTPClient)
}

// GetUserTokens Get all authentication tokens of the matching user.
func (t UserAPI) GetUserTokens(userLocator UserLocator, fields string) (value Tokens, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s/tokens", EncodeLocator(userLocator)),
		Values: Values{"fields": fields},
	}
	err = req.Get(t.HTTPClient, &value)
	return
}

// AddUserToken Create a new authentication token for the matching user.
func (t UserAPI) AddUserToken(userLocator string, data Token, fields string) (value Token, err error) {
	req := Request{
		Path:   fmt.Sprintf("/app/rest/users/%s/tokens", EncodeLocator(userLocator)),
		Values: Values{"fields": fields},
		Data:   data,
	}
	err = req.Post(t.HTTPClient, &value)
	return
}
