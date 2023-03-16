# \SpeechSynthesisApi

All URIs are relative to *<http://localhost:50021>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancellableSynthesisCancellableSynthesisPost**](SpeechSynthesisApi.md#CancellableSynthesisCancellableSynthesisPost) | **Post** /cancellable_synthesis | 音声合成する（キャンセル可能）
[**MorphableTargetsMorphableTargetsPost**](SpeechSynthesisApi.md#MorphableTargetsMorphableTargetsPost) | **Post** /morphable_targets | 指定した話者に対してエンジン内の話者がモーフィングが可能か判定する
[**MultiSynthesisMultiSynthesisPost**](SpeechSynthesisApi.md#MultiSynthesisMultiSynthesisPost) | **Post** /multi_synthesis | 複数まとめて音声合成する
[**SynthesisMorphingSynthesisMorphingPost**](SpeechSynthesisApi.md#SynthesisMorphingSynthesisMorphingPost) | **Post** /synthesis_morphing | 2人の話者でモーフィングした音声を合成する
[**SynthesisSynthesisPost**](SpeechSynthesisApi.md#SynthesisSynthesisPost) | **Post** /synthesis | 音声合成する

## CancellableSynthesisCancellableSynthesisPost

> *os.File CancellableSynthesisCancellableSynthesisPost(ctx).Speaker(speaker).AudioQuery(audioQuery).CoreVersion(coreVersion).Execute()

音声合成する（キャンセル可能）

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 |
    audioQuery := *openapiclient.NewAudioQuery([]openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))}, float32(123), float32(123), float32(123), float32(123), float32(123), float32(123), int32(123), false) // AudioQuery |
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpeechSynthesisApi.CancellableSynthesisCancellableSynthesisPost(context.Background()).Speaker(speaker).AudioQuery(audioQuery).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpeechSynthesisApi.CancellableSynthesisCancellableSynthesisPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancellableSynthesisCancellableSynthesisPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SpeechSynthesisApi.CancellableSynthesisCancellableSynthesisPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCancellableSynthesisCancellableSynthesisPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speaker** | **int32** |  |
 **audioQuery** | [**AudioQuery**](AudioQuery.md) |  |
 **coreVersion** | **string** |  |

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: audio/wav, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MorphableTargetsMorphableTargetsPost

> []map[string]MorphableTargetInfo MorphableTargetsMorphableTargetsPost(ctx).RequestBody(requestBody).CoreVersion(coreVersion).Execute()

指定した話者に対してエンジン内の話者がモーフィングが可能か判定する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shikanime/go-voicevox"
)

func main() {
    requestBody := []int32{int32(123)} // []int32 |
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpeechSynthesisApi.MorphableTargetsMorphableTargetsPost(context.Background()).RequestBody(requestBody).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpeechSynthesisApi.MorphableTargetsMorphableTargetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MorphableTargetsMorphableTargetsPost`: []map[string]MorphableTargetInfo
    fmt.Fprintf(os.Stdout, "Response from `SpeechSynthesisApi.MorphableTargetsMorphableTargetsPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiMorphableTargetsMorphableTargetsPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** |  |
 **coreVersion** | **string** |  |

### Return type

[**[]map[string]MorphableTargetInfo**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MultiSynthesisMultiSynthesisPost

> *os.File MultiSynthesisMultiSynthesisPost(ctx).Speaker(speaker).AudioQuery(audioQuery).CoreVersion(coreVersion).Execute()

複数まとめて音声合成する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 |
    audioQuery := []openapiclient.AudioQuery{*openapiclient.NewAudioQuery([]openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))}, float32(123), float32(123), float32(123), float32(123), float32(123), float32(123), int32(123), false)} // []AudioQuery |
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpeechSynthesisApi.MultiSynthesisMultiSynthesisPost(context.Background()).Speaker(speaker).AudioQuery(audioQuery).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpeechSynthesisApi.MultiSynthesisMultiSynthesisPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MultiSynthesisMultiSynthesisPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SpeechSynthesisApi.MultiSynthesisMultiSynthesisPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiMultiSynthesisMultiSynthesisPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speaker** | **int32** |  |
 **audioQuery** | [**[]AudioQuery**](AudioQuery.md) |  |
 **coreVersion** | **string** |  |

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SynthesisMorphingSynthesisMorphingPost

> *os.File SynthesisMorphingSynthesisMorphingPost(ctx).BaseSpeaker(baseSpeaker).TargetSpeaker(targetSpeaker).MorphRate(morphRate).AudioQuery(audioQuery).CoreVersion(coreVersion).Execute()

2人の話者でモーフィングした音声を合成する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shikanime/go-voicevox"
)

func main() {
    baseSpeaker := int32(56) // int32 |
    targetSpeaker := int32(56) // int32 |
    morphRate := float32(8.14) // float32 |
    audioQuery := *openapiclient.NewAudioQuery([]openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))}, float32(123), float32(123), float32(123), float32(123), float32(123), float32(123), int32(123), false) // AudioQuery |
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpeechSynthesisApi.SynthesisMorphingSynthesisMorphingPost(context.Background()).BaseSpeaker(baseSpeaker).TargetSpeaker(targetSpeaker).MorphRate(morphRate).AudioQuery(audioQuery).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpeechSynthesisApi.SynthesisMorphingSynthesisMorphingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SynthesisMorphingSynthesisMorphingPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SpeechSynthesisApi.SynthesisMorphingSynthesisMorphingPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSynthesisMorphingSynthesisMorphingPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseSpeaker** | **int32** |  |
 **targetSpeaker** | **int32** |  |
 **morphRate** | **float32** |  |
 **audioQuery** | [**AudioQuery**](AudioQuery.md) |  |
 **coreVersion** | **string** |  |

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: audio/wav, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SynthesisSynthesisPost

> *os.File SynthesisSynthesisPost(ctx).Speaker(speaker).AudioQuery(audioQuery).EnableInterrogativeUpspeak(enableInterrogativeUpspeak).CoreVersion(coreVersion).Execute()

音声合成する

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 |
    audioQuery := *openapiclient.NewAudioQuery([]openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))}, float32(123), float32(123), float32(123), float32(123), float32(123), float32(123), int32(123), false) // AudioQuery |
    enableInterrogativeUpspeak := true // bool | 疑問系のテキストが与えられたら語尾を自動調整する (optional) (default to true)
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpeechSynthesisApi.SynthesisSynthesisPost(context.Background()).Speaker(speaker).AudioQuery(audioQuery).EnableInterrogativeUpspeak(enableInterrogativeUpspeak).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpeechSynthesisApi.SynthesisSynthesisPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SynthesisSynthesisPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SpeechSynthesisApi.SynthesisSynthesisPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSynthesisSynthesisPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speaker** | **int32** |  |
 **audioQuery** | [**AudioQuery**](AudioQuery.md) |  |
 **enableInterrogativeUpspeak** | **bool** | 疑問系のテキストが与えられたら語尾を自動調整する | [default to true]
 **coreVersion** | **string** |  |

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: audio/wav, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
