# \GitDataApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV5ReposOwnerRepoGitBlobsSha**](GitDataApi.md#GetV5ReposOwnerRepoGitBlobsSha) | **Get** /v5/repos/{owner}/{repo}/git/blobs/{sha} | 获取文件Blob
[**GetV5ReposOwnerRepoGitTreesSha**](GitDataApi.md#GetV5ReposOwnerRepoGitTreesSha) | **Get** /v5/repos/{owner}/{repo}/git/trees/{sha} | 获取目录Tree


# **GetV5ReposOwnerRepoGitBlobsSha**
> Blob GetV5ReposOwnerRepoGitBlobsSha(ctx, owner, repo, sha, optional)
获取文件Blob

获取文件Blob

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **sha** | **string**| 文件的 Blob SHA，可通过 [获取仓库具体路径下的内容] API 获取 | 
 **optional** | ***GetV5ReposOwnerRepoGitBlobsShaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoGitBlobsShaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Blob**](Blob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoGitTreesSha**
> Tree GetV5ReposOwnerRepoGitTreesSha(ctx, owner, repo, sha, optional)
获取目录Tree

获取目录Tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **sha** | **string**| 可以是分支名(如master)、Commit或者目录Tree的SHA值 | 
 **optional** | ***GetV5ReposOwnerRepoGitTreesShaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoGitTreesShaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 
 **recursive** | **optional.Int32**| 赋值为1递归获取目录 | 

### Return type

[**Tree**](Tree.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

