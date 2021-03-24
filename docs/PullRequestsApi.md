# \PullRequestsApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5ReposOwnerRepoPullsCommentsId**](PullRequestsApi.md#DeleteV5ReposOwnerRepoPullsCommentsId) | **Delete** /v5/repos/{owner}/{repo}/pulls/comments/{id} | 删除评论
[**DeleteV5ReposOwnerRepoPullsLabel**](PullRequestsApi.md#DeleteV5ReposOwnerRepoPullsLabel) | **Delete** /v5/repos/{owner}/{repo}/pulls/{number}/labels/{name} | 删除 Pull Request 标签
[**DeleteV5ReposOwnerRepoPullsNumberAssignees**](PullRequestsApi.md#DeleteV5ReposOwnerRepoPullsNumberAssignees) | **Delete** /v5/repos/{owner}/{repo}/pulls/{number}/assignees | 取消用户审查 Pull Request
[**DeleteV5ReposOwnerRepoPullsNumberTesters**](PullRequestsApi.md#DeleteV5ReposOwnerRepoPullsNumberTesters) | **Delete** /v5/repos/{owner}/{repo}/pulls/{number}/testers | 取消用户测试 Pull Request
[**GetV5ReposOwnerRepoPulls**](PullRequestsApi.md#GetV5ReposOwnerRepoPulls) | **Get** /v5/repos/{owner}/{repo}/pulls | 获取Pull Request列表
[**GetV5ReposOwnerRepoPullsComments**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsComments) | **Get** /v5/repos/{owner}/{repo}/pulls/comments | 获取该仓库下的所有Pull Request评论
[**GetV5ReposOwnerRepoPullsCommentsId**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsCommentsId) | **Get** /v5/repos/{owner}/{repo}/pulls/comments/{id} | 获取Pull Request的某个评论
[**GetV5ReposOwnerRepoPullsNumber**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsNumber) | **Get** /v5/repos/{owner}/{repo}/pulls/{number} | 获取单个Pull Request
[**GetV5ReposOwnerRepoPullsNumberComments**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsNumberComments) | **Get** /v5/repos/{owner}/{repo}/pulls/{number}/comments | 获取某个Pull Request的所有评论
[**GetV5ReposOwnerRepoPullsNumberCommits**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsNumberCommits) | **Get** /v5/repos/{owner}/{repo}/pulls/{number}/commits | 获取某Pull Request的所有Commit信息。最多显示250条Commit
[**GetV5ReposOwnerRepoPullsNumberFiles**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsNumberFiles) | **Get** /v5/repos/{owner}/{repo}/pulls/{number}/files | Pull Request Commit文件列表。最多显示300条diff
[**GetV5ReposOwnerRepoPullsNumberIssues**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsNumberIssues) | **Get** /v5/repos/{owner}/{repo}/pulls/{number}/issues | 获取 Pull Request 关联的 issues
[**GetV5ReposOwnerRepoPullsNumberLabels**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsNumberLabels) | **Get** /v5/repos/{owner}/{repo}/pulls/{number}/labels | 获取某个 Pull Request 的所有标签
[**GetV5ReposOwnerRepoPullsNumberMerge**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsNumberMerge) | **Get** /v5/repos/{owner}/{repo}/pulls/{number}/merge | 判断Pull Request是否已经合并
[**GetV5ReposOwnerRepoPullsNumberOperateLogs**](PullRequestsApi.md#GetV5ReposOwnerRepoPullsNumberOperateLogs) | **Get** /v5/repos/{owner}/{repo}/pulls/{number}/operate_logs | 获取某个Pull Request的操作日志
[**PatchV5ReposOwnerRepoPullsCommentsId**](PullRequestsApi.md#PatchV5ReposOwnerRepoPullsCommentsId) | **Patch** /v5/repos/{owner}/{repo}/pulls/comments/{id} | 编辑评论
[**PatchV5ReposOwnerRepoPullsNumber**](PullRequestsApi.md#PatchV5ReposOwnerRepoPullsNumber) | **Patch** /v5/repos/{owner}/{repo}/pulls/{number} | 更新Pull Request信息
[**PostV5ReposOwnerRepoPulls**](PullRequestsApi.md#PostV5ReposOwnerRepoPulls) | **Post** /v5/repos/{owner}/{repo}/pulls | 创建Pull Request
[**PostV5ReposOwnerRepoPullsNumberAssignees**](PullRequestsApi.md#PostV5ReposOwnerRepoPullsNumberAssignees) | **Post** /v5/repos/{owner}/{repo}/pulls/{number}/assignees | 指派用户审查 Pull Request
[**PostV5ReposOwnerRepoPullsNumberComments**](PullRequestsApi.md#PostV5ReposOwnerRepoPullsNumberComments) | **Post** /v5/repos/{owner}/{repo}/pulls/{number}/comments | 提交Pull Request评论
[**PostV5ReposOwnerRepoPullsNumberLabels**](PullRequestsApi.md#PostV5ReposOwnerRepoPullsNumberLabels) | **Post** /v5/repos/{owner}/{repo}/pulls/{number}/labels | 创建 Pull Request 标签
[**PostV5ReposOwnerRepoPullsNumberTesters**](PullRequestsApi.md#PostV5ReposOwnerRepoPullsNumberTesters) | **Post** /v5/repos/{owner}/{repo}/pulls/{number}/testers | 指派用户测试 Pull Request
[**PutV5ReposOwnerRepoPullsNumberLabels**](PullRequestsApi.md#PutV5ReposOwnerRepoPullsNumberLabels) | **Put** /v5/repos/{owner}/{repo}/pulls/{number}/labels | 替换Pull Request 所有标签
[**PutV5ReposOwnerRepoPullsNumberMerge**](PullRequestsApi.md#PutV5ReposOwnerRepoPullsNumberMerge) | **Put** /v5/repos/{owner}/{repo}/pulls/{number}/merge | 合并Pull Request


# **DeleteV5ReposOwnerRepoPullsCommentsId**
> DeleteV5ReposOwnerRepoPullsCommentsId(ctx, owner, repo, id, optional)
删除评论

删除评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 评论的ID | 
 **optional** | ***DeleteV5ReposOwnerRepoPullsCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoPullsCommentsIdOpts struct

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

# **DeleteV5ReposOwnerRepoPullsLabel**
> DeleteV5ReposOwnerRepoPullsLabel(ctx, owner, repo, number, name, optional)
删除 Pull Request 标签

删除 Pull Request 标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **name** | **string**| 标签名称 | 
 **optional** | ***DeleteV5ReposOwnerRepoPullsLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoPullsLabelOpts struct

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

# **DeleteV5ReposOwnerRepoPullsNumberAssignees**
> PullRequest DeleteV5ReposOwnerRepoPullsNumberAssignees(ctx, owner, repo, number, assignees, optional)
取消用户审查 Pull Request

取消用户审查 Pull Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **assignees** | **string**| 用户的个人空间地址, 以 , 分隔 | 
 **optional** | ***DeleteV5ReposOwnerRepoPullsNumberAssigneesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoPullsNumberAssigneesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteV5ReposOwnerRepoPullsNumberTesters**
> PullRequest DeleteV5ReposOwnerRepoPullsNumberTesters(ctx, owner, repo, number, testers, optional)
取消用户测试 Pull Request

取消用户测试 Pull Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **testers** | **string**| 用户的个人空间地址, 以 , 分隔 | 
 **optional** | ***DeleteV5ReposOwnerRepoPullsNumberTestersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoPullsNumberTestersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPulls**
> []PullRequest GetV5ReposOwnerRepoPulls(ctx, owner, repo, optional)
获取Pull Request列表

获取Pull Request列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoPullsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **state** | **optional.String**| 可选。Pull Request 状态 | [default to open]
 **head** | **optional.String**| 可选。Pull Request 提交的源分支。格式：branch 或者：username:branch | 
 **base** | **optional.String**| 可选。Pull Request 提交目标分支的名称。 | 
 **sort** | **optional.String**| 可选。排序字段，默认按创建时间 | [default to created]
 **direction** | **optional.String**| 可选。升序/降序 | [default to desc]
 **milestoneNumber** | **optional.Int32**| 可选。里程碑序号(id) | 
 **labels** | **optional.String**| 用逗号分开的标签。如: bug,performance | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPullsComments**
> []PullRequestComments GetV5ReposOwnerRepoPullsComments(ctx, owner, repo, optional)
获取该仓库下的所有Pull Request评论

获取该仓库下的所有Pull Request评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoPullsCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **sort** | **optional.String**| 可选。按创建时间/更新时间排序 | [default to created]
 **direction** | **optional.String**| 可选。升序/降序 | [default to desc]
 **since** | **optional.String**| 起始的更新时间，要求时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]PullRequestComments**](PullRequestComments.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPullsCommentsId**
> PullRequestComments GetV5ReposOwnerRepoPullsCommentsId(ctx, owner, repo, id, optional)
获取Pull Request的某个评论

获取Pull Request的某个评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**|  | 
 **optional** | ***GetV5ReposOwnerRepoPullsCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsCommentsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**PullRequestComments**](PullRequestComments.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPullsNumber**
> PullRequest GetV5ReposOwnerRepoPullsNumber(ctx, owner, repo, number, optional)
获取单个Pull Request

获取单个Pull Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
 **optional** | ***GetV5ReposOwnerRepoPullsNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsNumberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPullsNumberComments**
> []PullRequestComments GetV5ReposOwnerRepoPullsNumberComments(ctx, owner, repo, number, optional)
获取某个Pull Request的所有评论

获取某个Pull Request的所有评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
 **optional** | ***GetV5ReposOwnerRepoPullsNumberCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsNumberCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]PullRequestComments**](PullRequestComments.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPullsNumberCommits**
> []PullRequestCommits GetV5ReposOwnerRepoPullsNumberCommits(ctx, owner, repo, number, optional)
获取某Pull Request的所有Commit信息。最多显示250条Commit

获取某Pull Request的所有Commit信息。最多显示250条Commit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
 **optional** | ***GetV5ReposOwnerRepoPullsNumberCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsNumberCommitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**[]PullRequestCommits**](PullRequestCommits.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPullsNumberFiles**
> []PullRequestFiles GetV5ReposOwnerRepoPullsNumberFiles(ctx, owner, repo, number, optional)
Pull Request Commit文件列表。最多显示300条diff

Pull Request Commit文件列表。最多显示300条diff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
 **optional** | ***GetV5ReposOwnerRepoPullsNumberFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsNumberFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**[]PullRequestFiles**](PullRequestFiles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPullsNumberIssues**
> []Issue GetV5ReposOwnerRepoPullsNumberIssues(ctx, owner, repo, number, optional)
获取 Pull Request 关联的 issues

获取 Pull Request 关联的 issues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**|  | 
 **optional** | ***GetV5ReposOwnerRepoPullsNumberIssuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsNumberIssuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPullsNumberLabels**
> []Label GetV5ReposOwnerRepoPullsNumberLabels(ctx, owner, repo, number, optional)
获取某个 Pull Request 的所有标签

获取某个 Pull Request 的所有标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
 **optional** | ***GetV5ReposOwnerRepoPullsNumberLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsNumberLabelsOpts struct

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

# **GetV5ReposOwnerRepoPullsNumberMerge**
> GetV5ReposOwnerRepoPullsNumberMerge(ctx, owner, repo, number, optional)
判断Pull Request是否已经合并

判断Pull Request是否已经合并

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
 **optional** | ***GetV5ReposOwnerRepoPullsNumberMergeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsNumberMergeOpts struct

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

# **GetV5ReposOwnerRepoPullsNumberOperateLogs**
> []OperateLog GetV5ReposOwnerRepoPullsNumberOperateLogs(ctx, owner, repo, number, optional)
获取某个Pull Request的操作日志

获取某个Pull Request的操作日志

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
 **optional** | ***GetV5ReposOwnerRepoPullsNumberOperateLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPullsNumberOperateLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 
 **sort** | **optional.String**| 按递增(asc)或递减(desc)排序，默认：递减 | [default to desc]

### Return type

[**[]OperateLog**](OperateLog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5ReposOwnerRepoPullsCommentsId**
> PullRequestComments PatchV5ReposOwnerRepoPullsCommentsId(ctx, owner, repo, id, body)
编辑评论

编辑评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 评论的ID | 
  **body** | [**PullRequestCommentPatchParam**](PullRequestCommentPatchParam.md)| 必填。评论内容 | 

### Return type

[**PullRequestComments**](PullRequestComments.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5ReposOwnerRepoPullsNumber**
> PullRequest PatchV5ReposOwnerRepoPullsNumber(ctx, owner, repo, number, body)
更新Pull Request信息

更新Pull Request信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **body** | [**PullRequestUpdateParam**](PullRequestUpdateParam.md)| 可选。Pull Request 内容 | 

### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoPulls**
> PullRequest PostV5ReposOwnerRepoPulls(ctx, owner, repo, body)
创建Pull Request

创建Pull Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **body** | [**CreatePullRequestParam**](CreatePullRequestParam.md)| pr的信息 | 

### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoPullsNumberAssignees**
> PullRequest PostV5ReposOwnerRepoPullsNumberAssignees(ctx, owner, repo, number, body)
指派用户审查 Pull Request

指派用户审查 Pull Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **body** | [**PullRequestAssigneePostParam**](PullRequestAssigneePostParam.md)| 必选，标签的内容 | 

### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoPullsNumberComments**
> PullRequestComments PostV5ReposOwnerRepoPullsNumberComments(ctx, owner, repo, number, body)
提交Pull Request评论

提交Pull Request评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **body** | [**PullRequestCommentPostParam**](PullRequestCommentPostParam.md)| 评论内容 | 

### Return type

[**PullRequestComments**](PullRequestComments.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoPullsNumberLabels**
> Label PostV5ReposOwnerRepoPullsNumberLabels(ctx, owner, repo, number, body)
创建 Pull Request 标签

创建 Pull Request 标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **body** | [**PullRequestLabelPostParam**](PullRequestLabelPostParam.md)| 必选，标签的内容 | 

### Return type

[**Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoPullsNumberTesters**
> PullRequest PostV5ReposOwnerRepoPullsNumberTesters(ctx, owner, repo, number, testers, optional)
指派用户测试 Pull Request

指派用户测试 Pull Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **testers** | **string**| 用户的个人空间地址, 以 , 分隔 | 
 **optional** | ***PostV5ReposOwnerRepoPullsNumberTestersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5ReposOwnerRepoPullsNumberTestersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5ReposOwnerRepoPullsNumberLabels**
> []Label PutV5ReposOwnerRepoPullsNumberLabels(ctx, owner, repo, number, body)
替换Pull Request 所有标签

替换Pull Request 所有标签  需要在请求的body里填上数组，元素为标签的名字。如: [\"performance\", \"bug\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **body** | [**PullRequestLabelPostParam**](PullRequestLabelPostParam.md)| 必选，标签的内容 | 

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5ReposOwnerRepoPullsNumberMerge**
> PutV5ReposOwnerRepoPullsNumberMerge(ctx, owner, repo, number, body)
合并Pull Request

合并Pull Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **number** | **int32**| 第几个PR，即本仓库PR的序数 | 
  **body** | [**PullRequestMergePutParam**](PullRequestMergePutParam.md)| PullRequest合入参数 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

