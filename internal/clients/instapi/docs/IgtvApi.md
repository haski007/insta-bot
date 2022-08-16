# \IgtvApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IgtvDownloadByUrlIgtvDownloadByUrlPost**](IgtvApi.md#IgtvDownloadByUrlIgtvDownloadByUrlPost) | **Post** /igtv/download/by_url | Igtv Download By Url
[**IgtvDownloadIgtvDownloadPost**](IgtvApi.md#IgtvDownloadIgtvDownloadPost) | **Post** /igtv/download | Igtv Download
[**IgtvUploadIgtvUploadPost**](IgtvApi.md#IgtvUploadIgtvUploadPost) | **Post** /igtv/upload | Igtv Upload



## IgtvDownloadByUrlIgtvDownloadByUrlPost

> interface{} IgtvDownloadByUrlIgtvDownloadByUrlPost(ctx).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()

Igtv Download By Url



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
    resp, r, err := apiClient.IgtvApi.IgtvDownloadByUrlIgtvDownloadByUrlPost(context.Background()).Sessionid(sessionid).Url(url).Filename(filename).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IgtvApi.IgtvDownloadByUrlIgtvDownloadByUrlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IgtvDownloadByUrlIgtvDownloadByUrlPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IgtvApi.IgtvDownloadByUrlIgtvDownloadByUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIgtvDownloadByUrlIgtvDownloadByUrlPostRequest struct via the builder pattern


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


## IgtvDownloadIgtvDownloadPost

> interface{} IgtvDownloadIgtvDownloadPost(ctx).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).ReturnFile(returnFile).Execute()

Igtv Download



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
    resp, r, err := apiClient.IgtvApi.IgtvDownloadIgtvDownloadPost(context.Background()).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).ReturnFile(returnFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IgtvApi.IgtvDownloadIgtvDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IgtvDownloadIgtvDownloadPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `IgtvApi.IgtvDownloadIgtvDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIgtvDownloadIgtvDownloadPostRequest struct via the builder pattern


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


## IgtvUploadIgtvUploadPost

> Media IgtvUploadIgtvUploadPost(ctx).Sessionid(sessionid).File(file).Title(title).Caption(caption).Thumbnail(thumbnail).Usertags(usertags).Location(location).Execute()

Igtv Upload



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
    title := "title_example" // string | 
    caption := "caption_example" // string | 
    thumbnail := os.NewFile(1234, "some_file") // *os.File |  (optional)
    usertags := []openapiclient.Usertag{*openapiclient.NewUsertag(*openapiclient.NewUserShort("Pk_example"), float32(123), float32(123))} // []Usertag |  (optional)
    location := *openapiclient.NewLocation("Name_example") // Location |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IgtvApi.IgtvUploadIgtvUploadPost(context.Background()).Sessionid(sessionid).File(file).Title(title).Caption(caption).Thumbnail(thumbnail).Usertags(usertags).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IgtvApi.IgtvUploadIgtvUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IgtvUploadIgtvUploadPost`: Media
    fmt.Fprintf(os.Stdout, "Response from `IgtvApi.IgtvUploadIgtvUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIgtvUploadIgtvUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **file** | ***os.File** |  | 
 **title** | **string** |  | 
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

