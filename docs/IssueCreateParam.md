# IssueCreateParam

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | 用户授权码 | [optional] [default to null]
**Repo** | **string** | 仓库路径(path) | [optional] [default to null]
**Title** | **string** | Issue标题 | [optional] [default to null]
**IssuType** | **string** | 企业自定义任务类型，非企业默认任务类型为“任务” | [optional] [default to null]
**Body** | **string** | Issue描述 | [optional] [default to null]
**Assignee** | **string** | Issue负责人的username | [optional] [default to null]
**Milestone** | **int32** | 里程碑序号 | [optional] [default to null]
**Labels** | **string** | 用逗号分开的标签，名称要求长度在 2-20 之间且非特殊字符。如: bug,performance | [optional] [default to null]
**Program** | **string** | 项目ID | [optional] [default to null]
**Collaborators** | **string** | Issue协助者的个人空间地址, 以 , 分隔 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


