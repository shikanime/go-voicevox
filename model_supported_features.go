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

// checks if the SupportedFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedFeatures{}

// SupportedFeatures エンジンが持つ機能の一覧
type SupportedFeatures struct {
	AdjustMoraPitch       bool `json:"adjust_mora_pitch"`
	AdjustPhonemeLength   bool `json:"adjust_phoneme_length"`
	AdjustSpeedScale      bool `json:"adjust_speed_scale"`
	AdjustPitchScale      bool `json:"adjust_pitch_scale"`
	AdjustIntonationScale bool `json:"adjust_intonation_scale"`
	AdjustVolumeScale     bool `json:"adjust_volume_scale"`
	InterrogativeUpspeak  bool `json:"interrogative_upspeak"`
	SynthesisMorphing     bool `json:"synthesis_morphing"`
	ManageLibrary         bool `json:"manage_library"`
}

// NewSupportedFeatures instantiates a new SupportedFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedFeatures(adjustMoraPitch bool, adjustPhonemeLength bool, adjustSpeedScale bool, adjustPitchScale bool, adjustIntonationScale bool, adjustVolumeScale bool, interrogativeUpspeak bool, synthesisMorphing bool, manageLibrary bool) *SupportedFeatures {
	this := SupportedFeatures{}
	this.AdjustMoraPitch = adjustMoraPitch
	this.AdjustPhonemeLength = adjustPhonemeLength
	this.AdjustSpeedScale = adjustSpeedScale
	this.AdjustPitchScale = adjustPitchScale
	this.AdjustIntonationScale = adjustIntonationScale
	this.AdjustVolumeScale = adjustVolumeScale
	this.InterrogativeUpspeak = interrogativeUpspeak
	this.SynthesisMorphing = synthesisMorphing
	this.ManageLibrary = manageLibrary
	return &this
}

// NewSupportedFeaturesWithDefaults instantiates a new SupportedFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedFeaturesWithDefaults() *SupportedFeatures {
	this := SupportedFeatures{}
	return &this
}

// GetAdjustMoraPitch returns the AdjustMoraPitch field value
func (o *SupportedFeatures) GetAdjustMoraPitch() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AdjustMoraPitch
}

// GetAdjustMoraPitchOk returns a tuple with the AdjustMoraPitch field value
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetAdjustMoraPitchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdjustMoraPitch, true
}

// SetAdjustMoraPitch sets field value
func (o *SupportedFeatures) SetAdjustMoraPitch(v bool) {
	o.AdjustMoraPitch = v
}

// GetAdjustPhonemeLength returns the AdjustPhonemeLength field value
func (o *SupportedFeatures) GetAdjustPhonemeLength() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AdjustPhonemeLength
}

// GetAdjustPhonemeLengthOk returns a tuple with the AdjustPhonemeLength field value
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetAdjustPhonemeLengthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdjustPhonemeLength, true
}

// SetAdjustPhonemeLength sets field value
func (o *SupportedFeatures) SetAdjustPhonemeLength(v bool) {
	o.AdjustPhonemeLength = v
}

// GetAdjustSpeedScale returns the AdjustSpeedScale field value
func (o *SupportedFeatures) GetAdjustSpeedScale() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AdjustSpeedScale
}

// GetAdjustSpeedScaleOk returns a tuple with the AdjustSpeedScale field value
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetAdjustSpeedScaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdjustSpeedScale, true
}

// SetAdjustSpeedScale sets field value
func (o *SupportedFeatures) SetAdjustSpeedScale(v bool) {
	o.AdjustSpeedScale = v
}

// GetAdjustPitchScale returns the AdjustPitchScale field value
func (o *SupportedFeatures) GetAdjustPitchScale() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AdjustPitchScale
}

// GetAdjustPitchScaleOk returns a tuple with the AdjustPitchScale field value
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetAdjustPitchScaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdjustPitchScale, true
}

// SetAdjustPitchScale sets field value
func (o *SupportedFeatures) SetAdjustPitchScale(v bool) {
	o.AdjustPitchScale = v
}

// GetAdjustIntonationScale returns the AdjustIntonationScale field value
func (o *SupportedFeatures) GetAdjustIntonationScale() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AdjustIntonationScale
}

// GetAdjustIntonationScaleOk returns a tuple with the AdjustIntonationScale field value
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetAdjustIntonationScaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdjustIntonationScale, true
}

// SetAdjustIntonationScale sets field value
func (o *SupportedFeatures) SetAdjustIntonationScale(v bool) {
	o.AdjustIntonationScale = v
}

// GetAdjustVolumeScale returns the AdjustVolumeScale field value
func (o *SupportedFeatures) GetAdjustVolumeScale() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AdjustVolumeScale
}

// GetAdjustVolumeScaleOk returns a tuple with the AdjustVolumeScale field value
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetAdjustVolumeScaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdjustVolumeScale, true
}

// SetAdjustVolumeScale sets field value
func (o *SupportedFeatures) SetAdjustVolumeScale(v bool) {
	o.AdjustVolumeScale = v
}

// GetInterrogativeUpspeak returns the InterrogativeUpspeak field value
func (o *SupportedFeatures) GetInterrogativeUpspeak() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InterrogativeUpspeak
}

// GetInterrogativeUpspeakOk returns a tuple with the InterrogativeUpspeak field value
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetInterrogativeUpspeakOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterrogativeUpspeak, true
}

// SetInterrogativeUpspeak sets field value
func (o *SupportedFeatures) SetInterrogativeUpspeak(v bool) {
	o.InterrogativeUpspeak = v
}

// GetSynthesisMorphing returns the SynthesisMorphing field value
func (o *SupportedFeatures) GetSynthesisMorphing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SynthesisMorphing
}

// GetSynthesisMorphingOk returns a tuple with the SynthesisMorphing field value
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetSynthesisMorphingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SynthesisMorphing, true
}

// SetSynthesisMorphing sets field value
func (o *SupportedFeatures) SetSynthesisMorphing(v bool) {
	o.SynthesisMorphing = v
}

// GetManageLibrary returns the ManageLibrary field value
func (o *SupportedFeatures) GetManageLibrary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ManageLibrary
}

// GetManageLibraryOk returns a tuple with the ManageLibrary field value
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetManageLibraryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManageLibrary, true
}

// SetManageLibrary sets field value
func (o *SupportedFeatures) SetManageLibrary(v bool) {
	o.ManageLibrary = v
}

func (o SupportedFeatures) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adjust_mora_pitch"] = o.AdjustMoraPitch
	toSerialize["adjust_phoneme_length"] = o.AdjustPhonemeLength
	toSerialize["adjust_speed_scale"] = o.AdjustSpeedScale
	toSerialize["adjust_pitch_scale"] = o.AdjustPitchScale
	toSerialize["adjust_intonation_scale"] = o.AdjustIntonationScale
	toSerialize["adjust_volume_scale"] = o.AdjustVolumeScale
	toSerialize["interrogative_upspeak"] = o.InterrogativeUpspeak
	toSerialize["synthesis_morphing"] = o.SynthesisMorphing
	toSerialize["manage_library"] = o.ManageLibrary
	return toSerialize, nil
}

type NullableSupportedFeatures struct {
	value *SupportedFeatures
	isSet bool
}

func (v NullableSupportedFeatures) Get() *SupportedFeatures {
	return v.value
}

func (v *NullableSupportedFeatures) Set(val *SupportedFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedFeatures(val *SupportedFeatures) *NullableSupportedFeatures {
	return &NullableSupportedFeatures{value: val, isSet: true}
}

func (v NullableSupportedFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
