# \IssuesApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5ReposOwnerRepoIssuesCommentsId**](IssuesApi.md#DeleteV5ReposOwnerRepoIssuesCommentsId) | **Delete** /v5/repos/{owner}/{repo}/issues/comments/{id} | 删除Issue某条评论
[**GetV5EnterprisesEnterpriseIssues**](IssuesApi.md#GetV5EnterprisesEnterpriseIssues) | **Get** /v5/enterprises/{enterprise}/issues | 获取某个企业的所有Issues
[**GetV5EnterprisesEnterpriseIssuesNumber**](IssuesApi.md#GetV5EnterprisesEnterpriseIssuesNumber) | **Get** /v5/enterprises/{enterprise}/issues/{number} | 获取企业的某个Issue
[**GetV5EnterprisesEnterpriseIssuesNumberComments**](IssuesApi.md#GetV5EnterprisesEnterpriseIssuesNumberComments) | **Get** /v5/enterprises/{enterprise}/issues/{number}/comments | 获取企业某个Issue所有评论
[**GetV5EnterprisesEnterpriseIssuesNumberLabels**](IssuesApi.md#GetV5EnterprisesEnterpriseIssuesNumberLabels) | **Get** /v5/enterprises/{enterprise}/issues/{number}/labels | 获取企业某个Issue所有标签
[**GetV5Issues**](IssuesApi.md#GetV5Issues) | **Get** /v5/issues | 获取当前授权用户的所有Issues
[**GetV5OrgsOrgIssues**](IssuesApi.md#GetV5OrgsOrgIssues) | **Get** /v5/orgs/{org}/issues | 获取当前用户某个组织的Issues
[**GetV5ReposOwnerIssuesNumberOperateLogs**](IssuesApi.md#GetV5ReposOwnerIssuesNumberOperateLogs) | **Get** /v5/repos/{owner}/issues/{number}/operate_logs | 获取某个Issue下的操作日志
[**GetV5ReposOwnerRepoIssues**](IssuesApi.md#GetV5ReposOwnerRepoIssues) | **Get** /v5/repos/{owner}/{repo}/issues | 仓库的所有Issues
[**GetV5ReposOwnerRepoIssuesComments**](IssuesApi.md#GetV5ReposOwnerRepoIssuesComments) | **Get** /v5/repos/{owner}/{repo}/issues/comments | 获取仓库所有Issue的评论
[**GetV5ReposOwnerRepoIssuesCommentsId**](IssuesApi.md#GetV5ReposOwnerRepoIssuesCommentsId) | **Get** /v5/repos/{owner}/{repo}/issues/comments/{id} | 获取仓库Issue某条评论
[**GetV5ReposOwnerRepoIssuesNumber**](IssuesApi.md#GetV5ReposOwnerRepoIssuesNumber) | **Get** /v5/repos/{owner}/{repo}/issues/{number} | 仓库的某个Issue
[**GetV5ReposOwnerRepoIssuesNumberComments**](IssuesApi.md#GetV5ReposOwnerRepoIssuesNumberComments) | **Get** /v5/repos/{owner}/{repo}/issues/{number}/comments | 获取仓库某个Issue所有的评论
[**GetV5UserIssues**](IssuesApi.md#GetV5UserIssues) | **Get** /v5/user/issues | 获取授权用户的所有Issues
[**PatchV5ReposOwnerIssuesNumber**](IssuesApi.md#PatchV5ReposOwnerIssuesNumber) | **Patch** /v5/repos/{owner}/issues/{number} | 更新Issue
[**PatchV5ReposOwnerRepoIssuesCommentsId**](IssuesApi.md#PatchV5ReposOwnerRepoIssuesCommentsId) | **Patch** /v5/repos/{owner}/{repo}/issues/comments/{id} | 更新Issue某条评论
[**PostV5ReposOwnerIssues**](IssuesApi.md#PostV5ReposOwnerIssues) | **Post** /v5/repos/{owner}/issues | 创建Issue
[**PostV5ReposOwnerRepoIssuesNumberComments**](IssuesApi.md#PostV5ReposOwnerRepoIssuesNumberComments) | **Post** /v5/repos/{owner}/{repo}/issues/{number}/comments | 创建某个Issue评论


# **DeleteV5ReposOwnerRepoIssuesCommentsId**
> DeleteV5ReposOwnerRepoIssuesCommentsId(ctx, owner, repo, id, optional)
删除Issue某条评论

删除Issue某条评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 评论的ID | 
 **optional** | ***DeleteV5ReposOwnerRepoIssuesCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoIssuesCommentsIdOpts struct

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

# **GetV5EnterprisesEnterpriseIssues**
> []Issue GetV5EnterprisesEnterpriseIssues(ctx, enterprise, optional)
获取某个企业的所有Issues

获取某个企业的所有Issues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
 **optional** | ***GetV5EnterprisesEnterpriseIssuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseIssuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **state** | **optional.String**| Issue的状态: open（开启的）, progressing(进行中), closed（关闭的）, rejected（拒绝的）。 默认: open | [default to open]
 **labels** | **optional.String**| 用逗号分开的标签。如: bug,performance | 
 **sort** | **optional.String**| 排序依据: 创建时间(created)，更新时间(updated_at)。默认: created_at | [default to created]
 **direction** | **optional.String**| 排序方式: 升序(asc)，降序(desc)。默认: desc | [default to desc]
 **since** | **optional.String**| 起始的更新时间，要求时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **schedule** | **optional.String**| 计划开始日期，格式：20181006T173008+80-20181007T173008+80（区间），或者 -20181007T173008+80（小于20181007T173008+80），或者 20181006T173008+80-（大于20181006T173008+80），要求时间格式为20181006T173008+80 | 
 **deadline** | **optional.String**| 计划截止日期，格式同上 | 
 **createdAt** | **optional.String**| 任务创建时间，格式同上 | 
 **finishedAt** | **optional.String**| 任务完成时间，即任务最后一次转为已完成状态的时间点。格式同上 | 
 **milestone** | **optional.String**| 根据里程碑标题。none为没里程碑的，*为所有带里程碑的 | 
 **assignee** | **optional.String**| 用户的username。 none为没指派者, *为所有带有指派者的 | 
 **creator** | **optional.String**| 创建Issues的用户username | 
 **program** | **optional.String**| 所属项目名称。none为没所属有项目的，*为所有带所属项目的 | 

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseIssuesNumber**
> Issue GetV5EnterprisesEnterpriseIssuesNumber(ctx, enterprise, number, optional)
获取企业的某个Issue

获取企业的某个Issue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
 **optional** | ***GetV5EnterprisesEnterpriseIssuesNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseIssuesNumberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseIssuesNumberComments**
> []Note GetV5EnterprisesEnterpriseIssuesNumberComments(ctx, enterprise, number, optional)
获取企业某个Issue所有评论

获取企业某个Issue所有评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
 **optional** | ***GetV5EnterprisesEnterpriseIssuesNumberCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseIssuesNumberCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseIssuesNumberLabels**
> []Label GetV5EnterprisesEnterpriseIssuesNumberLabels(ctx, enterprise, number, optional)
获取企业某个Issue所有标签

获取企业某个Issue所有标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
 **optional** | ***GetV5EnterprisesEnterpriseIssuesNumberLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseIssuesNumberLabelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5Issues**
> []Issue GetV5Issues(ctx, optional)
获取当前授权用户的所有Issues

获取当前授权用户的所有Issues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5IssuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5IssuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **filter** | **optional.String**| 筛选参数: 授权用户负责的(assigned)，授权用户创建的(created)，包含前两者的(all)。默认: assigned | [default to assigned]
 **state** | **optional.String**| Issue的状态: open（开启的）, progressing(进行中), closed（关闭的）, rejected（拒绝的）。 默认: open | [default to open]
 **labels** | **optional.String**| 用逗号分开的标签。如: bug,performance | 
 **sort** | **optional.String**| 排序依据: 创建时间(created)，更新时间(updated_at)。默认: created_at | [default to created]
 **direction** | **optional.String**| 排序方式: 升序(asc)，降序(desc)。默认: desc | [default to desc]
 **since** | **optional.String**| 起始的更新时间，要求时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **schedule** | **optional.String**| 计划开始日期，格式：20181006T173008+80-20181007T173008+80（区间），或者 -20181007T173008+80（小于20181007T173008+80），或者 20181006T173008+80-（大于20181006T173008+80），要求时间格式为20181006T173008+80 | 
 **deadline** | **optional.String**| 计划截止日期，格式同上 | 
 **createdAt** | **optional.String**| 任务创建时间，格式同上 | 
 **finishedAt** | **optional.String**| 任务完成时间，即任务最后一次转为已完成状态的时间点。格式同上 | 

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5OrgsOrgIssues**
> []Issue GetV5OrgsOrgIssues(ctx, org, optional)
获取当前用户某个组织的Issues

获取当前用户某个组织的Issues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***GetV5OrgsOrgIssuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5OrgsOrgIssuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **filter** | **optional.String**| 筛选参数: 授权用户负责的(assigned)，授权用户创建的(created)，包含前两者的(all)。默认: assigned | [default to assigned]
 **state** | **optional.String**| Issue的状态: open（开启的）, progressing(进行中), closed（关闭的）, rejected（拒绝的）。 默认: open | [default to open]
 **labels** | **optional.String**| 用逗号分开的标签。如: bug,performance | 
 **sort** | **optional.String**| 排序依据: 创建时间(created)，更新时间(updated_at)。默认: created_at | [default to created]
 **direction** | **optional.String**| 排序方式: 升序(asc)，降序(desc)。默认: desc | [default to desc]
 **since** | **optional.String**| 起始的更新时间，要求时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **schedule** | **optional.String**| 计划开始日期，格式：20181006T173008+80-20181007T173008+80（区间），或者 -20181007T173008+80（小于20181007T173008+80），或者 20181006T173008+80-（大于20181006T173008+80），要求时间格式为20181006T173008+80 | 
 **deadline** | **optional.String**| 计划截止日期，格式同上 | 
 **createdAt** | **optional.String**| 任务创建时间，格式同上 | 
 **finishedAt** | **optional.String**| 任务完成时间，即任务最后一次转为已完成状态的时间点。格式同上 | 

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerIssuesNumberOperateLogs**
> []OperateLog GetV5ReposOwnerIssuesNumberOperateLogs(ctx, owner, number, optional)
获取某个Issue下的操作日志

获取某个Issue下的操作日志

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
 **optional** | ***GetV5ReposOwnerIssuesNumberOperateLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerIssuesNumberOperateLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **repo** | **optional.String**| 仓库路径(path) | 
 **sort** | **optional.String**| 按递增(asc)或递减(desc)排序，默认：递减 | [default to desc]

### Return type

[**[]OperateLog**](OperateLog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoIssues**
> []Issue GetV5ReposOwnerRepoIssues(ctx, owner, repo, optional)
仓库的所有Issues

仓库的所有Issues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoIssuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoIssuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **state** | **optional.String**| Issue的状态: open（开启的）, progressing(进行中), closed（关闭的）, rejected（拒绝的）。 默认: open | [default to open]
 **labels** | **optional.String**| 用逗号分开的标签。如: bug,performance | 
 **sort** | **optional.String**| 排序依据: 创建时间(created)，更新时间(updated_at)。默认: created_at | [default to created]
 **direction** | **optional.String**| 排序方式: 升序(asc)，降序(desc)。默认: desc | [default to desc]
 **since** | **optional.String**| 起始的更新时间，要求时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **schedule** | **optional.String**| 计划开始日期，格式：20181006T173008+80-20181007T173008+80（区间），或者 -20181007T173008+80（小于20181007T173008+80），或者 20181006T173008+80-（大于20181006T173008+80），要求时间格式为20181006T173008+80 | 
 **deadline** | **optional.String**| 计划截止日期，格式同上 | 
 **createdAt** | **optional.String**| 任务创建时间，格式同上 | 
 **finishedAt** | **optional.String**| 任务完成时间，即任务最后一次转为已完成状态的时间点。格式同上 | 
 **milestone** | **optional.String**| 根据里程碑标题。none为没里程碑的，*为所有带里程碑的 | 
 **assignee** | **optional.String**| 用户的username。 none为没指派者, *为所有带有指派者的 | 
 **creator** | **optional.String**| 创建Issues的用户username | 
 **program** | **optional.String**| 所属项目名称。none为没有所属项目，*为所有带所属项目的 | 

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoIssuesComments**
> Note GetV5ReposOwnerRepoIssuesComments(ctx, owner, repo, optional)
获取仓库所有Issue的评论

获取仓库所有Issue的评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoIssuesCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoIssuesCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **sort** | **optional.String**| Either created or updated. Default: created | [default to created]
 **direction** | **optional.String**| Either asc or desc. Ignored without the sort parameter. | [default to asc]
 **since** | **optional.String**| Only comments updated at or after this time are returned.                                               This is a timestamp in ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoIssuesCommentsId**
> Note GetV5ReposOwnerRepoIssuesCommentsId(ctx, owner, repo, id, optional)
获取仓库Issue某条评论

获取仓库Issue某条评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 评论的ID | 
 **optional** | ***GetV5ReposOwnerRepoIssuesCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoIssuesCommentsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoIssuesNumber**
> Issue GetV5ReposOwnerRepoIssuesNumber(ctx, owner, repo, number, optional)
仓库的某个Issue

仓库的某个Issue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
 **optional** | ***GetV5ReposOwnerRepoIssuesNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoIssuesNumberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoIssuesNumberComments**
> []Note GetV5ReposOwnerRepoIssuesNumberComments(ctx, owner, repo, number, optional)
获取仓库某个Issue所有的评论

获取仓库某个Issue所有的评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
 **optional** | ***GetV5ReposOwnerRepoIssuesNumberCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoIssuesNumberCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 
 **since** | **optional.String**| Only comments updated at or after this time are returned.                                               This is a timestamp in ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserIssues**
> []Issue GetV5UserIssues(ctx, optional)
获取授权用户的所有Issues

获取授权用户的所有Issues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserIssuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserIssuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **filter** | **optional.String**| 筛选参数: 授权用户负责的(assigned)，授权用户创建的(created)，包含前两者的(all)。默认: assigned | [default to assigned]
 **state** | **optional.String**| Issue的状态: open（开启的）, progressing(进行中), closed（关闭的）, rejected（拒绝的）。 默认: open | [default to open]
 **labels** | **optional.String**| 用逗号分开的标签。如: bug,performance | 
 **sort** | **optional.String**| 排序依据: 创建时间(created)，更新时间(updated_at)。默认: created_at | [default to created]
 **direction** | **optional.String**| 排序方式: 升序(asc)，降序(desc)。默认: desc | [default to desc]
 **since** | **optional.String**| 起始的更新时间，要求时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **schedule** | **optional.String**| 计划开始日期，格式：20181006T173008+80-20181007T173008+80（区间），或者 -20181007T173008+80（小于20181007T173008+80），或者 20181006T173008+80-（大于20181006T173008+80），要求时间格式为20181006T173008+80 | 
 **deadline** | **optional.String**| 计划截止日期，格式同上 | 
 **createdAt** | **optional.String**| 任务创建时间，格式同上 | 
 **finishedAt** | **optional.String**| 任务完成时间，即任务最后一次转为已完成状态的时间点。格式同上 | 

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5ReposOwnerIssuesNumber**
> Issue PatchV5ReposOwnerIssuesNumber(ctx, owner, number, body)
更新Issue

更新Issue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
  **body** | [**IssueUpdateParam**](IssueUpdateParam.md)| 可选。Issue 内容 | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5ReposOwnerRepoIssuesCommentsId**
> Note PatchV5ReposOwnerRepoIssuesCommentsId(ctx, owner, repo, id, body)
更新Issue某条评论

更新Issue某条评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 评论的ID | 
  **body** | [**IssueCommentPatchParam**](IssueCommentPatchParam.md)| 必填。评论内容 | 

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerIssues**
> Issue PostV5ReposOwnerIssues(ctx, owner, body)
创建Issue

创建Issue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **body** | [**IssueCreateParam**](IssueCreateParam.md)| 可选。Issue 内容 | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoIssuesNumberComments**
> Note PostV5ReposOwnerRepoIssuesNumberComments(ctx, owner, repo, number, body)
创建某个Issue评论

创建某个Issue评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **string**| Issue 编号(区分大小写，无需添加 # 号) | 
  **body** | [**IssueCommentPostParam**](IssueCommentPostParam.md)| Issue comment内容 | 

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

