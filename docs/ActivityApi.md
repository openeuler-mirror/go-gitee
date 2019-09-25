# \ActivityApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5UserStarredOwnerRepo**](ActivityApi.md#DeleteV5UserStarredOwnerRepo) | **Delete** /v5/user/starred/{owner}/{repo} | 取消 star 一个仓库
[**DeleteV5UserSubscriptionsOwnerRepo**](ActivityApi.md#DeleteV5UserSubscriptionsOwnerRepo) | **Delete** /v5/user/subscriptions/{owner}/{repo} | 取消 watch 一个仓库
[**GetV5Events**](ActivityApi.md#GetV5Events) | **Get** /v5/events | 获取站内所有公开动态
[**GetV5NetworksOwnerRepoEvents**](ActivityApi.md#GetV5NetworksOwnerRepoEvents) | **Get** /v5/networks/{owner}/{repo}/events | 列出仓库的所有公开动态
[**GetV5NotificationsCount**](ActivityApi.md#GetV5NotificationsCount) | **Get** /v5/notifications/count | 获取授权用户的通知数
[**GetV5NotificationsMessages**](ActivityApi.md#GetV5NotificationsMessages) | **Get** /v5/notifications/messages | 列出授权用户的所有私信
[**GetV5NotificationsMessagesId**](ActivityApi.md#GetV5NotificationsMessagesId) | **Get** /v5/notifications/messages/{id} | 获取一条私信
[**GetV5NotificationsThreads**](ActivityApi.md#GetV5NotificationsThreads) | **Get** /v5/notifications/threads | 列出授权用户的所有通知
[**GetV5NotificationsThreadsId**](ActivityApi.md#GetV5NotificationsThreadsId) | **Get** /v5/notifications/threads/{id} | 获取一条通知
[**GetV5OrgsOrgEvents**](ActivityApi.md#GetV5OrgsOrgEvents) | **Get** /v5/orgs/{org}/events | 列出组织的公开动态
[**GetV5ReposOwnerRepoEvents**](ActivityApi.md#GetV5ReposOwnerRepoEvents) | **Get** /v5/repos/{owner}/{repo}/events | 列出仓库的所有动态
[**GetV5ReposOwnerRepoNotifications**](ActivityApi.md#GetV5ReposOwnerRepoNotifications) | **Get** /v5/repos/{owner}/{repo}/notifications | 列出一个仓库里的通知
[**GetV5ReposOwnerRepoStargazers**](ActivityApi.md#GetV5ReposOwnerRepoStargazers) | **Get** /v5/repos/{owner}/{repo}/stargazers | 列出 star 了仓库的用户
[**GetV5ReposOwnerRepoSubscribers**](ActivityApi.md#GetV5ReposOwnerRepoSubscribers) | **Get** /v5/repos/{owner}/{repo}/subscribers | 列出 watch 了仓库的用户
[**GetV5UserStarred**](ActivityApi.md#GetV5UserStarred) | **Get** /v5/user/starred | 列出授权用户 star 了的仓库
[**GetV5UserStarredOwnerRepo**](ActivityApi.md#GetV5UserStarredOwnerRepo) | **Get** /v5/user/starred/{owner}/{repo} | 检查授权用户是否 star 了一个仓库
[**GetV5UserSubscriptions**](ActivityApi.md#GetV5UserSubscriptions) | **Get** /v5/user/subscriptions | 列出授权用户 watch 了的仓库
[**GetV5UserSubscriptionsOwnerRepo**](ActivityApi.md#GetV5UserSubscriptionsOwnerRepo) | **Get** /v5/user/subscriptions/{owner}/{repo} | 检查授权用户是否 watch 了一个仓库
[**GetV5UsersUsernameEvents**](ActivityApi.md#GetV5UsersUsernameEvents) | **Get** /v5/users/{username}/events | 列出用户的动态
[**GetV5UsersUsernameEventsOrgsOrg**](ActivityApi.md#GetV5UsersUsernameEventsOrgsOrg) | **Get** /v5/users/{username}/events/orgs/{org} | 列出用户所属组织的动态
[**GetV5UsersUsernameEventsPublic**](ActivityApi.md#GetV5UsersUsernameEventsPublic) | **Get** /v5/users/{username}/events/public | 列出用户的公开动态
[**GetV5UsersUsernameReceivedEvents**](ActivityApi.md#GetV5UsersUsernameReceivedEvents) | **Get** /v5/users/{username}/received_events | 列出一个用户收到的动态
[**GetV5UsersUsernameReceivedEventsPublic**](ActivityApi.md#GetV5UsersUsernameReceivedEventsPublic) | **Get** /v5/users/{username}/received_events/public | 列出一个用户收到的公开动态
[**GetV5UsersUsernameStarred**](ActivityApi.md#GetV5UsersUsernameStarred) | **Get** /v5/users/{username}/starred | 列出用户 star 了的仓库
[**GetV5UsersUsernameSubscriptions**](ActivityApi.md#GetV5UsersUsernameSubscriptions) | **Get** /v5/users/{username}/subscriptions | 列出用户 watch 了的仓库
[**PatchV5NotificationsMessagesId**](ActivityApi.md#PatchV5NotificationsMessagesId) | **Patch** /v5/notifications/messages/{id} | 标记一条私信为已读
[**PatchV5NotificationsThreadsId**](ActivityApi.md#PatchV5NotificationsThreadsId) | **Patch** /v5/notifications/threads/{id} | 标记一条通知为已读
[**PostV5NotificationsMessages**](ActivityApi.md#PostV5NotificationsMessages) | **Post** /v5/notifications/messages | 发送私信给指定用户
[**PutV5NotificationsMessages**](ActivityApi.md#PutV5NotificationsMessages) | **Put** /v5/notifications/messages | 标记所有私信为已读
[**PutV5NotificationsThreads**](ActivityApi.md#PutV5NotificationsThreads) | **Put** /v5/notifications/threads | 标记所有通知为已读
[**PutV5ReposOwnerRepoNotifications**](ActivityApi.md#PutV5ReposOwnerRepoNotifications) | **Put** /v5/repos/{owner}/{repo}/notifications | 标记一个仓库里的通知为已读
[**PutV5UserStarredOwnerRepo**](ActivityApi.md#PutV5UserStarredOwnerRepo) | **Put** /v5/user/starred/{owner}/{repo} | star 一个仓库
[**PutV5UserSubscriptionsOwnerRepo**](ActivityApi.md#PutV5UserSubscriptionsOwnerRepo) | **Put** /v5/user/subscriptions/{owner}/{repo} | watch 一个仓库


# **DeleteV5UserStarredOwnerRepo**
> DeleteV5UserStarredOwnerRepo(ctx, owner, repo, optional)
取消 star 一个仓库

取消 star 一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***DeleteV5UserStarredOwnerRepoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5UserStarredOwnerRepoOpts struct

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

# **DeleteV5UserSubscriptionsOwnerRepo**
> DeleteV5UserSubscriptionsOwnerRepo(ctx, owner, repo, optional)
取消 watch 一个仓库

取消 watch 一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***DeleteV5UserSubscriptionsOwnerRepoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5UserSubscriptionsOwnerRepoOpts struct

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

# **GetV5Events**
> []Event GetV5Events(ctx, optional)
获取站内所有公开动态

获取站内所有公开动态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5EventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5NetworksOwnerRepoEvents**
> []Event GetV5NetworksOwnerRepoEvents(ctx, owner, repo, optional)
列出仓库的所有公开动态

列出仓库的所有公开动态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5NetworksOwnerRepoEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5NetworksOwnerRepoEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5NotificationsCount**
> UserNotificationCount GetV5NotificationsCount(ctx, optional)
获取授权用户的通知数

获取授权用户的通知数

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5NotificationsCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5NotificationsCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **unread** | **optional.Bool**| 是否只获取未读消息，默认：否 | 

### Return type

[**UserNotificationCount**](UserNotificationCount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5NotificationsMessages**
> []UserMessageList GetV5NotificationsMessages(ctx, optional)
列出授权用户的所有私信

列出授权用户的所有私信

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5NotificationsMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5NotificationsMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **unread** | **optional.Bool**| 是否只显示未读私信，默认：否 | 
 **since** | **optional.String**| 只获取在给定时间后更新的私信，要求时间格式为 ISO 8601 | 
 **before** | **optional.String**| 只获取在给定时间前更新的私信，要求时间格式为 ISO 8601 | 
 **ids** | **optional.String**| 指定一组私信 ID，以 , 分隔 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]UserMessageList**](UserMessageList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5NotificationsMessagesId**
> UserMessage GetV5NotificationsMessagesId(ctx, id, optional)
获取一条私信

获取一条私信

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 私信的 ID | 
 **optional** | ***GetV5NotificationsMessagesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5NotificationsMessagesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**UserMessage**](UserMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5NotificationsThreads**
> []UserNotificationList GetV5NotificationsThreads(ctx, optional)
列出授权用户的所有通知

列出授权用户的所有通知

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5NotificationsThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5NotificationsThreadsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **unread** | **optional.Bool**| 是否只获取未读消息，默认：否 | 
 **participating** | **optional.Bool**| 是否只获取自己直接参与的消息，默认：否 | 
 **type_** | **optional.String**| 筛选指定类型的通知，all：所有，event：事件通知，referer：@ 通知 | [default to all]
 **since** | **optional.String**| 只获取在给定时间后更新的消息，要求时间格式为 ISO 8601 | 
 **before** | **optional.String**| 只获取在给定时间前更新的消息，要求时间格式为 ISO 8601 | 
 **ids** | **optional.String**| 指定一组通知 ID，以 , 分隔 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]UserNotificationList**](UserNotificationList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5NotificationsThreadsId**
> UserNotification GetV5NotificationsThreadsId(ctx, id, optional)
获取一条通知

获取一条通知

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 通知的 ID | 
 **optional** | ***GetV5NotificationsThreadsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5NotificationsThreadsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**UserNotification**](UserNotification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5OrgsOrgEvents**
> []Event GetV5OrgsOrgEvents(ctx, org, optional)
列出组织的公开动态

列出组织的公开动态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***GetV5OrgsOrgEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5OrgsOrgEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoEvents**
> []Event GetV5ReposOwnerRepoEvents(ctx, owner, repo, optional)
列出仓库的所有动态

列出仓库的所有动态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoNotifications**
> []UserNotificationList GetV5ReposOwnerRepoNotifications(ctx, owner, repo, optional)
列出一个仓库里的通知

列出一个仓库里的通知

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoNotificationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **unread** | **optional.Bool**| 是否只获取未读消息，默认：否 | 
 **participating** | **optional.Bool**| 是否只获取自己直接参与的消息，默认：否 | 
 **type_** | **optional.String**| 筛选指定类型的通知，all：所有，event：事件通知，referer：@ 通知 | [default to all]
 **since** | **optional.String**| 只获取在给定时间后更新的消息，要求时间格式为 ISO 8601 | 
 **before** | **optional.String**| 只获取在给定时间前更新的消息，要求时间格式为 ISO 8601 | 
 **ids** | **optional.String**| 指定一组通知 ID，以 , 分隔 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]UserNotificationList**](UserNotificationList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoStargazers**
> []UserBasic GetV5ReposOwnerRepoStargazers(ctx, owner, repo, optional)
列出 star 了仓库的用户

列出 star 了仓库的用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoStargazersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoStargazersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]UserBasic**](UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoSubscribers**
> []UserBasic GetV5ReposOwnerRepoSubscribers(ctx, owner, repo, optional)
列出 watch 了仓库的用户

列出 watch 了仓库的用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoSubscribersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoSubscribersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]UserBasic**](UserBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserStarred**
> []Project GetV5UserStarred(ctx, optional)
列出授权用户 star 了的仓库

列出授权用户 star 了的仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserStarredOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserStarredOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **sort** | **optional.String**| 根据仓库创建时间(created)或最后推送时间(updated)进行排序，默认：创建时间 | [default to created]
 **direction** | **optional.String**| 按递增(asc)或递减(desc)排序，默认：递减 | [default to desc]
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserStarredOwnerRepo**
> GetV5UserStarredOwnerRepo(ctx, owner, repo, optional)
检查授权用户是否 star 了一个仓库

检查授权用户是否 star 了一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5UserStarredOwnerRepoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserStarredOwnerRepoOpts struct

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

# **GetV5UserSubscriptions**
> []Project GetV5UserSubscriptions(ctx, optional)
列出授权用户 watch 了的仓库

列出授权用户 watch 了的仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserSubscriptionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **sort** | **optional.String**| 根据仓库创建时间(created)或最后推送时间(updated)进行排序，默认：创建时间 | [default to created]
 **direction** | **optional.String**| 按递增(asc)或递减(desc)排序，默认：递减 | [default to desc]
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserSubscriptionsOwnerRepo**
> GetV5UserSubscriptionsOwnerRepo(ctx, owner, repo, optional)
检查授权用户是否 watch 了一个仓库

检查授权用户是否 watch 了一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5UserSubscriptionsOwnerRepoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserSubscriptionsOwnerRepoOpts struct

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

# **GetV5UsersUsernameEvents**
> []Event GetV5UsersUsernameEvents(ctx, username, optional)
列出用户的动态

列出用户的动态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameEventsOrgsOrg**
> []Event GetV5UsersUsernameEventsOrgsOrg(ctx, username, org, optional)
列出用户所属组织的动态

列出用户所属组织的动态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***GetV5UsersUsernameEventsOrgsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameEventsOrgsOrgOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameEventsPublic**
> []Event GetV5UsersUsernameEventsPublic(ctx, username, optional)
列出用户的公开动态

列出用户的公开动态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameEventsPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameEventsPublicOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameReceivedEvents**
> []Event GetV5UsersUsernameReceivedEvents(ctx, username, optional)
列出一个用户收到的动态

列出一个用户收到的动态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameReceivedEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameReceivedEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameReceivedEventsPublic**
> []Event GetV5UsersUsernameReceivedEventsPublic(ctx, username, optional)
列出一个用户收到的公开动态

列出一个用户收到的公开动态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameReceivedEventsPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameReceivedEventsPublicOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameStarred**
> []Project GetV5UsersUsernameStarred(ctx, username, optional)
列出用户 star 了的仓库

列出用户 star 了的仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameStarredOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameStarredOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **sort** | **optional.String**| 根据仓库创建时间(created)或最后推送时间(updated)进行排序，默认：创建时间 | [default to created]
 **direction** | **optional.String**| 按递增(asc)或递减(desc)排序，默认：递减 | [default to desc]

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameSubscriptions**
> []Project GetV5UsersUsernameSubscriptions(ctx, username, optional)
列出用户 watch 了的仓库

列出用户 watch 了的仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameSubscriptionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]
 **sort** | **optional.String**| 根据仓库创建时间(created)或最后推送时间(updated)进行排序，默认：创建时间 | [default to created]
 **direction** | **optional.String**| 按递增(asc)或递减(desc)排序，默认：递减 | [default to desc]

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5NotificationsMessagesId**
> PatchV5NotificationsMessagesId(ctx, id, optional)
标记一条私信为已读

标记一条私信为已读

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 私信的 ID | 
 **optional** | ***PatchV5NotificationsMessagesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5NotificationsMessagesIdOpts struct

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

# **PatchV5NotificationsThreadsId**
> PatchV5NotificationsThreadsId(ctx, id, optional)
标记一条通知为已读

标记一条通知为已读

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 通知的 ID | 
 **optional** | ***PatchV5NotificationsThreadsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5NotificationsThreadsIdOpts struct

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

# **PostV5NotificationsMessages**
> UserMessage PostV5NotificationsMessages(ctx, username, content, optional)
发送私信给指定用户

发送私信给指定用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
  **content** | **string**| 私信内容 | 
 **optional** | ***PostV5NotificationsMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5NotificationsMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**UserMessage**](UserMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5NotificationsMessages**
> PutV5NotificationsMessages(ctx, optional)
标记所有私信为已读

标记所有私信为已读

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PutV5NotificationsMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5NotificationsMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **ids** | **optional.String**| 指定一组私信 ID，以 , 分隔 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5NotificationsThreads**
> PutV5NotificationsThreads(ctx, optional)
标记所有通知为已读

标记所有通知为已读

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PutV5NotificationsThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5NotificationsThreadsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **ids** | **optional.String**| 指定一组通知 ID，以 , 分隔 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5ReposOwnerRepoNotifications**
> PutV5ReposOwnerRepoNotifications(ctx, owner, repo, optional)
标记一个仓库里的通知为已读

标记一个仓库里的通知为已读

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***PutV5ReposOwnerRepoNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5ReposOwnerRepoNotificationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **ids** | **optional.String**| 指定一组通知 ID，以 , 分隔 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5UserStarredOwnerRepo**
> PutV5UserStarredOwnerRepo(ctx, owner, repo, optional)
star 一个仓库

star 一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***PutV5UserStarredOwnerRepoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5UserStarredOwnerRepoOpts struct

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

# **PutV5UserSubscriptionsOwnerRepo**
> PutV5UserSubscriptionsOwnerRepo(ctx, owner, repo, watchType, optional)
watch 一个仓库

watch 一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **watchType** | **string**| watch策略, watching: 关注所有动态, releases_only: 仅关注版本发行动态, ignoring: 关注但不提醒动态 | [default to watching]
 **optional** | ***PutV5UserSubscriptionsOwnerRepoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5UserSubscriptionsOwnerRepoOpts struct

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

