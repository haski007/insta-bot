# \VideoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VideoDownloadByUrlVideoDownloadByUrlPost**](VideoApi.md#VideoDownloadByUrlVideoDownloadByUrlPost) | **Post** /video/download/by_url | Video Download By Url
[**VideoDownloadVideoDownloadPost**](VideoApi.md#VideoDownloadVideoDownloadPost) | **Post** /video/download | Video Download
[**VideoUploadToStoryByUrlVideoUploadToStoryByUrlPost**](VideoApi.md#VideoUploadToStoryByUrlVideoUploadToStoryByUrlPost) | **Post** /video/upload_to_story/by_url | Video Upload To Story By Url
[**VideoUploadToStoryVideoUploadToStoryPost**](VideoApi.md#VideoUploadToStoryVideoUploadToStoryPost) | **Post** /video/upload_to_story | Video Upload To Story
[**VideoUploadVideoUploadPost**](VideoApi.md#VideoUploadVideoUploadPost) | **Post** /video/upload | Video Upload



## VideoDownloadByUrlVideoDownloadByUrlPost

> interface{} VideoDownloadByUrlVideoDownloadByUrlPost(ctx).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()

Video Download By Url



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
    resp, r, err := apiClient.VideoApi.VideoDownloadByUrlVideoDownloadByUrlPost(context.Background()).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoApi.VideoDownloadByUrlVideoDownloadByUrlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VideoDownloadByUrlVideoDownloadByUrlPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `VideoApi.VideoDownloadByUrlVideoDownloadByUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVideoDownloadByUrlVideoDownloadByUrlPostRequest struct via the builder pattern


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


## VideoDownloadVideoDownloadPost

> interface{} VideoDownloadVideoDownloadPost(ctx).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).ReturnFile(returnFile).Execute()

Video Download



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
    folder := "folder_example" // string |  (optional) (default to "")
    returnFile := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VideoApi.VideoDownloadVideoDownloadPost(context.Background()).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoApi.VideoDownloadVideoDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VideoDownloadVideoDownloadPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `VideoApi.VideoDownloadVideoDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVideoDownloadVideoDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaPk** | **int32** |  | 
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


## VideoUploadToStoryByUrlVideoUploadToStoryByUrlPost

> Story VideoUploadToStoryByUrlVideoUploadToStoryByUrlPost(ctx).Sessionid(sessionid).Url(url).Caption(caption).Mentions(mentions).Locations(locations).Links(links).Hashtags(hashtags).Stickers(stickers).Execute()

Video Upload To Story By Url



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
    caption := "caption_example" // string |  (optional) (default to "")
    mentions := []openapiclient.StoryMention{*openapiclient.NewStoryMention(*openapiclient.NewUserShort("Pk_example"))} // []StoryMention |  (optional)
    locations := []openapiclient.StoryLocation{*openapiclient.NewStoryLocation(*openapiclient.NewLocation("Name_example"))} // []StoryLocation |  (optional)
    links := []openapiclient.StoryLink{*openapiclient.NewStoryLink("WebUri_example")} // []StoryLink |  (optional)
    hashtags := []openapiclient.StoryHashtag{*openapiclient.NewStoryHashtag("TODO")} // []StoryHashtag |  (optional)
    stickers := []openapiclient.StorySticker{*openapiclient.NewStorySticker(float32(123), float32(123), float32(123), float32(123), map[string]interface{}(123))} // []StorySticker |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VideoApi.VideoUploadToStoryByUrlVideoUploadToStoryByUrlPost(context.Background()).Sessionid(sessionid).Url(url).Caption(caption).Mentions(mentions).Locations(locations).Links(links).Hashtags(hashtags).Stickers(stickers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoApi.VideoUploadToStoryByUrlVideoUploadToStoryByUrlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VideoUploadToStoryByUrlVideoUploadToStoryByUrlPost`: Story
    fmt.Fprintf(os.Stdout, "Response from `VideoApi.VideoUploadToStoryByUrlVideoUploadToStoryByUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVideoUploadToStoryByUrlVideoUploadToStoryByUrlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **url** | **string** |  | 
 **caption** | **string** |  | [default to &quot;&quot;]
 **mentions** | [**[]StoryMention**](StoryMention.md) |  | 
 **locations** | [**[]StoryLocation**](StoryLocation.md) |  | 
 **links** | [**[]StoryLink**](StoryLink.md) |  | 
 **hashtags** | [**[]StoryHashtag**](StoryHashtag.md) |  | 
 **stickers** | [**[]StorySticker**](StorySticker.md) |  | 

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


## VideoUploadToStoryVideoUploadToStoryPost

> Story VideoUploadToStoryVideoUploadToStoryPost(ctx).Sessionid(sessionid).File(file).Caption(caption).Mentions(mentions).Locations(locations).Links(links).Hashtags(hashtags).Stickers(stickers).Execute()

Video Upload To Story



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
    file := os.NewFile(1234, "some_file") // *os.File | 
    caption := "caption_example" // string |  (optional) (default to "")
    mentions := []openapiclient.StoryMention{*openapiclient.NewStoryMention(*openapiclient.NewUserShort("Pk_example"))} // []StoryMention |  (optional)
    locations := []openapiclient.StoryLocation{*openapiclient.NewStoryLocation(*openapiclient.NewLocation("Name_example"))} // []StoryLocation |  (optional)
    links := []openapiclient.StoryLink{*openapiclient.NewStoryLink("WebUri_example")} // []StoryLink |  (optional)
    hashtags := []openapiclient.StoryHashtag{*openapiclient.NewStoryHashtag("TODO")} // []StoryHashtag |  (optional)
    stickers := []openapiclient.StorySticker{*openapiclient.NewStorySticker(float32(123), float32(123), float32(123), float32(123), map[string]interface{}(123))} // []StorySticker |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VideoApi.VideoUploadToStoryVideoUploadToStoryPost(context.Background()).Sessionid(sessionid).File(file).Caption(caption).Mentions(mentions).Locations(locations).Links(links).Hashtags(hashtags).Stickers(stickers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoApi.VideoUploadToStoryVideoUploadToStoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VideoUploadToStoryVideoUploadToStoryPost`: Story
    fmt.Fprintf(os.Stdout, "Response from `VideoApi.VideoUploadToStoryVideoUploadToStoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVideoUploadToStoryVideoUploadToStoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **file** | ***os.File** |  | 
 **caption** | **string** |  | [default to &quot;&quot;]
 **mentions** | [**[]StoryMention**](StoryMention.md) |  | 
 **locations** | [**[]StoryLocation**](StoryLocation.md) |  | 
 **links** | [**[]StoryLink**](StoryLink.md) |  | 
 **hashtags** | [**[]StoryHashtag**](StoryHashtag.md) |  | 
 **stickers** | [**[]StorySticker**](StorySticker.md) |  | 

### Return type

[**Story**](Story.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideoUploadVideoUploadPost

> Media VideoUploadVideoUploadPost(ctx).Sessionid(sessionid).File(file).Caption(caption).Thumbnail(thumbnail).Usertags(usertags).Location(location).Execute()

Video Upload



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
    file := os.NewFile(1234, "some_file") // *os.File | 
    caption := "caption_example" // string | 
    thumbnail := os.NewFile(1234, "some_file") // *os.File |  (optional)
    usertags := []openapiclient.Usertag{*openapiclient.NewUsertag(*openapiclient.NewUserShort("Pk_example"), float32(123), float32(123))} // []Usertag |  (optional)
    location := *openapiclient.NewLocation("Name_example") // Location |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VideoApi.VideoUploadVideoUploadPost(context.Background()).Sessionid(sessionid).File(file).Caption(caption).Thumbnail(thumbnail).Usertags(usertags).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoApi.VideoUploadVideoUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VideoUploadVideoUploadPost`: Media
    fmt.Fprintf(os.Stdout, "Response from `VideoApi.VideoUploadVideoUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVideoUploadVideoUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **file** | ***os.File** |  | 
 **caption** | **string** |  | 
 **thumbnail** | ***os.File** |  | 
 **usertags** | [**[]Usertag**](Usertag.md) |  | 
 **location** | [**Location**](Location.md) |  | 

### Return type

[**Media**](Media.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

