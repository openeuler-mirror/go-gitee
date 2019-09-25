# \MiscellaneousApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV5Emojis**](MiscellaneousApi.md#GetV5Emojis) | **Get** /v5/emojis | 列出可使用的 Emoji
[**GetV5GitignoreTemplates**](MiscellaneousApi.md#GetV5GitignoreTemplates) | **Get** /v5/gitignore/templates | 列出可使用的 .gitignore 模板
[**GetV5GitignoreTemplatesName**](MiscellaneousApi.md#GetV5GitignoreTemplatesName) | **Get** /v5/gitignore/templates/{name} | 获取一个 .gitignore 模板
[**GetV5GitignoreTemplatesNameRaw**](MiscellaneousApi.md#GetV5GitignoreTemplatesNameRaw) | **Get** /v5/gitignore/templates/{name}/raw | 获取一个 .gitignore 模板原始文件
[**GetV5Licenses**](MiscellaneousApi.md#GetV5Licenses) | **Get** /v5/licenses | 列出可使用的开源许可协议
[**GetV5LicensesLicense**](MiscellaneousApi.md#GetV5LicensesLicense) | **Get** /v5/licenses/{license} | 获取一个开源许可协议
[**GetV5LicensesLicenseRaw**](MiscellaneousApi.md#GetV5LicensesLicenseRaw) | **Get** /v5/licenses/{license}/raw | 获取一个开源许可协议原始文件
[**GetV5ReposOwnerRepoLicense**](MiscellaneousApi.md#GetV5ReposOwnerRepoLicense) | **Get** /v5/repos/{owner}/{repo}/license | 获取一个仓库使用的开源许可协议
[**PostV5Markdown**](MiscellaneousApi.md#PostV5Markdown) | **Post** /v5/markdown | 渲染 Markdown 文本


# **GetV5Emojis**
> GetV5Emojis(ctx, optional)
列出可使用的 Emoji

列出可使用的 Emoji

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5EmojisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EmojisOpts struct

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

# **GetV5GitignoreTemplates**
> GetV5GitignoreTemplates(ctx, optional)
列出可使用的 .gitignore 模板

列出可使用的 .gitignore 模板

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5GitignoreTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GitignoreTemplatesOpts struct

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

# **GetV5GitignoreTemplatesName**
> GetV5GitignoreTemplatesName(ctx, name, optional)
获取一个 .gitignore 模板

获取一个 .gitignore 模板

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| .gitignore 模板名 | 
 **optional** | ***GetV5GitignoreTemplatesNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GitignoreTemplatesNameOpts struct

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

# **GetV5GitignoreTemplatesNameRaw**
> GetV5GitignoreTemplatesNameRaw(ctx, name, optional)
获取一个 .gitignore 模板原始文件

获取一个 .gitignore 模板原始文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| .gitignore 模板名 | 
 **optional** | ***GetV5GitignoreTemplatesNameRawOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5GitignoreTemplatesNameRawOpts struct

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

# **GetV5Licenses**
> GetV5Licenses(ctx, optional)
列出可使用的开源许可协议

列出可使用的开源许可协议

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5LicensesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5LicensesOpts struct

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

# **GetV5LicensesLicense**
> GetV5LicensesLicense(ctx, license, optional)
获取一个开源许可协议

获取一个开源许可协议

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **license** | **string**| 协议名称 | 
 **optional** | ***GetV5LicensesLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5LicensesLicenseOpts struct

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

# **GetV5LicensesLicenseRaw**
> GetV5LicensesLicenseRaw(ctx, license, optional)
获取一个开源许可协议原始文件

获取一个开源许可协议原始文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **license** | **string**| 协议名称 | 
 **optional** | ***GetV5LicensesLicenseRawOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5LicensesLicenseRawOpts struct

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

# **GetV5ReposOwnerRepoLicense**
> GetV5ReposOwnerRepoLicense(ctx, owner, repo, optional)
获取一个仓库使用的开源许可协议

获取一个仓库使用的开源许可协议

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoLicenseOpts struct

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

# **PostV5Markdown**
> PostV5Markdown(ctx, text, optional)
渲染 Markdown 文本

渲染 Markdown 文本

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **text** | **string**| Markdown 文本 | 
 **optional** | ***PostV5MarkdownOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5MarkdownOpts struct

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

