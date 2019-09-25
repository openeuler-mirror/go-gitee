# \MilestonesApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5ReposOwnerRepoMilestonesNumber**](MilestonesApi.md#DeleteV5ReposOwnerRepoMilestonesNumber) | **Delete** /v5/repos/{owner}/{repo}/milestones/{number} | 删除仓库单个里程碑
[**GetV5ReposOwnerRepoMilestones**](MilestonesApi.md#GetV5ReposOwnerRepoMilestones) | **Get** /v5/repos/{owner}/{repo}/milestones | 获取仓库所有里程碑
[**GetV5ReposOwnerRepoMilestonesNumber**](MilestonesApi.md#GetV5ReposOwnerRepoMilestonesNumber) | **Get** /v5/repos/{owner}/{repo}/milestones/{number} | 获取仓库单个里程碑
[**PatchV5ReposOwnerRepoMilestonesNumber**](MilestonesApi.md#PatchV5ReposOwnerRepoMilestonesNumber) | **Patch** /v5/repos/{owner}/{repo}/milestones/{number} | 更新仓库里程碑
[**PostV5ReposOwnerRepoMilestones**](MilestonesApi.md#PostV5ReposOwnerRepoMilestones) | **Post** /v5/repos/{owner}/{repo}/milestones | 创建仓库里程碑


# **DeleteV5ReposOwnerRepoMilestonesNumber**
> DeleteV5ReposOwnerRepoMilestonesNumber(ctx, owner, repo, number, optional)
删除仓库单个里程碑

删除仓库单个里程碑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 里程碑序号(id) | 
 **optional** | ***DeleteV5ReposOwnerRepoMilestonesNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoMilestonesNumberOpts struct

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

# **GetV5ReposOwnerRepoMilestones**
> []Milestone GetV5ReposOwnerRepoMilestones(ctx, owner, repo, optional)
获取仓库所有里程碑

获取仓库所有里程碑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoMilestonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoMilestonesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **state** | **optional.String**| 里程碑状态: open, closed, all。默认: open | [default to open]
 **sort** | **optional.String**| 排序方式: due_on | [default to due_on]
 **direction** | **optional.String**| 升序(asc)或是降序(desc)。默认: asc | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Milestone**](Milestone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoMilestonesNumber**
> Milestone GetV5ReposOwnerRepoMilestonesNumber(ctx, owner, repo, number, optional)
获取仓库单个里程碑

获取仓库单个里程碑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 里程碑序号(id) | 
 **optional** | ***GetV5ReposOwnerRepoMilestonesNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoMilestonesNumberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5ReposOwnerRepoMilestonesNumber**
> Milestone PatchV5ReposOwnerRepoMilestonesNumber(ctx, owner, repo, number, title, dueOn, optional)
更新仓库里程碑

更新仓库里程碑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 里程碑序号(id) | 
  **title** | **string**| 里程碑标题 | 
  **dueOn** | **string**| 里程碑的截止日期 | 
 **optional** | ***PatchV5ReposOwnerRepoMilestonesNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5ReposOwnerRepoMilestonesNumberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **accessToken** | **optional.String**| 用户授权码 | 
 **state** | **optional.String**| 里程碑状态: open, closed, all。默认: open | [default to open]
 **description** | **optional.String**| 里程碑具体描述 | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoMilestones**
> Milestone PostV5ReposOwnerRepoMilestones(ctx, owner, repo, title, dueOn, optional)
创建仓库里程碑

创建仓库里程碑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **title** | **string**| 里程碑标题 | 
  **dueOn** | **string**| 里程碑的截止日期 | 
 **optional** | ***PostV5ReposOwnerRepoMilestonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5ReposOwnerRepoMilestonesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessToken** | **optional.String**| 用户授权码 | 
 **state** | **optional.String**| 里程碑状态: open, closed, all。默认: open | [default to open]
 **description** | **optional.String**| 里程碑具体描述 | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

