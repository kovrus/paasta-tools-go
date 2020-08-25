/*
 * Paasta API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package paastaapi

import (
	"encoding/json"
)

// SmartstackLocation struct for SmartstackLocation
type SmartstackLocation struct {
	// Smartstack backends running in this location
	Backends *[]SmartstackBackend `json:"backends,omitempty"`
	// Name of the location
	Name *string `json:"name,omitempty"`
	// Number of running backends for the service in this location
	RunningBackendsCount *int32 `json:"running_backends_count,omitempty"`
}

// NewSmartstackLocation instantiates a new SmartstackLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartstackLocation() *SmartstackLocation {
	this := SmartstackLocation{}
	return &this
}

// NewSmartstackLocationWithDefaults instantiates a new SmartstackLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartstackLocationWithDefaults() *SmartstackLocation {
	this := SmartstackLocation{}
	return &this
}

// GetBackends returns the Backends field value if set, zero value otherwise.
func (o *SmartstackLocation) GetBackends() []SmartstackBackend {
	if o == nil || o.Backends == nil {
		var ret []SmartstackBackend
		return ret
	}
	return *o.Backends
}

// GetBackendsOk returns a tuple with the Backends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartstackLocation) GetBackendsOk() (*[]SmartstackBackend, bool) {
	if o == nil || o.Backends == nil {
		return nil, false
	}
	return o.Backends, true
}

// HasBackends returns a boolean if a field has been set.
func (o *SmartstackLocation) HasBackends() bool {
	if o != nil && o.Backends != nil {
		return true
	}

	return false
}

// SetBackends gets a reference to the given []SmartstackBackend and assigns it to the Backends field.
func (o *SmartstackLocation) SetBackends(v []SmartstackBackend) {
	o.Backends = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SmartstackLocation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartstackLocation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SmartstackLocation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SmartstackLocation) SetName(v string) {
	o.Name = &v
}

// GetRunningBackendsCount returns the RunningBackendsCount field value if set, zero value otherwise.
func (o *SmartstackLocation) GetRunningBackendsCount() int32 {
	if o == nil || o.RunningBackendsCount == nil {
		var ret int32
		return ret
	}
	return *o.RunningBackendsCount
}

// GetRunningBackendsCountOk returns a tuple with the RunningBackendsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartstackLocation) GetRunningBackendsCountOk() (*int32, bool) {
	if o == nil || o.RunningBackendsCount == nil {
		return nil, false
	}
	return o.RunningBackendsCount, true
}

// HasRunningBackendsCount returns a boolean if a field has been set.
func (o *SmartstackLocation) HasRunningBackendsCount() bool {
	if o != nil && o.RunningBackendsCount != nil {
		return true
	}

	return false
}

// SetRunningBackendsCount gets a reference to the given int32 and assigns it to the RunningBackendsCount field.
func (o *SmartstackLocation) SetRunningBackendsCount(v int32) {
	o.RunningBackendsCount = &v
}

func (o SmartstackLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Backends != nil {
		toSerialize["backends"] = o.Backends
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RunningBackendsCount != nil {
		toSerialize["running_backends_count"] = o.RunningBackendsCount
	}
	return json.Marshal(toSerialize)
}

type NullableSmartstackLocation struct {
	value *SmartstackLocation
	isSet bool
}

func (v NullableSmartstackLocation) Get() *SmartstackLocation {
	return v.value
}

func (v *NullableSmartstackLocation) Set(val *SmartstackLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartstackLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartstackLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartstackLocation(val *SmartstackLocation) *NullableSmartstackLocation {
	return &NullableSmartstackLocation{value: val, isSet: true}
}

func (v NullableSmartstackLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartstackLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


