/*
VOICEVOX Engine

VOICEVOXの音声合成エンジンです。

API version: latest
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package voicevox

import (
	"encoding/json"
)

// checks if the AudioQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudioQuery{}

// AudioQuery 音声合成用のクエリ
type AudioQuery struct {
	AccentPhrases      []AccentPhrase `json:"accent_phrases"`
	SpeedScale         float32        `json:"speedScale"`
	PitchScale         float32        `json:"pitchScale"`
	IntonationScale    float32        `json:"intonationScale"`
	VolumeScale        float32        `json:"volumeScale"`
	PrePhonemeLength   float32        `json:"prePhonemeLength"`
	PostPhonemeLength  float32        `json:"postPhonemeLength"`
	OutputSamplingRate int32          `json:"outputSamplingRate"`
	OutputStereo       bool           `json:"outputStereo"`
	Kana               *string        `json:"kana,omitempty"`
}

// NewAudioQuery instantiates a new AudioQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudioQuery(accentPhrases []AccentPhrase, speedScale float32, pitchScale float32, intonationScale float32, volumeScale float32, prePhonemeLength float32, postPhonemeLength float32, outputSamplingRate int32, outputStereo bool) *AudioQuery {
	this := AudioQuery{}
	this.AccentPhrases = accentPhrases
	this.SpeedScale = speedScale
	this.PitchScale = pitchScale
	this.IntonationScale = intonationScale
	this.VolumeScale = volumeScale
	this.PrePhonemeLength = prePhonemeLength
	this.PostPhonemeLength = postPhonemeLength
	this.OutputSamplingRate = outputSamplingRate
	this.OutputStereo = outputStereo
	return &this
}

// NewAudioQueryWithDefaults instantiates a new AudioQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudioQueryWithDefaults() *AudioQuery {
	this := AudioQuery{}
	return &this
}

// GetAccentPhrases returns the AccentPhrases field value
func (o *AudioQuery) GetAccentPhrases() []AccentPhrase {
	if o == nil {
		var ret []AccentPhrase
		return ret
	}

	return o.AccentPhrases
}

// GetAccentPhrasesOk returns a tuple with the AccentPhrases field value
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetAccentPhrasesOk() ([]AccentPhrase, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccentPhrases, true
}

// SetAccentPhrases sets field value
func (o *AudioQuery) SetAccentPhrases(v []AccentPhrase) {
	o.AccentPhrases = v
}

// GetSpeedScale returns the SpeedScale field value
func (o *AudioQuery) GetSpeedScale() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpeedScale
}

// GetSpeedScaleOk returns a tuple with the SpeedScale field value
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetSpeedScaleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpeedScale, true
}

// SetSpeedScale sets field value
func (o *AudioQuery) SetSpeedScale(v float32) {
	o.SpeedScale = v
}

// GetPitchScale returns the PitchScale field value
func (o *AudioQuery) GetPitchScale() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PitchScale
}

// GetPitchScaleOk returns a tuple with the PitchScale field value
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetPitchScaleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PitchScale, true
}

// SetPitchScale sets field value
func (o *AudioQuery) SetPitchScale(v float32) {
	o.PitchScale = v
}

// GetIntonationScale returns the IntonationScale field value
func (o *AudioQuery) GetIntonationScale() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IntonationScale
}

// GetIntonationScaleOk returns a tuple with the IntonationScale field value
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetIntonationScaleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntonationScale, true
}

// SetIntonationScale sets field value
func (o *AudioQuery) SetIntonationScale(v float32) {
	o.IntonationScale = v
}

// GetVolumeScale returns the VolumeScale field value
func (o *AudioQuery) GetVolumeScale() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.VolumeScale
}

// GetVolumeScaleOk returns a tuple with the VolumeScale field value
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetVolumeScaleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeScale, true
}

// SetVolumeScale sets field value
func (o *AudioQuery) SetVolumeScale(v float32) {
	o.VolumeScale = v
}

// GetPrePhonemeLength returns the PrePhonemeLength field value
func (o *AudioQuery) GetPrePhonemeLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PrePhonemeLength
}

// GetPrePhonemeLengthOk returns a tuple with the PrePhonemeLength field value
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetPrePhonemeLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrePhonemeLength, true
}

// SetPrePhonemeLength sets field value
func (o *AudioQuery) SetPrePhonemeLength(v float32) {
	o.PrePhonemeLength = v
}

// GetPostPhonemeLength returns the PostPhonemeLength field value
func (o *AudioQuery) GetPostPhonemeLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PostPhonemeLength
}

// GetPostPhonemeLengthOk returns a tuple with the PostPhonemeLength field value
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetPostPhonemeLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostPhonemeLength, true
}

// SetPostPhonemeLength sets field value
func (o *AudioQuery) SetPostPhonemeLength(v float32) {
	o.PostPhonemeLength = v
}

// GetOutputSamplingRate returns the OutputSamplingRate field value
func (o *AudioQuery) GetOutputSamplingRate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OutputSamplingRate
}

// GetOutputSamplingRateOk returns a tuple with the OutputSamplingRate field value
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetOutputSamplingRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputSamplingRate, true
}

// SetOutputSamplingRate sets field value
func (o *AudioQuery) SetOutputSamplingRate(v int32) {
	o.OutputSamplingRate = v
}

// GetOutputStereo returns the OutputStereo field value
func (o *AudioQuery) GetOutputStereo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OutputStereo
}

// GetOutputStereoOk returns a tuple with the OutputStereo field value
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetOutputStereoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputStereo, true
}

// SetOutputStereo sets field value
func (o *AudioQuery) SetOutputStereo(v bool) {
	o.OutputStereo = v
}

// GetKana returns the Kana field value if set, zero value otherwise.
func (o *AudioQuery) GetKana() string {
	if o == nil || IsNil(o.Kana) {
		var ret string
		return ret
	}
	return *o.Kana
}

// GetKanaOk returns a tuple with the Kana field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioQuery) GetKanaOk() (*string, bool) {
	if o == nil || IsNil(o.Kana) {
		return nil, false
	}
	return o.Kana, true
}

// HasKana returns a boolean if a field has been set.
func (o *AudioQuery) HasKana() bool {
	if o != nil && !IsNil(o.Kana) {
		return true
	}

	return false
}

// SetKana gets a reference to the given string and assigns it to the Kana field.
func (o *AudioQuery) SetKana(v string) {
	o.Kana = &v
}

func (o AudioQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudioQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accent_phrases"] = o.AccentPhrases
	toSerialize["speedScale"] = o.SpeedScale
	toSerialize["pitchScale"] = o.PitchScale
	toSerialize["intonationScale"] = o.IntonationScale
	toSerialize["volumeScale"] = o.VolumeScale
	toSerialize["prePhonemeLength"] = o.PrePhonemeLength
	toSerialize["postPhonemeLength"] = o.PostPhonemeLength
	toSerialize["outputSamplingRate"] = o.OutputSamplingRate
	toSerialize["outputStereo"] = o.OutputStereo
	if !IsNil(o.Kana) {
		toSerialize["kana"] = o.Kana
	}
	return toSerialize, nil
}

type NullableAudioQuery struct {
	value *AudioQuery
	isSet bool
}

func (v NullableAudioQuery) Get() *AudioQuery {
	return v.value
}

func (v *NullableAudioQuery) Set(val *AudioQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableAudioQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableAudioQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudioQuery(val *AudioQuery) *NullableAudioQuery {
	return &NullableAudioQuery{value: val, isSet: true}
}

func (v NullableAudioQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudioQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
