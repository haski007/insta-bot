# \ClipApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClipDownloadByUrlClipDownloadByUrlPost**](ClipApi.md#ClipDownloadByUrlClipDownloadByUrlPost) | **Post** /clip/download/by_url | Clip Download By Url
[**ClipDownloadClipDownloadPost**](ClipApi.md#ClipDownloadClipDownloadPost) | **Post** /clip/download | Clip Download
[**ClipUploadClipUploadPost**](ClipApi.md#ClipUploadClipUploadPost) | **Post** /clip/upload | Clip Upload



## ClipDownloadByUrlClipDownloadByUrlPost

> interface{} ClipDownloadByUrlClipDownloadByUrlPost(ctx).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()

Clip Download By Url



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
    resp, r, err := apiClient.ClipApi.ClipDownloadByUrlClipDownloadByUrlPost(context.Background()).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.ClipDownloadByUrlClipDownloadByUrlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClipDownloadByUrlClipDownloadByUrlPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClipApi.ClipDownloadByUrlClipDownloadByUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClipDownloadByUrlClipDownloadByUrlPostRequest struct via the builder pattern


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


## ClipDownloadClipDownloadPost

> interface{} ClipDownloadClipDownloadPost(ctx).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).ReturnFile(returnFile).Execute()

Clip Download



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
    resp, r, err := apiClient.ClipApi.ClipDownloadClipDownloadPost(context.Background()).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.ClipDownloadClipDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClipDownloadClipDownloadPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClipApi.ClipDownloadClipDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClipDownloadClipDownloadPostRequest struct via the builder pattern


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


## ClipUploadClipUploadPost

> Media ClipUploadClipUploadPost(ctx).Sessionid(sessionid).File(file).Caption(caption).Thumbnail(thumbnail).Usertags(usertags).Location(location).Execute()

Clip Upload



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
    resp, r, err := apiClient.ClipApi.ClipUploadClipUploadPost(context.Background()).Sessionid(sessionid).File(file).Caption(caption).Thumbnail(thumbnail).Usertags(usertags).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClipApi.ClipUploadClipUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClipUploadClipUploadPost`: Media
    fmt.Fprintf(os.Stdout, "Response from `ClipApi.ClipUploadClipUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClipUploadClipUploadPostRequest struct via the builder pattern


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

