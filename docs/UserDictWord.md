# UserDictWord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Surface** | **string** |  |
**Priority** | **int32** |  |
**ContextId** | Pointer to **int32** |  | [optional] [default to 1348]
**PartOfSpeech** | **string** |  |
**PartOfSpeechDetail1** | **string** |  |
**PartOfSpeechDetail2** | **string** |  |
**PartOfSpeechDetail3** | **string** |  |
**InflectionalType** | **string** |  |
**InflectionalForm** | **string** |  |
**Stem** | **string** |  |
**Yomi** | **string** |  |
**Pronunciation** | **string** |  |
**AccentType** | **int32** |  |
**MoraCount** | Pointer to **int32** |  | [optional]
**AccentAssociativeRule** | **string** |  |

## Methods

### NewUserDictWord

`func NewUserDictWord(surface string, priority int32, partOfSpeech string, partOfSpeechDetail1 string, partOfSpeechDetail2 string, partOfSpeechDetail3 string, inflectionalType string, inflectionalForm string, stem string, yomi string, pronunciation string, accentType int32, accentAssociativeRule string, ) *UserDictWord`

NewUserDictWord instantiates a new UserDictWord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDictWordWithDefaults

`func NewUserDictWordWithDefaults() *UserDictWord`

NewUserDictWordWithDefaults instantiates a new UserDictWord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSurface

`func (o *UserDictWord) GetSurface() string`

GetSurface returns the Surface field if non-nil, zero value otherwise.

### GetSurfaceOk

`func (o *UserDictWord) GetSurfaceOk() (*string, bool)`

GetSurfaceOk returns a tuple with the Surface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurface

`func (o *UserDictWord) SetSurface(v string)`

SetSurface sets Surface field to given value.

### GetPriority

`func (o *UserDictWord) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UserDictWord) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UserDictWord) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### GetContextId

`func (o *UserDictWord) GetContextId() int32`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *UserDictWord) GetContextIdOk() (*int32, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *UserDictWord) SetContextId(v int32)`

SetContextId sets ContextId field to given value.

### HasContextId

`func (o *UserDictWord) HasContextId() bool`

HasContextId returns a boolean if a field has been set.

### GetPartOfSpeech

`func (o *UserDictWord) GetPartOfSpeech() string`

GetPartOfSpeech returns the PartOfSpeech field if non-nil, zero value otherwise.

### GetPartOfSpeechOk

`func (o *UserDictWord) GetPartOfSpeechOk() (*string, bool)`

GetPartOfSpeechOk returns a tuple with the PartOfSpeech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOfSpeech

`func (o *UserDictWord) SetPartOfSpeech(v string)`

SetPartOfSpeech sets PartOfSpeech field to given value.

### GetPartOfSpeechDetail1

`func (o *UserDictWord) GetPartOfSpeechDetail1() string`

GetPartOfSpeechDetail1 returns the PartOfSpeechDetail1 field if non-nil, zero value otherwise.

### GetPartOfSpeechDetail1Ok

`func (o *UserDictWord) GetPartOfSpeechDetail1Ok() (*string, bool)`

GetPartOfSpeechDetail1Ok returns a tuple with the PartOfSpeechDetail1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOfSpeechDetail1

`func (o *UserDictWord) SetPartOfSpeechDetail1(v string)`

SetPartOfSpeechDetail1 sets PartOfSpeechDetail1 field to given value.

### GetPartOfSpeechDetail2

`func (o *UserDictWord) GetPartOfSpeechDetail2() string`

GetPartOfSpeechDetail2 returns the PartOfSpeechDetail2 field if non-nil, zero value otherwise.

### GetPartOfSpeechDetail2Ok

`func (o *UserDictWord) GetPartOfSpeechDetail2Ok() (*string, bool)`

GetPartOfSpeechDetail2Ok returns a tuple with the PartOfSpeechDetail2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOfSpeechDetail2

`func (o *UserDictWord) SetPartOfSpeechDetail2(v string)`

SetPartOfSpeechDetail2 sets PartOfSpeechDetail2 field to given value.

### GetPartOfSpeechDetail3

`func (o *UserDictWord) GetPartOfSpeechDetail3() string`

GetPartOfSpeechDetail3 returns the PartOfSpeechDetail3 field if non-nil, zero value otherwise.

