# \TableApi

All URIs are relative to *http://example.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTable**](TableApi.md#GetTable) | **Get** /api/table/{id} | Get a table
[**GetTableMetadata**](TableApi.md#GetTableMetadata) | **Get** /api/table/{id}/query_metadata | Get table metadata
[**UpdateTable**](TableApi.md#UpdateTable) | **Put** /api/table/{id} | Update a table


# **GetTable**
> Table GetTable(ctx, id)
Get a table

Fetch one table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableMetadata**
> TableMetadata GetTableMetadata(ctx, id, optional)
Get table metadata

Fetch a table's metadata. (fields, db, foreign keys)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
 **optional** | ***GetTableMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTableMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSensitiveFields** | **optional.Bool**| include fields marked visibility_type &#x3D; \&quot;sensitive\&quot; | [default to false]

### Return type

[**TableMetadata**](TableMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTable**
> Table UpdateTable(ctx, id, table)
Update a table

Update a table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
  **table** | [**Table**](Table.md)| Table | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

