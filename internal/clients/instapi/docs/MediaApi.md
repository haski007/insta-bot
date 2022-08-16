# \MediaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MediaArchiveMediaArchivePost**](MediaApi.md#MediaArchiveMediaArchivePost) | **Post** /media/archive | Media Archive
[**MediaDeleteMediaDeletePost**](MediaApi.md#MediaDeleteMediaDeletePost) | **Post** /media/delete | Media Delete
[**MediaEditMediaEditPost**](MediaApi.md#MediaEditMediaEditPost) | **Post** /media/edit | Media Edit
[**MediaIdMediaIdGet**](MediaApi.md#MediaIdMediaIdGet) | **Get** /media/id | Media Id
[**MediaInfoMediaInfoPost**](MediaApi.md#MediaInfoMediaInfoPost) | **Post** /media/info | Media Info
[**MediaLikeMediaLikePost**](MediaApi.md#MediaLikeMediaLikePost) | **Post** /media/like | Media Like
[**MediaLikersMediaLikersPost**](MediaApi.md#MediaLikersMediaLikersPost) | **Post** /media/likers | Media Likers
[**MediaOembedMediaOembedPost**](MediaApi.md#MediaOembedMediaOembedPost) | **Post** /media/oembed | Media Oembed
[**MediaPkFromCodeMediaPkFromCodeGet**](MediaApi.md#MediaPkFromCodeMediaPkFromCodeGet) | **Get** /media/pk_from_code | Media Pk From Code
[**MediaPkFromUrlMediaPkFromUrlGet**](MediaApi.md#MediaPkFromUrlMediaPkFromUrlGet) | **Get** /media/pk_from_url | Media Pk From Url
[**MediaPkMediaPkGet**](MediaApi.md#MediaPkMediaPkGet) | **Get** /media/pk | Media Pk
[**MediaSeenMediaSeenPost**](MediaApi.md#MediaSeenMediaSeenPost) | **Post** /media/seen | Media Seen
[**MediaUnarchiveMediaUnarchivePost**](MediaApi.md#MediaUnarchiveMediaUnarchivePost) | **Post** /media/unarchive | Media Unarchive
[**MediaUnlikeMediaUnlikePost**](MediaApi.md#MediaUnlikeMediaUnlikePost) | **Post** /media/unlike | Media Unlike
[**MediaUserMediaUserPost**](MediaApi.md#MediaUserMediaUserPost) | **Post** /media/user | Media User
[**UserMediasMediaUserMediasPost**](MediaApi.md#UserMediasMediaUserMediasPost) | **Post** /media/user_medias | User Medias



## MediaArchiveMediaArchivePost

> bool MediaArchiveMediaArchivePost(ctx).Sessionid(sessionid).MediaId(mediaId).Revert(revert).Execute()

Media Archive



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
    mediaId := "mediaId_example" // string | 
    revert := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaArchiveMediaArchivePost(context.Background()).Sessionid(sessionid).MediaId(mediaId).Revert(revert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaArchiveMediaArchivePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaArchiveMediaArchivePost`: bool
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaArchiveMediaArchivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaArchiveMediaArchivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaId** | **string** |  | 
 **revert** | **bool** |  | [default to false]

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


## MediaDeleteMediaDeletePost

> bool MediaDeleteMediaDeletePost(ctx).Sessionid(sessionid).MediaId(mediaId).Execute()

Media Delete



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
    mediaId := "mediaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaDeleteMediaDeletePost(context.Background()).Sessionid(sessionid).MediaId(mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaDeleteMediaDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaDeleteMediaDeletePost`: bool
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaDeleteMediaDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaDeleteMediaDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaId** | **string** |  | 

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


## MediaEditMediaEditPost

> map[string]interface{} MediaEditMediaEditPost(ctx).Sessionid(sessionid).MediaId(mediaId).Caption(caption).Title(title).Usertags(usertags).Location(location).Execute()

Media Edit



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
    mediaId := "mediaId_example" // string | 
    caption := "caption_example" // string | 
    title := "title_example" // string |  (optional) (default to "")
    usertags := []openapiclient.Usertag{*openapiclient.NewUsertag(*openapiclient.NewUserShort("Pk_example"), float32(123), float32(123))} // []Usertag |  (optional)
    location := *openapiclient.NewLocation("Name_example") // Location |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaEditMediaEditPost(context.Background()).Sessionid(sessionid).MediaId(mediaId).Caption(caption).Title(title).Usertags(usertags).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaEditMediaEditPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaEditMediaEditPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaEditMediaEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaEditMediaEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaId** | **string** |  | 
 **caption** | **string** |  | 
 **title** | **string** |  | [default to &quot;&quot;]
 **usertags** | [**[]Usertag**](Usertag.md) |  | 
 **location** | [**Location**](Location.md) |  | 

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


## MediaIdMediaIdGet

> interface{} MediaIdMediaIdGet(ctx).MediaPk(mediaPk).Execute()

Media Id



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
    mediaPk := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaIdMediaIdGet(context.Background()).MediaPk(mediaPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaIdMediaIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaIdMediaIdGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaIdMediaIdGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaIdMediaIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaPk** | **int32** |  | 

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


## MediaInfoMediaInfoPost

> Media MediaInfoMediaInfoPost(ctx).Sessionid(sessionid).Pk(pk).UseCache(useCache).Execute()

Media Info



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
    pk := int32(56) // int32 | 
    useCache := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaInfoMediaInfoPost(context.Background()).Sessionid(sessionid).Pk(pk).UseCache(useCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaInfoMediaInfoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaInfoMediaInfoPost`: Media
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaInfoMediaInfoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaInfoMediaInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **pk** | **int32** |  | 
 **useCache** | **bool** |  | [default to true]

### Return type

[**Media**](Media.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaLikeMediaLikePost

> bool MediaLikeMediaLikePost(ctx).Sessionid(sessionid).MediaId(mediaId).Revert(revert).Execute()

Media Like



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
    mediaId := "mediaId_example" // string | 
    revert := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaLikeMediaLikePost(context.Background()).Sessionid(sessionid).MediaId(mediaId).Revert(revert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaLikeMediaLikePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaLikeMediaLikePost`: bool
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaLikeMediaLikePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaLikeMediaLikePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaId** | **string** |  | 
 **revert** | **bool** |  | [default to false]

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


## MediaLikersMediaLikersPost

> []UserShort MediaLikersMediaLikersPost(ctx).Sessionid(sessionid).MediaId(mediaId).Execute()

Media Likers



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
    mediaId := "mediaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaLikersMediaLikersPost(context.Background()).Sessionid(sessionid).MediaId(mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaLikersMediaLikersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaLikersMediaLikersPost`: []UserShort
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaLikersMediaLikersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaLikersMediaLikersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaId** | **string** |  | 

### Return type

[**[]UserShort**](UserShort.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaOembedMediaOembedPost

> map[string]interface{} MediaOembedMediaOembedPost(ctx).Sessionid(sessionid).Url(url).Execute()

Media Oembed



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaOembedMediaOembedPost(context.Background()).Sessionid(sessionid).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaOembedMediaOembedPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaOembedMediaOembedPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaOembedMediaOembedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaOembedMediaOembedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **url** | **string** |  | 

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


## MediaPkFromCodeMediaPkFromCodeGet

> interface{} MediaPkFromCodeMediaPkFromCodeGet(ctx).Code(code).Execute()

Media Pk From Code



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
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaPkFromCodeMediaPkFromCodeGet(context.Background()).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaPkFromCodeMediaPkFromCodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaPkFromCodeMediaPkFromCodeGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaPkFromCodeMediaPkFromCodeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaPkFromCodeMediaPkFromCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 

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


## MediaPkFromUrlMediaPkFromUrlGet

> interface{} MediaPkFromUrlMediaPkFromUrlGet(ctx).Url(url).Execute()

Media Pk From Url



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
    resp, r, err := apiClient.MediaApi.MediaPkFromUrlMediaPkFromUrlGet(context.Background()).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaPkFromUrlMediaPkFromUrlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaPkFromUrlMediaPkFromUrlGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaPkFromUrlMediaPkFromUrlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaPkFromUrlMediaPkFromUrlGetRequest struct via the builder pattern


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


## MediaPkMediaPkGet

> interface{} MediaPkMediaPkGet(ctx).MediaId(mediaId).Execute()

Media Pk



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
    mediaId := "mediaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaPkMediaPkGet(context.Background()).MediaId(mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaPkMediaPkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaPkMediaPkGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaPkMediaPkGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaPkMediaPkGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string** |  | 

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


## MediaSeenMediaSeenPost

> bool MediaSeenMediaSeenPost(ctx).Sessionid(sessionid).MediaIds(mediaIds).SkippedMediaIds(skippedMediaIds).Execute()

Media Seen



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
    mediaIds := []string{"Inner_example"} // []string | 
    skippedMediaIds := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaSeenMediaSeenPost(context.Background()).Sessionid(sessionid).MediaIds(mediaIds).SkippedMediaIds(skippedMediaIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaSeenMediaSeenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaSeenMediaSeenPost`: bool
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaSeenMediaSeenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaSeenMediaSeenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaIds** | **[]string** |  | 
 **skippedMediaIds** | **[]string** |  | 

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


## MediaUnarchiveMediaUnarchivePost

> bool MediaUnarchiveMediaUnarchivePost(ctx).Sessionid(sessionid).MediaId(mediaId).Execute()

Media Unarchive



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
    mediaId := "mediaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaUnarchiveMediaUnarchivePost(context.Background()).Sessionid(sessionid).MediaId(mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUnarchiveMediaUnarchivePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUnarchiveMediaUnarchivePost`: bool
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUnarchiveMediaUnarchivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUnarchiveMediaUnarchivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaId** | **string** |  | 

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


## MediaUnlikeMediaUnlikePost

> bool MediaUnlikeMediaUnlikePost(ctx).Sessionid(sessionid).MediaId(mediaId).Execute()

Media Unlike



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
    mediaId := "mediaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaUnlikeMediaUnlikePost(context.Background()).Sessionid(sessionid).MediaId(mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUnlikeMediaUnlikePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUnlikeMediaUnlikePost`: bool
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUnlikeMediaUnlikePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUnlikeMediaUnlikePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaId** | **string** |  | 

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


## MediaUserMediaUserPost

> UserShort MediaUserMediaUserPost(ctx).Sessionid(sessionid).MediaPk(mediaPk).Execute()

Media User



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
    resp, r, err := apiClient.MediaApi.MediaUserMediaUserPost(context.Background()).Sessionid(sessionid).MediaPk(mediaPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUserMediaUserPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUserMediaUserPost`: UserShort
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUserMediaUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUserMediaUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **mediaPk** | **int32** |  | 

### Return type

[**UserShort**](UserShort.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMediasMediaUserMediasPost

> []Media UserMediasMediaUserMediasPost(ctx).Sessionid(sessionid).UserId(userId).Amount(amount).Execute()

User Medias



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
    userId := int32(56) // int32 | 
    amount := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.UserMediasMediaUserMediasPost(context.Background()).Sessionid(sessionid).UserId(userId).Amount(amount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.UserMediasMediaUserMediasPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserMediasMediaUserMediasPost`: []Media
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.UserMediasMediaUserMediasPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserMediasMediaUserMediasPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **int32** |  | 
 **amount** | **int32** |  | [default to 50]

### Return type

[**[]Media**](Media.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

