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

// TaskTailLines struct for TaskTailLines
type TaskTailLines struct {
	// Error message when fetching tail lines fails
	ErrorMessage *string `json:"error_message,omitempty"`
	// The requested number of lines from the task's stderr
	Stderr *[]string `json:"stderr,omitempty"`
	// The requested number of lines from the task's stdout
	Stdout *[]string `json:"stdout,omitempty"`
}

// NewTaskTailLines instantiates a new TaskTailLines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskTailLines() *TaskTailLines {
	this := TaskTailLines{}
	return &this
}

// NewTaskTailLinesWithDefaults instantiates a new TaskTailLines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskTailLinesWithDefaults() *TaskTailLines {
	this := TaskTailLines{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *TaskTailLines) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTailLines) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TaskTailLines) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *TaskTailLines) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetStderr returns the Stderr field value if set, zero value otherwise.
func (o *TaskTailLines) GetStderr() []string {
	if o == nil || o.Stderr == nil {
		var ret []string
		return ret
	}
	return *o.Stderr
}

// GetStderrOk returns a tuple with the Stderr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTailLines) GetStderrOk() (*[]string, bool) {
	if o == nil || o.Stderr == nil {
		return nil, false
	}
	return o.Stderr, true
}

// HasStderr returns a boolean if a field has been set.
func (o *TaskTailLines) HasStderr() bool {
	if o != nil && o.Stderr != nil {
		return true
	}

	return false
}

// SetStderr gets a reference to the given []string and assigns it to the Stderr field.
func (o *TaskTailLines) SetStderr(v []string) {
	o.Stderr = &v
}

// GetStdout returns the Stdout field value if set, zero value otherwise.
func (o *TaskTailLines) GetStdout() []string {
	if o == nil || o.Stdout == nil {
		var ret []string
		return ret
	}
	return *o.Stdout
}

// GetStdoutOk returns a tuple with the Stdout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTailLines) GetStdoutOk() (*[]string, bool) {
	if o == nil || o.Stdout == nil {
		return nil, false
	}
	return o.Stdout, true
}

// HasStdout returns a boolean if a field has been set.
func (o *TaskTailLines) HasStdout() bool {
	if o != nil && o.Stdout != nil {
		return true
	}

	return false
}

// SetStdout gets a reference to the given []string and assigns it to the Stdout field.
func (o *TaskTailLines) SetStdout(v []string) {
	o.Stdout = &v
}

func (o TaskTailLines) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Stderr != nil {
		toSerialize["stderr"] = o.Stderr
	}
	if o.Stdout != nil {
		toSerialize["stdout"] = o.Stdout
	}
	return json.Marshal(toSerialize)
}

type NullableTaskTailLines struct {
	value *TaskTailLines
	isSet bool
}

func (v NullableTaskTailLines) Get() *TaskTailLines {
	return v.value
}

func (v *NullableTaskTailLines) Set(val *TaskTailLines) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskTailLines) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskTailLines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskTailLines(val *TaskTailLines) *NullableTaskTailLines {
	return &NullableTaskTailLines{value: val, isSet: true}
}

func (v NullableTaskTailLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskTailLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


