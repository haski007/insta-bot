# \AlbumApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlbumDownloadAlbumDownloadPost**](AlbumApi.md#AlbumDownloadAlbumDownloadPost) | **Post** /album/download | Album Download
[**AlbumDownloadByUrlsAlbumDownloadByUrlsPost**](AlbumApi.md#AlbumDownloadByUrlsAlbumDownloadByUrlsPost) | **Post** /album/download/by_urls | Album Download By Urls
[**AlbumUploadAlbumUploadPost**](AlbumApi.md#AlbumUploadAlbumUploadPost) | **Post** /album/upload | Album Upload



## AlbumDownloadAlbumDownloadPost

> []string AlbumDownloadAlbumDownloadPost(ctx).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).Execute()

Album Download



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlbumApi.AlbumDownloadAlbumDownloadPost(context.Background()).Sessionid(sessionid).MediaPk(mediaPk).Folder(folder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.AlbumDownloadAlbumDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlbumDownloadAlbumDownloadPost`: []string
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.AlbumDownloadAlbumDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlbumDownloadAlbumDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaPk** | **int32** |  | 
 **folder** | **string** |  | [default to &quot;&quot;]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlbumDownloadByUrlsAlbumDownloadByUrlsPost

> []string AlbumDownloadByUrlsAlbumDownloadByUrlsPost(ctx).Sessionid(sessionid).Urls(urls).Folder(folder).Execute()

Album Download By Urls



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
    urls := []string{"Inner_example"} // []string | 
    folder := "folder_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlbumApi.AlbumDownloadByUrlsAlbumDownloadByUrlsPost(context.Background()).Sessionid(sessionid).Urls(urls).Folder(folder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.AlbumDownloadByUrlsAlbumDownloadByUrlsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlbumDownloadByUrlsAlbumDownloadByUrlsPost`: []string
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.AlbumDownloadByUrlsAlbumDownloadByUrlsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlbumDownloadByUrlsAlbumDownloadByUrlsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **urls** | **[]string** |  | 
 **folder** | **string** |  | [default to &quot;&quot;]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlbumUploadAlbumUploadPost

> Media AlbumUploadAlbumUploadPost(ctx).Sessionid(sessionid).Files(files).Caption(caption).Usertags(usertags).Location(location).Execute()

Album Upload



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
    files := []*os.File{"TODO"} // []*os.File | 
    caption := "caption_example" // string | 
    usertags := []openapiclient.Usertag{*openapiclient.NewUsertag(*openapiclient.NewUserShort("Pk_example"), float32(123), float32(123))} // []Usertag |  (optional)
    location := *openapiclient.NewLocation("Name_example") // Location |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlbumApi.AlbumUploadAlbumUploadPost(context.Background()).Sessionid(sessionid).Files(files).Caption(caption).Usertags(usertags).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.AlbumUploadAlbumUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlbumUploadAlbumUploadPost`: Media
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.AlbumUploadAlbumUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlbumUploadAlbumUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **files** | **[]*os.File** |  | 
 **caption** | **string** |  | 
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

