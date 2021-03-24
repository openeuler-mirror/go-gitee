# RepositoryPostParam

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | 用户授权码 | [optional] [default to null]
**Name** | **string** | 仓库名称 | [optional] [default to null]
**Description** | **string** | 仓库描述 | [optional] [default to null]
**Homepage** | **string** | 主页(eg: https://gitee.com) | [optional] [default to null]
**HasIssues** | **bool** | 允许提Issue与否。默认: 允许(true) | [optional] [default to null]
**HasWiki** | **bool** | 提供Wiki与否。默认: 提供(true) | [optional] [default to null]
**CanComment** | **bool** | 允许用户对仓库进行评论。默认： 允许(true) | [optional] [default to null]
**Public** | **int32** | 仓库开源类型。0(私有), 1(外部开源), 2(内部开源)，注：与private互斥，以public为主。 | [optional] [default to null]
**Private** | **bool** | 仓库公开或私有。默认: 公开(false)，注：与public互斥，以public为主。 | [optional] [default to null]
**AutoInit** | **bool** | 值为true时则会用README初始化仓库。默认: 不初始化(false) | [optional] [default to null]
**GitignoreTemplate** | **string** | Git Ingore模版 | [optional] [default to null]
**LicenseTemplate** | **string** | License模版 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


