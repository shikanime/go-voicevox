# DownloadableLibrary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Uuid** | **string** |  | 
**Version** | **string** |  | 
**DownloadUrl** | **string** |  | 
**Bytes** | **int32** |  | 
**Speakers** | [**[]LibrarySpeaker**](LibrarySpeaker.md) |  | 

## Methods

### NewDownloadableLibrary

`func NewDownloadableLibrary(name string, uuid string, version string, downloadUrl string, bytes int32, speakers []LibrarySpeaker, ) *DownloadableLibrary`

NewDownloadableLibrary instantiates a new DownloadableLibrary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadableLibraryWithDefaults

`func NewDownloadableLibraryWithDefaults() *DownloadableLibrary`

NewDownloadableLibraryWithDefaults instantiates a new DownloadableLibrary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DownloadableLibrary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DownloadableLibrary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DownloadableLibrary) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *DownloadableLibrary) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DownloadableLibrary) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DownloadableLibrary) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetVersion

`func (o *DownloadableLibrary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DownloadableLibrary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DownloadableLibrary) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDownloadUrl

`func (o *DownloadableLibrary) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *DownloadableLibrary) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *DownloadableLibrary) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### GetBytes

`func (o *DownloadableLibrary) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *DownloadableLibrary) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *DownloadableLibrary) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetSpeakers

`func (o *DownloadableLibrary) GetSpeakers() []LibrarySpeaker`

GetSpeakers returns the Speakers field if non-nil, zero value otherwise.

### GetSpeakersOk

`func (o *DownloadableLibrary) GetSpeakersOk() (*[]LibrarySpeaker, bool)`

GetSpeakersOk returns a tuple with the Speakers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakers

`func (o *DownloadableLibrary) SetSpeakers(v []LibrarySpeaker)`

SetSpeakers sets Speakers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


