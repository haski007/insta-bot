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

// User struct for User
type User struct {
	Pk string `json:"pk"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	IsPrivate bool `json:"is_private"`
	ProfilePicUrl string `json:"profile_pic_url"`
	ProfilePicUrlHd *string `json:"profile_pic_url_hd,omitempty"`
	IsVerified bool `json:"is_verified"`
	MediaCount int32 `json:"media_count"`
	FollowerCount int32 `json:"follower_count"`
	FollowingCount int32 `json:"following_count"`
	Biography *string `json:"biography,omitempty"`
	ExternalUrl *string `json:"external_url,omitempty"`
	IsBusiness bool `json:"is_business"`
	PublicEmail *string `json:"public_email,omitempty"`
	ContactPhoneNumber *string `json:"contact_phone_number,omitempty"`
	BusinessContactMethod *string `json:"business_contact_method,omitempty"`
	BusinessCategoryName *string `json:"business_category_name,omitempty"`
	CategoryName *string `json:"category_name,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(pk string, username string, fullName string, isPrivate bool, profilePicUrl string, isVerified bool, mediaCount int32, followerCount int32, followingCount int32, isBusiness bool) *User {
	this := User{}
	this.Pk = pk
	this.Username = username
	this.FullName = fullName
	this.IsPrivate = isPrivate
	this.ProfilePicUrl = profilePicUrl
	this.IsVerified = isVerified
	this.MediaCount = mediaCount
	this.FollowerCount = followerCount
	this.FollowingCount = followingCount
	var biography string = ""
	this.Biography = &biography
	this.IsBusiness = isBusiness
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	var biography string = ""
	this.Biography = &biography
	return &this
}

// GetPk returns the Pk field value
func (o *User) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *User) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *User) SetPk(v string) {
	o.Pk = v
}

// GetUsername returns the Username field value
func (o *User) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *User) SetUsername(v string) {
	o.Username = v
}

// GetFullName returns the FullName field value
func (o *User) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *User) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *User) SetFullName(v string) {
	o.FullName = v
}

// GetIsPrivate returns the IsPrivate field value
func (o *User) GetIsPrivate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value
// and a boolean to check if the value has been set.
func (o *User) GetIsPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrivate, true
}

// SetIsPrivate sets field value
func (o *User) SetIsPrivate(v bool) {
	o.IsPrivate = v
}

// GetProfilePicUrl returns the ProfilePicUrl field value
func (o *User) GetProfilePicUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfilePicUrl
}

// GetProfilePicUrlOk returns a tuple with the ProfilePicUrl field value
// and a boolean to check if the value has been set.
func (o *User) GetProfilePicUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilePicUrl, true
}

// SetProfilePicUrl sets field value
func (o *User) SetProfilePicUrl(v string) {
	o.ProfilePicUrl = v
}

// GetProfilePicUrlHd returns the ProfilePicUrlHd field value if set, zero value otherwise.
func (o *User) GetProfilePicUrlHd() string {
	if o == nil || o.ProfilePicUrlHd == nil {
		var ret string
		return ret
	}
	return *o.ProfilePicUrlHd
}

// GetProfilePicUrlHdOk returns a tuple with the ProfilePicUrlHd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetProfilePicUrlHdOk() (*string, bool) {
	if o == nil || o.ProfilePicUrlHd == nil {
		return nil, false
	}
	return o.ProfilePicUrlHd, true
}

// HasProfilePicUrlHd returns a boolean if a field has been set.
func (o *User) HasProfilePicUrlHd() bool {
	if o != nil && o.ProfilePicUrlHd != nil {
		return true
	}

	return false
}

// SetProfilePicUrlHd gets a reference to the given string and assigns it to the ProfilePicUrlHd field.
func (o *User) SetProfilePicUrlHd(v string) {
	o.ProfilePicUrlHd = &v
}

// GetIsVerified returns the IsVerified field value
func (o *User) GetIsVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsVerified
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value
// and a boolean to check if the value has been set.
func (o *User) GetIsVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsVerified, true
}

