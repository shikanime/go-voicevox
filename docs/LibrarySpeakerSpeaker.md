# LibrarySpeakerSpeaker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to [**SpeakerSupportedFeatures**](SpeakerSupportedFeatures.md) |  | [optional] 
**Name** | **string** |  | 
**SpeakerUuid** | **string** |  | 
**Styles** | [**[]SpeakerStyle**](SpeakerStyle.md) |  | 
**Version** | Pointer to **string** |  | [optional] [default to "スピーカーのバージョン"]

## Methods

### NewLibrarySpeakerSpeaker

`func NewLibrarySpeakerSpeaker(name string, speakerUuid string, styles []SpeakerStyle, ) *LibrarySpeakerSpeaker`

NewLibrarySpeakerSpeaker instantiates a new LibrarySpeakerSpeaker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibrarySpeakerSpeakerWithDefaults

`func NewLibrarySpeakerSpeakerWithDefaults() *LibrarySpeakerSpeaker`

NewLibrarySpeakerSpeakerWithDefaults instantiates a new LibrarySpeakerSpeaker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *LibrarySpeakerSpeaker) GetSupportedFeatures() SpeakerSupportedFeatures`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *LibrarySpeakerSpeaker) GetSupportedFeaturesOk() (*SpeakerSupportedFeatures, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *LibrarySpeakerSpeaker) SetSupportedFeatures(v SpeakerSupportedFeatures)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *LibrarySpeakerSpeaker) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetName

`func (o *LibrarySpeakerSpeaker) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibrarySpeakerSpeaker) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibrarySpeakerSpeaker) SetName(v string)`

SetName sets Name field to given value.


### GetSpeakerUuid

`func (o *LibrarySpeakerSpeaker) GetSpeakerUuid() string`

GetSpeakerUuid returns the SpeakerUuid field if non-nil, zero value otherwise.

### GetSpeakerUuidOk

`func (o *LibrarySpeakerSpeaker) GetSpeakerUuidOk() (*string, bool)`

GetSpeakerUuidOk returns a tuple with the SpeakerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerUuid

`func (o *LibrarySpeakerSpeaker) SetSpeakerUuid(v string)`

SetSpeakerUuid sets SpeakerUuid field to given value.


### GetStyles

`func (o *LibrarySpeakerSpeaker) GetStyles() []SpeakerStyle`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *LibrarySpeakerSpeaker) GetStylesOk() (*[]SpeakerStyle, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *LibrarySpeakerSpeaker) SetStyles(v []SpeakerStyle)`

SetStyles sets Styles field to given value.


### GetVersion

`func (o *LibrarySpeakerSpeaker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LibrarySpeakerSpeaker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LibrarySpeakerSpeaker) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LibrarySpeakerSpeaker) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


