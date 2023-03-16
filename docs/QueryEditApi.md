# \QueryEditApi

All URIs are relative to *<http://localhost:50021>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccentPhrasesAccentPhrasesPost**](QueryEditApi.md#AccentPhrasesAccentPhrasesPost) | **Post** /accent_phrases | テキストからアクセント句を得る
[**MoraDataMoraDataPost**](QueryEditApi.md#MoraDataMoraDataPost) | **Post** /mora_data | アクセント句から音高・音素長を得る
[**MoraLengthMoraLengthPost**](QueryEditApi.md#MoraLengthMoraLengthPost) | **Post** /mora_length | アクセント句から音素長を得る
[**MoraPitchMoraPitchPost**](QueryEditApi.md#MoraPitchMoraPitchPost) | **Post** /mora_pitch | アクセント句から音高を得る

## AccentPhrasesAccentPhrasesPost

> []AccentPhrase AccentPhrasesAccentPhrasesPost(ctx).Text(text).Speaker(speaker).IsKana(isKana).CoreVersion(coreVersion).Execute()

テキストからアクセント句を得る

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
    text := "text_example" // string | 
    speaker := int32(56) // int32 | 
    isKana := true // bool |  (optional) (default to false)
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryEditApi.AccentPhrasesAccentPhrasesPost(context.Background()).Text(text).Speaker(speaker).IsKana(isKana).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryEditApi.AccentPhrasesAccentPhrasesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccentPhrasesAccentPhrasesPost`: []AccentPhrase
    fmt.Fprintf(os.Stdout, "Response from `QueryEditApi.AccentPhrasesAccentPhrasesPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiAccentPhrasesAccentPhrasesPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** |  |
 **speaker** | **int32** |  |
 **isKana** | **bool** |  | [default to false]
 **coreVersion** | **string** |  |

### Return type

[**[]AccentPhrase**](AccentPhrase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MoraDataMoraDataPost

> []AccentPhrase MoraDataMoraDataPost(ctx).Speaker(speaker).AccentPhrase(accentPhrase).CoreVersion(coreVersion).Execute()

アクセント句から音高・音素長を得る

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
    speaker := int32(56) // int32 | 
    accentPhrase := []openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))} // []AccentPhrase | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryEditApi.MoraDataMoraDataPost(context.Background()).Speaker(speaker).AccentPhrase(accentPhrase).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryEditApi.MoraDataMoraDataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoraDataMoraDataPost`: []AccentPhrase
    fmt.Fprintf(os.Stdout, "Response from `QueryEditApi.MoraDataMoraDataPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiMoraDataMoraDataPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speaker** | **int32** |  |
 **accentPhrase** | [**[]AccentPhrase**](AccentPhrase.md) |  |
 **coreVersion** | **string** |  |

### Return type

[**[]AccentPhrase**](AccentPhrase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MoraLengthMoraLengthPost

> []AccentPhrase MoraLengthMoraLengthPost(ctx).Speaker(speaker).AccentPhrase(accentPhrase).CoreVersion(coreVersion).Execute()

アクセント句から音素長を得る

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
    speaker := int32(56) // int32 | 
    accentPhrase := []openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))} // []AccentPhrase | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryEditApi.MoraLengthMoraLengthPost(context.Background()).Speaker(speaker).AccentPhrase(accentPhrase).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryEditApi.MoraLengthMoraLengthPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoraLengthMoraLengthPost`: []AccentPhrase
    fmt.Fprintf(os.Stdout, "Response from `QueryEditApi.MoraLengthMoraLengthPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiMoraLengthMoraLengthPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speaker** | **int32** |  |
 **accentPhrase** | [**[]AccentPhrase**](AccentPhrase.md) |  |
 **coreVersion** | **string** |  |

### Return type

[**[]AccentPhrase**](AccentPhrase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MoraPitchMoraPitchPost

> []AccentPhrase MoraPitchMoraPitchPost(ctx).Speaker(speaker).AccentPhrase(accentPhrase).CoreVersion(coreVersion).Execute()

アクセント句から音高を得る

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
    speaker := int32(56) // int32 | 
    accentPhrase := []openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))} // []AccentPhrase | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryEditApi.MoraPitchMoraPitchPost(context.Background()).Speaker(speaker).AccentPhrase(accentPhrase).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryEditApi.MoraPitchMoraPitchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoraPitchMoraPitchPost`: []AccentPhrase
    fmt.Fprintf(os.Stdout, "Response from `QueryEditApi.MoraPitchMoraPitchPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiMoraPitchMoraPitchPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speaker** | **int32** |  |
 **accentPhrase** | [**[]AccentPhrase**](AccentPhrase.md) |  |
 **coreVersion** | **string** |  |

### Return type

[**[]AccentPhrase**](AccentPhrase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
