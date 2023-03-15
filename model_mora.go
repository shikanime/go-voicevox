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

// checks if the Mora type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mora{}

// Mora モーラ（子音＋母音）ごとの情報
type Mora struct {
	Text            string   `json:"text"`
	Consonant       *string  `json:"consonant,omitempty"`
	ConsonantLength *float32 `json:"consonant_length,omitempty"`
	Vowel           string   `json:"vowel"`
	VowelLength     float32  `json:"vowel_length"`
	Pitch           float32  `json:"pitch"`
}

// NewMora instantiates a new Mora object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMora(text string, vowel string, vowelLength float32, pitch float32) *Mora {
	this := Mora{}
	this.Text = text
	this.Vowel = vowel
	this.VowelLength = vowelLength
	this.Pitch = pitch
	return &this
}

// NewMoraWithDefaults instantiates a new Mora object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoraWithDefaults() *Mora {
	this := Mora{}
	return &this
}

// GetText returns the Text field value
func (o *Mora) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *Mora) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *Mora) SetText(v string) {
	o.Text = v
}

// GetConsonant returns the Consonant field value if set, zero value otherwise.
func (o *Mora) GetConsonant() string {
	if o == nil || IsNil(o.Consonant) {
		var ret string
		return ret
	}
	return *o.Consonant
}

// GetConsonantOk returns a tuple with the Consonant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mora) GetConsonantOk() (*string, bool) {
	if o == nil || IsNil(o.Consonant) {
		return nil, false
	}
	return o.Consonant, true
}

// HasConsonant returns a boolean if a field has been set.
func (o *Mora) HasConsonant() bool {
	if o != nil && !IsNil(o.Consonant) {
		return true
	}

	return false
}

// SetConsonant gets a reference to the given string and assigns it to the Consonant field.
func (o *Mora) SetConsonant(v string) {
	o.Consonant = &v
}

// GetConsonantLength returns the ConsonantLength field value if set, zero value otherwise.
func (o *Mora) GetConsonantLength() float32 {
	if o == nil || IsNil(o.ConsonantLength) {
		var ret float32
		return ret
	}
	return *o.ConsonantLength
}

// GetConsonantLengthOk returns a tuple with the ConsonantLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mora) GetConsonantLengthOk() (*float32, bool) {
	if o == nil || IsNil(o.ConsonantLength) {
		return nil, false
	}
	return o.ConsonantLength, true
}

// HasConsonantLength returns a boolean if a field has been set.
func (o *Mora) HasConsonantLength() bool {
	if o != nil && !IsNil(o.ConsonantLength) {
		return true
	}

	return false
}

// SetConsonantLength gets a reference to the given float32 and assigns it to the ConsonantLength field.
func (o *Mora) SetConsonantLength(v float32) {
	o.ConsonantLength = &v
}

// GetVowel returns the Vowel field value
func (o *Mora) GetVowel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vowel
}

// GetVowelOk returns a tuple with the Vowel field value
// and a boolean to check if the value has been set.
func (o *Mora) GetVowelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vowel, true
}

// SetVowel sets field value
func (o *Mora) SetVowel(v string) {
	o.Vowel = v
}

// GetVowelLength returns the VowelLength field value
func (o *Mora) GetVowelLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.VowelLength
}

// GetVowelLengthOk returns a tuple with the VowelLength field value
// and a boolean to check if the value has been set.
func (o *Mora) GetVowelLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VowelLength, true
}

// SetVowelLength sets field value
func (o *Mora) SetVowelLength(v float32) {
	o.VowelLength = v
}

// GetPitch returns the Pitch field value
func (o *Mora) GetPitch() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Pitch
}

// GetPitchOk returns a tuple with the Pitch field value
// and a boolean to check if the value has been set.
func (o *Mora) GetPitchOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pitch, true
}

// SetPitch sets field value
func (o *Mora) SetPitch(v float32) {
	o.Pitch = v
}

func (o Mora) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mora) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if !IsNil(o.Consonant) {
		toSerialize["consonant"] = o.Consonant
	}
	if !IsNil(o.ConsonantLength) {
		toSerialize["consonant_length"] = o.ConsonantLength
	}
	toSerialize["vowel"] = o.Vowel
	toSerialize["vowel_length"] = o.VowelLength
	toSerialize["pitch"] = o.Pitch
	return toSerialize, nil
}

type NullableMora struct {
	value *Mora
	isSet bool
}

func (v NullableMora) Get() *Mora {
	return v.value
}

func (v *NullableMora) Set(val *Mora) {
	v.value = val
	v.isSet = true
}

func (v NullableMora) IsSet() bool {
	return v.isSet
}

func (v *NullableMora) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMora(val *Mora) *NullableMora {
	return &NullableMora{value: val, isSet: true}
}

func (v NullableMora) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMora) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}