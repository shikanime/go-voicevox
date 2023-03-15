# Speaker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to [**SpeakerSupportedFeatures**](SpeakerSupportedFeatures.md) |  | [optional] 
**Name** | **string** |  | 
**SpeakerUuid** | **string** |  | 
**Styles** | [**[]SpeakerStyle**](SpeakerStyle.md) |  | 
**Version** | Pointer to **string** |  | [optional] [default to "スピーカーのバージョン"]

## Methods

### NewSpeaker

`func NewSpeaker(name string, speakerUuid string, styles []SpeakerStyle, ) *Speaker`

NewSpeaker instantiates a new Speaker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeakerWithDefaults

`func NewSpeakerWithDefaults() *Speaker`

NewSpeakerWithDefaults instantiates a new Speaker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *Speaker) GetSupportedFeatures() SpeakerSupportedFeatures`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *Speaker) GetSupportedFeaturesOk() (*SpeakerSupportedFeatures, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *Speaker) SetSupportedFeatures(v SpeakerSupportedFeatures)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *Speaker) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetName

`func (o *Speaker) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Speaker) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Speaker) SetName(v string)`

SetName sets Name field to given value.


### GetSpeakerUuid

`func (o *Speaker) GetSpeakerUuid() string`

GetSpeakerUuid returns the SpeakerUuid field if non-nil, zero value otherwise.

### GetSpeakerUuidOk

`func (o *Speaker) GetSpeakerUuidOk() (*string, bool)`

GetSpeakerUuidOk returns a tuple with the SpeakerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerUuid

`func (o *Speaker) SetSpeakerUuid(v string)`

SetSpeakerUuid sets SpeakerUuid field to given value.


### GetStyles

`func (o *Speaker) GetStyles() []SpeakerStyle`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *Speaker) GetStylesOk() (*[]SpeakerStyle, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *Speaker) SetStyles(v []SpeakerStyle)`

SetStyles sets Styles field to given value.


### GetVersion

`func (o *Speaker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Speaker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Speaker) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Speaker) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


