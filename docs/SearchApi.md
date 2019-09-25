# \SearchApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV5SearchGists**](SearchApi.md#GetV5SearchGists) | **Get** /v5/search/gists | 搜索代码片段
[**GetV5SearchIssues**](SearchApi.md#GetV5SearchIssues) | **Get** /v5/search/issues | 搜索 Issues
[**GetV5SearchRepositories**](SearchApi.md#GetV5SearchRepositories) | **Get** /v5/search/repositories | 搜索仓库
[**GetV5SearchUsers**](SearchApi.md#GetV5SearchUsers) | **Get** /v5/search/users | 搜索用户


# **GetV5SearchGists**
> []Code GetV5SearchGists(ctx, q, optional)
搜索代码片段

搜索代码片段

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| 搜索关键字 | 
 **optional** | ***GetV5SearchGistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5SearchGistsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **language** | **optional.String**| 筛选指定语言的代码片段 | 
 **owner** | **optional.String**| 筛选所属用户 (username/login) 的代码片段 | 
 **sort** | **optional.String**| 排序字段，created_at(创建时间)、updated_at(更新时间)、notes_count(评论数)、stars_count(收藏数)、forks_count(Fork 数)，默认为最佳匹配 | 
 **order** | **optional.String**| 排序顺序: desc(default)、asc | [default to desc]

### Return type

[**[]Code**](Code.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5SearchIssues**
> []Issue GetV5SearchIssues(ctx, q, optional)
搜索 Issues

搜索 Issues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| 搜索关键字 | 
 **optional** | ***GetV5SearchIssuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5SearchIssuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **repo** | **optional.String**| 筛选指定仓库 (path, e.g. oschina/git-osc) 的 issues | 
 **language** | **optional.String**| 筛选指定语言的 issues | 
 **label** | **optional.String**| 筛选指定标签的 issues | 
 **state** | **optional.String**| 筛选指定状态的 issues, open(开启)、closed(完成)、rejected(拒绝) | 
 **author** | **optional.String**| 筛选指定创建者 (username/login) 的 issues | 
 **assignee** | **optional.String**| 筛选指定负责人 (username/login) 的 issues | 
 **sort** | **optional.String**| 排序字段，created_at(创建时间)、last_push_at(更新时间)、notes_count(评论数)，默认为最佳匹配 | 
 **order** | **optional.String**| 排序顺序: desc(default)、asc | [default to desc]

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5SearchRepositories**
> []Project GetV5SearchRepositories(ctx, q, optional)
搜索仓库

搜索仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| 搜索关键字 | 
 **optional** | ***GetV5SearchRepositoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5SearchRepositoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **owner** | **optional.String**| 筛选指定空间地址(企业、组织或个人的地址 path) 的仓库 | 
 **fork** | **optional.Bool**| 是否搜索含 fork 的仓库，默认：否 | 
 **language** | **optional.String**| 筛选指定语言的仓库 | 
 **sort** | **optional.String**| 排序字段，created_at(创建时间)、last_push_at(更新时间)、stars_count(收藏数)、forks_count(Fork 数)、watches_count(关注数)，默认为最佳匹配 | 
 **order** | **optional.String**| 排序顺序: desc(default)、asc | [default to desc]

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5SearchUsers**
> []User GetV5SearchUsers(ctx, q, optional)
搜索用户

搜索用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| 搜索关键字 | 
 **optional** | ***GetV5SearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5SearchUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **sort** | **optional.String**| 排序字段，joined_at(注册时间)，默认为最佳匹配 | 
 **order** | **optional.String**| 排序顺序: desc(default)、asc | [default to desc]

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

