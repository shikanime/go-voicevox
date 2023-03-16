# SpeakerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** |  |
**Portrait** | **string** |  |
**StyleInfos** | [**[]StyleInfo**](StyleInfo.md) |  |

## Methods

### NewSpeakerInfo

`func NewSpeakerInfo(policy string, portrait string, styleInfos []StyleInfo, ) *SpeakerInfo`

NewSpeakerInfo instantiates a new SpeakerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeakerInfoWithDefaults

`func NewSpeakerInfoWithDefaults() *SpeakerInfo`

NewSpeakerInfoWithDefaults instantiates a new SpeakerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *SpeakerInfo) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SpeakerInfo) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SpeakerInfo) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### GetPortrait

`func (o *SpeakerInfo) GetPortrait() string`

GetPortrait returns the Portrait field if non-nil, zero value otherwise.

### GetPortraitOk

`func (o *SpeakerInfo) GetPortraitOk() (*string, bool)`

GetPortraitOk returns a tuple with the Portrait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortrait

`func (o *SpeakerInfo) SetPortrait(v string)`

SetPortrait sets Portrait field to given value.

### GetStyleInfos

`func (o *SpeakerInfo) GetStyleInfos() []StyleInfo`

GetStyleInfos returns the StyleInfos field if non-nil, zero value otherwise.

### GetStyleInfosOk

`func (o *SpeakerInfo) GetStyleInfosOk() (*[]StyleInfo, bool)`

GetStyleInfosOk returns a tuple with the StyleInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleInfos

`func (o *SpeakerInfo) SetStyleInfos(v []StyleInfo)`

SetStyleInfos sets StyleInfos field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
