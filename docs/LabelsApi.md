# \LabelsApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5ReposOwnerRepoIssuesNumberLabels**](LabelsApi.md#DeleteV5ReposOwnerRepoIssuesNumberLabels) | **Delete** /v5/repos/{owner}/{repo}/issues/{number}/labels | 删除Issue所有标签
[**DeleteV5ReposOwnerRepoIssuesNumberLabelsName**](LabelsApi.md#DeleteV5ReposOwnerRepoIssuesNumberLabelsName) | **Delete** /v5/repos/{owner}/{repo}/issues/{number}/labels/{name} | 删除Issue标签
[**DeleteV5ReposOwnerRepoLabelsName**](LabelsApi.md#DeleteV5ReposOwnerRepoLabelsName) | **Delete** /v5/repos/{owner}/{repo}/labels/{name} | 删除一个仓库任务标签
[**GetV5EnterprisesEnterpriseLabels**](LabelsApi.md#GetV5EnterprisesEnterpriseLabels) | **Get** /v5/enterprises/{enterprise}/labels | 获取企业所有标签
[**GetV5EnterprisesEnterpriseLabelsName**](LabelsApi.md#GetV5EnterprisesEnterpriseLabelsName) | **Get** /v5/enterprises/{enterprise}/labels/{name} | 获取企业某个标签
[**GetV5ReposOwnerRepoIssuesNumberLabels**](LabelsApi.md#GetV5ReposOwnerRepoIssuesNumberLabels) | **Get** /v5/repos/{owner}/{repo}/issues/{number}/labels | 获取仓库任务的所有标签
[**GetV5ReposOwnerRepoLabels**](LabelsApi.md#GetV5ReposOwnerRepoLabels) | **Get** /v5/repos/{owner}/{repo}/labels | 获取仓库所有任务标签
[**GetV5ReposOwnerRepoLabelsName**](LabelsApi.md#GetV5ReposOwnerRepoLabelsName) | **Get** /v5/repos/{owner}/{repo}/labels/{name} | 根据标签名称获取单个标签
[**PatchV5ReposOwnerRepoLabelsOriginalName**](LabelsApi.md#PatchV5ReposOwnerRepoLabelsOriginalName) | **Patch** /v5/repos/{owner}/{repo}/labels/{original_name} | 更新一个仓库任务标签
[**PostV5ReposOwnerRepoIssuesNumberLabels**](LabelsApi.md#PostV5ReposOwnerRepoIssuesNumberLabels) | **Post** /v5/repos/{owner}/{repo}/issues/{number}/labels | 创建Issue标签
[**PostV5ReposOwnerRepoLabels**](LabelsApi.md#PostV5ReposOwnerRepoLabels) | **Post** /v5/repos/{owner}/{repo}/labels | 创建仓库任务标签
[**PostV5ReposOwnerRepoProjectLabels**](LabelsApi.md#PostV5ReposOwnerRepoProjectLabels) | **Post** /v5/repos/{owner}/{repo}/project_labels | 添加仓库标签
[**PutV5ReposOwnerRepoIssuesNumberLabels**](LabelsApi.md#PutV5ReposOwnerRepoIssuesNumberLabels) | **Put** /v5/repos/{owner}/{repo}/issues/{number}/labels | 替换Issue所有标签
[**PutV5ReposOwnerRepoProjectLabels**](LabelsApi.md#PutV5ReposOwnerRepoProjectLabels) | **Put** /v5/repos/{owner}/{repo}/project_labels | 替换所有仓库标签


# **DeleteV5ReposOwnerRepoIssuesNumberLabels**
> DeleteV5ReposOwnerRepoIssuesNumberLabels(ctx, owner, repo, number, optional)
删除Issue所有标签

删除Issue所有标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
 **optional** | ***DeleteV5ReposOwnerRepoIssuesNumberLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoIssuesNumberLabelsOpts struct

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

# **DeleteV5ReposOwnerRepoIssuesNumberLabelsName**
> DeleteV5ReposOwnerRepoIssuesNumberLabelsName(ctx, owner, repo, number, name, optional)
删除Issue标签

删除Issue标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
  **name** | **string**| 标签名称 | 
 **optional** | ***DeleteV5ReposOwnerRepoIssuesNumberLabelsNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoIssuesNumberLabelsNameOpts struct

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

# **DeleteV5ReposOwnerRepoLabelsName**
> DeleteV5ReposOwnerRepoLabelsName(ctx, owner, repo, name, optional)
删除一个仓库任务标签

删除一个仓库任务标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **name** | **string**| 标签名称 | 
 **optional** | ***DeleteV5ReposOwnerRepoLabelsNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoLabelsNameOpts struct

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

# **GetV5EnterprisesEnterpriseLabels**
> []Label GetV5EnterprisesEnterpriseLabels(ctx, enterprise, optional)
获取企业所有标签

获取企业所有标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
 **optional** | ***GetV5EnterprisesEnterpriseLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseLabelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseLabelsName**
> Label GetV5EnterprisesEnterpriseLabelsName(ctx, enterprise, name, optional)
获取企业某个标签

获取企业某个标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **name** | **string**| 标签名称 | 
 **optional** | ***GetV5EnterprisesEnterpriseLabelsNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseLabelsNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoIssuesNumberLabels**
> []Label GetV5ReposOwnerRepoIssuesNumberLabels(ctx, owner, repo, number, optional)
获取仓库任务的所有标签

获取仓库任务的所有标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
 **optional** | ***GetV5ReposOwnerRepoIssuesNumberLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoIssuesNumberLabelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoLabels**
> []Label GetV5ReposOwnerRepoLabels(ctx, owner, repo, optional)
获取仓库所有任务标签

获取仓库所有任务标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoLabelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoLabelsName**
> Label GetV5ReposOwnerRepoLabelsName(ctx, owner, repo, name, optional)
根据标签名称获取单个标签

根据标签名称获取单个标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **name** | **string**| 标签名称 | 
 **optional** | ***GetV5ReposOwnerRepoLabelsNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoLabelsNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5ReposOwnerRepoLabelsOriginalName**
> Label PatchV5ReposOwnerRepoLabelsOriginalName(ctx, owner, repo, originalName, optional)
更新一个仓库任务标签

更新一个仓库任务标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **originalName** | **string**| 标签原有名称 | 
 **optional** | ***PatchV5ReposOwnerRepoLabelsOriginalNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5ReposOwnerRepoLabelsOriginalNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 
 **name** | **optional.String**| 标签新名称 | 
 **color** | **optional.String**| 标签新颜色 | 

### Return type

[**Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoIssuesNumberLabels**
> []Label PostV5ReposOwnerRepoIssuesNumberLabels(ctx, owner, repo, number, body)
创建Issue标签

创建Issue标签  需要在请求的body里填上数组，元素为标签的名字。如: [\"performance\", \"bug\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
  **body** | [**PullRequestLabelPostParam**](PullRequestLabelPostParam.md)| 必选，标签的内容 | 

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoLabels**
> Label PostV5ReposOwnerRepoLabels(ctx, owner, repo, body)
创建仓库任务标签

创建仓库任务标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **body** | [**LabelPostParam**](LabelPostParam.md)| 必选，标签的内容 | 

### Return type

[**Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoProjectLabels**
> []ProjectLabel PostV5ReposOwnerRepoProjectLabels(ctx, owner, repo, body)
添加仓库标签

添加仓库标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **body** | [**PullRequestLabelPostParam**](PullRequestLabelPostParam.md)| 必选，标签的内容 | 

### Return type

[**[]ProjectLabel**](ProjectLabel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5ReposOwnerRepoIssuesNumberLabels**
> []Label PutV5ReposOwnerRepoIssuesNumberLabels(ctx, owner, repo, number, body)
替换Issue所有标签

替换Issue所有标签  需要在请求的body里填上数组，元素为标签的名字。如: [\"performance\", \"bug\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
  **body** | [**PullRequestLabelPostParam**](PullRequestLabelPostParam.md)| 必选，标签的内容 | 

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5ReposOwnerRepoProjectLabels**
> []ProjectLabel PutV5ReposOwnerRepoProjectLabels(ctx, owner, repo, body)
替换所有仓库标签

替换所有仓库标签  需要在请求的body里填上数组，元素为标签的名字。如: [\"feat\", \"bug\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **body** | [**PullRequestLabelPostParam**](PullRequestLabelPostParam.md)| 必选，标签的内容 | 

### Return type

[**[]ProjectLabel**](ProjectLabel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

