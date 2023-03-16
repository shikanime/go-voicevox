# \QueryCreationApi

All URIs are relative to *http://localhost:50021*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AudioQueryAudioQueryPost**](QueryCreationApi.md#AudioQueryAudioQueryPost) | **Post** /audio_query | 音声合成用のクエリを作成する
[**AudioQueryFromPresetAudioQueryFromPresetPost**](QueryCreationApi.md#AudioQueryFromPresetAudioQueryFromPresetPost) | **Post** /audio_query_from_preset | 音声合成用のクエリをプリセットを用いて作成する



## AudioQueryAudioQueryPost

> AudioQuery AudioQueryAudioQueryPost(ctx).Text(text).Speaker(speaker).CoreVersion(coreVersion).Execute()

音声合成用のクエリを作成する



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
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryCreationApi.AudioQueryAudioQueryPost(context.Background()).Text(text).Speaker(speaker).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryCreationApi.AudioQueryAudioQueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AudioQueryAudioQueryPost`: AudioQuery
    fmt.Fprintf(os.Stdout, "Response from `QueryCreationApi.AudioQueryAudioQueryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAudioQueryAudioQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** |  | 
 **speaker** | **int32** |  | 
 **coreVersion** | **string** |  | 

### Return type

[**AudioQuery**](AudioQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AudioQueryFromPresetAudioQueryFromPresetPost

> AudioQuery AudioQueryFromPresetAudioQueryFromPresetPost(ctx).Text(text).PresetId(presetId).CoreVersion(coreVersion).Execute()

音声合成用のクエリをプリセットを用いて作成する



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
    presetId := int32(56) // int32 | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryCreationApi.AudioQueryFromPresetAudioQueryFromPresetPost(context.Background()).Text(text).PresetId(presetId).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryCreationApi.AudioQueryFromPresetAudioQueryFromPresetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AudioQueryFromPresetAudioQueryFromPresetPost`: AudioQuery
    fmt.Fprintf(os.Stdout, "Response from `QueryCreationApi.AudioQueryFromPresetAudioQueryFromPresetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAudioQueryFromPresetAudioQueryFromPresetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** |  | 
 **presetId** | **int32** |  | 
 **coreVersion** | **string** |  | 

### Return type

[**AudioQuery**](AudioQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

