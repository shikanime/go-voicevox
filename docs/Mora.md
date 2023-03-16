# Mora

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  |
**Consonant** | Pointer to **string** |  | [optional]
**ConsonantLength** | Pointer to **float32** |  | [optional]
**Vowel** | **string** |  |
**VowelLength** | **float32** |  |
**Pitch** | **float32** |  |

## Methods

### NewMora

`func NewMora(text string, vowel string, vowelLength float32, pitch float32, ) *Mora`

NewMora instantiates a new Mora object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoraWithDefaults

`func NewMoraWithDefaults() *Mora`

NewMoraWithDefaults instantiates a new Mora object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *Mora) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Mora) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Mora) SetText(v string)`

SetText sets Text field to given value.

### GetConsonant

`func (o *Mora) GetConsonant() string`

GetConsonant returns the Consonant field if non-nil, zero value otherwise.

### GetConsonantOk

`func (o *Mora) GetConsonantOk() (*string, bool)`

GetConsonantOk returns a tuple with the Consonant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsonant

`func (o *Mora) SetConsonant(v string)`

SetConsonant sets Consonant field to given value.

### HasConsonant

`func (o *Mora) HasConsonant() bool`

HasConsonant returns a boolean if a field has been set.

### GetConsonantLength

`func (o *Mora) GetConsonantLength() float32`

GetConsonantLength returns the ConsonantLength field if non-nil, zero value otherwise.

### GetConsonantLengthOk

`func (o *Mora) GetConsonantLengthOk() (*float32, bool)`

GetConsonantLengthOk returns a tuple with the ConsonantLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsonantLength

`func (o *Mora) SetConsonantLength(v float32)`

SetConsonantLength sets ConsonantLength field to given value.

### HasConsonantLength

`func (o *Mora) HasConsonantLength() bool`

HasConsonantLength returns a boolean if a field has been set.

### GetVowel

`func (o *Mora) GetVowel() string`

GetVowel returns the Vowel field if non-nil, zero value otherwise.

### GetVowelOk

`func (o *Mora) GetVowelOk() (*string, bool)`

GetVowelOk returns a tuple with the Vowel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVowel

`func (o *Mora) SetVowel(v string)`

SetVowel sets Vowel field to given value.

### GetVowelLength

`func (o *Mora) GetVowelLength() float32`

GetVowelLength returns the VowelLength field if non-nil, zero value otherwise.

### GetVowelLengthOk

`func (o *Mora) GetVowelLengthOk() (*float32, bool)`

GetVowelLengthOk returns a tuple with the VowelLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVowelLength

`func (o *Mora) SetVowelLength(v float32)`

SetVowelLength sets VowelLength field to given value.

### GetPitch

`func (o *Mora) GetPitch() float32`

GetPitch returns the Pitch field if non-nil, zero value otherwise.

### GetPitchOk

`func (o *Mora) GetPitchOk() (*float32, bool)`

GetPitchOk returns a tuple with the Pitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitch

`func (o *Mora) SetPitch(v float32)`

SetPitch sets Pitch field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
