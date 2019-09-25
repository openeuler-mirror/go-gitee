# \GistsApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5GistsGistIdCommentsId**](GistsApi.md#DeleteV5GistsGistIdCommentsId) | **Delete** /v5/gists/{gist_id}/comments/{id} | 删除代码片段的评论
[**DeleteV5GistsId**](GistsApi.md#DeleteV5GistsId) | **Delete** /v5/gists/{id} | 删除指定代码片段
[**DeleteV5GistsIdStar**](GistsApi.md#DeleteV5GistsIdStar) | **Delete** /v5/gists/{id}/star | 取消Star代码片段
[**GetV5Gists**](GistsApi.md#GetV5Gists) | **Get** /v5/gists | 获取代码片段
[**GetV5GistsGistIdComments**](GistsApi.md#GetV5GistsGistIdComments) | **Get** /v5/gists/{gist_id}/comments | 获取代码片段的评论
[**GetV5GistsGistIdCommentsId**](GistsApi.md#GetV5GistsGistIdCommentsId) | **Get** /v5/gists/{gist_id}/comments/{id} | 获取单条代码片段的评论
[**GetV5GistsId**](GistsApi.md#GetV5GistsId) | **Get** /v5/gists/{id} | 获取单条代码片段
[**GetV5GistsIdCommits**](GistsApi.md#GetV5GistsIdCommits) | **Get** /v5/gists/{id}/commits | 获取代码片段的commit
[**GetV5GistsIdForks**](GistsApi.md#GetV5GistsIdForks) | **Get** /v5/gists/{id}/forks | 获取 Fork 了指定代码片段的列表
[**GetV5GistsIdStar**](GistsApi.md#GetV5GistsIdStar) | **Get** /v5/gists/{id}/star | 判断代码片段是否已Star
[**GetV5GistsPublic**](GistsApi.md#GetV5GistsPublic) | **Get** /v5/gists/public | 获取公开的代码片段
[**GetV5GistsStarred**](GistsApi.md#GetV5GistsStarred) | **Get** /v5/gists/starred | 获取用户Star的代码片段
[**GetV5UsersUsernameGists**](GistsApi.md#GetV5UsersUsernameGists) | **Get** /v5/users/{username}/gists | 获取指定用户的公开代码片段
[**PatchV5GistsGistIdCommentsId**](GistsApi.md#PatchV5GistsGistIdCommentsId) | **Patch** /v5/gists/{gist_id}/comments/{id} | 修改代码片段的评论
[**PatchV5GistsId**](GistsApi.md#PatchV5GistsId) | **Patch** /v5/gists/{id} | 修改代码片段
[**PostV5Gists**](GistsApi.md#PostV5Gists) | **Post** /v5/gists | 创建代码片段
[**PostV5GistsGistIdComments**](GistsApi.md#PostV5GistsGistIdComments) | **Post** /v5/gists/{gist_id}/comments | 增加代码片段的评论
[**PostV5GistsIdForks**](GistsApi.md#PostV5GistsIdForks) | **Post** /v5/gists/{id}/forks | Fork代码片段
[**PutV5GistsIdStar**](GistsApi.md#PutV5GistsIdStar) | **Put** /v5/gists/{id}/star | Star代码片段


# **DeleteV5GistsGistIdCommentsId**
> DeleteV5GistsGistIdCommentsId(ctx, gistId, id, optional)
删除代码片段的评论

删除代码片段的评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gistId** | **string**| 代码片段的ID | 
  **id** | **int32**| 评论的ID | 
 **optional** | ***DeleteV5GistsGistIdCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5GistsGistIdCommentsIdOpts struct

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

# **DeleteV5GistsId**
> DeleteV5GistsId(ctx, id, optional)
删除指定代码片段

删除指定代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 代码片段的ID | 
 **optional** | ***DeleteV5GistsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5GistsIdOpts struct

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

# **DeleteV5GistsIdStar**
> DeleteV5GistsIdStar(ctx, id, optional)
取消Star代码片段

取消Star代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 代码片段的ID | 
 **optional** | ***DeleteV5GistsIdStarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5GistsIdStarOpts struct

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

# **GetV5Gists**
> []Code GetV5Gists(ctx, optional)
获取代码片段

获取代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5GistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GistsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **since** | **optional.String**| 起始的更新时间，要求时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Code**](Code.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5GistsGistIdComments**
> []CodeComment GetV5GistsGistIdComments(ctx, gistId, optional)
获取代码片段的评论

获取代码片段的评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gistId** | **string**| 代码片段的ID | 
 **optional** | ***GetV5GistsGistIdCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GistsGistIdCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]CodeComment**](CodeComment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5GistsGistIdCommentsId**
> CodeComment GetV5GistsGistIdCommentsId(ctx, gistId, id, optional)
获取单条代码片段的评论

获取单条代码片段的评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gistId** | **string**| 代码片段的ID | 
  **id** | **int32**| 评论的ID | 
 **optional** | ***GetV5GistsGistIdCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GistsGistIdCommentsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**CodeComment**](CodeComment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5GistsId**
> CodeForksHistory GetV5GistsId(ctx, id, optional)
获取单条代码片段

获取单条代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 代码片段的ID | 
 **optional** | ***GetV5GistsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GistsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**CodeForksHistory**](CodeForksHistory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5GistsIdCommits**
> CodeForksHistory GetV5GistsIdCommits(ctx, id, optional)
获取代码片段的commit

获取代码片段的commit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 代码片段的ID | 
 **optional** | ***GetV5GistsIdCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GistsIdCommitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**CodeForksHistory**](CodeForksHistory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5GistsIdForks**
> CodeForks GetV5GistsIdForks(ctx, id, optional)
获取 Fork 了指定代码片段的列表

获取 Fork 了指定代码片段的列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 代码片段的ID | 
 **optional** | ***GetV5GistsIdForksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GistsIdForksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**CodeForks**](CodeForks.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5GistsIdStar**
> GetV5GistsIdStar(ctx, id, optional)
判断代码片段是否已Star

判断代码片段是否已Star

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 代码片段的ID | 
 **optional** | ***GetV5GistsIdStarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GistsIdStarOpts struct

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

# **GetV5GistsPublic**
> []Code GetV5GistsPublic(ctx, optional)
获取公开的代码片段

获取公开的代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5GistsPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GistsPublicOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **since** | **optional.String**| 起始的更新时间，要求时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Code**](Code.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5GistsStarred**
> []Code GetV5GistsStarred(ctx, optional)
获取用户Star的代码片段

获取用户Star的代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5GistsStarredOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GistsStarredOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **since** | **optional.String**| 起始的更新时间，要求时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Code**](Code.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameGists**
> []Code GetV5UsersUsernameGists(ctx, username, optional)
获取指定用户的公开代码片段

获取指定用户的公开代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameGistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameGistsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Code**](Code.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5GistsGistIdCommentsId**
> CodeComment PatchV5GistsGistIdCommentsId(ctx, gistId, id, body, optional)
修改代码片段的评论

修改代码片段的评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gistId** | **string**| 代码片段的ID | 
  **id** | **int32**| 评论的ID | 
  **body** | **string**| 评论内容 | 
 **optional** | ***PatchV5GistsGistIdCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5GistsGistIdCommentsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**CodeComment**](CodeComment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5GistsId**
> CodeForksHistory PatchV5GistsId(ctx, id, optional)
修改代码片段

修改代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 代码片段的ID | 
 **optional** | ***PatchV5GistsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5GistsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **files** | **optional.Interface of *os.File**| Hash形式的代码片段文件名以及文件内容。如: { \&quot;file1.txt\&quot;: { \&quot;content\&quot;: \&quot;String file contents\&quot; } } | 
 **description** | **optional.String**| 代码片段描述，1~30个字符 | 
 **public** | **optional.Bool**| 公开/私有，默认: 私有 | 

### Return type

[**CodeForksHistory**](CodeForksHistory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5Gists**
> []CodeForksHistory PostV5Gists(ctx, files, description, optional)
创建代码片段

创建代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **files** | ***os.File**| Hash形式的代码片段文件名以及文件内容。如: { \&quot;file1.txt\&quot;: { \&quot;content\&quot;: \&quot;String file contents\&quot; } } | 
  **description** | **string**| 代码片段描述，1~30个字符 | 
 **optional** | ***PostV5GistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5GistsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **public** | **optional.Bool**| 公开/私有，默认: 私有 | 

### Return type

[**[]CodeForksHistory**](CodeForksHistory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5GistsGistIdComments**
> CodeComment PostV5GistsGistIdComments(ctx, gistId, body, optional)
增加代码片段的评论

增加代码片段的评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gistId** | **string**| 代码片段的ID | 
  **body** | **string**| 评论内容 | 
 **optional** | ***PostV5GistsGistIdCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5GistsGistIdCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**CodeComment**](CodeComment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5GistsIdForks**
> PostV5GistsIdForks(ctx, id, optional)
Fork代码片段

Fork代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 代码片段的ID | 
 **optional** | ***PostV5GistsIdForksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5GistsIdForksOpts struct

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

# **PutV5GistsIdStar**
> PutV5GistsIdStar(ctx, id, optional)
Star代码片段

Star代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 代码片段的ID | 
 **optional** | ***PutV5GistsIdStarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5GistsIdStarOpts struct

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

