# \OrganizationsApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5OrgsOrgMembershipsUsername**](OrganizationsApi.md#DeleteV5OrgsOrgMembershipsUsername) | **Delete** /v5/orgs/{org}/memberships/{username} | 移除授权用户所管理组织中的成员
[**DeleteV5UserMembershipsOrgsOrg**](OrganizationsApi.md#DeleteV5UserMembershipsOrgsOrg) | **Delete** /v5/user/memberships/orgs/{org} | 退出一个组织
[**GetV5OrgsOrg**](OrganizationsApi.md#GetV5OrgsOrg) | **Get** /v5/orgs/{org} | 获取一个组织
[**GetV5OrgsOrgMembers**](OrganizationsApi.md#GetV5OrgsOrgMembers) | **Get** /v5/orgs/{org}/members | 列出一个组织的所有成员
[**GetV5OrgsOrgMembershipsUsername**](OrganizationsApi.md#GetV5OrgsOrgMembershipsUsername) | **Get** /v5/orgs/{org}/memberships/{username} | 获取授权用户所属组织的一个成员
[**GetV5UserMembershipsOrgs**](OrganizationsApi.md#GetV5UserMembershipsOrgs) | **Get** /v5/user/memberships/orgs | 列出授权用户在所属组织的成员资料
[**GetV5UserMembershipsOrgsOrg**](OrganizationsApi.md#GetV5UserMembershipsOrgsOrg) | **Get** /v5/user/memberships/orgs/{org} | 获取授权用户在一个组织的成员资料
[**GetV5UserOrgs**](OrganizationsApi.md#GetV5UserOrgs) | **Get** /v5/user/orgs | 列出授权用户所属的组织
[**GetV5UsersUsernameOrgs**](OrganizationsApi.md#GetV5UsersUsernameOrgs) | **Get** /v5/users/{username}/orgs | 列出用户所属的组织
[**PatchV5OrgsOrg**](OrganizationsApi.md#PatchV5OrgsOrg) | **Patch** /v5/orgs/{org} | 更新授权用户所管理的组织资料
[**PatchV5UserMembershipsOrgsOrg**](OrganizationsApi.md#PatchV5UserMembershipsOrgsOrg) | **Patch** /v5/user/memberships/orgs/{org} | 更新授权用户在一个组织的成员资料
[**PostV5UsersOrganization**](OrganizationsApi.md#PostV5UsersOrganization) | **Post** /v5/users/organization | 创建组织
[**PutV5OrgsOrgMembershipsUsername**](OrganizationsApi.md#PutV5OrgsOrgMembershipsUsername) | **Put** /v5/orgs/{org}/memberships/{username} | 增加或更新授权用户所管理组织的成员


# **DeleteV5OrgsOrgMembershipsUsername**
> DeleteV5OrgsOrgMembershipsUsername(ctx, org, username, optional)
移除授权用户所管理组织中的成员

移除授权用户所管理组织中的成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***DeleteV5OrgsOrgMembershipsUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5OrgsOrgMembershipsUsernameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteV5UserMembershipsOrgsOrg**
> DeleteV5UserMembershipsOrgsOrg(ctx, org, optional)
退出一个组织

退出一个组织

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***DeleteV5UserMembershipsOrgsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5UserMembershipsOrgsOrgOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5OrgsOrg**
> Group GetV5OrgsOrg(ctx, org, optional)
获取一个组织

获取一个组织

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***GetV5OrgsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5OrgsOrgOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5OrgsOrgMembers**
> []UserBasic GetV5OrgsOrgMembers(ctx, org, optional)
列出一个组织的所有成员

列出一个组织的所有成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***GetV5OrgsOrgMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5OrgsOrgMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **role** | **optional.String**| 根据角色筛选成员 | [default to all]

### Return type

[**[]UserBasic**](UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5OrgsOrgMembershipsUsername**
> GroupMember GetV5OrgsOrgMembershipsUsername(ctx, org, username, optional)
获取授权用户所属组织的一个成员

获取授权用户所属组织的一个成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5OrgsOrgMembershipsUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5OrgsOrgMembershipsUsernameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**GroupMember**](GroupMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserMembershipsOrgs**
> []GroupMember GetV5UserMembershipsOrgs(ctx, optional)
列出授权用户在所属组织的成员资料

列出授权用户在所属组织的成员资料

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserMembershipsOrgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserMembershipsOrgsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **active** | **optional.Bool**| 根据成员是否已激活进行筛选资料，缺省返回所有资料 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserMembershipsOrgsOrg**
> GroupMember GetV5UserMembershipsOrgsOrg(ctx, org, optional)
获取授权用户在一个组织的成员资料

获取授权用户在一个组织的成员资料

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***GetV5UserMembershipsOrgsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserMembershipsOrgsOrgOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**GroupMember**](GroupMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserOrgs**
> []Group GetV5UserOrgs(ctx, optional)
列出授权用户所属的组织

列出授权用户所属的组织

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserOrgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserOrgsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **admin** | **optional.Bool**| 只列出授权用户管理的组织 | 

### Return type

[**[]Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameOrgs**
> []Group GetV5UsersUsernameOrgs(ctx, username, optional)
列出用户所属的组织

列出用户所属的组织

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameOrgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameOrgsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5OrgsOrg**
> GroupDetail PatchV5OrgsOrg(ctx, org, optional)
更新授权用户所管理的组织资料

更新授权用户所管理的组织资料

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***PatchV5OrgsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5OrgsOrgOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **email** | **optional.String**| 组织公开的邮箱地址 | 
 **location** | **optional.String**| 组织所在地 | 
 **name** | **optional.String**| 组织名称 | 
 **description** | **optional.String**| 组织简介 | 
 **htmlUrl** | **optional.String**| 组织站点 | 

### Return type

[**GroupDetail**](GroupDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5UserMembershipsOrgsOrg**
> GroupMember PatchV5UserMembershipsOrgsOrg(ctx, org, optional)
更新授权用户在一个组织的成员资料

更新授权用户在一个组织的成员资料

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***PatchV5UserMembershipsOrgsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5UserMembershipsOrgsOrgOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **remark** | **optional.String**| 在组织中的备注信息 | 

### Return type

[**GroupMember**](GroupMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5UsersOrganization**
> Group PostV5UsersOrganization(ctx, name, org, optional)
创建组织

创建组织

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 组织名称 | 
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***PostV5UsersOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5UsersOrganizationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **description** | **optional.String**| 组织描述 | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5OrgsOrgMembershipsUsername**
> GroupMember PutV5OrgsOrgMembershipsUsername(ctx, org, username, optional)
增加或更新授权用户所管理组织的成员

增加或更新授权用户所管理组织的成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***PutV5OrgsOrgMembershipsUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5OrgsOrgMembershipsUsernameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **role** | **optional.String**| 设置用户在组织的角色 | [default to member]

### Return type

[**GroupMember**](GroupMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

