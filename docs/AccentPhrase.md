# AccentPhrase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Moras** | [**[]Mora**](Mora.md) |  | 
**Accent** | **int32** |  | 
**PauseMora** | Pointer to [**AccentPhrasePauseMora**](AccentPhrasePauseMora.md) |  | [optional] 
**IsInterrogative** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAccentPhrase

`func NewAccentPhrase(moras []Mora, accent int32, ) *AccentPhrase`

NewAccentPhrase instantiates a new AccentPhrase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccentPhraseWithDefaults

`func NewAccentPhraseWithDefaults() *AccentPhrase`

NewAccentPhraseWithDefaults instantiates a new AccentPhrase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoras

`func (o *AccentPhrase) GetMoras() []Mora`

GetMoras returns the Moras field if non-nil, zero value otherwise.

### GetMorasOk

`func (o *AccentPhrase) GetMorasOk() (*[]Mora, bool)`

GetMorasOk returns a tuple with the Moras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoras

`func (o *AccentPhrase) SetMoras(v []Mora)`

SetMoras sets Moras field to given value.


### GetAccent

`func (o *AccentPhrase) GetAccent() int32`

GetAccent returns the Accent field if non-nil, zero value otherwise.

### GetAccentOk

`func (o *AccentPhrase) GetAccentOk() (*int32, bool)`

GetAccentOk returns a tuple with the Accent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccent

`func (o *AccentPhrase) SetAccent(v int32)`

SetAccent sets Accent field to given value.


### GetPauseMora

`func (o *AccentPhrase) GetPauseMora() AccentPhrasePauseMora`

GetPauseMora returns the PauseMora field if non-nil, zero value otherwise.

### GetPauseMoraOk

`func (o *AccentPhrase) GetPauseMoraOk() (*AccentPhrasePauseMora, bool)`

GetPauseMoraOk returns a tuple with the PauseMora field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseMora

`func (o *AccentPhrase) SetPauseMora(v AccentPhrasePauseMora)`

SetPauseMora sets PauseMora field to given value.

### HasPauseMora

`func (o *AccentPhrase) HasPauseMora() bool`

HasPauseMora returns a boolean if a field has been set.

### GetIsInterrogative

`func (o *AccentPhrase) GetIsInterrogative() bool`

GetIsInterrogative returns the IsInterrogative field if non-nil, zero value otherwise.

### GetIsInterrogativeOk

`func (o *AccentPhrase) GetIsInterrogativeOk() (*bool, bool)`

GetIsInterrogativeOk returns a tuple with the IsInterrogative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterrogative

`func (o *AccentPhrase) SetIsInterrogative(v bool)`

SetIsInterrogative sets IsInterrogative field to given value.

### HasIsInterrogative

`func (o *AccentPhrase) HasIsInterrogative() bool`

HasIsInterrogative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


