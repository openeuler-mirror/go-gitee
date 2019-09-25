# \UsersApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5UserFollowingUsername**](UsersApi.md#DeleteV5UserFollowingUsername) | **Delete** /v5/user/following/{username} | 取消关注一个用户
[**DeleteV5UserKeysId**](UsersApi.md#DeleteV5UserKeysId) | **Delete** /v5/user/keys/{id} | 删除一个公钥
[**GetV5User**](UsersApi.md#GetV5User) | **Get** /v5/user | 获取授权用户的资料
[**GetV5UserFollowers**](UsersApi.md#GetV5UserFollowers) | **Get** /v5/user/followers | 列出授权用户的关注者
[**GetV5UserFollowing**](UsersApi.md#GetV5UserFollowing) | **Get** /v5/user/following | 列出授权用户正关注的用户
[**GetV5UserFollowingUsername**](UsersApi.md#GetV5UserFollowingUsername) | **Get** /v5/user/following/{username} | 检查授权用户是否关注了一个用户
[**GetV5UserKeys**](UsersApi.md#GetV5UserKeys) | **Get** /v5/user/keys | 列出授权用户的所有公钥
[**GetV5UserKeysId**](UsersApi.md#GetV5UserKeysId) | **Get** /v5/user/keys/{id} | 获取一个公钥
[**GetV5UserNamespace**](UsersApi.md#GetV5UserNamespace) | **Get** /v5/user/namespace | 获取授权用户的一个 Namespace
[**GetV5UserNamespaces**](UsersApi.md#GetV5UserNamespaces) | **Get** /v5/user/namespaces | 列出授权用户所有的 Namespace
[**GetV5UsersUsername**](UsersApi.md#GetV5UsersUsername) | **Get** /v5/users/{username} | 获取一个用户
[**GetV5UsersUsernameFollowers**](UsersApi.md#GetV5UsersUsernameFollowers) | **Get** /v5/users/{username}/followers | 列出指定用户的关注者
[**GetV5UsersUsernameFollowing**](UsersApi.md#GetV5UsersUsernameFollowing) | **Get** /v5/users/{username}/following | 列出指定用户正在关注的用户
[**GetV5UsersUsernameFollowingTargetUser**](UsersApi.md#GetV5UsersUsernameFollowingTargetUser) | **Get** /v5/users/{username}/following/{target_user} | 检查指定用户是否关注目标用户
[**GetV5UsersUsernameKeys**](UsersApi.md#GetV5UsersUsernameKeys) | **Get** /v5/users/{username}/keys | 列出指定用户的所有公钥
[**PatchV5User**](UsersApi.md#PatchV5User) | **Patch** /v5/user | 更新授权用户的资料
[**PostV5UserKeys**](UsersApi.md#PostV5UserKeys) | **Post** /v5/user/keys | 添加一个公钥
[**PutV5UserFollowingUsername**](UsersApi.md#PutV5UserFollowingUsername) | **Put** /v5/user/following/{username} | 关注一个用户


# **DeleteV5UserFollowingUsername**
> DeleteV5UserFollowingUsername(ctx, username, optional)
取消关注一个用户

取消关注一个用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***DeleteV5UserFollowingUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5UserFollowingUsernameOpts struct

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

# **DeleteV5UserKeysId**
> DeleteV5UserKeysId(ctx, id, optional)
删除一个公钥

删除一个公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| 公钥 ID | 
 **optional** | ***DeleteV5UserKeysIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5UserKeysIdOpts struct

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

# **GetV5User**
> User GetV5User(ctx, optional)
获取授权用户的资料

获取授权用户的资料

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserFollowers**
> []UserBasic GetV5UserFollowers(ctx, optional)
列出授权用户的关注者

列出授权用户的关注者

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserFollowersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserFollowersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]UserBasic**](UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserFollowing**
> []UserBasic GetV5UserFollowing(ctx, optional)
列出授权用户正关注的用户

列出授权用户正关注的用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserFollowingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserFollowingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]UserBasic**](UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserFollowingUsername**
> GetV5UserFollowingUsername(ctx, username, optional)
检查授权用户是否关注了一个用户

检查授权用户是否关注了一个用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UserFollowingUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserFollowingUsernameOpts struct

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

# **GetV5UserKeys**
> []SshKey GetV5UserKeys(ctx, optional)
列出授权用户的所有公钥

列出授权用户的所有公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]SshKey**](SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserKeysId**
> SshKey GetV5UserKeysId(ctx, id, optional)
获取一个公钥

获取一个公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| 公钥 ID | 
 **optional** | ***GetV5UserKeysIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserKeysIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserNamespace**
> []Namespace GetV5UserNamespace(ctx, path, optional)
获取授权用户的一个 Namespace

获取授权用户的一个 Namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Namespace path | 
 **optional** | ***GetV5UserNamespaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserNamespaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**[]Namespace**](Namespace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserNamespaces**
> []Namespace GetV5UserNamespaces(ctx, optional)
列出授权用户所有的 Namespace

列出授权用户所有的 Namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserNamespacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserNamespacesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **mode** | **optional.String**| 参与方式: project(所有参与仓库的namepsce)、intrant(所加入的namespace)、all(包含前两者)，默认(intrant) | 

### Return type

[**[]Namespace**](Namespace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsername**
> User GetV5UsersUsername(ctx, username, optional)
获取一个用户

获取一个用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameFollowers**
> []UserBasic GetV5UsersUsernameFollowers(ctx, username, optional)
列出指定用户的关注者

列出指定用户的关注者

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameFollowersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameFollowersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]UserBasic**](UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameFollowing**
> []UserBasic GetV5UsersUsernameFollowing(ctx, username, optional)
列出指定用户正在关注的用户

列出指定用户正在关注的用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameFollowingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameFollowingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]UserBasic**](UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameFollowingTargetUser**
> GetV5UsersUsernameFollowingTargetUser(ctx, username, targetUser, optional)
检查指定用户是否关注目标用户

检查指定用户是否关注目标用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
  **targetUser** | **string**| 目标用户的用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameFollowingTargetUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameFollowingTargetUserOpts struct

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

# **GetV5UsersUsernameKeys**
> []SshKeyBasic GetV5UsersUsernameKeys(ctx, username, optional)
列出指定用户的所有公钥

列出指定用户的所有公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]SshKeyBasic**](SSHKeyBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5User**
> User PatchV5User(ctx, optional)
更新授权用户的资料

更新授权用户的资料

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PatchV5UserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5UserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **name** | **optional.String**| 昵称 | 
 **blog** | **optional.String**| 微博链接 | 
 **weibo** | **optional.String**| 博客站点 | 
 **bio** | **optional.String**| 自我介绍 | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5UserKeys**
> SshKey PostV5UserKeys(ctx, key, title, optional)
添加一个公钥

添加一个公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| 公钥内容 | 
  **title** | **string**| 公钥名称 | 
 **optional** | ***PostV5UserKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5UserKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5UserFollowingUsername**
> PutV5UserFollowingUsername(ctx, username, optional)
关注一个用户

关注一个用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***PutV5UserFollowingUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5UserFollowingUsernameOpts struct

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

