/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 列出指定用户的所有公钥
type SshKeyBasic struct {
	Id  string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
}
