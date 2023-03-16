# \AudioLibraryManagementApi

All URIs are relative to *http://localhost:50021*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadableLibrariesDownloadableLibrariesGet**](AudioLibraryManagementApi.md#DownloadableLibrariesDownloadableLibrariesGet) | **Get** /downloadable_libraries | Downloadable Libraries
[**InstallLibraryInstallLibraryLibraryUuidPost**](AudioLibraryManagementApi.md#InstallLibraryInstallLibraryLibraryUuidPost) | **Post** /install_library/{library_uuid} | Install Library
[**InstalledLibrariesInstalledLibrariesGet**](AudioLibraryManagementApi.md#InstalledLibrariesInstalledLibrariesGet) | **Get** /installed_libraries | Installed Libraries



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
    openapiclient "github.com/infinity-blackhole/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudioLibraryManagementApi.DownloadableLibrariesDownloadableLibrariesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioLibraryManagementApi.DownloadableLibrariesDownloadableLibrariesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadableLibrariesDownloadableLibrariesGet`: []DownloadableLibrary
    fmt.Fprintf(os.Stdout, "Response from `AudioLibraryManagementApi.DownloadableLibrariesDownloadableLibrariesGet`: %v\n", resp)
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
    openapiclient "github.com/infinity-blackhole/go-voicevox"
)

func main() {
    libraryUuid := "libraryUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioLibraryManagementApi.InstallLibraryInstallLibraryLibraryUuidPost(context.Background(), libraryUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioLibraryManagementApi.InstallLibraryInstallLibraryLibraryUuidPost``: %v\n", err)
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
    openapiclient "github.com/infinity-blackhole/go-voicevox"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudioLibraryManagementApi.InstalledLibrariesInstalledLibrariesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioLibraryManagementApi.InstalledLibrariesInstalledLibrariesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstalledLibrariesInstalledLibrariesGet`: []DownloadableLibrary
    fmt.Fprintf(os.Stdout, "Response from `AudioLibraryManagementApi.InstalledLibrariesInstalledLibrariesGet`: %v\n", resp)
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