// SetIsVerified sets field value
func (o *User) SetIsVerified(v bool) {
	o.IsVerified = v
}

// GetMediaCount returns the MediaCount field value
func (o *User) GetMediaCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MediaCount
}

// GetMediaCountOk returns a tuple with the MediaCount field value
// and a boolean to check if the value has been set.
func (o *User) GetMediaCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaCount, true
}

// SetMediaCount sets field value
func (o *User) SetMediaCount(v int32) {
	o.MediaCount = v
}

// GetFollowerCount returns the FollowerCount field value
func (o *User) GetFollowerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowerCount
}

// GetFollowerCountOk returns a tuple with the FollowerCount field value
// and a boolean to check if the value has been set.
func (o *User) GetFollowerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowerCount, true
}

// SetFollowerCount sets field value
func (o *User) SetFollowerCount(v int32) {
	o.FollowerCount = v
}

// GetFollowingCount returns the FollowingCount field value
func (o *User) GetFollowingCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowingCount
}

// GetFollowingCountOk returns a tuple with the FollowingCount field value
// and a boolean to check if the value has been set.
func (o *User) GetFollowingCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingCount, true
}

// SetFollowingCount sets field value
func (o *User) SetFollowingCount(v int32) {
	o.FollowingCount = v
}

// GetBiography returns the Biography field value if set, zero value otherwise.
func (o *User) GetBiography() string {
	if o == nil || o.Biography == nil {
		var ret string
		return ret
	}
	return *o.Biography
}

// GetBiographyOk returns a tuple with the Biography field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBiographyOk() (*string, bool) {
	if o == nil || o.Biography == nil {
		return nil, false
	}
	return o.Biography, true
}

// HasBiography returns a boolean if a field has been set.
func (o *User) HasBiography() bool {
	if o != nil && o.Biography != nil {
		return true
	}

	return false
}

// SetBiography gets a reference to the given string and assigns it to the Biography field.
func (o *User) SetBiography(v string) {
	o.Biography = &v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *User) GetExternalUrl() string {
	if o == nil || o.ExternalUrl == nil {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetExternalUrlOk() (*string, bool) {
	if o == nil || o.ExternalUrl == nil {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *User) HasExternalUrl() bool {
	if o != nil && o.ExternalUrl != nil {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *User) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetIsBusiness returns the IsBusiness field value
func (o *User) GetIsBusiness() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBusiness
}

// GetIsBusinessOk returns a tuple with the IsBusiness field value
// and a boolean to check if the value has been set.
func (o *User) GetIsBusinessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBusiness, true
}

// SetIsBusiness sets field value
func (o *User) SetIsBusiness(v bool) {
	o.IsBusiness = v
}

// GetPublicEmail returns the PublicEmail field value if set, zero value otherwise.
func (o *User) GetPublicEmail() string {
	if o == nil || o.PublicEmail == nil {
		var ret string
		return ret
	}
	return *o.PublicEmail
}

// GetPublicEmailOk returns a tuple with the PublicEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPublicEmailOk() (*string, bool) {
	if o == nil || o.PublicEmail == nil {
		return nil, false
	}
	return o.PublicEmail, true
}

// HasPublicEmail returns a boolean if a field has been set.
func (o *User) HasPublicEmail() bool {
	if o != nil && o.PublicEmail != nil {
		return true
	}

	return false
}

// SetPublicEmail gets a reference to the given string and assigns it to the PublicEmail field.
func (o *User) SetPublicEmail(v string) {
	o.PublicEmail = &v
}

// GetContactPhoneNumber returns the ContactPhoneNumber field value if set, zero value otherwise.
func (o *User) GetContactPhoneNumber() string {
	if o == nil || o.ContactPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.ContactPhoneNumber
}

// GetContactPhoneNumberOk returns a tuple with the ContactPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetContactPhoneNumberOk() (*string, bool) {
	if o == nil || o.ContactPhoneNumber == nil {
		return nil, false
	}
	return o.ContactPhoneNumber, true
}

