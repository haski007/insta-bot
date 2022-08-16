# \UserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MutePostsFromFollowUserMutePostsFromFollowPost**](UserApi.md#MutePostsFromFollowUserMutePostsFromFollowPost) | **Post** /user/mute_posts_from_follow | Mute Posts From Follow
[**MuteStoriesFromFollowUserMuteStoriesFromFollowPost**](UserApi.md#MuteStoriesFromFollowUserMuteStoriesFromFollowPost) | **Post** /user/mute_stories_from_follow | Mute Stories From Follow
[**UnmutePostsFromFollowUserUnmutePostsFromFollowPost**](UserApi.md#UnmutePostsFromFollowUserUnmutePostsFromFollowPost) | **Post** /user/unmute_posts_from_follow | Unmute Posts From Follow
[**UnmuteStoriesFromFollowUserUnmuteStoriesFromFollowPost**](UserApi.md#UnmuteStoriesFromFollowUserUnmuteStoriesFromFollowPost) | **Post** /user/unmute_stories_from_follow | Unmute Stories From Follow
[**UserFollowUserFollowPost**](UserApi.md#UserFollowUserFollowPost) | **Post** /user/follow | User Follow
[**UserFollowersUserFollowersPost**](UserApi.md#UserFollowersUserFollowersPost) | **Post** /user/followers | User Followers
[**UserFollowingUserFollowingPost**](UserApi.md#UserFollowingUserFollowingPost) | **Post** /user/following | User Following
[**UserIdFromUsernameUserIdFromUsernamePost**](UserApi.md#UserIdFromUsernameUserIdFromUsernamePost) | **Post** /user/id_from_username | User Id From Username
[**UserInfoByUsernameUserInfoByUsernamePost**](UserApi.md#UserInfoByUsernameUserInfoByUsernamePost) | **Post** /user/info_by_username | User Info By Username
[**UserInfoUserInfoPost**](UserApi.md#UserInfoUserInfoPost) | **Post** /user/info | User Info
[**UserRemoveFollowerUserRemoveFollowerPost**](UserApi.md#UserRemoveFollowerUserRemoveFollowerPost) | **Post** /user/remove_follower | User Remove Follower
[**UserUnfollowUserUnfollowPost**](UserApi.md#UserUnfollowUserUnfollowPost) | **Post** /user/unfollow | User Unfollow
[**UsernameFromUserIdUserUsernameFromIdPost**](UserApi.md#UsernameFromUserIdUserUsernameFromIdPost) | **Post** /user/username_from_id | Username From User Id



## MutePostsFromFollowUserMutePostsFromFollowPost

> bool MutePostsFromFollowUserMutePostsFromFollowPost(ctx).Sessionid(sessionid).UserId(userId).Revert(revert).Execute()

Mute Posts From Follow



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
    revert := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.MutePostsFromFollowUserMutePostsFromFollowPost(context.Background()).Sessionid(sessionid).UserId(userId).Revert(revert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.MutePostsFromFollowUserMutePostsFromFollowPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutePostsFromFollowUserMutePostsFromFollowPost`: bool
    fmt.Fprintf(os.Stdout, "Response from `UserApi.MutePostsFromFollowUserMutePostsFromFollowPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMutePostsFromFollowUserMutePostsFromFollowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **int32** |  | 
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


## MuteStoriesFromFollowUserMuteStoriesFromFollowPost

> bool MuteStoriesFromFollowUserMuteStoriesFromFollowPost(ctx).Sessionid(sessionid).UserId(userId).Revert(revert).Execute()

Mute Stories From Follow



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
    revert := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.MuteStoriesFromFollowUserMuteStoriesFromFollowPost(context.Background()).Sessionid(sessionid).UserId(userId).Revert(revert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.MuteStoriesFromFollowUserMuteStoriesFromFollowPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MuteStoriesFromFollowUserMuteStoriesFromFollowPost`: bool
    fmt.Fprintf(os.Stdout, "Response from `UserApi.MuteStoriesFromFollowUserMuteStoriesFromFollowPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMuteStoriesFromFollowUserMuteStoriesFromFollowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **int32** |  | 
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


## UnmutePostsFromFollowUserUnmutePostsFromFollowPost

> bool UnmutePostsFromFollowUserUnmutePostsFromFollowPost(ctx).Sessionid(sessionid).UserId(userId).Execute()

Unmute Posts From Follow



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UnmutePostsFromFollowUserUnmutePostsFromFollowPost(context.Background()).Sessionid(sessionid).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UnmutePostsFromFollowUserUnmutePostsFromFollowPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnmutePostsFromFollowUserUnmutePostsFromFollowPost`: bool
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UnmutePostsFromFollowUserUnmutePostsFromFollowPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnmutePostsFromFollowUserUnmutePostsFromFollowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **int32** |  | 

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


## UnmuteStoriesFromFollowUserUnmuteStoriesFromFollowPost

> bool UnmuteStoriesFromFollowUserUnmuteStoriesFromFollowPost(ctx).Sessionid(sessionid).UserId(userId).Execute()

Unmute Stories From Follow



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UnmuteStoriesFromFollowUserUnmuteStoriesFromFollowPost(context.Background()).Sessionid(sessionid).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UnmuteStoriesFromFollowUserUnmuteStoriesFromFollowPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnmuteStoriesFromFollowUserUnmuteStoriesFromFollowPost`: bool
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UnmuteStoriesFromFollowUserUnmuteStoriesFromFollowPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteStoriesFromFollowUserUnmuteStoriesFromFollowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **int32** |  | 

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


## UserFollowUserFollowPost

> bool UserFollowUserFollowPost(ctx).Sessionid(sessionid).UserId(userId).Execute()

User Follow



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserFollowUserFollowPost(context.Background()).Sessionid(sessionid).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserFollowUserFollowPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserFollowUserFollowPost`: bool
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserFollowUserFollowPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserFollowUserFollowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **int32** |  | 

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


## UserFollowersUserFollowersPost

> map[string]UserShort UserFollowersUserFollowersPost(ctx).Sessionid(sessionid).UserId(userId).UseCache(useCache).Amount(amount).Execute()

User Followers



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
    useCache := true // bool |  (optional) (default to true)
    amount := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserFollowersUserFollowersPost(context.Background()).Sessionid(sessionid).UserId(userId).UseCache(useCache).Amount(amount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserFollowersUserFollowersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserFollowersUserFollowersPost`: map[string]UserShort
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserFollowersUserFollowersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserFollowersUserFollowersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **string** |  | 
 **useCache** | **bool** |  | [default to true]
 **amount** | **int32** |  | [default to 0]

### Return type

[**map[string]UserShort**](UserShort.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserFollowingUserFollowingPost

> map[string]UserShort UserFollowingUserFollowingPost(ctx).Sessionid(sessionid).UserId(userId).UseCache(useCache).Amount(amount).Execute()

User Following



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
    useCache := true // bool |  (optional) (default to true)
    amount := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserFollowingUserFollowingPost(context.Background()).Sessionid(sessionid).UserId(userId).UseCache(useCache).Amount(amount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserFollowingUserFollowingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserFollowingUserFollowingPost`: map[string]UserShort
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserFollowingUserFollowingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserFollowingUserFollowingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **string** |  | 
 **useCache** | **bool** |  | [default to true]
 **amount** | **int32** |  | [default to 0]

### Return type

[**map[string]UserShort**](UserShort.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserIdFromUsernameUserIdFromUsernamePost

> int32 UserIdFromUsernameUserIdFromUsernamePost(ctx).Sessionid(sessionid).Username(username).Execute()

User Id From Username



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
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserIdFromUsernameUserIdFromUsernamePost(context.Background()).Sessionid(sessionid).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserIdFromUsernameUserIdFromUsernamePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserIdFromUsernameUserIdFromUsernamePost`: int32
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserIdFromUsernameUserIdFromUsernamePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserIdFromUsernameUserIdFromUsernamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **username** | **string** |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInfoByUsernameUserInfoByUsernamePost

> User UserInfoByUsernameUserInfoByUsernamePost(ctx).Sessionid(sessionid).Username(username).UseCache(useCache).Execute()

User Info By Username



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
    username := "username_example" // string | 
    useCache := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserInfoByUsernameUserInfoByUsernamePost(context.Background()).Sessionid(sessionid).Username(username).UseCache(useCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserInfoByUsernameUserInfoByUsernamePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserInfoByUsernameUserInfoByUsernamePost`: User
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserInfoByUsernameUserInfoByUsernamePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserInfoByUsernameUserInfoByUsernamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **username** | **string** |  | 
 **useCache** | **bool** |  | [default to true]

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInfoUserInfoPost

> User UserInfoUserInfoPost(ctx).Sessionid(sessionid).UserId(userId).UseCache(useCache).Execute()

User Info



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
    useCache := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserInfoUserInfoPost(context.Background()).Sessionid(sessionid).UserId(userId).UseCache(useCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserInfoUserInfoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserInfoUserInfoPost`: User
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserInfoUserInfoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserInfoUserInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **string** |  | 
 **useCache** | **bool** |  | [default to true]

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserRemoveFollowerUserRemoveFollowerPost

> bool UserRemoveFollowerUserRemoveFollowerPost(ctx).Sessionid(sessionid).UserId(userId).Execute()

User Remove Follower



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserRemoveFollowerUserRemoveFollowerPost(context.Background()).Sessionid(sessionid).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserRemoveFollowerUserRemoveFollowerPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserRemoveFollowerUserRemoveFollowerPost`: bool
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserRemoveFollowerUserRemoveFollowerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserRemoveFollowerUserRemoveFollowerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **int32** |  | 

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


## UserUnfollowUserUnfollowPost

> bool UserUnfollowUserUnfollowPost(ctx).Sessionid(sessionid).UserId(userId).Execute()

User Unfollow



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserUnfollowUserUnfollowPost(context.Background()).Sessionid(sessionid).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserUnfollowUserUnfollowPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUnfollowUserUnfollowPost`: bool
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserUnfollowUserUnfollowPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserUnfollowUserUnfollowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **int32** |  | 

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


## UsernameFromUserIdUserUsernameFromIdPost

> string UsernameFromUserIdUserUsernameFromIdPost(ctx).Sessionid(sessionid).UserId(userId).Execute()

Username From User Id



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UsernameFromUserIdUserUsernameFromIdPost(context.Background()).Sessionid(sessionid).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UsernameFromUserIdUserUsernameFromIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsernameFromUserIdUserUsernameFromIdPost`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UsernameFromUserIdUserUsernameFromIdPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsernameFromUserIdUserUsernameFromIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 
 **userId** | **int32** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

