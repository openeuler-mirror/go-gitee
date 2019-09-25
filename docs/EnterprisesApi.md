# \EnterprisesApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5EnterprisesEnterpriseMembersUsername**](EnterprisesApi.md#DeleteV5EnterprisesEnterpriseMembersUsername) | **Delete** /v5/enterprises/{enterprise}/members/{username} | 移除企业成员
[**DeleteV5EnterprisesEnterpriseWeekReportsReportIdCommentsId**](EnterprisesApi.md#DeleteV5EnterprisesEnterpriseWeekReportsReportIdCommentsId) | **Delete** /v5/enterprises/{enterprise}/week_reports/{report_id}/comments/{id} | 删除周报某个评论
[**GetV5EnterprisesEnterprise**](EnterprisesApi.md#GetV5EnterprisesEnterprise) | **Get** /v5/enterprises/{enterprise} | 获取一个企业
[**GetV5EnterprisesEnterpriseMembers**](EnterprisesApi.md#GetV5EnterprisesEnterpriseMembers) | **Get** /v5/enterprises/{enterprise}/members | 列出企业的所有成员
[**GetV5EnterprisesEnterpriseMembersUsername**](EnterprisesApi.md#GetV5EnterprisesEnterpriseMembersUsername) | **Get** /v5/enterprises/{enterprise}/members/{username} | 获取企业的一个成员
[**GetV5EnterprisesEnterpriseUsersUsernameWeekReports**](EnterprisesApi.md#GetV5EnterprisesEnterpriseUsersUsernameWeekReports) | **Get** /v5/enterprises/{enterprise}/users/{username}/week_reports | 个人周报列表
[**GetV5EnterprisesEnterpriseWeekReports**](EnterprisesApi.md#GetV5EnterprisesEnterpriseWeekReports) | **Get** /v5/enterprises/{enterprise}/week_reports | 企业成员周报列表
[**GetV5EnterprisesEnterpriseWeekReportsId**](EnterprisesApi.md#GetV5EnterprisesEnterpriseWeekReportsId) | **Get** /v5/enterprises/{enterprise}/week_reports/{id} | 周报详情
[**GetV5EnterprisesEnterpriseWeekReportsIdComments**](EnterprisesApi.md#GetV5EnterprisesEnterpriseWeekReportsIdComments) | **Get** /v5/enterprises/{enterprise}/week_reports/{id}/comments | 某个周报评论列表
[**GetV5UserEnterprises**](EnterprisesApi.md#GetV5UserEnterprises) | **Get** /v5/user/enterprises | 列出授权用户所属的企业
[**PatchV5EnterprisesEnterpriseWeekReportId**](EnterprisesApi.md#PatchV5EnterprisesEnterpriseWeekReportId) | **Patch** /v5/enterprises/{enterprise}/week_report/{id} | 编辑周报
[**PostV5EnterprisesEnterpriseMembers**](EnterprisesApi.md#PostV5EnterprisesEnterpriseMembers) | **Post** /v5/enterprises/{enterprise}/members | 添加或邀请企业成员
[**PostV5EnterprisesEnterpriseWeekReport**](EnterprisesApi.md#PostV5EnterprisesEnterpriseWeekReport) | **Post** /v5/enterprises/{enterprise}/week_report | 新建周报
[**PostV5EnterprisesEnterpriseWeekReportsIdComment**](EnterprisesApi.md#PostV5EnterprisesEnterpriseWeekReportsIdComment) | **Post** /v5/enterprises/{enterprise}/week_reports/{id}/comment | 评论周报
[**PutV5EnterprisesEnterpriseMembersUsername**](EnterprisesApi.md#PutV5EnterprisesEnterpriseMembersUsername) | **Put** /v5/enterprises/{enterprise}/members/{username} | 修改企业成员权限或备注


# **DeleteV5EnterprisesEnterpriseMembersUsername**
> DeleteV5EnterprisesEnterpriseMembersUsername(ctx, enterprise, username, optional)
移除企业成员

移除企业成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***DeleteV5EnterprisesEnterpriseMembersUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5EnterprisesEnterpriseMembersUsernameOpts struct

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

# **DeleteV5EnterprisesEnterpriseWeekReportsReportIdCommentsId**
> DeleteV5EnterprisesEnterpriseWeekReportsReportIdCommentsId(ctx, enterprise, reportId, id, optional)
删除周报某个评论

删除周报某个评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **reportId** | **int32**| 周报ID | 
  **id** | **int32**| 评论ID | 
 **optional** | ***DeleteV5EnterprisesEnterpriseWeekReportsReportIdCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5EnterprisesEnterpriseWeekReportsReportIdCommentsIdOpts struct

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

# **GetV5EnterprisesEnterprise**
> EnterpriseBasic GetV5EnterprisesEnterprise(ctx, enterprise, optional)
获取一个企业

获取一个企业

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
 **optional** | ***GetV5EnterprisesEnterpriseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**EnterpriseBasic**](EnterpriseBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseMembers**
> []EnterpriseMember GetV5EnterprisesEnterpriseMembers(ctx, enterprise, optional)
列出企业的所有成员

列出企业的所有成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
 **optional** | ***GetV5EnterprisesEnterpriseMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **role** | **optional.String**| 根据角色筛选成员 | [default to all]

### Return type

[**[]EnterpriseMember**](EnterpriseMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseMembersUsername**
> EnterpriseMember GetV5EnterprisesEnterpriseMembersUsername(ctx, enterprise, username, optional)
获取企业的一个成员

获取企业的一个成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5EnterprisesEnterpriseMembersUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseMembersUsernameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**EnterpriseMember**](EnterpriseMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseUsersUsernameWeekReports**
> []WeekReport GetV5EnterprisesEnterpriseUsersUsernameWeekReports(ctx, enterprise, username, optional)
个人周报列表

个人周报列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5EnterprisesEnterpriseUsersUsernameWeekReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseUsersUsernameWeekReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]WeekReport**](WeekReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseWeekReports**
> []WeekReport GetV5EnterprisesEnterpriseWeekReports(ctx, enterprise, optional)
企业成员周报列表

企业成员周报列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
 **optional** | ***GetV5EnterprisesEnterpriseWeekReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseWeekReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **username** | **optional.String**| 用户名(username/login) | 
 **year** | **optional.Int32**| 周报所属年 | 
 **weekIndex** | **optional.Int32**| 周报所属周 | 
 **date** | **optional.String**| 周报日期(格式：2019-03-25) | 

### Return type

[**[]WeekReport**](WeekReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseWeekReportsId**
> WeekReport GetV5EnterprisesEnterpriseWeekReportsId(ctx, enterprise, id, optional)
周报详情

周报详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **id** | **int32**| 周报ID | 
 **optional** | ***GetV5EnterprisesEnterpriseWeekReportsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseWeekReportsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**WeekReport**](WeekReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5EnterprisesEnterpriseWeekReportsIdComments**
> []Note GetV5EnterprisesEnterpriseWeekReportsIdComments(ctx, enterprise, id, optional)
某个周报评论列表

某个周报评论列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **id** | **int32**| 周报ID | 
 **optional** | ***GetV5EnterprisesEnterpriseWeekReportsIdCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseWeekReportsIdCommentsOpts struct

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

# **GetV5UserEnterprises**
> []EnterpriseBasic GetV5UserEnterprises(ctx, optional)
列出授权用户所属的企业

列出授权用户所属的企业

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserEnterprisesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserEnterprisesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **admin** | **optional.Bool**| 只列出授权用户管理的企业 | [default to true]

### Return type

[**[]EnterpriseBasic**](EnterpriseBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5EnterprisesEnterpriseWeekReportId**
> WeekReport PatchV5EnterprisesEnterpriseWeekReportId(ctx, enterprise, id, content, optional)
编辑周报

编辑周报

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **id** | **int32**| 周报ID | 
  **content** | **string**| 周报内容 | 
 **optional** | ***PatchV5EnterprisesEnterpriseWeekReportIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5EnterprisesEnterpriseWeekReportIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**WeekReport**](WeekReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5EnterprisesEnterpriseMembers**
> PostV5EnterprisesEnterpriseMembers(ctx, enterprise, optional)
添加或邀请企业成员

添加或邀请企业成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
 **optional** | ***PostV5EnterprisesEnterpriseMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5EnterprisesEnterpriseMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **username** | **optional.String**| 需要邀请的码云用户名(username/login)，username,email至少填写一个 | 
 **email** | **optional.String**| 要添加邮箱地址，若该邮箱未注册则自动创建帐号。username,email至少填写一个 | 
 **outsourced** | **optional.Bool**| 是否企业外包成员，默认：否 | 
 **role** | **optional.String**| 企业角色，默认普通成员 | [default to member]
 **name** | **optional.String**| 企业成员真实姓名（备注） | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5EnterprisesEnterpriseWeekReport**
> WeekReport PostV5EnterprisesEnterpriseWeekReport(ctx, enterprise, year, content, weekIndex, username, optional)
新建周报

新建周报

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **year** | **int32**| 周报所属年 | 
  **content** | **string**| 周报内容 | 
  **weekIndex** | **int32**| 周报所属周 | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***PostV5EnterprisesEnterpriseWeekReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5EnterprisesEnterpriseWeekReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **accessToken** | **optional.String**| 用户授权码 | 
 **date** | **optional.String**| 周报日期(格式：2019-03-25) | 

### Return type

[**WeekReport**](WeekReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5EnterprisesEnterpriseWeekReportsIdComment**
> Note PostV5EnterprisesEnterpriseWeekReportsIdComment(ctx, enterprise, id, body, optional)
评论周报

评论周报

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **id** | **int32**| 周报ID | 
  **body** | **string**| 评论的内容 | 
 **optional** | ***PostV5EnterprisesEnterpriseWeekReportsIdCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5EnterprisesEnterpriseWeekReportsIdCommentOpts struct

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

# **PutV5EnterprisesEnterpriseMembersUsername**
> EnterpriseMember PutV5EnterprisesEnterpriseMembersUsername(ctx, enterprise, username, optional)
修改企业成员权限或备注

修改企业成员权限或备注

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***PutV5EnterprisesEnterpriseMembersUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5EnterprisesEnterpriseMembersUsernameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **outsourced** | **optional.Bool**| 是否企业外包成员，默认：否 | 
 **role** | **optional.String**| 企业角色，默认普通成员 | [default to member]
 **active** | **optional.Bool**| 是否可访问企业资源，默认:是。（若选否则禁止该用户访问企业资源） | [default to true]
 **name** | **optional.String**| 企业成员真实姓名（备注） | 

### Return type

[**EnterpriseMember**](EnterpriseMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

