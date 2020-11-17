/*
 * Paasta API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paastaapi

import (
	"encoding/json"
)

// MarathonAutoscalingInfo struct for MarathonAutoscalingInfo
type MarathonAutoscalingInfo struct {
	// The number of instances of the service currently running
	CurrentInstances *int32 `json:"current_instances,omitempty"`
	// The current utilization of the instances' allocated resources
	CurrentUtilization NullableFloat32 `json:"current_utilization,omitempty"`
	// The maximum number of instances that the autoscaler will scale to
	MaxInstances *int32 `json:"max_instances,omitempty"`
	// The minimum number of instances that the autoscaler will scale to
	MinInstances *int32 `json:"min_instances,omitempty"`
	// The autoscaler's current target number of instances of this service to run
	TargetInstances NullableInt32 `json:"target_instances,omitempty"`
}

// NewMarathonAutoscalingInfo instantiates a new MarathonAutoscalingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarathonAutoscalingInfo() *MarathonAutoscalingInfo {
	this := MarathonAutoscalingInfo{}
	return &this
}

// NewMarathonAutoscalingInfoWithDefaults instantiates a new MarathonAutoscalingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarathonAutoscalingInfoWithDefaults() *MarathonAutoscalingInfo {
	this := MarathonAutoscalingInfo{}
	return &this
}

// GetCurrentInstances returns the CurrentInstances field value if set, zero value otherwise.
func (o *MarathonAutoscalingInfo) GetCurrentInstances() int32 {
	if o == nil || o.CurrentInstances == nil {
		var ret int32
		return ret
	}
	return *o.CurrentInstances
}

// GetCurrentInstancesOk returns a tuple with the CurrentInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarathonAutoscalingInfo) GetCurrentInstancesOk() (*int32, bool) {
	if o == nil || o.CurrentInstances == nil {
		return nil, false
	}
	return o.CurrentInstances, true
}

// HasCurrentInstances returns a boolean if a field has been set.
func (o *MarathonAutoscalingInfo) HasCurrentInstances() bool {
	if o != nil && o.CurrentInstances != nil {
		return true
	}

	return false
}

// SetCurrentInstances gets a reference to the given int32 and assigns it to the CurrentInstances field.
func (o *MarathonAutoscalingInfo) SetCurrentInstances(v int32) {
	o.CurrentInstances = &v
}

// GetCurrentUtilization returns the CurrentUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarathonAutoscalingInfo) GetCurrentUtilization() float32 {
	if o == nil || o.CurrentUtilization.Get() == nil {
		var ret float32
		return ret
	}
	return *o.CurrentUtilization.Get()
}

// GetCurrentUtilizationOk returns a tuple with the CurrentUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarathonAutoscalingInfo) GetCurrentUtilizationOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentUtilization.Get(), o.CurrentUtilization.IsSet()
}

// HasCurrentUtilization returns a boolean if a field has been set.
func (o *MarathonAutoscalingInfo) HasCurrentUtilization() bool {
	if o != nil && o.CurrentUtilization.IsSet() {
		return true
	}

	return false
}

// SetCurrentUtilization gets a reference to the given NullableFloat32 and assigns it to the CurrentUtilization field.
func (o *MarathonAutoscalingInfo) SetCurrentUtilization(v float32) {
	o.CurrentUtilization.Set(&v)
}
// SetCurrentUtilizationNil sets the value for CurrentUtilization to be an explicit nil
func (o *MarathonAutoscalingInfo) SetCurrentUtilizationNil() {
	o.CurrentUtilization.Set(nil)
}

// UnsetCurrentUtilization ensures that no value is present for CurrentUtilization, not even an explicit nil
func (o *MarathonAutoscalingInfo) UnsetCurrentUtilization() {
	o.CurrentUtilization.Unset()
}

// GetMaxInstances returns the MaxInstances field value if set, zero value otherwise.
func (o *MarathonAutoscalingInfo) GetMaxInstances() int32 {
	if o == nil || o.MaxInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxInstances
}

// GetMaxInstancesOk returns a tuple with the MaxInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarathonAutoscalingInfo) GetMaxInstancesOk() (*int32, bool) {
	if o == nil || o.MaxInstances == nil {
		return nil, false
	}
	return o.MaxInstances, true
}

// HasMaxInstances returns a boolean if a field has been set.
func (o *MarathonAutoscalingInfo) HasMaxInstances() bool {
	if o != nil && o.MaxInstances != nil {
		return true
	}

	return false
}

// SetMaxInstances gets a reference to the given int32 and assigns it to the MaxInstances field.
func (o *MarathonAutoscalingInfo) SetMaxInstances(v int32) {
	o.MaxInstances = &v
}

// GetMinInstances returns the MinInstances field value if set, zero value otherwise.
func (o *MarathonAutoscalingInfo) GetMinInstances() int32 {
	if o == nil || o.MinInstances == nil {
		var ret int32
		return ret
	}
	return *o.MinInstances
}

// GetMinInstancesOk returns a tuple with the MinInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarathonAutoscalingInfo) GetMinInstancesOk() (*int32, bool) {
	if o == nil || o.MinInstances == nil {
		return nil, false
	}
	return o.MinInstances, true
}

// HasMinInstances returns a boolean if a field has been set.
func (o *MarathonAutoscalingInfo) HasMinInstances() bool {
	if o != nil && o.MinInstances != nil {
		return true
	}

	return false
}

// SetMinInstances gets a reference to the given int32 and assigns it to the MinInstances field.
func (o *MarathonAutoscalingInfo) SetMinInstances(v int32) {
	o.MinInstances = &v
}

// GetTargetInstances returns the TargetInstances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarathonAutoscalingInfo) GetTargetInstances() int32 {
	if o == nil || o.TargetInstances.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TargetInstances.Get()
}

// GetTargetInstancesOk returns a tuple with the TargetInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarathonAutoscalingInfo) GetTargetInstancesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetInstances.Get(), o.TargetInstances.IsSet()
}

// HasTargetInstances returns a boolean if a field has been set.
func (o *MarathonAutoscalingInfo) HasTargetInstances() bool {
	if o != nil && o.TargetInstances.IsSet() {
		return true
	}

	return false
}

// SetTargetInstances gets a reference to the given NullableInt32 and assigns it to the TargetInstances field.
func (o *MarathonAutoscalingInfo) SetTargetInstances(v int32) {
	o.TargetInstances.Set(&v)
}
// SetTargetInstancesNil sets the value for TargetInstances to be an explicit nil
func (o *MarathonAutoscalingInfo) SetTargetInstancesNil() {
	o.TargetInstances.Set(nil)
}

// UnsetTargetInstances ensures that no value is present for TargetInstances, not even an explicit nil
func (o *MarathonAutoscalingInfo) UnsetTargetInstances() {
	o.TargetInstances.Unset()
}

func (o MarathonAutoscalingInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentInstances != nil {
		toSerialize["current_instances"] = o.CurrentInstances
	}
	if o.CurrentUtilization.IsSet() {
		toSerialize["current_utilization"] = o.CurrentUtilization.Get()
	}
	if o.MaxInstances != nil {
		toSerialize["max_instances"] = o.MaxInstances
	}
	if o.MinInstances != nil {
		toSerialize["min_instances"] = o.MinInstances
	}
	if o.TargetInstances.IsSet() {
		toSerialize["target_instances"] = o.TargetInstances.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMarathonAutoscalingInfo struct {
	value *MarathonAutoscalingInfo
	isSet bool
}

func (v NullableMarathonAutoscalingInfo) Get() *MarathonAutoscalingInfo {
	return v.value
}

func (v *NullableMarathonAutoscalingInfo) Set(val *MarathonAutoscalingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMarathonAutoscalingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMarathonAutoscalingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarathonAutoscalingInfo(val *MarathonAutoscalingInfo) *NullableMarathonAutoscalingInfo {
	return &NullableMarathonAutoscalingInfo{value: val, isSet: true}
}

func (v NullableMarathonAutoscalingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarathonAutoscalingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


