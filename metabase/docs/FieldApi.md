# \FieldApi

All URIs are relative to *http://example.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDimension**](FieldApi.md#DeleteDimension) | **Delete** /api/field/{id}/dimension | Delete field dimension
[**GetField**](FieldApi.md#GetField) | **Get** /api/field/{id} | Get field
[**UpdateDimension**](FieldApi.md#UpdateDimension) | **Post** /api/field/{id}/dimension | Update field dimension
[**UpdateField**](FieldApi.md#UpdateField) | **Put** /api/field/{id} | Update field


# **DeleteDimension**
> DeleteDimension(ctx, id)
Delete field dimension

Delete field dimension

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetField**
> Field GetField(ctx, id)
Get field

Get field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDimension**
> FieldDimension UpdateDimension(ctx, id, fieldDimension)
Update field dimension

Update field dimension

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
  **fieldDimension** | [**FieldDimension**](FieldDimension.md)| Dimension | 

### Return type

[**FieldDimension**](FieldDimension.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateField**
> Field UpdateField(ctx, id, field)
Update field

Update field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
  **field** | [**Field**](Field.md)| Field | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