### GetPartOfSpeechDetail3Ok

`func (o *UserDictWord) GetPartOfSpeechDetail3Ok() (*string, bool)`

GetPartOfSpeechDetail3Ok returns a tuple with the PartOfSpeechDetail3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOfSpeechDetail3

`func (o *UserDictWord) SetPartOfSpeechDetail3(v string)`

SetPartOfSpeechDetail3 sets PartOfSpeechDetail3 field to given value.

### GetInflectionalType

`func (o *UserDictWord) GetInflectionalType() string`

GetInflectionalType returns the InflectionalType field if non-nil, zero value otherwise.

### GetInflectionalTypeOk

`func (o *UserDictWord) GetInflectionalTypeOk() (*string, bool)`

GetInflectionalTypeOk returns a tuple with the InflectionalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflectionalType

`func (o *UserDictWord) SetInflectionalType(v string)`

SetInflectionalType sets InflectionalType field to given value.

### GetInflectionalForm

`func (o *UserDictWord) GetInflectionalForm() string`

GetInflectionalForm returns the InflectionalForm field if non-nil, zero value otherwise.

### GetInflectionalFormOk

`func (o *UserDictWord) GetInflectionalFormOk() (*string, bool)`

GetInflectionalFormOk returns a tuple with the InflectionalForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflectionalForm

`func (o *UserDictWord) SetInflectionalForm(v string)`

SetInflectionalForm sets InflectionalForm field to given value.

### GetStem

`func (o *UserDictWord) GetStem() string`

GetStem returns the Stem field if non-nil, zero value otherwise.

### GetStemOk

`func (o *UserDictWord) GetStemOk() (*string, bool)`

GetStemOk returns a tuple with the Stem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStem

`func (o *UserDictWord) SetStem(v string)`

SetStem sets Stem field to given value.

### GetYomi

`func (o *UserDictWord) GetYomi() string`

GetYomi returns the Yomi field if non-nil, zero value otherwise.

### GetYomiOk

`func (o *UserDictWord) GetYomiOk() (*string, bool)`

GetYomiOk returns a tuple with the Yomi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYomi

`func (o *UserDictWord) SetYomi(v string)`

SetYomi sets Yomi field to given value.

### GetPronunciation

`func (o *UserDictWord) GetPronunciation() string`

GetPronunciation returns the Pronunciation field if non-nil, zero value otherwise.

### GetPronunciationOk

`func (o *UserDictWord) GetPronunciationOk() (*string, bool)`

GetPronunciationOk returns a tuple with the Pronunciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronunciation

`func (o *UserDictWord) SetPronunciation(v string)`

SetPronunciation sets Pronunciation field to given value.

### GetAccentType

`func (o *UserDictWord) GetAccentType() int32`

GetAccentType returns the AccentType field if non-nil, zero value otherwise.

### GetAccentTypeOk

`func (o *UserDictWord) GetAccentTypeOk() (*int32, bool)`

GetAccentTypeOk returns a tuple with the AccentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentType

`func (o *UserDictWord) SetAccentType(v int32)`

SetAccentType sets AccentType field to given value.

### GetMoraCount

`func (o *UserDictWord) GetMoraCount() int32`

GetMoraCount returns the MoraCount field if non-nil, zero value otherwise.

### GetMoraCountOk

`func (o *UserDictWord) GetMoraCountOk() (*int32, bool)`

GetMoraCountOk returns a tuple with the MoraCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoraCount

`func (o *UserDictWord) SetMoraCount(v int32)`

SetMoraCount sets MoraCount field to given value.

### HasMoraCount

`func (o *UserDictWord) HasMoraCount() bool`

HasMoraCount returns a boolean if a field has been set.

### GetAccentAssociativeRule

`func (o *UserDictWord) GetAccentAssociativeRule() string`

GetAccentAssociativeRule returns the AccentAssociativeRule field if non-nil, zero value otherwise.

### GetAccentAssociativeRuleOk

`func (o *UserDictWord) GetAccentAssociativeRuleOk() (*string, bool)`

GetAccentAssociativeRuleOk returns a tuple with the AccentAssociativeRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentAssociativeRule

`func (o *UserDictWord) SetAccentAssociativeRule(v string)`

SetAccentAssociativeRule sets AccentAssociativeRule field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
