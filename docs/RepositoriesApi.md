# \RepositoriesApi

All URIs are relative to *https://gitee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV5ReposOwnerRepo**](RepositoriesApi.md#DeleteV5ReposOwnerRepo) | **Delete** /v5/repos/{owner}/{repo} | 删除一个仓库
[**DeleteV5ReposOwnerRepoBranchesBranchProtection**](RepositoriesApi.md#DeleteV5ReposOwnerRepoBranchesBranchProtection) | **Delete** /v5/repos/{owner}/{repo}/branches/{branch}/protection | 取消保护分支的设置
[**DeleteV5ReposOwnerRepoCollaboratorsUsername**](RepositoriesApi.md#DeleteV5ReposOwnerRepoCollaboratorsUsername) | **Delete** /v5/repos/{owner}/{repo}/collaborators/{username} | 移除仓库成员
[**DeleteV5ReposOwnerRepoCommentsId**](RepositoriesApi.md#DeleteV5ReposOwnerRepoCommentsId) | **Delete** /v5/repos/{owner}/{repo}/comments/{id} | 删除Commit评论
[**DeleteV5ReposOwnerRepoContentsPath**](RepositoriesApi.md#DeleteV5ReposOwnerRepoContentsPath) | **Delete** /v5/repos/{owner}/{repo}/contents/{path} | 删除文件
[**DeleteV5ReposOwnerRepoKeysEnableId**](RepositoriesApi.md#DeleteV5ReposOwnerRepoKeysEnableId) | **Delete** /v5/repos/{owner}/{repo}/keys/enable/{id} | 停用仓库公钥
[**DeleteV5ReposOwnerRepoKeysId**](RepositoriesApi.md#DeleteV5ReposOwnerRepoKeysId) | **Delete** /v5/repos/{owner}/{repo}/keys/{id} | 删除一个仓库公钥
[**DeleteV5ReposOwnerRepoReleasesId**](RepositoriesApi.md#DeleteV5ReposOwnerRepoReleasesId) | **Delete** /v5/repos/{owner}/{repo}/releases/{id} | 删除仓库Release
[**GetV5EnterprisesEnterpriseRepos**](RepositoriesApi.md#GetV5EnterprisesEnterpriseRepos) | **Get** /v5/enterprises/{enterprise}/repos | 获取企业的所有仓库
[**GetV5OrgsOrgRepos**](RepositoriesApi.md#GetV5OrgsOrgRepos) | **Get** /v5/orgs/{org}/repos | 获取一个组织的仓库
[**GetV5ReposOwnerRepo**](RepositoriesApi.md#GetV5ReposOwnerRepo) | **Get** /v5/repos/{owner}/{repo} | 获取用户的某个仓库
[**GetV5ReposOwnerRepoBranches**](RepositoriesApi.md#GetV5ReposOwnerRepoBranches) | **Get** /v5/repos/{owner}/{repo}/branches | 获取所有分支
[**GetV5ReposOwnerRepoBranchesBranch**](RepositoriesApi.md#GetV5ReposOwnerRepoBranchesBranch) | **Get** /v5/repos/{owner}/{repo}/branches/{branch} | 获取单个分支
[**GetV5ReposOwnerRepoCollaborators**](RepositoriesApi.md#GetV5ReposOwnerRepoCollaborators) | **Get** /v5/repos/{owner}/{repo}/collaborators | 获取仓库的所有成员
[**GetV5ReposOwnerRepoCollaboratorsUsername**](RepositoriesApi.md#GetV5ReposOwnerRepoCollaboratorsUsername) | **Get** /v5/repos/{owner}/{repo}/collaborators/{username} | 判断用户是否为仓库成员
[**GetV5ReposOwnerRepoCollaboratorsUsernamePermission**](RepositoriesApi.md#GetV5ReposOwnerRepoCollaboratorsUsernamePermission) | **Get** /v5/repos/{owner}/{repo}/collaborators/{username}/permission | 查看仓库成员的权限
[**GetV5ReposOwnerRepoComments**](RepositoriesApi.md#GetV5ReposOwnerRepoComments) | **Get** /v5/repos/{owner}/{repo}/comments | 获取仓库的Commit评论
[**GetV5ReposOwnerRepoCommentsId**](RepositoriesApi.md#GetV5ReposOwnerRepoCommentsId) | **Get** /v5/repos/{owner}/{repo}/comments/{id} | 获取仓库的某条Commit评论
[**GetV5ReposOwnerRepoCommits**](RepositoriesApi.md#GetV5ReposOwnerRepoCommits) | **Get** /v5/repos/{owner}/{repo}/commits | 仓库的所有提交
[**GetV5ReposOwnerRepoCommitsRefComments**](RepositoriesApi.md#GetV5ReposOwnerRepoCommitsRefComments) | **Get** /v5/repos/{owner}/{repo}/commits/{ref}/comments | 获取单个Commit的评论
[**GetV5ReposOwnerRepoCommitsSha**](RepositoriesApi.md#GetV5ReposOwnerRepoCommitsSha) | **Get** /v5/repos/{owner}/{repo}/commits/{sha} | 仓库的某个提交
[**GetV5ReposOwnerRepoCompareBaseHead**](RepositoriesApi.md#GetV5ReposOwnerRepoCompareBaseHead) | **Get** /v5/repos/{owner}/{repo}/compare/{base}...{head} | 两个Commits之间对比的版本差异
[**GetV5ReposOwnerRepoContentsPath**](RepositoriesApi.md#GetV5ReposOwnerRepoContentsPath) | **Get** /v5/repos/{owner}/{repo}/contents/{path} | 获取仓库具体路径下的内容
[**GetV5ReposOwnerRepoContributors**](RepositoriesApi.md#GetV5ReposOwnerRepoContributors) | **Get** /v5/repos/{owner}/{repo}/contributors | 获取仓库贡献者
[**GetV5ReposOwnerRepoForks**](RepositoriesApi.md#GetV5ReposOwnerRepoForks) | **Get** /v5/repos/{owner}/{repo}/forks | 查看仓库的Forks
[**GetV5ReposOwnerRepoKeys**](RepositoriesApi.md#GetV5ReposOwnerRepoKeys) | **Get** /v5/repos/{owner}/{repo}/keys | 获取仓库已部署的公钥
[**GetV5ReposOwnerRepoKeysAvailable**](RepositoriesApi.md#GetV5ReposOwnerRepoKeysAvailable) | **Get** /v5/repos/{owner}/{repo}/keys/available | 获取仓库可部署的公钥
[**GetV5ReposOwnerRepoKeysId**](RepositoriesApi.md#GetV5ReposOwnerRepoKeysId) | **Get** /v5/repos/{owner}/{repo}/keys/{id} | 获取仓库的单个公钥
[**GetV5ReposOwnerRepoPages**](RepositoriesApi.md#GetV5ReposOwnerRepoPages) | **Get** /v5/repos/{owner}/{repo}/pages | 获取Pages信息
[**GetV5ReposOwnerRepoReadme**](RepositoriesApi.md#GetV5ReposOwnerRepoReadme) | **Get** /v5/repos/{owner}/{repo}/readme | 获取仓库README
[**GetV5ReposOwnerRepoReleases**](RepositoriesApi.md#GetV5ReposOwnerRepoReleases) | **Get** /v5/repos/{owner}/{repo}/releases | 获取仓库的所有Releases
[**GetV5ReposOwnerRepoReleasesId**](RepositoriesApi.md#GetV5ReposOwnerRepoReleasesId) | **Get** /v5/repos/{owner}/{repo}/releases/{id} | 获取仓库的单个Releases
[**GetV5ReposOwnerRepoReleasesLatest**](RepositoriesApi.md#GetV5ReposOwnerRepoReleasesLatest) | **Get** /v5/repos/{owner}/{repo}/releases/latest | 获取仓库的最后更新的Release
[**GetV5ReposOwnerRepoReleasesTagsTag**](RepositoriesApi.md#GetV5ReposOwnerRepoReleasesTagsTag) | **Get** /v5/repos/{owner}/{repo}/releases/tags/{tag} | 根据Tag名称获取仓库的Release
[**GetV5ReposOwnerRepoTags**](RepositoriesApi.md#GetV5ReposOwnerRepoTags) | **Get** /v5/repos/{owner}/{repo}/tags | 列出仓库所有的tags
[**GetV5UserRepos**](RepositoriesApi.md#GetV5UserRepos) | **Get** /v5/user/repos | 列出授权用户的所有仓库
[**GetV5UsersUsernameRepos**](RepositoriesApi.md#GetV5UsersUsernameRepos) | **Get** /v5/users/{username}/repos | 获取某个用户的公开仓库
[**PatchV5ReposOwnerRepo**](RepositoriesApi.md#PatchV5ReposOwnerRepo) | **Patch** /v5/repos/{owner}/{repo} | 更新仓库设置
[**PatchV5ReposOwnerRepoCommentsId**](RepositoriesApi.md#PatchV5ReposOwnerRepoCommentsId) | **Patch** /v5/repos/{owner}/{repo}/comments/{id} | 更新Commit评论
[**PatchV5ReposOwnerRepoReleasesId**](RepositoriesApi.md#PatchV5ReposOwnerRepoReleasesId) | **Patch** /v5/repos/{owner}/{repo}/releases/{id} | 更新仓库Release
[**PostV5EnterprisesEnterpriseRepos**](RepositoriesApi.md#PostV5EnterprisesEnterpriseRepos) | **Post** /v5/enterprises/{enterprise}/repos | 创建企业仓库
[**PostV5OrgsOrgRepos**](RepositoriesApi.md#PostV5OrgsOrgRepos) | **Post** /v5/orgs/{org}/repos | 创建组织仓库
[**PostV5ReposOwnerRepoBranches**](RepositoriesApi.md#PostV5ReposOwnerRepoBranches) | **Post** /v5/repos/{owner}/{repo}/branches | 创建分支
[**PostV5ReposOwnerRepoCommitsShaComments**](RepositoriesApi.md#PostV5ReposOwnerRepoCommitsShaComments) | **Post** /v5/repos/{owner}/{repo}/commits/{sha}/comments | 创建Commit评论
[**PostV5ReposOwnerRepoContentsPath**](RepositoriesApi.md#PostV5ReposOwnerRepoContentsPath) | **Post** /v5/repos/{owner}/{repo}/contents/{path} | 新建文件
[**PostV5ReposOwnerRepoForks**](RepositoriesApi.md#PostV5ReposOwnerRepoForks) | **Post** /v5/repos/{owner}/{repo}/forks | Fork一个仓库
[**PostV5ReposOwnerRepoKeys**](RepositoriesApi.md#PostV5ReposOwnerRepoKeys) | **Post** /v5/repos/{owner}/{repo}/keys | 为仓库添加公钥
[**PostV5ReposOwnerRepoPagesBuilds**](RepositoriesApi.md#PostV5ReposOwnerRepoPagesBuilds) | **Post** /v5/repos/{owner}/{repo}/pages/builds | 请求建立Pages
[**PostV5ReposOwnerRepoReleases**](RepositoriesApi.md#PostV5ReposOwnerRepoReleases) | **Post** /v5/repos/{owner}/{repo}/releases | 创建仓库Release
[**PostV5UserRepos**](RepositoriesApi.md#PostV5UserRepos) | **Post** /v5/user/repos | 创建一个仓库
[**PutV5ReposOwnerRepoBranchesBranchProtection**](RepositoriesApi.md#PutV5ReposOwnerRepoBranchesBranchProtection) | **Put** /v5/repos/{owner}/{repo}/branches/{branch}/protection | 设置分支保护
[**PutV5ReposOwnerRepoClear**](RepositoriesApi.md#PutV5ReposOwnerRepoClear) | **Put** /v5/repos/{owner}/{repo}/clear | 清空一个仓库
[**PutV5ReposOwnerRepoCollaboratorsUsername**](RepositoriesApi.md#PutV5ReposOwnerRepoCollaboratorsUsername) | **Put** /v5/repos/{owner}/{repo}/collaborators/{username} | 添加仓库成员
[**PutV5ReposOwnerRepoContentsPath**](RepositoriesApi.md#PutV5ReposOwnerRepoContentsPath) | **Put** /v5/repos/{owner}/{repo}/contents/{path} | 更新文件
[**PutV5ReposOwnerRepoKeysEnableId**](RepositoriesApi.md#PutV5ReposOwnerRepoKeysEnableId) | **Put** /v5/repos/{owner}/{repo}/keys/enable/{id} | 启用仓库公钥
[**PutV5ReposOwnerRepoReviewer**](RepositoriesApi.md#PutV5ReposOwnerRepoReviewer) | **Put** /v5/repos/{owner}/{repo}/reviewer | 修改代码审查设置


# **DeleteV5ReposOwnerRepo**
> DeleteV5ReposOwnerRepo(ctx, owner, repo, optional)
删除一个仓库

删除一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***DeleteV5ReposOwnerRepoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoOpts struct

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

# **DeleteV5ReposOwnerRepoBranchesBranchProtection**
> DeleteV5ReposOwnerRepoBranchesBranchProtection(ctx, owner, repo, branch, optional)
取消保护分支的设置

取消保护分支的设置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **branch** | **string**| 分支名称 | 
 **optional** | ***DeleteV5ReposOwnerRepoBranchesBranchProtectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoBranchesBranchProtectionOpts struct

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

# **DeleteV5ReposOwnerRepoCollaboratorsUsername**
> DeleteV5ReposOwnerRepoCollaboratorsUsername(ctx, owner, repo, username, optional)
移除仓库成员

移除仓库成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***DeleteV5ReposOwnerRepoCollaboratorsUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoCollaboratorsUsernameOpts struct

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

# **DeleteV5ReposOwnerRepoCommentsId**
> DeleteV5ReposOwnerRepoCommentsId(ctx, owner, repo, id, optional)
删除Commit评论

删除Commit评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 评论的ID | 
 **optional** | ***DeleteV5ReposOwnerRepoCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoCommentsIdOpts struct

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

# **DeleteV5ReposOwnerRepoContentsPath**
> CommitContent DeleteV5ReposOwnerRepoContentsPath(ctx, owner, repo, path, sha, message, optional)
删除文件

删除文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **path** | **string**| 文件的路径 | 
  **sha** | **string**| 文件的 Blob SHA，可通过 [获取仓库具体路径下的内容] API 获取 | 
  **message** | **string**| 提交信息 | 
 **optional** | ***DeleteV5ReposOwnerRepoContentsPathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoContentsPathOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **accessToken** | **optional.String**| 用户授权码 | 
 **branch** | **optional.String**| 分支名称。默认为仓库对默认分支 | 
 **committerName** | **optional.String**| Committer的名字，默认为当前用户的名字 | 
 **committerEmail** | **optional.String**| Committer的邮箱，默认为当前用户的邮箱 | 
 **authorName** | **optional.String**| Author的名字，默认为当前用户的名字 | 
 **authorEmail** | **optional.String**| Author的邮箱，默认为当前用户的邮箱 | 

### Return type

[**CommitContent**](CommitContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteV5ReposOwnerRepoKeysEnableId**
> DeleteV5ReposOwnerRepoKeysEnableId(ctx, owner, repo, id, optional)
停用仓库公钥

停用仓库公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 公钥 ID | 
 **optional** | ***DeleteV5ReposOwnerRepoKeysEnableIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoKeysEnableIdOpts struct

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

# **DeleteV5ReposOwnerRepoKeysId**
> DeleteV5ReposOwnerRepoKeysId(ctx, owner, repo, id, optional)
删除一个仓库公钥

删除一个仓库公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 公钥 ID | 
 **optional** | ***DeleteV5ReposOwnerRepoKeysIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoKeysIdOpts struct

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

# **DeleteV5ReposOwnerRepoReleasesId**
> DeleteV5ReposOwnerRepoReleasesId(ctx, owner, repo, id, optional)
删除仓库Release

删除仓库Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**|  | 
 **optional** | ***DeleteV5ReposOwnerRepoReleasesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteV5ReposOwnerRepoReleasesIdOpts struct

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

# **GetV5EnterprisesEnterpriseRepos**
> Project GetV5EnterprisesEnterpriseRepos(ctx, enterprise, optional)
获取企业的所有仓库

获取企业的所有仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enterprise** | **string**| 企业的路径(path/login) | 
 **optional** | ***GetV5EnterprisesEnterpriseReposOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5EnterprisesEnterpriseReposOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **type_** | **optional.String**| 筛选仓库的类型，可以是 all, public, internal, private。默认: all | [default to all]
 **direct** | **optional.Bool**| 只获取直属仓库，默认: false | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5OrgsOrgRepos**
> []Project GetV5OrgsOrgRepos(ctx, org, optional)
获取一个组织的仓库

获取一个组织的仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
 **optional** | ***GetV5OrgsOrgReposOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5OrgsOrgReposOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **type_** | **optional.String**| 筛选仓库的类型，可以是 all, public, private。默认: all | [default to all]
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

# **GetV5ReposOwnerRepo**
> Project GetV5ReposOwnerRepo(ctx, owner, repo, optional)
获取用户的某个仓库

获取用户的某个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoBranches**
> []Branch GetV5ReposOwnerRepoBranches(ctx, owner, repo, optional)
获取所有分支

获取所有分支

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoBranchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoBranchesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**[]Branch**](Branch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoBranchesBranch**
> Branch GetV5ReposOwnerRepoBranchesBranch(ctx, owner, repo, branch, optional)
获取单个分支

获取单个分支

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **branch** | **string**| 分支名称 | 
 **optional** | ***GetV5ReposOwnerRepoBranchesBranchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoBranchesBranchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Branch**](Branch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoCollaborators**
> []ProjectMember GetV5ReposOwnerRepoCollaborators(ctx, owner, repo, optional)
获取仓库的所有成员

获取仓库的所有成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoCollaboratorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoCollaboratorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]ProjectMember**](ProjectMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoCollaboratorsUsername**
> GetV5ReposOwnerRepoCollaboratorsUsername(ctx, owner, repo, username, optional)
判断用户是否为仓库成员

判断用户是否为仓库成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5ReposOwnerRepoCollaboratorsUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoCollaboratorsUsernameOpts struct

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

# **GetV5ReposOwnerRepoCollaboratorsUsernamePermission**
> ProjectMemberPermission GetV5ReposOwnerRepoCollaboratorsUsernamePermission(ctx, owner, repo, username, optional)
查看仓库成员的权限

查看仓库成员的权限

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5ReposOwnerRepoCollaboratorsUsernamePermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoCollaboratorsUsernamePermissionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**ProjectMemberPermission**](ProjectMemberPermission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoComments**
> Note GetV5ReposOwnerRepoComments(ctx, owner, repo, optional)
获取仓库的Commit评论

获取仓库的Commit评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
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

# **GetV5ReposOwnerRepoCommentsId**
> Note GetV5ReposOwnerRepoCommentsId(ctx, owner, repo, id, optional)
获取仓库的某条Commit评论

获取仓库的某条Commit评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 评论的ID | 
 **optional** | ***GetV5ReposOwnerRepoCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoCommentsIdOpts struct

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

# **GetV5ReposOwnerRepoCommits**
> []RepoCommit GetV5ReposOwnerRepoCommits(ctx, owner, repo, optional)
仓库的所有提交

仓库的所有提交

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoCommitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **sha** | **optional.String**| 提交起始的SHA值或者分支名. 默认: 仓库的默认分支 | 
 **path** | **optional.String**| 包含该文件的提交 | 
 **author** | **optional.String**| 提交作者的邮箱或个人空间地址(username/login) | 
 **since** | **optional.String**| 提交的起始时间，时间格式为 ISO 8601 | 
 **until** | **optional.String**| 提交的最后时间，时间格式为 ISO 8601 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]RepoCommit**](RepoCommit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoCommitsRefComments**
> Note GetV5ReposOwnerRepoCommitsRefComments(ctx, owner, repo, ref, optional)
获取单个Commit的评论

获取单个Commit的评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **ref** | **string**| Commit的Reference | 
 **optional** | ***GetV5ReposOwnerRepoCommitsRefCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoCommitsRefCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 
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

# **GetV5ReposOwnerRepoCommitsSha**
> RepoCommit GetV5ReposOwnerRepoCommitsSha(ctx, owner, repo, sha, optional)
仓库的某个提交

仓库的某个提交

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **sha** | **string**| 提交的SHA值或者分支名 | 
 **optional** | ***GetV5ReposOwnerRepoCommitsShaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoCommitsShaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**RepoCommit**](RepoCommit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoCompareBaseHead**
> Compare GetV5ReposOwnerRepoCompareBaseHead(ctx, owner, repo, base, head, optional)
两个Commits之间对比的版本差异

两个Commits之间对比的版本差异

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **base** | **string**| Commit提交的SHA值或者分支名作为对比起点 | 
  **head** | **string**| Commit提交的SHA值或者分支名作为对比终点 | 
 **optional** | ***GetV5ReposOwnerRepoCompareBaseHeadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoCompareBaseHeadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Compare**](Compare.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoContentsPath**
> Content GetV5ReposOwnerRepoContentsPath(ctx, owner, repo, path, optional)
获取仓库具体路径下的内容

获取仓库具体路径下的内容

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **path** | **string**| 文件的路径 | 
 **optional** | ***GetV5ReposOwnerRepoContentsPathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoContentsPathOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 
 **ref** | **optional.String**| 分支、tag或commit。默认: 仓库的默认分支(通常是master) | 

### Return type

[**Content**](Content.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoContributors**
> Contributor GetV5ReposOwnerRepoContributors(ctx, owner, repo, optional)
获取仓库贡献者

获取仓库贡献者

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoContributorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoContributorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Contributor**](Contributor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoForks**
> Project GetV5ReposOwnerRepoForks(ctx, owner, repo, optional)
查看仓库的Forks

查看仓库的Forks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoForksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoForksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **sort** | **optional.String**| 排序方式: fork的时间(newest, oldest)，star的人数(stargazers) | [default to newest]
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoKeys**
> []SshKey GetV5ReposOwnerRepoKeys(ctx, owner, repo, optional)
获取仓库已部署的公钥

获取仓库已部署的公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]SshKey**](SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoKeysAvailable**
> []SshKeyBasic GetV5ReposOwnerRepoKeysAvailable(ctx, owner, repo, optional)
获取仓库可部署的公钥

获取仓库可部署的公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoKeysAvailableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoKeysAvailableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]SshKeyBasic**](SSHKeyBasic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoKeysId**
> SshKey GetV5ReposOwnerRepoKeysId(ctx, owner, repo, id, optional)
获取仓库的单个公钥

获取仓库的单个公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 公钥 ID | 
 **optional** | ***GetV5ReposOwnerRepoKeysIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoKeysIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoPages**
> GetV5ReposOwnerRepoPages(ctx, owner, repo, optional)
获取Pages信息

获取Pages信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoPagesOpts struct

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

# **GetV5ReposOwnerRepoReadme**
> Content GetV5ReposOwnerRepoReadme(ctx, owner, repo, optional)
获取仓库README

获取仓库README

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoReadmeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoReadmeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **ref** | **optional.String**| 分支、tag或commit。默认: 仓库的默认分支(通常是master) | 

### Return type

[**Content**](Content.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoReleases**
> []Release GetV5ReposOwnerRepoReleases(ctx, owner, repo, optional)
获取仓库的所有Releases

获取仓库的所有Releases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoReleasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**[]Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoReleasesId**
> Release GetV5ReposOwnerRepoReleasesId(ctx, owner, repo, id, optional)
获取仓库的单个Releases

获取仓库的单个Releases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 发行版本的ID | 
 **optional** | ***GetV5ReposOwnerRepoReleasesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoReleasesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoReleasesLatest**
> Release GetV5ReposOwnerRepoReleasesLatest(ctx, owner, repo, optional)
获取仓库的最后更新的Release

获取仓库的最后更新的Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoReleasesLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoReleasesLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoReleasesTagsTag**
> Release GetV5ReposOwnerRepoReleasesTagsTag(ctx, owner, repo, tag, optional)
根据Tag名称获取仓库的Release

根据Tag名称获取仓库的Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **tag** | **string**| Tag 名称 | 
 **optional** | ***GetV5ReposOwnerRepoReleasesTagsTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoReleasesTagsTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5ReposOwnerRepoTags**
> Tag GetV5ReposOwnerRepoTags(ctx, owner, repo, optional)
列出仓库所有的tags

列出仓库所有的tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***GetV5ReposOwnerRepoTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5ReposOwnerRepoTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UserRepos**
> Project GetV5UserRepos(ctx, optional)
列出授权用户的所有仓库

列出授权用户的所有仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetV5UserReposOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UserReposOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**| 用户授权码 | 
 **visibility** | **optional.String**| 公开(public)、私有(private)或者所有(all)，默认: 所有(all) | 
 **affiliation** | **optional.String**| owner(授权用户拥有的仓库)、collaborator(授权用户为仓库成员)、organization_member(授权用户为仓库所在组织并有访问仓库权限)、enterprise_member(授权用户所在企业并有访问仓库权限)、admin(所有有权限的，包括所管理的组织中所有仓库、所管理的企业的所有仓库)。                    可以用逗号分隔符组合。如: owner, organization_member 或 owner, collaborator, organization_member | 
 **type_** | **optional.String**| 筛选用户仓库: 其创建(owner)、个人(personal)、其为成员(member)、公开(public)、私有(private)，不能与 visibility 或 affiliation 参数一并使用，否则会报 422 错误 | 
 **sort** | **optional.String**| 排序方式: 创建时间(created)，更新时间(updated)，最后推送时间(pushed)，仓库所属与名称(full_name)。默认: full_name | [default to full_name]
 **direction** | **optional.String**| 如果sort参数为full_name，用升序(asc)。否则降序(desc) | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV5UsersUsernameRepos**
> Project GetV5UsersUsernameRepos(ctx, username, optional)
获取某个用户的公开仓库

获取某个用户的公开仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名(username/login) | 
 **optional** | ***GetV5UsersUsernameReposOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetV5UsersUsernameReposOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **type_** | **optional.String**| 用户创建的仓库(owner)，用户个人仓库(personal)，用户为仓库成员(member)，所有(all)。默认: 所有(all) | [default to all]
 **sort** | **optional.String**| 排序方式: 创建时间(created)，更新时间(updated)，最后推送时间(pushed)，仓库所属与名称(full_name)。默认: full_name | [default to full_name]
 **direction** | **optional.String**| 如果sort参数为full_name，用升序(asc)。否则降序(desc) | 
 **page** | **optional.Int32**| 当前的页码 | [default to 1]
 **perPage** | **optional.Int32**| 每页的数量，最大为 100 | [default to 20]

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5ReposOwnerRepo**
> Project PatchV5ReposOwnerRepo(ctx, owner, repo, body)
更新仓库设置

更新仓库设置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **body** | [**RepoPatchParam**](RepoPatchParam.md)| repo patch param | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchV5ReposOwnerRepoCommentsId**
> Note PatchV5ReposOwnerRepoCommentsId(ctx, owner, repo, id, body, optional)
更新Commit评论

更新Commit评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 评论的ID | 
  **body** | **string**| 评论的内容 | 
 **optional** | ***PatchV5ReposOwnerRepoCommentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5ReposOwnerRepoCommentsIdOpts struct

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

# **PatchV5ReposOwnerRepoReleasesId**
> Release PatchV5ReposOwnerRepoReleasesId(ctx, owner, repo, tagName, name, body, id, optional)
更新仓库Release

更新仓库Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **tagName** | **string**| Tag 名称, 提倡以v字母为前缀做为Release名称，例如v1.0或者v2.3.4 | 
  **name** | **string**| Release 名称 | 
  **body** | **string**| Release 描述 | 
  **id** | **int32**|  | 
 **optional** | ***PatchV5ReposOwnerRepoReleasesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchV5ReposOwnerRepoReleasesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **accessToken** | **optional.String**| 用户授权码 | 
 **prerelease** | **optional.Bool**| 是否为预览版本。默认: false（非预览版本） | 

### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5EnterprisesEnterpriseRepos**
> Project PostV5EnterprisesEnterpriseRepos(ctx, name, enterprise, optional)
创建企业仓库

创建企业仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 仓库名称 | 
  **enterprise** | **string**| 企业的路径(path/login) | 
 **optional** | ***PostV5EnterprisesEnterpriseReposOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5EnterprisesEnterpriseReposOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **description** | **optional.String**| 仓库描述 | 
 **homepage** | **optional.String**| 主页(eg: https://gitee.com) | 
 **hasIssues** | **optional.Bool**| 允许提Issue与否。默认: 允许(true) | [default to true]
 **hasWiki** | **optional.Bool**| 提供Wiki与否。默认: 提供(true) | [default to true]
 **canComment** | **optional.Bool**| 允许用户对仓库进行评论。默认： 允许(true) | [default to true]
 **autoInit** | **optional.Bool**| 值为true时则会用README初始化仓库。默认: 不初始化(false) | 
 **gitignoreTemplate** | **optional.String**| Git Ingore模版 | 
 **licenseTemplate** | **optional.String**| License模版 | 
 **private** | **optional.Int32**| 仓库开源类型。0(私有), 1(外部开源), 2(内部开源)。默认: 0 | [default to 0]
 **outsourced** | **optional.Bool**| 值为true值为外包仓库, false值为内部仓库。默认: 内部仓库(false) | 
 **projectCreator** | **optional.String**| 负责人的username | 
 **members** | **optional.String**| 用逗号分开的仓库成员。如: member1,member2 | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5OrgsOrgRepos**
> Project PostV5OrgsOrgRepos(ctx, org, body)
创建组织仓库

创建组织仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| 组织的路径(path/login) | 
  **body** | [**RepositoryPostParam**](RepositoryPostParam.md)| Repositorie 内容 | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoBranches**
> CompleteBranch PostV5ReposOwnerRepoBranches(ctx, owner, repo, body)
创建分支

创建分支

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **body** | [**CreateBranchParam**](CreateBranchParam.md)| 新建分支内容 | 

### Return type

[**CompleteBranch**](CompleteBranch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoCommitsShaComments**
> Note PostV5ReposOwnerRepoCommitsShaComments(ctx, owner, repo, sha, body, optional)
创建Commit评论

创建Commit评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **sha** | **string**| 评论的sha值 | 
  **body** | **string**| 评论的内容 | 
 **optional** | ***PostV5ReposOwnerRepoCommitsShaCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5ReposOwnerRepoCommitsShaCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessToken** | **optional.String**| 用户授权码 | 
 **path** | **optional.String**| 文件的相对路径 | 
 **position** | **optional.Int32**| Diff的相对行数 | 

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoContentsPath**
> CommitContent PostV5ReposOwnerRepoContentsPath(ctx, owner, repo, path, body)
新建文件

新建文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **path** | **string**| 文件的路径 | 
  **body** | [**NewFileParam**](NewFileParam.md)| 描述文件信息 | 

### Return type

[**CommitContent**](CommitContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoForks**
> Project PostV5ReposOwnerRepoForks(ctx, owner, repo, optional)
Fork一个仓库

Fork一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***PostV5ReposOwnerRepoForksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5ReposOwnerRepoForksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **optional.String**| 用户授权码 | 
 **organization** | **optional.String**| 组织空间地址，不填写默认Fork到用户个人空间地址 | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoKeys**
> SshKey PostV5ReposOwnerRepoKeys(ctx, owner, repo, key, title, optional)
为仓库添加公钥

为仓库添加公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **key** | **string**| 公钥内容 | 
  **title** | **string**| 公钥名称 | 
 **optional** | ***PostV5ReposOwnerRepoKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5ReposOwnerRepoKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessToken** | **optional.String**| 用户授权码 | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5ReposOwnerRepoPagesBuilds**
> PostV5ReposOwnerRepoPagesBuilds(ctx, owner, repo, optional)
请求建立Pages

请求建立Pages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***PostV5ReposOwnerRepoPagesBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5ReposOwnerRepoPagesBuildsOpts struct

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

# **PostV5ReposOwnerRepoReleases**
> Release PostV5ReposOwnerRepoReleases(ctx, owner, repo, tagName, name, body, targetCommitish, optional)
创建仓库Release

创建仓库Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **tagName** | **string**| Tag 名称, 提倡以v字母为前缀做为Release名称，例如v1.0或者v2.3.4 | 
  **name** | **string**| Release 名称 | 
  **body** | **string**| Release 描述 | 
  **targetCommitish** | **string**| 分支名称或者commit SHA, 默认是当前默认分支 | 
 **optional** | ***PostV5ReposOwnerRepoReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5ReposOwnerRepoReleasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **accessToken** | **optional.String**| 用户授权码 | 
 **prerelease** | **optional.Bool**| 是否为预览版本。默认: false（非预览版本） | 

### Return type

[**Release**](Release.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostV5UserRepos**
> Project PostV5UserRepos(ctx, name, optional)
创建一个仓库

创建一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 仓库名称 | 
 **optional** | ***PostV5UserReposOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostV5UserReposOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **optional.String**| 用户授权码 | 
 **description** | **optional.String**| 仓库描述 | 
 **homepage** | **optional.String**| 主页(eg: https://gitee.com) | 
 **hasIssues** | **optional.Bool**| 允许提Issue与否。默认: 允许(true) | [default to true]
 **hasWiki** | **optional.Bool**| 提供Wiki与否。默认: 提供(true) | [default to true]
 **canComment** | **optional.Bool**| 允许用户对仓库进行评论。默认: 允许(true) | [default to true]
 **autoInit** | **optional.Bool**| 值为true时则会用README初始化仓库。默认: 不初始化(false) | 
 **gitignoreTemplate** | **optional.String**| Git Ingore模版 | 
 **licenseTemplate** | **optional.String**| License模版 | 
 **private** | **optional.Bool**| 仓库公开或私有。默认: 公开(false) | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5ReposOwnerRepoBranchesBranchProtection**
> CompleteBranch PutV5ReposOwnerRepoBranchesBranchProtection(ctx, owner, repo, branch, body)
设置分支保护

设置分支保护

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **branch** | **string**| 分支名称 | 
  **body** | [**BranchProtectionPutParam**](BranchProtectionPutParam.md)| 设置分支保护参数 | 

### Return type

[**CompleteBranch**](CompleteBranch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5ReposOwnerRepoClear**
> PutV5ReposOwnerRepoClear(ctx, owner, repo, optional)
清空一个仓库

清空一个仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
 **optional** | ***PutV5ReposOwnerRepoClearOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5ReposOwnerRepoClearOpts struct

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

# **PutV5ReposOwnerRepoCollaboratorsUsername**
> ProjectMember PutV5ReposOwnerRepoCollaboratorsUsername(ctx, owner, repo, username, body)
添加仓库成员

添加仓库成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **username** | **string**| 用户名(username/login) | 
  **body** | [**ProjectMemberPutParam**](ProjectMemberPutParam.md)| 仓库成员内容 | 

### Return type

[**ProjectMember**](ProjectMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5ReposOwnerRepoContentsPath**
> CommitContent PutV5ReposOwnerRepoContentsPath(ctx, owner, repo, path, content, sha, message, optional)
更新文件

更新文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **path** | **string**| 文件的路径 | 
  **content** | **string**| 文件内容, 要用 base64 编码 | 
  **sha** | **string**| 文件的 Blob SHA，可通过 [获取仓库具体路径下的内容] API 获取 | 
  **message** | **string**| 提交信息 | 
 **optional** | ***PutV5ReposOwnerRepoContentsPathOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5ReposOwnerRepoContentsPathOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **accessToken** | **optional.String**| 用户授权码 | 
 **branch** | **optional.String**| 分支名称。默认为仓库对默认分支 | 
 **committerName** | **optional.String**| Committer的名字，默认为当前用户的名字 | 
 **committerEmail** | **optional.String**| Committer的邮箱，默认为当前用户的邮箱 | 
 **authorName** | **optional.String**| Author的名字，默认为当前用户的名字 | 
 **authorEmail** | **optional.String**| Author的邮箱，默认为当前用户的邮箱 | 

### Return type

[**CommitContent**](CommitContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutV5ReposOwnerRepoKeysEnableId**
> PutV5ReposOwnerRepoKeysEnableId(ctx, owner, repo, id, optional)
启用仓库公钥

启用仓库公钥

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **id** | **int32**| 公钥 ID | 
 **optional** | ***PutV5ReposOwnerRepoKeysEnableIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutV5ReposOwnerRepoKeysEnableIdOpts struct

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

# **PutV5ReposOwnerRepoReviewer**
> PutV5ReposOwnerRepoReviewer(ctx, owner, repo, body)
修改代码审查设置

修改代码审查设置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**| 仓库所属空间地址(企业、组织或个人的地址path) | 
  **repo** | **string**| 仓库路径(path) | 
  **body** | [**SetRepoReviewer**](SetRepoReviewer.md)| 修改代码审查的信息 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

