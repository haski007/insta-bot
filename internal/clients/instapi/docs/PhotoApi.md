# \PhotoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PhotoDownloadByUrlPhotoDownloadByUrlPost**](PhotoApi.md#PhotoDownloadByUrlPhotoDownloadByUrlPost) | **Post** /photo/download/by_url | Photo Download By Url
[**PhotoDownloadPhotoDownloadPost**](PhotoApi.md#PhotoDownloadPhotoDownloadPost) | **Post** /photo/download | Photo Download
[**PhotoUploadPhotoUploadPost**](PhotoApi.md#PhotoUploadPhotoUploadPost) | **Post** /photo/upload | Photo Upload
[**PhotoUploadToStoryByUrlPhotoUploadToStoryByUrlPost**](PhotoApi.md#PhotoUploadToStoryByUrlPhotoUploadToStoryByUrlPost) | **Post** /photo/upload_to_story/by_url | Photo Upload To Story By Url
[**PhotoUploadToStoryPhotoUploadToStoryPost**](PhotoApi.md#PhotoUploadToStoryPhotoUploadToStoryPost) | **Post** /photo/upload_to_story | Photo Upload To Story



## PhotoDownloadByUrlPhotoDownloadByUrlPost

> interface{} PhotoDownloadByUrlPhotoDownloadByUrlPost(ctx).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()

Photo Download By Url



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
    resp, r, err := apiClient.PhotoApi.PhotoDownloadByUrlPhotoDownloadByUrlPost(context.Background()).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhotoApi.PhotoDownloadByUrlPhotoDownloadByUrlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PhotoDownloadByUrlPhotoDownloadByUrlPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PhotoApi.PhotoDownloadByUrlPhotoDownloadByUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPhotoDownloadByUrlPhotoDownloadByUrlPostRequest struct via the builder pattern


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


## PhotoDownloadPhotoDownloadPost

> interface{} PhotoDownloadPhotoDownloadPost(ctx).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).ReturnFile(returnFile).Execute()

Photo Download



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
    resp, r, err := apiClient.PhotoApi.PhotoDownloadPhotoDownloadPost(context.Background()).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhotoApi.PhotoDownloadPhotoDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PhotoDownloadPhotoDownloadPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PhotoApi.PhotoDownloadPhotoDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPhotoDownloadPhotoDownloadPostRequest struct via the builder pattern


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


## PhotoUploadPhotoUploadPost

> Media PhotoUploadPhotoUploadPost(ctx).Sessionid(sessionid).File(file).Caption(caption).UploadId(uploadId).Usertags(usertags).Location(location).Execute()

Photo Upload



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
    uploadId := "uploadId_example" // string |  (optional) (default to "")
    usertags := []openapiclient.Usertag{*openapiclient.NewUsertag(*openapiclient.NewUserShort("Pk_example"), float32(123), float32(123))} // []Usertag |  (optional)
    location := *openapiclient.NewLocation("Name_example") // Location |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhotoApi.PhotoUploadPhotoUploadPost(context.Background()).Sessionid(sessionid).File(file).Caption(caption).UploadId(uploadId).Usertags(usertags).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhotoApi.PhotoUploadPhotoUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PhotoUploadPhotoUploadPost`: Media
    fmt.Fprintf(os.Stdout, "Response from `PhotoApi.PhotoUploadPhotoUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPhotoUploadPhotoUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **file** | ***os.File** |  | 
 **caption** | **string** |  | 
 **uploadId** | **string** |  | [default to &quot;&quot;]
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


## PhotoUploadToStoryByUrlPhotoUploadToStoryByUrlPost

> Story PhotoUploadToStoryByUrlPhotoUploadToStoryByUrlPost(ctx).Sessionid(sessionid).Url(url).AsVideo(asVideo).Caption(caption).Mentions(mentions).Locations(locations).Links(links).Hashtags(hashtags).Stickers(stickers).Execute()

Photo Upload To Story By Url



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
    asVideo := true // bool |  (optional) (default to false)
    caption := "caption_example" // string |  (optional) (default to "")
    mentions := []openapiclient.StoryMention{*openapiclient.NewStoryMention(*openapiclient.NewUserShort("Pk_example"))} // []StoryMention |  (optional)
    locations := []openapiclient.StoryLocation{*openapiclient.NewStoryLocation(*openapiclient.NewLocation("Name_example"))} // []StoryLocation |  (optional)
    links := []openapiclient.StoryLink{*openapiclient.NewStoryLink("WebUri_example")} // []StoryLink |  (optional)
    hashtags := []openapiclient.StoryHashtag{*openapiclient.NewStoryHashtag("TODO")} // []StoryHashtag |  (optional)
    stickers := []openapiclient.StorySticker{*openapiclient.NewStorySticker(float32(123), float32(123), float32(123), float32(123), map[string]interface{}(123))} // []StorySticker |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhotoApi.PhotoUploadToStoryByUrlPhotoUploadToStoryByUrlPost(context.Background()).Sessionid(sessionid).Url(url).AsVideo(asVideo).Caption(caption).Mentions(mentions).Locations(locations).Links(links).Hashtags(hashtags).Stickers(stickers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhotoApi.PhotoUploadToStoryByUrlPhotoUploadToStoryByUrlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PhotoUploadToStoryByUrlPhotoUploadToStoryByUrlPost`: Story
    fmt.Fprintf(os.Stdout, "Response from `PhotoApi.PhotoUploadToStoryByUrlPhotoUploadToStoryByUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPhotoUploadToStoryByUrlPhotoUploadToStoryByUrlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **url** | **string** |  | 
 **asVideo** | **bool** |  | [default to false]
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


## PhotoUploadToStoryPhotoUploadToStoryPost

> Story PhotoUploadToStoryPhotoUploadToStoryPost(ctx).Sessionid(sessionid).File(file).AsVideo(asVideo).Caption(caption).Mentions(mentions).Locations(locations).Links(links).Hashtags(hashtags).Stickers(stickers).Execute()

Photo Upload To Story



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
    asVideo := true // bool |  (optional) (default to false)
    caption := "caption_example" // string |  (optional) (default to "")
    mentions := []openapiclient.StoryMention{*openapiclient.NewStoryMention(*openapiclient.NewUserShort("Pk_example"))} // []StoryMention |  (optional)
    locations := []openapiclient.StoryLocation{*openapiclient.NewStoryLocation(*openapiclient.NewLocation("Name_example"))} // []StoryLocation |  (optional)
    links := []openapiclient.StoryLink{*openapiclient.NewStoryLink("WebUri_example")} // []StoryLink |  (optional)
    hashtags := []openapiclient.StoryHashtag{*openapiclient.NewStoryHashtag("TODO")} // []StoryHashtag |  (optional)
    stickers := []openapiclient.StorySticker{*openapiclient.NewStorySticker(float32(123), float32(123), float32(123), float32(123), map[string]interface{}(123))} // []StorySticker |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhotoApi.PhotoUploadToStoryPhotoUploadToStoryPost(context.Background()).Sessionid(sessionid).File(file).AsVideo(asVideo).Caption(caption).Mentions(mentions).Locations(locations).Links(links).Hashtags(hashtags).Stickers(stickers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhotoApi.PhotoUploadToStoryPhotoUploadToStoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PhotoUploadToStoryPhotoUploadToStoryPost`: Story
    fmt.Fprintf(os.Stdout, "Response from `PhotoApi.PhotoUploadToStoryPhotoUploadToStoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPhotoUploadToStoryPhotoUploadToStoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **file** | ***os.File** |  | 
 **asVideo** | **bool** |  | [default to false]
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

