# StyleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Icon** | **string** |  | 
**Portrait** | Pointer to **string** |  | [optional] 
**VoiceSamples** | **[]string** |  | 

## Methods

### NewStyleInfo

`func NewStyleInfo(id int32, icon string, voiceSamples []string, ) *StyleInfo`

NewStyleInfo instantiates a new StyleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStyleInfoWithDefaults

`func NewStyleInfoWithDefaults() *StyleInfo`

NewStyleInfoWithDefaults instantiates a new StyleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StyleInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StyleInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StyleInfo) SetId(v int32)`

SetId sets Id field to given value.


### GetIcon

`func (o *StyleInfo) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *StyleInfo) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *StyleInfo) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetPortrait

`func (o *StyleInfo) GetPortrait() string`

GetPortrait returns the Portrait field if non-nil, zero value otherwise.

### GetPortraitOk

`func (o *StyleInfo) GetPortraitOk() (*string, bool)`

GetPortraitOk returns a tuple with the Portrait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortrait

`func (o *StyleInfo) SetPortrait(v string)`

SetPortrait sets Portrait field to given value.

### HasPortrait

`func (o *StyleInfo) HasPortrait() bool`

HasPortrait returns a boolean if a field has been set.

### GetVoiceSamples

`func (o *StyleInfo) GetVoiceSamples() []string`

GetVoiceSamples returns the VoiceSamples field if non-nil, zero value otherwise.

### GetVoiceSamplesOk

`func (o *StyleInfo) GetVoiceSamplesOk() (*[]string, bool)`

GetVoiceSamplesOk returns a tuple with the VoiceSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSamples

`func (o *StyleInfo) SetVoiceSamples(v []string)`

SetVoiceSamples sets VoiceSamples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