// HasContactPhoneNumber returns a boolean if a field has been set.
func (o *User) HasContactPhoneNumber() bool {
	if o != nil && o.ContactPhoneNumber != nil {
		return true
	}

	return false
}

// SetContactPhoneNumber gets a reference to the given string and assigns it to the ContactPhoneNumber field.
func (o *User) SetContactPhoneNumber(v string) {
	o.ContactPhoneNumber = &v
}

// GetBusinessContactMethod returns the BusinessContactMethod field value if set, zero value otherwise.
func (o *User) GetBusinessContactMethod() string {
	if o == nil || o.BusinessContactMethod == nil {
		var ret string
		return ret
	}
	return *o.BusinessContactMethod
}

// GetBusinessContactMethodOk returns a tuple with the BusinessContactMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBusinessContactMethodOk() (*string, bool) {
	if o == nil || o.BusinessContactMethod == nil {
		return nil, false
	}
	return o.BusinessContactMethod, true
}

// HasBusinessContactMethod returns a boolean if a field has been set.
func (o *User) HasBusinessContactMethod() bool {
	if o != nil && o.BusinessContactMethod != nil {
		return true
	}

	return false
}

// SetBusinessContactMethod gets a reference to the given string and assigns it to the BusinessContactMethod field.
func (o *User) SetBusinessContactMethod(v string) {
	o.BusinessContactMethod = &v
}

// GetBusinessCategoryName returns the BusinessCategoryName field value if set, zero value otherwise.
func (o *User) GetBusinessCategoryName() string {
	if o == nil || o.BusinessCategoryName == nil {
		var ret string
		return ret
	}
	return *o.BusinessCategoryName
}

// GetBusinessCategoryNameOk returns a tuple with the BusinessCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBusinessCategoryNameOk() (*string, bool) {
	if o == nil || o.BusinessCategoryName == nil {
		return nil, false
	}
	return o.BusinessCategoryName, true
}

// HasBusinessCategoryName returns a boolean if a field has been set.
func (o *User) HasBusinessCategoryName() bool {
	if o != nil && o.BusinessCategoryName != nil {
		return true
	}

	return false
}

// SetBusinessCategoryName gets a reference to the given string and assigns it to the BusinessCategoryName field.
func (o *User) SetBusinessCategoryName(v string) {
	o.BusinessCategoryName = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *User) GetCategoryName() string {
	if o == nil || o.CategoryName == nil {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCategoryNameOk() (*string, bool) {
	if o == nil || o.CategoryName == nil {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *User) HasCategoryName() bool {
	if o != nil && o.CategoryName != nil {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *User) SetCategoryName(v string) {
	o.CategoryName = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["full_name"] = o.FullName
	}
	if true {
		toSerialize["is_private"] = o.IsPrivate
	}
	if true {
		toSerialize["profile_pic_url"] = o.ProfilePicUrl
	}
	if o.ProfilePicUrlHd != nil {
		toSerialize["profile_pic_url_hd"] = o.ProfilePicUrlHd
	}
	if true {
		toSerialize["is_verified"] = o.IsVerified
	}
	if true {
		toSerialize["media_count"] = o.MediaCount
	}
	if true {
		toSerialize["follower_count"] = o.FollowerCount
	}
	if true {
		toSerialize["following_count"] = o.FollowingCount
	}
	if o.Biography != nil {
		toSerialize["biography"] = o.Biography
	}
	if o.ExternalUrl != nil {
		toSerialize["external_url"] = o.ExternalUrl
	}
	if true {
		toSerialize["is_business"] = o.IsBusiness
	}
	if o.PublicEmail != nil {
		toSerialize["public_email"] = o.PublicEmail
	}
	if o.ContactPhoneNumber != nil {
		toSerialize["contact_phone_number"] = o.ContactPhoneNumber
	}
	if o.BusinessContactMethod != nil {
		toSerialize["business_contact_method"] = o.BusinessContactMethod
	}
	if o.BusinessCategoryName != nil {
		toSerialize["business_category_name"] = o.BusinessCategoryName
	}
	if o.CategoryName != nil {
		toSerialize["category_name"] = o.CategoryName
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


