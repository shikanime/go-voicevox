# AudioQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccentPhrases** | [**[]AccentPhrase**](AccentPhrase.md) |  |
**SpeedScale** | **float32** |  |
**PitchScale** | **float32** |  |
**IntonationScale** | **float32** |  |
**VolumeScale** | **float32** |  |
**PrePhonemeLength** | **float32** |  |
**PostPhonemeLength** | **float32** |  |
**OutputSamplingRate** | **int32** |  |
**OutputStereo** | **bool** |  |
**Kana** | Pointer to **string** |  | [optional]

## Methods

### NewAudioQuery

`func NewAudioQuery(accentPhrases []AccentPhrase, speedScale float32, pitchScale float32, intonationScale float32, volumeScale float32, prePhonemeLength float32, postPhonemeLength float32, outputSamplingRate int32, outputStereo bool, ) *AudioQuery`

NewAudioQuery instantiates a new AudioQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioQueryWithDefaults

`func NewAudioQueryWithDefaults() *AudioQuery`

NewAudioQueryWithDefaults instantiates a new AudioQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccentPhrases

`func (o *AudioQuery) GetAccentPhrases() []AccentPhrase`

GetAccentPhrases returns the AccentPhrases field if non-nil, zero value otherwise.

### GetAccentPhrasesOk

`func (o *AudioQuery) GetAccentPhrasesOk() (*[]AccentPhrase, bool)`

GetAccentPhrasesOk returns a tuple with the AccentPhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentPhrases

`func (o *AudioQuery) SetAccentPhrases(v []AccentPhrase)`

SetAccentPhrases sets AccentPhrases field to given value.

### GetSpeedScale

`func (o *AudioQuery) GetSpeedScale() float32`

GetSpeedScale returns the SpeedScale field if non-nil, zero value otherwise.

### GetSpeedScaleOk

`func (o *AudioQuery) GetSpeedScaleOk() (*float32, bool)`

GetSpeedScaleOk returns a tuple with the SpeedScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedScale

`func (o *AudioQuery) SetSpeedScale(v float32)`

SetSpeedScale sets SpeedScale field to given value.

### GetPitchScale

`func (o *AudioQuery) GetPitchScale() float32`

GetPitchScale returns the PitchScale field if non-nil, zero value otherwise.

### GetPitchScaleOk

`func (o *AudioQuery) GetPitchScaleOk() (*float32, bool)`

GetPitchScaleOk returns a tuple with the PitchScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitchScale

`func (o *AudioQuery) SetPitchScale(v float32)`

SetPitchScale sets PitchScale field to given value.

### GetIntonationScale

`func (o *AudioQuery) GetIntonationScale() float32`

GetIntonationScale returns the IntonationScale field if non-nil, zero value otherwise.

### GetIntonationScaleOk

`func (o *AudioQuery) GetIntonationScaleOk() (*float32, bool)`

GetIntonationScaleOk returns a tuple with the IntonationScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntonationScale

`func (o *AudioQuery) SetIntonationScale(v float32)`

SetIntonationScale sets IntonationScale field to given value.

### GetVolumeScale

`func (o *AudioQuery) GetVolumeScale() float32`

GetVolumeScale returns the VolumeScale field if non-nil, zero value otherwise.

### GetVolumeScaleOk

`func (o *AudioQuery) GetVolumeScaleOk() (*float32, bool)`

GetVolumeScaleOk returns a tuple with the VolumeScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeScale

`func (o *AudioQuery) SetVolumeScale(v float32)`

SetVolumeScale sets VolumeScale field to given value.

### GetPrePhonemeLength

`func (o *AudioQuery) GetPrePhonemeLength() float32`

GetPrePhonemeLength returns the PrePhonemeLength field if non-nil, zero value otherwise.

### GetPrePhonemeLengthOk

`func (o *AudioQuery) GetPrePhonemeLengthOk() (*float32, bool)`

GetPrePhonemeLengthOk returns a tuple with the PrePhonemeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePhonemeLength

`func (o *AudioQuery) SetPrePhonemeLength(v float32)`

SetPrePhonemeLength sets PrePhonemeLength field to given value.

### GetPostPhonemeLength

`func (o *AudioQuery) GetPostPhonemeLength() float32`

GetPostPhonemeLength returns the PostPhonemeLength field if non-nil, zero value otherwise.

### GetPostPhonemeLengthOk

`func (o *AudioQuery) GetPostPhonemeLengthOk() (*float32, bool)`

GetPostPhonemeLengthOk returns a tuple with the PostPhonemeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPhonemeLength

`func (o *AudioQuery) SetPostPhonemeLength(v float32)`

SetPostPhonemeLength sets PostPhonemeLength field to given value.

### GetOutputSamplingRate

`func (o *AudioQuery) GetOutputSamplingRate() int32`

GetOutputSamplingRate returns the OutputSamplingRate field if non-nil, zero value otherwise.

### GetOutputSamplingRateOk

`func (o *AudioQuery) GetOutputSamplingRateOk() (*int32, bool)`

GetOutputSamplingRateOk returns a tuple with the OutputSamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSamplingRate

`func (o *AudioQuery) SetOutputSamplingRate(v int32)`

SetOutputSamplingRate sets OutputSamplingRate field to given value.

### GetOutputStereo

`func (o *AudioQuery) GetOutputStereo() bool`

GetOutputStereo returns the OutputStereo field if non-nil, zero value otherwise.

### GetOutputStereoOk

`func (o *AudioQuery) GetOutputStereoOk() (*bool, bool)`

GetOutputStereoOk returns a tuple with the OutputStereo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputStereo

`func (o *AudioQuery) SetOutputStereo(v bool)`

SetOutputStereo sets OutputStereo field to given value.

### GetKana

`func (o *AudioQuery) GetKana() string`

GetKana returns the Kana field if non-nil, zero value otherwise.

### GetKanaOk

`func (o *AudioQuery) GetKanaOk() (*string, bool)`

GetKanaOk returns a tuple with the Kana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKana

`func (o *AudioQuery) SetKana(v string)`

SetKana sets Kana field to given value.

### HasKana

`func (o *AudioQuery) HasKana() bool`

HasKana returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
