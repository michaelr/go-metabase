# \FieldApi

All URIs are relative to *http://example.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateField**](FieldApi.md#UpdateField) | **Put** /api/field/{id} | Update field


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

