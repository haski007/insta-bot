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

type Hashtag struct {
}

// StoryHashtag struct for StoryHashtag
type StoryHashtag struct {
	Hashtag Hashtag  `json:"hashtag"`
	X       *float32 `json:"x,omitempty"`
	Y       *float32 `json:"y,omitempty"`
	Width   *float32 `json:"width,omitempty"`
	Height  *float32 `json:"height,omitempty"`
}

// NewStoryHashtag instantiates a new StoryHashtag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoryHashtag(hashtag Hashtag) *StoryHashtag {
	this := StoryHashtag{}
	this.Hashtag = hashtag
	return &this
}

// NewStoryHashtagWithDefaults instantiates a new StoryHashtag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryHashtagWithDefaults() *StoryHashtag {
	this := StoryHashtag{}
	return &this
}

// GetHashtag returns the Hashtag field value
func (o *StoryHashtag) GetHashtag() Hashtag {
	if o == nil {
		var ret Hashtag
		return ret
	}

	return o.Hashtag
}

// GetHashtagOk returns a tuple with the Hashtag field value
// and a boolean to check if the value has been set.
func (o *StoryHashtag) GetHashtagOk() (*Hashtag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hashtag, true
}

// SetHashtag sets field value
func (o *StoryHashtag) SetHashtag(v Hashtag) {
	o.Hashtag = v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *StoryHashtag) GetX() float32 {
	if o == nil || o.X == nil {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryHashtag) GetXOk() (*float32, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *StoryHashtag) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *StoryHashtag) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *StoryHashtag) GetY() float32 {
	if o == nil || o.Y == nil {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryHashtag) GetYOk() (*float32, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *StoryHashtag) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *StoryHashtag) SetY(v float32) {
	o.Y = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *StoryHashtag) GetWidth() float32 {
	if o == nil || o.Width == nil {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryHashtag) GetWidthOk() (*float32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *StoryHashtag) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *StoryHashtag) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *StoryHashtag) GetHeight() float32 {
	if o == nil || o.Height == nil {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryHashtag) GetHeightOk() (*float32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *StoryHashtag) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *StoryHashtag) SetHeight(v float32) {
	o.Height = &v
}

func (o StoryHashtag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hashtag"] = o.Hashtag
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

type NullableStoryHashtag struct {
	value *StoryHashtag
	isSet bool
}

func (v NullableStoryHashtag) Get() *StoryHashtag {
	return v.value
}

func (v *NullableStoryHashtag) Set(val *StoryHashtag) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryHashtag) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryHashtag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryHashtag(val *StoryHashtag) *NullableStoryHashtag {
	return &NullableStoryHashtag{value: val, isSet: true}
}

func (v NullableStoryHashtag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryHashtag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
