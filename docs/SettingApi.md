# \SettingApi

All URIs are relative to *http://localhost:50021*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingGetSettingGet**](SettingApi.md#SettingGetSettingGet) | **Get** /setting | Setting Get
[**SettingPostSettingPost**](SettingApi.md#SettingPostSettingPost) | **Post** /setting | Setting Post



## SettingGetSettingGet

> string SettingGetSettingGet(ctx).Execute()

Setting Get

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/infinity-blackhole/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingApi.SettingGetSettingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingApi.SettingGetSettingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingGetSettingGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SettingApi.SettingGetSettingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingGetSettingGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingPostSettingPost

> string SettingPostSettingPost(ctx).CorsPolicyMode(corsPolicyMode).AllowOrigin(allowOrigin).Execute()

Setting Post

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/infinity-blackhole/go-voicevox"
)

func main() {
    corsPolicyMode := "corsPolicyMode_example" // string |  (optional)
    allowOrigin := "allowOrigin_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingApi.SettingPostSettingPost(context.Background()).CorsPolicyMode(corsPolicyMode).AllowOrigin(allowOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingApi.SettingPostSettingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingPostSettingPost`: string
    fmt.Fprintf(os.Stdout, "Response from `SettingApi.SettingPostSettingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettingPostSettingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corsPolicyMode** | **string** |  | 
 **allowOrigin** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

