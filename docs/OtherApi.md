# \OtherApi

All URIs are relative to *http://localhost:50021*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPresetAddPresetPost**](OtherApi.md#AddPresetAddPresetPost) | **Post** /add_preset | Add Preset
[**ConnectWavesConnectWavesPost**](OtherApi.md#ConnectWavesConnectWavesPost) | **Post** /connect_waves | base64エンコードされた複数のwavデータを一つに結合する
[**CoreVersionsCoreVersionsGet**](OtherApi.md#CoreVersionsCoreVersionsGet) | **Get** /core_versions | Core Versions
[**DeletePresetDeletePresetPost**](OtherApi.md#DeletePresetDeletePresetPost) | **Post** /delete_preset | Delete Preset
[**EngineManifestEngineManifestGet**](OtherApi.md#EngineManifestEngineManifestGet) | **Get** /engine_manifest | Engine Manifest
[**GetPresetsPresetsGet**](OtherApi.md#GetPresetsPresetsGet) | **Get** /presets | Get Presets
[**InitializeSpeakerInitializeSpeakerPost**](OtherApi.md#InitializeSpeakerInitializeSpeakerPost) | **Post** /initialize_speaker | Initialize Speaker
[**IsInitializedSpeakerIsInitializedSpeakerGet**](OtherApi.md#IsInitializedSpeakerIsInitializedSpeakerGet) | **Get** /is_initialized_speaker | Is Initialized Speaker
[**SpeakerInfoSpeakerInfoGet**](OtherApi.md#SpeakerInfoSpeakerInfoGet) | **Get** /speaker_info | Speaker Info
[**SpeakersSpeakersGet**](OtherApi.md#SpeakersSpeakersGet) | **Get** /speakers | Speakers
[**SupportedDevicesSupportedDevicesGet**](OtherApi.md#SupportedDevicesSupportedDevicesGet) | **Get** /supported_devices | Supported Devices
[**UpdatePresetUpdatePresetPost**](OtherApi.md#UpdatePresetUpdatePresetPost) | **Post** /update_preset | Update Preset
[**VersionVersionGet**](OtherApi.md#VersionVersionGet) | **Get** /version | Version



## AddPresetAddPresetPost

> int32 AddPresetAddPresetPost(ctx).Preset(preset).Execute()

Add Preset



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
    preset := *openapiclient.NewPreset(int32(123), "Name_example", "SpeakerUuid_example", int32(123), float32(123), float32(123), float32(123), float32(123), float32(123), float32(123)) // Preset | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherApi.AddPresetAddPresetPost(context.Background()).Preset(preset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.AddPresetAddPresetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPresetAddPresetPost`: int32
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.AddPresetAddPresetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPresetAddPresetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preset** | [**Preset**](Preset.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectWavesConnectWavesPost

> *os.File ConnectWavesConnectWavesPost(ctx).RequestBody(requestBody).Execute()

base64エンコードされた複数のwavデータを一つに結合する



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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherApi.ConnectWavesConnectWavesPost(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.ConnectWavesConnectWavesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectWavesConnectWavesPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.ConnectWavesConnectWavesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectWavesConnectWavesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

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


## CoreVersionsCoreVersionsGet

> []string CoreVersionsCoreVersionsGet(ctx).Execute()

Core Versions

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
    resp, r, err := apiClient.OtherApi.CoreVersionsCoreVersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.CoreVersionsCoreVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreVersionsCoreVersionsGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.CoreVersionsCoreVersionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCoreVersionsCoreVersionsGetRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePresetDeletePresetPost

> DeletePresetDeletePresetPost(ctx).Id(id).Execute()

Delete Preset



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OtherApi.DeletePresetDeletePresetPost(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.DeletePresetDeletePresetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePresetDeletePresetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineManifestEngineManifestGet

> EngineManifest EngineManifestEngineManifestGet(ctx).Execute()

Engine Manifest

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
    resp, r, err := apiClient.OtherApi.EngineManifestEngineManifestGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.EngineManifestEngineManifestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EngineManifestEngineManifestGet`: EngineManifest
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.EngineManifestEngineManifestGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEngineManifestEngineManifestGetRequest struct via the builder pattern


### Return type

[**EngineManifest**](EngineManifest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPresetsPresetsGet

> []Preset GetPresetsPresetsGet(ctx).Execute()

Get Presets



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
    resp, r, err := apiClient.OtherApi.GetPresetsPresetsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.GetPresetsPresetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPresetsPresetsGet`: []Preset
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.GetPresetsPresetsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPresetsPresetsGetRequest struct via the builder pattern


### Return type

[**[]Preset**](Preset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeSpeakerInitializeSpeakerPost

> InitializeSpeakerInitializeSpeakerPost(ctx).Speaker(speaker).SkipReinit(skipReinit).CoreVersion(coreVersion).Execute()

Initialize Speaker



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
    skipReinit := true // bool | 既に初期化済みの話者の再初期化をスキップするかどうか (optional) (default to false)
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OtherApi.InitializeSpeakerInitializeSpeakerPost(context.Background()).Speaker(speaker).SkipReinit(skipReinit).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.InitializeSpeakerInitializeSpeakerPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitializeSpeakerInitializeSpeakerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speaker** | **int32** |  | 
 **skipReinit** | **bool** | 既に初期化済みの話者の再初期化をスキップするかどうか | [default to false]
 **coreVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsInitializedSpeakerIsInitializedSpeakerGet

> bool IsInitializedSpeakerIsInitializedSpeakerGet(ctx).Speaker(speaker).CoreVersion(coreVersion).Execute()

Is Initialized Speaker



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
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherApi.IsInitializedSpeakerIsInitializedSpeakerGet(context.Background()).Speaker(speaker).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.IsInitializedSpeakerIsInitializedSpeakerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsInitializedSpeakerIsInitializedSpeakerGet`: bool
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.IsInitializedSpeakerIsInitializedSpeakerGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsInitializedSpeakerIsInitializedSpeakerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speaker** | **int32** |  | 
 **coreVersion** | **string** |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeakerInfoSpeakerInfoGet

> SpeakerInfo SpeakerInfoSpeakerInfoGet(ctx).SpeakerUuid(speakerUuid).CoreVersion(coreVersion).Execute()

Speaker Info



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
    speakerUuid := "speakerUuid_example" // string | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherApi.SpeakerInfoSpeakerInfoGet(context.Background()).SpeakerUuid(speakerUuid).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.SpeakerInfoSpeakerInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpeakerInfoSpeakerInfoGet`: SpeakerInfo
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.SpeakerInfoSpeakerInfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpeakerInfoSpeakerInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speakerUuid** | **string** |  | 
 **coreVersion** | **string** |  | 

### Return type

[**SpeakerInfo**](SpeakerInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeakersSpeakersGet

> []Speaker SpeakersSpeakersGet(ctx).CoreVersion(coreVersion).Execute()

Speakers

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
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherApi.SpeakersSpeakersGet(context.Background()).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.SpeakersSpeakersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpeakersSpeakersGet`: []Speaker
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.SpeakersSpeakersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpeakersSpeakersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coreVersion** | **string** |  | 

### Return type

[**[]Speaker**](Speaker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportedDevicesSupportedDevicesGet

> SupportedDevicesInfo SupportedDevicesSupportedDevicesGet(ctx).CoreVersion(coreVersion).Execute()

Supported Devices

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
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherApi.SupportedDevicesSupportedDevicesGet(context.Background()).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.SupportedDevicesSupportedDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportedDevicesSupportedDevicesGet`: SupportedDevicesInfo
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.SupportedDevicesSupportedDevicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupportedDevicesSupportedDevicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coreVersion** | **string** |  | 

### Return type

[**SupportedDevicesInfo**](SupportedDevicesInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePresetUpdatePresetPost

> int32 UpdatePresetUpdatePresetPost(ctx).Preset(preset).Execute()

Update Preset



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
    preset := *openapiclient.NewPreset(int32(123), "Name_example", "SpeakerUuid_example", int32(123), float32(123), float32(123), float32(123), float32(123), float32(123), float32(123)) // Preset | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherApi.UpdatePresetUpdatePresetPost(context.Background()).Preset(preset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.UpdatePresetUpdatePresetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePresetUpdatePresetPost`: int32
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.UpdatePresetUpdatePresetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePresetUpdatePresetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preset** | [**Preset**](Preset.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionVersionGet

> interface{} VersionVersionGet(ctx).Execute()

Version

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
    resp, r, err := apiClient.OtherApi.VersionVersionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.VersionVersionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VersionVersionGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.VersionVersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionVersionGetRequest struct via the builder pattern


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

