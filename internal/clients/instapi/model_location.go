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

// Location struct for Location
type Location struct {
	Pk *int32 `json:"pk,omitempty"`
	Name string `json:"name"`
	Phone *string `json:"phone,omitempty"`
	Website *string `json:"website,omitempty"`
	Category *string `json:"category,omitempty"`
	Hours map[string]interface{} `json:"hours,omitempty"`
	Address *string `json:"address,omitempty"`
	City *string `json:"city,omitempty"`
	Zip *string `json:"zip,omitempty"`
	Lng *float32 `json:"lng,omitempty"`
	Lat *float32 `json:"lat,omitempty"`
	ExternalId *int32 `json:"external_id,omitempty"`
	ExternalIdSource *string `json:"external_id_source,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation(name string) *Location {
	this := Location{}
	this.Name = name
	var phone string = ""
	this.Phone = &phone
	var website string = ""
	this.Website = &website
	var category string = ""
	this.Category = &category
	var address string = ""
	this.Address = &address
	var city string = ""
	this.City = &city
	var zip string = ""
	this.Zip = &zip
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	var phone string = ""
	this.Phone = &phone
	var website string = ""
	this.Website = &website
	var category string = ""
	this.Category = &category
	var address string = ""
	this.Address = &address
	var city string = ""
	this.City = &city
	var zip string = ""
	this.Zip = &zip
	return &this
}

// GetPk returns the Pk field value if set, zero value otherwise.
func (o *Location) GetPk() int32 {
	if o == nil || o.Pk == nil {
		var ret int32
		return ret
	}
	return *o.Pk
}

// GetPkOk returns a tuple with the Pk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPkOk() (*int32, bool) {
	if o == nil || o.Pk == nil {
		return nil, false
	}
	return o.Pk, true
}

// HasPk returns a boolean if a field has been set.
func (o *Location) HasPk() bool {
	if o != nil && o.Pk != nil {
		return true
	}

	return false
}

// SetPk gets a reference to the given int32 and assigns it to the Pk field.
func (o *Location) SetPk(v int32) {
	o.Pk = &v
}

// GetName returns the Name field value
func (o *Location) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Location) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Location) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Location) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Location) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Location) SetPhone(v string) {
	o.Phone = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *Location) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *Location) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *Location) SetWebsite(v string) {
	o.Website = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Location) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Location) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Location) SetCategory(v string) {
	o.Category = &v
}

// GetHours returns the Hours field value if set, zero value otherwise.
func (o *Location) GetHours() map[string]interface{} {
	if o == nil || o.Hours == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetHoursOk() (map[string]interface{}, bool) {
	if o == nil || o.Hours == nil {
		return nil, false
	}
	return o.Hours, true
}

// HasHours returns a boolean if a field has been set.
func (o *Location) HasHours() bool {
	if o != nil && o.Hours != nil {
		return true
	}

	return false
}

// SetHours gets a reference to the given map[string]interface{} and assigns it to the Hours field.
func (o *Location) SetHours(v map[string]interface{}) {
	o.Hours = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Location) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Location) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Location) SetAddress(v string) {
	o.Address = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Location) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Location) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Location) SetCity(v string) {
	o.City = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *Location) GetZip() string {
	if o == nil || o.Zip == nil {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetZipOk() (*string, bool) {
	if o == nil || o.Zip == nil {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *Location) HasZip() bool {
	if o != nil && o.Zip != nil {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *Location) SetZip(v string) {
	o.Zip = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *Location) GetLng() float32 {
	if o == nil || o.Lng == nil {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLngOk() (*float32, bool) {
	if o == nil || o.Lng == nil {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *Location) HasLng() bool {
	if o != nil && o.Lng != nil {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *Location) SetLng(v float32) {
	o.Lng = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *Location) GetLat() float32 {
	if o == nil || o.Lat == nil {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLatOk() (*float32, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *Location) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *Location) SetLat(v float32) {
	o.Lat = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Location) GetExternalId() int32 {
	if o == nil || o.ExternalId == nil {
		var ret int32
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetExternalIdOk() (*int32, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Location) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given int32 and assigns it to the ExternalId field.
func (o *Location) SetExternalId(v int32) {
	o.ExternalId = &v
}

// GetExternalIdSource returns the ExternalIdSource field value if set, zero value otherwise.
func (o *Location) GetExternalIdSource() string {
	if o == nil || o.ExternalIdSource == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdSource
}

// GetExternalIdSourceOk returns a tuple with the ExternalIdSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetExternalIdSourceOk() (*string, bool) {
	if o == nil || o.ExternalIdSource == nil {
		return nil, false
	}
	return o.ExternalIdSource, true
}

// HasExternalIdSource returns a boolean if a field has been set.
func (o *Location) HasExternalIdSource() bool {
	if o != nil && o.ExternalIdSource != nil {
		return true
	}

	return false
}

// SetExternalIdSource gets a reference to the given string and assigns it to the ExternalIdSource field.
func (o *Location) SetExternalIdSource(v string) {
	o.ExternalIdSource = &v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pk != nil {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Hours != nil {
		toSerialize["hours"] = o.Hours
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Zip != nil {
		toSerialize["zip"] = o.Zip
	}
	if o.Lng != nil {
		toSerialize["lng"] = o.Lng
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.ExternalIdSource != nil {
		toSerialize["external_id_source"] = o.ExternalIdSource
	}
	return json.Marshal(toSerialize)
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


