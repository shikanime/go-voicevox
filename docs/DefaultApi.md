# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccentPhrasesAccentPhrasesPost**](DefaultApi.md#AccentPhrasesAccentPhrasesPost) | **Post** /accent_phrases | テキストからアクセント句を得る
[**AddPresetAddPresetPost**](DefaultApi.md#AddPresetAddPresetPost) | **Post** /add_preset | Add Preset
[**AddUserDictWordUserDictWordPost**](DefaultApi.md#AddUserDictWordUserDictWordPost) | **Post** /user_dict_word | Add User Dict Word
[**AudioQueryAudioQueryPost**](DefaultApi.md#AudioQueryAudioQueryPost) | **Post** /audio_query | 音声合成用のクエリを作成する
[**AudioQueryFromPresetAudioQueryFromPresetPost**](DefaultApi.md#AudioQueryFromPresetAudioQueryFromPresetPost) | **Post** /audio_query_from_preset | 音声合成用のクエリをプリセットを用いて作成する
[**CancellableSynthesisCancellableSynthesisPost**](DefaultApi.md#CancellableSynthesisCancellableSynthesisPost) | **Post** /cancellable_synthesis | 音声合成する（キャンセル可能）
[**ConnectWavesConnectWavesPost**](DefaultApi.md#ConnectWavesConnectWavesPost) | **Post** /connect_waves | base64エンコードされた複数のwavデータを一つに結合する
[**CoreVersionsCoreVersionsGet**](DefaultApi.md#CoreVersionsCoreVersionsGet) | **Get** /core_versions | Core Versions
[**DeletePresetDeletePresetPost**](DefaultApi.md#DeletePresetDeletePresetPost) | **Post** /delete_preset | Delete Preset
[**DeleteUserDictWordUserDictWordWordUuidDelete**](DefaultApi.md#DeleteUserDictWordUserDictWordWordUuidDelete) | **Delete** /user_dict_word/{word_uuid} | Delete User Dict Word
[**DownloadableLibrariesDownloadableLibrariesGet**](DefaultApi.md#DownloadableLibrariesDownloadableLibrariesGet) | **Get** /downloadable_libraries | Downloadable Libraries
[**EngineManifestEngineManifestGet**](DefaultApi.md#EngineManifestEngineManifestGet) | **Get** /engine_manifest | Engine Manifest
[**GetPresetsPresetsGet**](DefaultApi.md#GetPresetsPresetsGet) | **Get** /presets | Get Presets
[**GetUserDictWordsUserDictGet**](DefaultApi.md#GetUserDictWordsUserDictGet) | **Get** /user_dict | Get User Dict Words
[**ImportUserDictWordsImportUserDictPost**](DefaultApi.md#ImportUserDictWordsImportUserDictPost) | **Post** /import_user_dict | Import User Dict Words
[**InitializeSpeakerInitializeSpeakerPost**](DefaultApi.md#InitializeSpeakerInitializeSpeakerPost) | **Post** /initialize_speaker | Initialize Speaker
[**InstallLibraryInstallLibraryLibraryUuidPost**](DefaultApi.md#InstallLibraryInstallLibraryLibraryUuidPost) | **Post** /install_library/{library_uuid} | Install Library
[**InstalledLibrariesInstalledLibrariesGet**](DefaultApi.md#InstalledLibrariesInstalledLibrariesGet) | **Get** /installed_libraries | Installed Libraries
[**IsInitializedSpeakerIsInitializedSpeakerGet**](DefaultApi.md#IsInitializedSpeakerIsInitializedSpeakerGet) | **Get** /is_initialized_speaker | Is Initialized Speaker
[**MoraDataMoraDataPost**](DefaultApi.md#MoraDataMoraDataPost) | **Post** /mora_data | アクセント句から音高・音素長を得る
[**MoraLengthMoraLengthPost**](DefaultApi.md#MoraLengthMoraLengthPost) | **Post** /mora_length | アクセント句から音素長を得る
[**MoraPitchMoraPitchPost**](DefaultApi.md#MoraPitchMoraPitchPost) | **Post** /mora_pitch | アクセント句から音高を得る
[**MorphableTargetsMorphableTargetsPost**](DefaultApi.md#MorphableTargetsMorphableTargetsPost) | **Post** /morphable_targets | 指定した話者に対してエンジン内の話者がモーフィングが可能か判定する
[**MultiSynthesisMultiSynthesisPost**](DefaultApi.md#MultiSynthesisMultiSynthesisPost) | **Post** /multi_synthesis | 複数まとめて音声合成する
[**RewriteUserDictWordUserDictWordWordUuidPut**](DefaultApi.md#RewriteUserDictWordUserDictWordWordUuidPut) | **Put** /user_dict_word/{word_uuid} | Rewrite User Dict Word
[**SettingGetSettingGet**](DefaultApi.md#SettingGetSettingGet) | **Get** /setting | Setting Get
[**SettingPostSettingPost**](DefaultApi.md#SettingPostSettingPost) | **Post** /setting | Setting Post
[**SpeakerInfoSpeakerInfoGet**](DefaultApi.md#SpeakerInfoSpeakerInfoGet) | **Get** /speaker_info | Speaker Info
[**SpeakersSpeakersGet**](DefaultApi.md#SpeakersSpeakersGet) | **Get** /speakers | Speakers
[**SupportedDevicesSupportedDevicesGet**](DefaultApi.md#SupportedDevicesSupportedDevicesGet) | **Get** /supported_devices | Supported Devices
[**SynthesisMorphingSynthesisMorphingPost**](DefaultApi.md#SynthesisMorphingSynthesisMorphingPost) | **Post** /synthesis_morphing | 2人の話者でモーフィングした音声を合成する
[**SynthesisSynthesisPost**](DefaultApi.md#SynthesisSynthesisPost) | **Post** /synthesis | 音声合成する
[**UpdatePresetUpdatePresetPost**](DefaultApi.md#UpdatePresetUpdatePresetPost) | **Post** /update_preset | Update Preset
[**VersionVersionGet**](DefaultApi.md#VersionVersionGet) | **Get** /version | Version



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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    text := "text_example" // string | 
    speaker := int32(56) // int32 | 
    isKana := true // bool |  (optional) (default to false)
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AccentPhrasesAccentPhrasesPost(context.Background()).Text(text).Speaker(speaker).IsKana(isKana).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AccentPhrasesAccentPhrasesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccentPhrasesAccentPhrasesPost`: []AccentPhrase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AccentPhrasesAccentPhrasesPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    preset := *openapiclient.NewPreset(int32(123), "Name_example", "SpeakerUuid_example", int32(123), float32(123), float32(123), float32(123), float32(123), float32(123), float32(123)) // Preset | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddPresetAddPresetPost(context.Background()).Preset(preset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddPresetAddPresetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPresetAddPresetPost`: int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddPresetAddPresetPost`: %v\n", resp)
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


## AddUserDictWordUserDictWordPost

> string AddUserDictWordUserDictWordPost(ctx).Surface(surface).Pronunciation(pronunciation).AccentType(accentType).WordType(wordType).Priority(priority).Execute()

Add User Dict Word



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    surface := "surface_example" // string | 
    pronunciation := "pronunciation_example" // string | 
    accentType := int32(56) // int32 | 
    wordType := openapiclient.WordTypes("PROPER_NOUN") // WordTypes |  (optional)
    priority := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddUserDictWordUserDictWordPost(context.Background()).Surface(surface).Pronunciation(pronunciation).AccentType(accentType).WordType(wordType).Priority(priority).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddUserDictWordUserDictWordPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserDictWordUserDictWordPost`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddUserDictWordUserDictWordPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserDictWordUserDictWordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **surface** | **string** |  | 
 **pronunciation** | **string** |  | 
 **accentType** | **int32** |  | 
 **wordType** | [**WordTypes**](WordTypes.md) |  | 
 **priority** | **int32** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    text := "text_example" // string | 
    speaker := int32(56) // int32 | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AudioQueryAudioQueryPost(context.Background()).Text(text).Speaker(speaker).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AudioQueryAudioQueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AudioQueryAudioQueryPost`: AudioQuery
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AudioQueryAudioQueryPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    text := "text_example" // string | 
    presetId := int32(56) // int32 | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AudioQueryFromPresetAudioQueryFromPresetPost(context.Background()).Text(text).PresetId(presetId).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AudioQueryFromPresetAudioQueryFromPresetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AudioQueryFromPresetAudioQueryFromPresetPost`: AudioQuery
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AudioQueryFromPresetAudioQueryFromPresetPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 | 
    audioQuery := *openapiclient.NewAudioQuery([]openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))}, float32(123), float32(123), float32(123), float32(123), float32(123), float32(123), int32(123), false) // AudioQuery | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CancellableSynthesisCancellableSynthesisPost(context.Background()).Speaker(speaker).AudioQuery(audioQuery).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CancellableSynthesisCancellableSynthesisPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancellableSynthesisCancellableSynthesisPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CancellableSynthesisCancellableSynthesisPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectWavesConnectWavesPost(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectWavesConnectWavesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectWavesConnectWavesPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectWavesConnectWavesPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CoreVersionsCoreVersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CoreVersionsCoreVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreVersionsCoreVersionsGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CoreVersionsCoreVersionsGet`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeletePresetDeletePresetPost(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePresetDeletePresetPost``: %v\n", err)
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


## DeleteUserDictWordUserDictWordWordUuidDelete

> DeleteUserDictWordUserDictWordWordUuidDelete(ctx, wordUuid).Execute()

Delete User Dict Word



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    wordUuid := "wordUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteUserDictWordUserDictWordWordUuidDelete(context.Background(), wordUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUserDictWordUserDictWordWordUuidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wordUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserDictWordUserDictWordWordUuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DownloadableLibrariesDownloadableLibrariesGet

> []DownloadableLibrary DownloadableLibrariesDownloadableLibrariesGet(ctx).Execute()

Downloadable Libraries



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github/shikanime/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DownloadableLibrariesDownloadableLibrariesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadableLibrariesDownloadableLibrariesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadableLibrariesDownloadableLibrariesGet`: []DownloadableLibrary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadableLibrariesDownloadableLibrariesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadableLibrariesDownloadableLibrariesGetRequest struct via the builder pattern


### Return type

[**[]DownloadableLibrary**](DownloadableLibrary.md)

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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EngineManifestEngineManifestGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EngineManifestEngineManifestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EngineManifestEngineManifestGet`: EngineManifest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EngineManifestEngineManifestGet`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPresetsPresetsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPresetsPresetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPresetsPresetsGet`: []Preset
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPresetsPresetsGet`: %v\n", resp)
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


## GetUserDictWordsUserDictGet

> map[string]UserDictWord GetUserDictWordsUserDictGet(ctx).Execute()

Get User Dict Words



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github/shikanime/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetUserDictWordsUserDictGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserDictWordsUserDictGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserDictWordsUserDictGet`: map[string]UserDictWord
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserDictWordsUserDictGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDictWordsUserDictGetRequest struct via the builder pattern


### Return type

[**map[string]UserDictWord**](UserDictWord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportUserDictWordsImportUserDictPost

> ImportUserDictWordsImportUserDictPost(ctx).Override(override).RequestBody(requestBody).Execute()

Import User Dict Words



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    override := true // bool | 
    requestBody := map[string]UserDictWord{"key": *openapiclient.NewUserDictWord("Surface_example", int32(123), "PartOfSpeech_example", "PartOfSpeechDetail1_example", "PartOfSpeechDetail2_example", "PartOfSpeechDetail3_example", "InflectionalType_example", "InflectionalForm_example", "Stem_example", "Yomi_example", "Pronunciation_example", int32(123), "AccentAssociativeRule_example")} // map[string]UserDictWord | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ImportUserDictWordsImportUserDictPost(context.Background()).Override(override).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportUserDictWordsImportUserDictPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportUserDictWordsImportUserDictPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **override** | **bool** |  | 
 **requestBody** | [**map[string]UserDictWord**](UserDictWord.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 | 
    skipReinit := true // bool | 既に初期化済みの話者の再初期化をスキップするかどうか (optional) (default to false)
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.InitializeSpeakerInitializeSpeakerPost(context.Background()).Speaker(speaker).SkipReinit(skipReinit).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InitializeSpeakerInitializeSpeakerPost``: %v\n", err)
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


## InstallLibraryInstallLibraryLibraryUuidPost

> InstallLibraryInstallLibraryLibraryUuidPost(ctx, libraryUuid).Execute()

Install Library



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    libraryUuid := "libraryUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.InstallLibraryInstallLibraryLibraryUuidPost(context.Background(), libraryUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InstallLibraryInstallLibraryLibraryUuidPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallLibraryInstallLibraryLibraryUuidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## InstalledLibrariesInstalledLibrariesGet

> []DownloadableLibrary InstalledLibrariesInstalledLibrariesGet(ctx).Execute()

Installed Libraries



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github/shikanime/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.InstalledLibrariesInstalledLibrariesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InstalledLibrariesInstalledLibrariesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstalledLibrariesInstalledLibrariesGet`: []DownloadableLibrary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InstalledLibrariesInstalledLibrariesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstalledLibrariesInstalledLibrariesGetRequest struct via the builder pattern


### Return type

[**[]DownloadableLibrary**](DownloadableLibrary.md)

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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.IsInitializedSpeakerIsInitializedSpeakerGet(context.Background()).Speaker(speaker).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IsInitializedSpeakerIsInitializedSpeakerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsInitializedSpeakerIsInitializedSpeakerGet`: bool
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IsInitializedSpeakerIsInitializedSpeakerGet`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 | 
    accentPhrase := []openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))} // []AccentPhrase | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MoraDataMoraDataPost(context.Background()).Speaker(speaker).AccentPhrase(accentPhrase).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MoraDataMoraDataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoraDataMoraDataPost`: []AccentPhrase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MoraDataMoraDataPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 | 
    accentPhrase := []openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))} // []AccentPhrase | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MoraLengthMoraLengthPost(context.Background()).Speaker(speaker).AccentPhrase(accentPhrase).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MoraLengthMoraLengthPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoraLengthMoraLengthPost`: []AccentPhrase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MoraLengthMoraLengthPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 | 
    accentPhrase := []openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))} // []AccentPhrase | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MoraPitchMoraPitchPost(context.Background()).Speaker(speaker).AccentPhrase(accentPhrase).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MoraPitchMoraPitchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoraPitchMoraPitchPost`: []AccentPhrase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MoraPitchMoraPitchPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    requestBody := []int32{int32(123)} // []int32 | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MorphableTargetsMorphableTargetsPost(context.Background()).RequestBody(requestBody).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MorphableTargetsMorphableTargetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MorphableTargetsMorphableTargetsPost`: []map[string]MorphableTargetInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MorphableTargetsMorphableTargetsPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 | 
    audioQuery := []openapiclient.AudioQuery{*openapiclient.NewAudioQuery([]openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))}, float32(123), float32(123), float32(123), float32(123), float32(123), float32(123), int32(123), false)} // []AudioQuery | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MultiSynthesisMultiSynthesisPost(context.Background()).Speaker(speaker).AudioQuery(audioQuery).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MultiSynthesisMultiSynthesisPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MultiSynthesisMultiSynthesisPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MultiSynthesisMultiSynthesisPost`: %v\n", resp)
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


## RewriteUserDictWordUserDictWordWordUuidPut

> RewriteUserDictWordUserDictWordWordUuidPut(ctx, wordUuid).Surface(surface).Pronunciation(pronunciation).AccentType(accentType).WordType(wordType).Priority(priority).Execute()

Rewrite User Dict Word



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    wordUuid := "wordUuid_example" // string | 
    surface := "surface_example" // string | 
    pronunciation := "pronunciation_example" // string | 
    accentType := int32(56) // int32 | 
    wordType := openapiclient.WordTypes("PROPER_NOUN") // WordTypes |  (optional)
    priority := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.RewriteUserDictWordUserDictWordWordUuidPut(context.Background(), wordUuid).Surface(surface).Pronunciation(pronunciation).AccentType(accentType).WordType(wordType).Priority(priority).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RewriteUserDictWordUserDictWordWordUuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wordUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRewriteUserDictWordUserDictWordWordUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **surface** | **string** |  | 
 **pronunciation** | **string** |  | 
 **accentType** | **int32** |  | 
 **wordType** | [**WordTypes**](WordTypes.md) |  | 
 **priority** | **int32** |  | 

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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SettingGetSettingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SettingGetSettingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingGetSettingGet`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SettingGetSettingGet`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    corsPolicyMode := "corsPolicyMode_example" // string |  (optional)
    allowOrigin := "allowOrigin_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SettingPostSettingPost(context.Background()).CorsPolicyMode(corsPolicyMode).AllowOrigin(allowOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SettingPostSettingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingPostSettingPost`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SettingPostSettingPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    speakerUuid := "speakerUuid_example" // string | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SpeakerInfoSpeakerInfoGet(context.Background()).SpeakerUuid(speakerUuid).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SpeakerInfoSpeakerInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpeakerInfoSpeakerInfoGet`: SpeakerInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SpeakerInfoSpeakerInfoGet`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SpeakersSpeakersGet(context.Background()).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SpeakersSpeakersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpeakersSpeakersGet`: []Speaker
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SpeakersSpeakersGet`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SupportedDevicesSupportedDevicesGet(context.Background()).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SupportedDevicesSupportedDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportedDevicesSupportedDevicesGet`: SupportedDevicesInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SupportedDevicesSupportedDevicesGet`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    baseSpeaker := int32(56) // int32 | 
    targetSpeaker := int32(56) // int32 | 
    morphRate := float32(8.14) // float32 | 
    audioQuery := *openapiclient.NewAudioQuery([]openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))}, float32(123), float32(123), float32(123), float32(123), float32(123), float32(123), int32(123), false) // AudioQuery | 
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SynthesisMorphingSynthesisMorphingPost(context.Background()).BaseSpeaker(baseSpeaker).TargetSpeaker(targetSpeaker).MorphRate(morphRate).AudioQuery(audioQuery).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SynthesisMorphingSynthesisMorphingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SynthesisMorphingSynthesisMorphingPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SynthesisMorphingSynthesisMorphingPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    speaker := int32(56) // int32 | 
    audioQuery := *openapiclient.NewAudioQuery([]openapiclient.AccentPhrase{*openapiclient.NewAccentPhrase([]openapiclient.Mora{*openapiclient.NewMora("Text_example", "Vowel_example", float32(123), float32(123))}, int32(123))}, float32(123), float32(123), float32(123), float32(123), float32(123), float32(123), int32(123), false) // AudioQuery | 
    enableInterrogativeUpspeak := true // bool | 疑問系のテキストが与えられたら語尾を自動調整する (optional) (default to true)
    coreVersion := "coreVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SynthesisSynthesisPost(context.Background()).Speaker(speaker).AudioQuery(audioQuery).EnableInterrogativeUpspeak(enableInterrogativeUpspeak).CoreVersion(coreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SynthesisSynthesisPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SynthesisSynthesisPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SynthesisSynthesisPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {
    preset := *openapiclient.NewPreset(int32(123), "Name_example", "SpeakerUuid_example", int32(123), float32(123), float32(123), float32(123), float32(123), float32(123), float32(123)) // Preset | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdatePresetUpdatePresetPost(context.Background()).Preset(preset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdatePresetUpdatePresetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePresetUpdatePresetPost`: int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdatePresetUpdatePresetPost`: %v\n", resp)
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
    openapiclient "github/shikanime/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VersionVersionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VersionVersionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VersionVersionGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VersionVersionGet`: %v\n", resp)
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

