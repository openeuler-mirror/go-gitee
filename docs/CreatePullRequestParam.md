# CreatePullRequestParam

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | 用户授权码 | [optional] [default to null]
**Title** | **string** | 必填。Pull Request 标题 | [optional] [default to null]
**Head** | **string** | 必填。Pull Request 提交的源分支。格式：branch 或者：username:branch | [optional] [default to null]
**Base** | **string** | 必填。Pull Request 提交目标分支的名称 | [optional] [default to null]
**Body** | **string** | 可选。Pull Request 内容 | [optional] [default to null]
**MilestoneNumber** | **int32** | 可选。里程碑序号(id) | [optional] [default to null]
**Labels** | **string** | 用逗号分开的标签，名称要求长度在 2-20 之间且非特殊字符。如: bug,performance | [optional] [default to null]
**Issue** | **string** | 可选。Pull Request的标题和内容可以根据指定的Issue Id自动填充 | [optional] [default to null]
**Assignees** | **string** | 可选。审查人员username，可多个，半角逗号分隔，如：(username1,username2) | [optional] [default to null]
**Testers** | **string** | 可选。测试人员username，可多个，半角逗号分隔，如：(username1,username2) | [optional] [default to null]
**PruneSourceBranch** | **bool** | 可选。合并PR后是否删除源分支，默认false（不删除） | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


