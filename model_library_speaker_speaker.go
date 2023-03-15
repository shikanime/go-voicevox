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

// checks if the LibrarySpeakerSpeaker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibrarySpeakerSpeaker{}

// LibrarySpeakerSpeaker struct for LibrarySpeakerSpeaker
type LibrarySpeakerSpeaker struct {
	SupportedFeatures *SpeakerSupportedFeatures `json:"supported_features,omitempty"`
	Name              string                    `json:"name"`
	SpeakerUuid       string                    `json:"speaker_uuid"`
	Styles            []SpeakerStyle            `json:"styles"`
	Version           *string                   `json:"version,omitempty"`
}

// NewLibrarySpeakerSpeaker instantiates a new LibrarySpeakerSpeaker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibrarySpeakerSpeaker(name string, speakerUuid string, styles []SpeakerStyle) *LibrarySpeakerSpeaker {
	this := LibrarySpeakerSpeaker{}
	this.Name = name
	this.SpeakerUuid = speakerUuid
	this.Styles = styles
	var version string = "スピーカーのバージョン"
	this.Version = &version
	return &this
}

// NewLibrarySpeakerSpeakerWithDefaults instantiates a new LibrarySpeakerSpeaker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibrarySpeakerSpeakerWithDefaults() *LibrarySpeakerSpeaker {
	this := LibrarySpeakerSpeaker{}
	var version string = "スピーカーのバージョン"
	this.Version = &version
	return &this
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *LibrarySpeakerSpeaker) GetSupportedFeatures() SpeakerSupportedFeatures {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret SpeakerSupportedFeatures
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibrarySpeakerSpeaker) GetSupportedFeaturesOk() (*SpeakerSupportedFeatures, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *LibrarySpeakerSpeaker) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given SpeakerSupportedFeatures and assigns it to the SupportedFeatures field.
func (o *LibrarySpeakerSpeaker) SetSupportedFeatures(v SpeakerSupportedFeatures) {
	o.SupportedFeatures = &v
}

// GetName returns the Name field value
func (o *LibrarySpeakerSpeaker) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LibrarySpeakerSpeaker) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LibrarySpeakerSpeaker) SetName(v string) {
	o.Name = v
}

// GetSpeakerUuid returns the SpeakerUuid field value
func (o *LibrarySpeakerSpeaker) GetSpeakerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpeakerUuid
}

// GetSpeakerUuidOk returns a tuple with the SpeakerUuid field value
// and a boolean to check if the value has been set.
func (o *LibrarySpeakerSpeaker) GetSpeakerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpeakerUuid, true
}

// SetSpeakerUuid sets field value
func (o *LibrarySpeakerSpeaker) SetSpeakerUuid(v string) {
	o.SpeakerUuid = v
}

// GetStyles returns the Styles field value
func (o *LibrarySpeakerSpeaker) GetStyles() []SpeakerStyle {
	if o == nil {
		var ret []SpeakerStyle
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *LibrarySpeakerSpeaker) GetStylesOk() ([]SpeakerStyle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Styles, true
}

// SetStyles sets field value
func (o *LibrarySpeakerSpeaker) SetStyles(v []SpeakerStyle) {
	o.Styles = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LibrarySpeakerSpeaker) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibrarySpeakerSpeaker) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LibrarySpeakerSpeaker) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LibrarySpeakerSpeaker) SetVersion(v string) {
	o.Version = &v
}

func (o LibrarySpeakerSpeaker) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibrarySpeakerSpeaker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supported_features"] = o.SupportedFeatures
	}
	toSerialize["name"] = o.Name
	toSerialize["speaker_uuid"] = o.SpeakerUuid
	toSerialize["styles"] = o.Styles
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableLibrarySpeakerSpeaker struct {
	value *LibrarySpeakerSpeaker
	isSet bool
}

func (v NullableLibrarySpeakerSpeaker) Get() *LibrarySpeakerSpeaker {
	return v.value
}

func (v *NullableLibrarySpeakerSpeaker) Set(val *LibrarySpeakerSpeaker) {
	v.value = val
	v.isSet = true
}

func (v NullableLibrarySpeakerSpeaker) IsSet() bool {
	return v.isSet
}

func (v *NullableLibrarySpeakerSpeaker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibrarySpeakerSpeaker(val *LibrarySpeakerSpeaker) *NullableLibrarySpeakerSpeaker {
	return &NullableLibrarySpeakerSpeaker{value: val, isSet: true}
}

func (v NullableLibrarySpeakerSpeaker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibrarySpeakerSpeaker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}