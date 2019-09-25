# \WebhooksApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5ReposOwnerRepoHooksId**](WebhooksApi.md#DeleteV5ReposOwnerRepoHooksId) | **Delete** /v5/repos/{owner}/{repo}/hooks/{id} | 删除一个仓库WebHook
[**GetV5ReposOwnerRepoHooks**](WebhooksApi.md#GetV5ReposOwnerRepoHooks) | **Get** /v5/repos/{owner}/{repo}/hooks | 列出仓库的WebHooks
[**GetV5ReposOwnerRepoHooksId**](WebhooksApi.md#GetV5ReposOwnerRepoHooksId) | **Get** /v5/repos/{owner}/{repo}/hooks/{id} | 获取仓库单个WebHook
[**PatchV5ReposOwnerRepoHooksId**](WebhooksApi.md#PatchV5ReposOwnerRepoHooksId) | **Patch** /v5/repos/{owner}/{repo}/hooks/{id} | 更新一个仓库WebHook
[**PostV5ReposOwnerRepoHooks**](WebhooksApi.md#PostV5ReposOwnerRepoHooks) | **Post** /v5/repos/{owner}/{repo}/hooks | 创建一个仓库WebHook
[**PostV5ReposOwnerRepoHooksIdTests**](WebhooksApi.md#PostV5ReposOwnerRepoHooksIdTests) | **Post** /v5/repos/{owner}/{repo}/hooks/{id}/tests | 测试WebHook是否发送成功


# **DeleteV5ReposOwnerRepoHooksId**
> DeleteV5ReposOwnerRepoHooksId(ctx, owner, repo, id, optional)
删除一个仓库WebHook

删除一个仓库WebHook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| Webhook的ID | 
 **optional** | ***DeleteV5ReposOwnerRepoHooksIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoHooksIdOpts struct

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

# **GetV5ReposOwnerRepoHooks**
> []Hook GetV5ReposOwnerRepoHooks(ctx, owner, repo, optional)
列出仓库的WebHooks

列出仓库的WebHooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoHooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoHooksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoHooksId**
> Hook GetV5ReposOwnerRepoHooksId(ctx, owner, repo, id, optional)
获取仓库单个WebHook

获取仓库单个WebHook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| Webhook的ID | 
 **optional** | ***GetV5ReposOwnerRepoHooksIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoHooksIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5ReposOwnerRepoHooksId**
> Hook PatchV5ReposOwnerRepoHooksId(ctx, owner, repo, id, url, optional)
更新一个仓库WebHook

更新一个仓库WebHook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| Webhook的ID | 
  **url** | **string**| 远程HTTP URL | 
 **optional** | ***PatchV5ReposOwnerRepoHooksIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5ReposOwnerRepoHooksIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessToken** | **optional.String**| 用户授权码 | 
 **password** | **optional.String**| 请求URL时会带上该密码，防止URL被恶意请求 | 
 **pushEvents** | **optional.Bool**| Push代码到仓库 | [default to true]
 **tagPushEvents** | **optional.Bool**| 提交Tag到仓库 | 
 **issuesEvents** | **optional.Bool**| 创建/关闭Issue | 
 **noteEvents** | **optional.Bool**| 评论了Issue/代码等等 | 
 **mergeRequestsEvents** | **optional.Bool**| 合并请求和合并后 | 

### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoHooks**
> Hook PostV5ReposOwnerRepoHooks(ctx, owner, repo, url, optional)
创建一个仓库WebHook

创建一个仓库WebHook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **url** | **string**| 远程HTTP URL | 
 **optional** | ***PostV5ReposOwnerRepoHooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5ReposOwnerRepoHooksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 
 **password** | **optional.String**| 请求URL时会带上该密码，防止URL被恶意请求 | 
 **pushEvents** | **optional.Bool**| Push代码到仓库 | [default to true]
 **tagPushEvents** | **optional.Bool**| 提交Tag到仓库 | 
 **issuesEvents** | **optional.Bool**| 创建/关闭Issue | 
 **noteEvents** | **optional.Bool**| 评论了Issue/代码等等 | 
 **mergeRequestsEvents** | **optional.Bool**| 合并请求和合并后 | 

### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoHooksIdTests**
> PostV5ReposOwnerRepoHooksIdTests(ctx, owner, repo, id, optional)
测试WebHook是否发送成功

测试WebHook是否发送成功

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| Webhook的ID | 
 **optional** | ***PostV5ReposOwnerRepoHooksIdTestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5ReposOwnerRepoHooksIdTestsOpts struct

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

