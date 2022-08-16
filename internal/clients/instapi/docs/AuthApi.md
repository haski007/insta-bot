# \AuthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthLoginAuthLoginPost**](AuthApi.md#AuthLoginAuthLoginPost) | **Post** /auth/login | Auth Login
[**AuthReloginAuthReloginPost**](AuthApi.md#AuthReloginAuthReloginPost) | **Post** /auth/relogin | Auth Relogin
[**SettingsGetAuthSettingsGetGet**](AuthApi.md#SettingsGetAuthSettingsGetGet) | **Get** /auth/settings/get | Settings Get
[**SettingsSetAuthSettingsSetPost**](AuthApi.md#SettingsSetAuthSettingsSetPost) | **Post** /auth/settings/set | Settings Set
[**TimelineFeedAuthTimelineFeedGet**](AuthApi.md#TimelineFeedAuthTimelineFeedGet) | **Get** /auth/timeline_feed | Timeline Feed



## AuthLoginAuthLoginPost

> interface{} AuthLoginAuthLoginPost(ctx).Username(username).Password(password).VerificationCode(verificationCode).Proxy(proxy).Locale(locale).Timezone(timezone).Execute()

Auth Login



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
    username := "username_example" // string | 
    password := "password_example" // string | 
    verificationCode := "verificationCode_example" // string |  (optional) (default to "")
    proxy := "proxy_example" // string |  (optional) (default to "")
    locale := "locale_example" // string |  (optional) (default to "")
    timezone := "timezone_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.AuthLoginAuthLoginPost(context.Background()).Username(username).Password(password).VerificationCode(verificationCode).Proxy(proxy).Locale(locale).Timezone(timezone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthLoginAuthLoginPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthLoginAuthLoginPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthLoginAuthLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthLoginAuthLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** |  | 
 **password** | **string** |  | 
 **verificationCode** | **string** |  | [default to &quot;&quot;]
 **proxy** | **string** |  | [default to &quot;&quot;]
 **locale** | **string** |  | [default to &quot;&quot;]
 **timezone** | **string** |  | [default to &quot;&quot;]

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


## AuthReloginAuthReloginPost

> interface{} AuthReloginAuthReloginPost(ctx).Sessionid(sessionid).Execute()

Auth Relogin



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
    resp, r, err := apiClient.AuthApi.AuthReloginAuthReloginPost(context.Background()).Sessionid(sessionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthReloginAuthReloginPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthReloginAuthReloginPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthReloginAuthReloginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthReloginAuthReloginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 

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


## SettingsGetAuthSettingsGetGet

> interface{} SettingsGetAuthSettingsGetGet(ctx).Sessionid(sessionid).Execute()

Settings Get



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
    resp, r, err := apiClient.AuthApi.SettingsGetAuthSettingsGetGet(context.Background()).Sessionid(sessionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.SettingsGetAuthSettingsGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsGetAuthSettingsGetGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.SettingsGetAuthSettingsGetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettingsGetAuthSettingsGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 

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


## SettingsSetAuthSettingsSetPost

> interface{} SettingsSetAuthSettingsSetPost(ctx).Settings(settings).Sessionid(sessionid).Execute()

Settings Set



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
    settings := "settings_example" // string | 
    sessionid := "sessionid_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.SettingsSetAuthSettingsSetPost(context.Background()).Settings(settings).Sessionid(sessionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.SettingsSetAuthSettingsSetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsSetAuthSettingsSetPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.SettingsSetAuthSettingsSetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSetAuthSettingsSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settings** | **string** |  | 
 **sessionid** | **string** |  | [default to &quot;&quot;]

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


## TimelineFeedAuthTimelineFeedGet

> interface{} TimelineFeedAuthTimelineFeedGet(ctx).Sessionid(sessionid).Execute()

Timeline Feed



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
    resp, r, err := apiClient.AuthApi.TimelineFeedAuthTimelineFeedGet(context.Background()).Sessionid(sessionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.TimelineFeedAuthTimelineFeedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TimelineFeedAuthTimelineFeedGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.TimelineFeedAuthTimelineFeedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimelineFeedAuthTimelineFeedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionid** | **string** |  | 

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

