/*
 * Metabase
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metabase

type DatasetQueryResultsCol struct {
	Description string `json:"description,omitempty"`
	TableId     int64  `json:"table_id,omitempty"`
	SchemaName  string `json:"schema_name,omitempty"`
	SpecialType string `json:"special_type,omitempty"`
	Name        string `json:"name,omitempty"`
	Source      string `json:"source,omitempty"`
	// unknown type
	RemappedFrom string `json:"remapped_from,omitempty"`
	// can be '{\"target_table_id\":517}'
	ExtraInfo map[string]interface{} `json:"extra_info,omitempty"`
	// unknown type
	FkFieldId string `json:"fk_field_id,omitempty"`
	// unknown type
	RemappedTo     string                            `json:"remapped_to,omitempty"`
	Id             int64                             `json:"id,omitempty"`
	VisibilityType string                            `json:"visibility_type,omitempty"`
	Target         DatasetQueryResultsColTarget      `json:"target,omitempty"`
	DisplayName    string                            `json:"display_name,omitempty"`
	Fingerprint    DatasetQueryResultsColFingerprint `json:"fingerprint,omitempty"`
	BaseType       string                            `json:"base_type,omitempty"`
}
