# \StoryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StoryDeleteStoryDeletePost**](StoryApi.md#StoryDeleteStoryDeletePost) | **Post** /story/delete | Story Delete
[**StoryDownloadByUrlStoryDownloadByUrlPost**](StoryApi.md#StoryDownloadByUrlStoryDownloadByUrlPost) | **Post** /story/download/by_url | Story Download By Url
[**StoryDownloadStoryDownloadPost**](StoryApi.md#StoryDownloadStoryDownloadPost) | **Post** /story/download | Story Download
[**StoryInfoStoryInfoPost**](StoryApi.md#StoryInfoStoryInfoPost) | **Post** /story/info | Story Info
[**StoryPkFromUrlStoryPkFromUrlGet**](StoryApi.md#StoryPkFromUrlStoryPkFromUrlGet) | **Get** /story/pk_from_url | Story Pk From Url
[**StorySeenStorySeenPost**](StoryApi.md#StorySeenStorySeenPost) | **Post** /story/seen | Story Seen
[**StoryUserStoriesStoryUserStoriesPost**](StoryApi.md#StoryUserStoriesStoryUserStoriesPost) | **Post** /story/user_stories | Story User Stories



## StoryDeleteStoryDeletePost

> bool StoryDeleteStoryDeletePost(ctx).Sessionid(sessionid).StoryPk(storyPk).Execute()

Story Delete



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
    storyPk := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoryApi.StoryDeleteStoryDeletePost(context.Background()).Sessionid(sessionid).StoryPk(storyPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoryApi.StoryDeleteStoryDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoryDeleteStoryDeletePost`: bool
    fmt.Fprintf(os.Stdout, "Response from `StoryApi.StoryDeleteStoryDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoryDeleteStoryDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **storyPk** | **int32** |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoryDownloadByUrlStoryDownloadByUrlPost

> interface{} StoryDownloadByUrlStoryDownloadByUrlPost(ctx).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()

Story Download By Url



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
    url := "url_example" // string | 
    filename := "filename_example" // string |  (optional) (default to "")
    folder := "folder_example" // string |  (optional) (default to "")
    returnFile := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoryApi.StoryDownloadByUrlStoryDownloadByUrlPost(context.Background()).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoryApi.StoryDownloadByUrlStoryDownloadByUrlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoryDownloadByUrlStoryDownloadByUrlPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StoryApi.StoryDownloadByUrlStoryDownloadByUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoryDownloadByUrlStoryDownloadByUrlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **url** | **string** |  | 
 **filename** | **string** |  | [default to &quot;&quot;]
 **folder** | **string** |  | [default to &quot;&quot;]
 **returnFile** | **bool** |  | [default to true]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoryDownloadStoryDownloadPost

> interface{} StoryDownloadStoryDownloadPost(ctx).Sessionid(sessionid).StoryPk(storyPk).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()

Story Download



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
    storyPk := int32(56) // int32 | 
    filename := "filename_example" // string |  (optional) (default to "")
    folder := "folder_example" // string |  (optional) (default to "")
    returnFile := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoryApi.StoryDownloadStoryDownloadPost(context.Background()).Sessionid(sessionid).StoryPk(storyPk).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoryApi.StoryDownloadStoryDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoryDownloadStoryDownloadPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StoryApi.StoryDownloadStoryDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoryDownloadStoryDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **storyPk** | **int32** |  | 
 **filename** | **string** |  | [default to &quot;&quot;]
 **folder** | **string** |  | [default to &quot;&quot;]
 **returnFile** | **bool** |  | [default to true]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoryInfoStoryInfoPost

> Story StoryInfoStoryInfoPost(ctx).Sessionid(sessionid).StoryPk(storyPk).UseCache(useCache).Execute()

Story Info



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
    storyPk := int32(56) // int32 | 
    useCache := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoryApi.StoryInfoStoryInfoPost(context.Background()).Sessionid(sessionid).StoryPk(storyPk).UseCache(useCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoryApi.StoryInfoStoryInfoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoryInfoStoryInfoPost`: Story
    fmt.Fprintf(os.Stdout, "Response from `StoryApi.StoryInfoStoryInfoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoryInfoStoryInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **storyPk** | **int32** |  | 
 **useCache** | **bool** |  | [default to true]

### Return type

[**Story**](Story.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoryPkFromUrlStoryPkFromUrlGet

> interface{} StoryPkFromUrlStoryPkFromUrlGet(ctx).Url(url).Execute()

Story Pk From Url



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
    url := "url_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoryApi.StoryPkFromUrlStoryPkFromUrlGet(context.Background()).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoryApi.StoryPkFromUrlStoryPkFromUrlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoryPkFromUrlStoryPkFromUrlGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StoryApi.StoryPkFromUrlStoryPkFromUrlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoryPkFromUrlStoryPkFromUrlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorySeenStorySeenPost

> bool StorySeenStorySeenPost(ctx).Sessionid(sessionid).StoryPks(storyPks).SkippedStoryPks(skippedStoryPks).Execute()

Story Seen



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
    storyPks := []int32{int32(123)} // []int32 | 
    skippedStoryPks := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoryApi.StorySeenStorySeenPost(context.Background()).Sessionid(sessionid).StoryPks(storyPks).SkippedStoryPks(skippedStoryPks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoryApi.StorySeenStorySeenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorySeenStorySeenPost`: bool
    fmt.Fprintf(os.Stdout, "Response from `StoryApi.StorySeenStorySeenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorySeenStorySeenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **storyPks** | **[]int32** |  | 
 **skippedStoryPks** | **[]int32** |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoryUserStoriesStoryUserStoriesPost

> []Story StoryUserStoriesStoryUserStoriesPost(ctx).Sessionid(sessionid).UserId(userId).Amount(amount).Execute()

Story User Stories



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
    userId := "userId_example" // string | 
    amount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoryApi.StoryUserStoriesStoryUserStoriesPost(context.Background()).Sessionid(sessionid).UserId(userId).Amount(amount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoryApi.StoryUserStoriesStoryUserStoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoryUserStoriesStoryUserStoriesPost`: []Story
    fmt.Fprintf(os.Stdout, "Response from `StoryApi.StoryUserStoriesStoryUserStoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoryUserStoriesStoryUserStoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **string** |  | 
 **amount** | **int32** |  | 

### Return type

[**[]Story**](Story.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

