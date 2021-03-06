/*
 * Metabase
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metabase

import (
	"time"
)

type Database struct {
	Id                       int64    `json:"id"`
	Name                     string   `json:"name,omitempty"`
	Description              string   `json:"description,omitempty"`
	Features                 []string `json:"features,omitempty"`
	IsFullSync               bool     `json:"is_full_sync,omitempty"`
	IsSample                 bool     `json:"is_sample,omitempty"`
	CacheFieldValuesSchedule string   `json:"cache_field_values_schedule,omitempty"`
	MetadataSyncSchedule     string   `json:"metadata_sync_schedule,omitempty"`
	// type unknown
	Caveats           string    `json:"caveats,omitempty"`
	Engine            string    `json:"engine,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
	NativePermissions string    `json:"native_permissions,omitempty"`
	// type unknown
	PointsOfInterest string          `json:"points_of_interest,omitempty"`
	Details          DatabaseDetails `json:"details,omitempty"`
	Tables           []DatabaseTable `json:"tables,omitempty"`
}
