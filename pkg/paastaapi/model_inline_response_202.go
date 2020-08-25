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

// InlineResponse202 struct for InlineResponse202
type InlineResponse202 struct {
	DesiredInstances *int32 `json:"desired_instances,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewInlineResponse202 instantiates a new InlineResponse202 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse202() *InlineResponse202 {
	this := InlineResponse202{}
	return &this
}

// NewInlineResponse202WithDefaults instantiates a new InlineResponse202 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse202WithDefaults() *InlineResponse202 {
	this := InlineResponse202{}
	return &this
}

// GetDesiredInstances returns the DesiredInstances field value if set, zero value otherwise.
func (o *InlineResponse202) GetDesiredInstances() int32 {
	if o == nil || o.DesiredInstances == nil {
		var ret int32
		return ret
	}
	return *o.DesiredInstances
}

// GetDesiredInstancesOk returns a tuple with the DesiredInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse202) GetDesiredInstancesOk() (*int32, bool) {
	if o == nil || o.DesiredInstances == nil {
		return nil, false
	}
	return o.DesiredInstances, true
}

// HasDesiredInstances returns a boolean if a field has been set.
func (o *InlineResponse202) HasDesiredInstances() bool {
	if o != nil && o.DesiredInstances != nil {
		return true
	}

	return false
}

// SetDesiredInstances gets a reference to the given int32 and assigns it to the DesiredInstances field.
func (o *InlineResponse202) SetDesiredInstances(v int32) {
	o.DesiredInstances = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse202) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse202) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse202) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse202) SetStatus(v string) {
	o.Status = &v
}

func (o InlineResponse202) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DesiredInstances != nil {
		toSerialize["desired_instances"] = o.DesiredInstances
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse202 struct {
	value *InlineResponse202
	isSet bool
}

func (v NullableInlineResponse202) Get() *InlineResponse202 {
	return v.value
}

func (v *NullableInlineResponse202) Set(val *InlineResponse202) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse202) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse202) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse202(val *InlineResponse202) *NullableInlineResponse202 {
	return &NullableInlineResponse202{value: val, isSet: true}
}

func (v NullableInlineResponse202) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse202) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

