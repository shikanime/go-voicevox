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

// checks if the SpeakerSupportedFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpeakerSupportedFeatures{}

// SpeakerSupportedFeatures 話者の対応機能の情報
type SpeakerSupportedFeatures struct {
	PermittedSynthesisMorphing *SpeakerSupportPermittedSynthesisMorphing `json:"permitted_synthesis_morphing,omitempty"`
}

// NewSpeakerSupportedFeatures instantiates a new SpeakerSupportedFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpeakerSupportedFeatures() *SpeakerSupportedFeatures {
	this := SpeakerSupportedFeatures{}
	var permittedSynthesisMorphing SpeakerSupportPermittedSynthesisMorphing = "ALL"
	this.PermittedSynthesisMorphing = &permittedSynthesisMorphing
	return &this
}

// NewSpeakerSupportedFeaturesWithDefaults instantiates a new SpeakerSupportedFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpeakerSupportedFeaturesWithDefaults() *SpeakerSupportedFeatures {
	this := SpeakerSupportedFeatures{}
	var permittedSynthesisMorphing SpeakerSupportPermittedSynthesisMorphing = "ALL"
	this.PermittedSynthesisMorphing = &permittedSynthesisMorphing
	return &this
}

// GetPermittedSynthesisMorphing returns the PermittedSynthesisMorphing field value if set, zero value otherwise.
func (o *SpeakerSupportedFeatures) GetPermittedSynthesisMorphing() SpeakerSupportPermittedSynthesisMorphing {
	if o == nil || IsNil(o.PermittedSynthesisMorphing) {
		var ret SpeakerSupportPermittedSynthesisMorphing
		return ret
	}
	return *o.PermittedSynthesisMorphing
}

// GetPermittedSynthesisMorphingOk returns a tuple with the PermittedSynthesisMorphing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpeakerSupportedFeatures) GetPermittedSynthesisMorphingOk() (*SpeakerSupportPermittedSynthesisMorphing, bool) {
	if o == nil || IsNil(o.PermittedSynthesisMorphing) {
		return nil, false
	}
	return o.PermittedSynthesisMorphing, true
}

// HasPermittedSynthesisMorphing returns a boolean if a field has been set.
func (o *SpeakerSupportedFeatures) HasPermittedSynthesisMorphing() bool {
	if o != nil && !IsNil(o.PermittedSynthesisMorphing) {
		return true
	}

	return false
}

// SetPermittedSynthesisMorphing gets a reference to the given SpeakerSupportPermittedSynthesisMorphing and assigns it to the PermittedSynthesisMorphing field.
func (o *SpeakerSupportedFeatures) SetPermittedSynthesisMorphing(v SpeakerSupportPermittedSynthesisMorphing) {
	o.PermittedSynthesisMorphing = &v
}

func (o SpeakerSupportedFeatures) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpeakerSupportedFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PermittedSynthesisMorphing) {
		toSerialize["permitted_synthesis_morphing"] = o.PermittedSynthesisMorphing
	}
	return toSerialize, nil
}

type NullableSpeakerSupportedFeatures struct {
	value *SpeakerSupportedFeatures
	isSet bool
}

func (v NullableSpeakerSupportedFeatures) Get() *SpeakerSupportedFeatures {
	return v.value
}

func (v *NullableSpeakerSupportedFeatures) Set(val *SpeakerSupportedFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableSpeakerSupportedFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableSpeakerSupportedFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpeakerSupportedFeatures(val *SpeakerSupportedFeatures) *NullableSpeakerSupportedFeatures {
	return &NullableSpeakerSupportedFeatures{value: val, isSet: true}
}

func (v NullableSpeakerSupportedFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpeakerSupportedFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
