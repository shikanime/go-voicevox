# \UserDictionaryApi

All URIs are relative to *http://localhost:50021*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserDictWordUserDictWordPost**](UserDictionaryApi.md#AddUserDictWordUserDictWordPost) | **Post** /user_dict_word | Add User Dict Word
[**DeleteUserDictWordUserDictWordWordUuidDelete**](UserDictionaryApi.md#DeleteUserDictWordUserDictWordWordUuidDelete) | **Delete** /user_dict_word/{word_uuid} | Delete User Dict Word
[**GetUserDictWordsUserDictGet**](UserDictionaryApi.md#GetUserDictWordsUserDictGet) | **Get** /user_dict | Get User Dict Words
[**ImportUserDictWordsImportUserDictPost**](UserDictionaryApi.md#ImportUserDictWordsImportUserDictPost) | **Post** /import_user_dict | Import User Dict Words
[**RewriteUserDictWordUserDictWordWordUuidPut**](UserDictionaryApi.md#RewriteUserDictWordUserDictWordWordUuidPut) | **Put** /user_dict_word/{word_uuid} | Rewrite User Dict Word



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
    openapiclient "github.com/infinity-blackhole/go-voicevox"
)

func main() {
    surface := "surface_example" // string | 
    pronunciation := "pronunciation_example" // string | 
    accentType := int32(56) // int32 | 
    wordType := openapiclient.WordTypes("PROPER_NOUN") // WordTypes |  (optional)
    priority := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserDictionaryApi.AddUserDictWordUserDictWordPost(context.Background()).Surface(surface).Pronunciation(pronunciation).AccentType(accentType).WordType(wordType).Priority(priority).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserDictionaryApi.AddUserDictWordUserDictWordPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserDictWordUserDictWordPost`: string
    fmt.Fprintf(os.Stdout, "Response from `UserDictionaryApi.AddUserDictWordUserDictWordPost`: %v\n", resp)
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
    openapiclient "github.com/infinity-blackhole/go-voicevox"
)

func main() {
    wordUuid := "wordUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserDictionaryApi.DeleteUserDictWordUserDictWordWordUuidDelete(context.Background(), wordUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserDictionaryApi.DeleteUserDictWordUserDictWordWordUuidDelete``: %v\n", err)
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
    openapiclient "github.com/infinity-blackhole/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserDictionaryApi.GetUserDictWordsUserDictGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserDictionaryApi.GetUserDictWordsUserDictGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserDictWordsUserDictGet`: map[string]UserDictWord
    fmt.Fprintf(os.Stdout, "Response from `UserDictionaryApi.GetUserDictWordsUserDictGet`: %v\n", resp)
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
    openapiclient "github.com/infinity-blackhole/go-voicevox"
)

func main() {
    override := true // bool | 
    requestBody := map[string]UserDictWord{"key": *openapiclient.NewUserDictWord("Surface_example", int32(123), "PartOfSpeech_example", "PartOfSpeechDetail1_example", "PartOfSpeechDetail2_example", "PartOfSpeechDetail3_example", "InflectionalType_example", "InflectionalForm_example", "Stem_example", "Yomi_example", "Pronunciation_example", int32(123), "AccentAssociativeRule_example")} // map[string]UserDictWord | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserDictionaryApi.ImportUserDictWordsImportUserDictPost(context.Background()).Override(override).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserDictionaryApi.ImportUserDictWordsImportUserDictPost``: %v\n", err)
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
    openapiclient "github.com/infinity-blackhole/go-voicevox"
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
    r, err := apiClient.UserDictionaryApi.RewriteUserDictWordUserDictWordWordUuidPut(context.Background(), wordUuid).Surface(surface).Pronunciation(pronunciation).AccentType(accentType).WordType(wordType).Priority(priority).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserDictionaryApi.RewriteUserDictWordUserDictWordWordUuidPut``: %v\n", err)
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

