/*
instagrapi-rest

RESTful API Service for instagrapi

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instapi

import (
	"encoding/json"
)

// StoryMention struct for StoryMention
type StoryMention struct {
	User UserShort `json:"user"`
	X *float32 `json:"x,omitempty"`
	Y *float32 `json:"y,omitempty"`
	Width *float32 `json:"width,omitempty"`
	Height *float32 `json:"height,omitempty"`
}

// NewStoryMention instantiates a new StoryMention object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoryMention(user UserShort) *StoryMention {
	this := StoryMention{}
	this.User = user
	return &this
}

// NewStoryMentionWithDefaults instantiates a new StoryMention object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryMentionWithDefaults() *StoryMention {
	this := StoryMention{}
	return &this
}

// GetUser returns the User field value
func (o *StoryMention) GetUser() UserShort {
	if o == nil {
		var ret UserShort
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *StoryMention) GetUserOk() (*UserShort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *StoryMention) SetUser(v UserShort) {
	o.User = v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *StoryMention) GetX() float32 {
	if o == nil || o.X == nil {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryMention) GetXOk() (*float32, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *StoryMention) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *StoryMention) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *StoryMention) GetY() float32 {
	if o == nil || o.Y == nil {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryMention) GetYOk() (*float32, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *StoryMention) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *StoryMention) SetY(v float32) {
	o.Y = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *StoryMention) GetWidth() float32 {
	if o == nil || o.Width == nil {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryMention) GetWidthOk() (*float32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *StoryMention) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *StoryMention) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *StoryMention) GetHeight() float32 {
	if o == nil || o.Height == nil {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryMention) GetHeightOk() (*float32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *StoryMention) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *StoryMention) SetHeight(v float32) {
	o.Height = &v
}

func (o StoryMention) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	return json.Marshal(toSerialize)
}

type NullableStoryMention struct {
	value *StoryMention
	isSet bool
}

func (v NullableStoryMention) Get() *StoryMention {
	return v.value
}

func (v *NullableStoryMention) Set(val *StoryMention) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryMention) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryMention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryMention(val *StoryMention) *NullableStoryMention {
	return &NullableStoryMention{value: val, isSet: true}
}

func (v NullableStoryMention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryMention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


