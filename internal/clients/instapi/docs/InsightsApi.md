# \InsightsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountInsightsAccountPost**](InsightsApi.md#AccountInsightsAccountPost) | **Post** /insights/account | Account
[**MediaFeedAllInsightsMediaFeedAllPost**](InsightsApi.md#MediaFeedAllInsightsMediaFeedAllPost) | **Post** /insights/media_feed_all | Media Feed All
[**MediaInsightsMediaPost**](InsightsApi.md#MediaInsightsMediaPost) | **Post** /insights/media | Media



## AccountInsightsAccountPost

> map[string]interface{} AccountInsightsAccountPost(ctx).Sessionid(sessionid).Execute()

Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionid := "sessionid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.AccountInsightsAccountPost(context.Background()).Sessionid(sessionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.AccountInsightsAccountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountInsightsAccountPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.AccountInsightsAccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountInsightsAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaFeedAllInsightsMediaFeedAllPost

> []map[string]interface{} MediaFeedAllInsightsMediaFeedAllPost(ctx).Sessionid(sessionid).PostType(postType).TimeFrame(timeFrame).DataOrdering(dataOrdering).Count(count).Execute()

Media Feed All



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionid := "sessionid_example" // string | 
    postType := "postType_example" // string |  (optional) (default to "ALL")
    timeFrame := "timeFrame_example" // string |  (optional) (default to "TWO_YEARS")
    dataOrdering := "dataOrdering_example" // string |  (optional) (default to "REACH_COUNT")
    count := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.MediaFeedAllInsightsMediaFeedAllPost(context.Background()).Sessionid(sessionid).PostType(postType).TimeFrame(timeFrame).DataOrdering(dataOrdering).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.MediaFeedAllInsightsMediaFeedAllPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaFeedAllInsightsMediaFeedAllPost`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.MediaFeedAllInsightsMediaFeedAllPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaFeedAllInsightsMediaFeedAllPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **postType** | **string** |  | [default to &quot;ALL&quot;]
 **timeFrame** | **string** |  | [default to &quot;TWO_YEARS&quot;]
 **dataOrdering** | **string** |  | [default to &quot;REACH_COUNT&quot;]
 **count** | **int32** |  | [default to 0]

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaInsightsMediaPost

> map[string]interface{} MediaInsightsMediaPost(ctx).Sessionid(sessionid).MediaPk(mediaPk).Execute()

Media



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionid := "sessionid_example" // string | 
    mediaPk := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.MediaInsightsMediaPost(context.Background()).Sessionid(sessionid).MediaPk(mediaPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.MediaInsightsMediaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaInsightsMediaPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.MediaInsightsMediaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaInsightsMediaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaPk** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

