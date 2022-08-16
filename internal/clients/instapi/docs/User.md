# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | 
**Username** | **string** |  | 
**FullName** | **string** |  | 
**IsPrivate** | **bool** |  | 
**ProfilePicUrl** | **string** |  | 
**ProfilePicUrlHd** | Pointer to **string** |  | [optional] 
**IsVerified** | **bool** |  | 
**MediaCount** | **int32** |  | 
**FollowerCount** | **int32** |  | 
**FollowingCount** | **int32** |  | 
**Biography** | Pointer to **string** |  | [optional] [default to ""]
**ExternalUrl** | Pointer to **string** |  | [optional] 
**IsBusiness** | **bool** |  | 
**PublicEmail** | Pointer to **string** |  | [optional] 
**ContactPhoneNumber** | Pointer to **string** |  | [optional] 
**BusinessContactMethod** | Pointer to **string** |  | [optional] 
**BusinessCategoryName** | Pointer to **string** |  | [optional] 
**CategoryName** | Pointer to **string** |  | [optional] 

## Methods

### NewUser

`func NewUser(pk string, username string, fullName string, isPrivate bool, profilePicUrl string, isVerified bool, mediaCount int32, followerCount int32, followingCount int32, isBusiness bool, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *User) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *User) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *User) SetPk(v string)`

SetPk sets Pk field to given value.


### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFullName

`func (o *User) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *User) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *User) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetIsPrivate

`func (o *User) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *User) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *User) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.


### GetProfilePicUrl

`func (o *User) GetProfilePicUrl() string`

GetProfilePicUrl returns the ProfilePicUrl field if non-nil, zero value otherwise.

### GetProfilePicUrlOk

`func (o *User) GetProfilePicUrlOk() (*string, bool)`

GetProfilePicUrlOk returns a tuple with the ProfilePicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicUrl

`func (o *User) SetProfilePicUrl(v string)`

SetProfilePicUrl sets ProfilePicUrl field to given value.


### GetProfilePicUrlHd

`func (o *User) GetProfilePicUrlHd() string`

GetProfilePicUrlHd returns the ProfilePicUrlHd field if non-nil, zero value otherwise.

### GetProfilePicUrlHdOk

`func (o *User) GetProfilePicUrlHdOk() (*string, bool)`

GetProfilePicUrlHdOk returns a tuple with the ProfilePicUrlHd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicUrlHd

`func (o *User) SetProfilePicUrlHd(v string)`

SetProfilePicUrlHd sets ProfilePicUrlHd field to given value.

### HasProfilePicUrlHd

`func (o *User) HasProfilePicUrlHd() bool`

HasProfilePicUrlHd returns a boolean if a field has been set.

### GetIsVerified

`func (o *User) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *User) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *User) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.


### GetMediaCount

`func (o *User) GetMediaCount() int32`

GetMediaCount returns the MediaCount field if non-nil, zero value otherwise.

### GetMediaCountOk

`func (o *User) GetMediaCountOk() (*int32, bool)`

GetMediaCountOk returns a tuple with the MediaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaCount

`func (o *User) SetMediaCount(v int32)`

SetMediaCount sets MediaCount field to given value.


### GetFollowerCount

`func (o *User) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *User) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *User) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.


### GetFollowingCount

`func (o *User) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *User) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *User) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.


### GetBiography

`func (o *User) GetBiography() string`

GetBiography returns the Biography field if non-nil, zero value otherwise.

### GetBiographyOk

`func (o *User) GetBiographyOk() (*string, bool)`

GetBiographyOk returns a tuple with the Biography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiography

`func (o *User) SetBiography(v string)`

SetBiography sets Biography field to given value.

### HasBiography

`func (o *User) HasBiography() bool`

HasBiography returns a boolean if a field has been set.

### GetExternalUrl

`func (o *User) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *User) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *User) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *User) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetIsBusiness

`func (o *User) GetIsBusiness() bool`

GetIsBusiness returns the IsBusiness field if non-nil, zero value otherwise.

### GetIsBusinessOk

`func (o *User) GetIsBusinessOk() (*bool, bool)`

GetIsBusinessOk returns a tuple with the IsBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusiness

`func (o *User) SetIsBusiness(v bool)`

SetIsBusiness sets IsBusiness field to given value.


### GetPublicEmail

`func (o *User) GetPublicEmail() string`

GetPublicEmail returns the PublicEmail field if non-nil, zero value otherwise.

### GetPublicEmailOk

`func (o *User) GetPublicEmailOk() (*string, bool)`

GetPublicEmailOk returns a tuple with the PublicEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEmail

`func (o *User) SetPublicEmail(v string)`

SetPublicEmail sets PublicEmail field to given value.

### HasPublicEmail

`func (o *User) HasPublicEmail() bool`

HasPublicEmail returns a boolean if a field has been set.

### GetContactPhoneNumber

`func (o *User) GetContactPhoneNumber() string`

GetContactPhoneNumber returns the ContactPhoneNumber field if non-nil, zero value otherwise.

### GetContactPhoneNumberOk

`func (o *User) GetContactPhoneNumberOk() (*string, bool)`

GetContactPhoneNumberOk returns a tuple with the ContactPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhoneNumber

`func (o *User) SetContactPhoneNumber(v string)`

SetContactPhoneNumber sets ContactPhoneNumber field to given value.

### HasContactPhoneNumber

`func (o *User) HasContactPhoneNumber() bool`

HasContactPhoneNumber returns a boolean if a field has been set.

### GetBusinessContactMethod

`func (o *User) GetBusinessContactMethod() string`

GetBusinessContactMethod returns the BusinessContactMethod field if non-nil, zero value otherwise.

### GetBusinessContactMethodOk

`func (o *User) GetBusinessContactMethodOk() (*string, bool)`

GetBusinessContactMethodOk returns a tuple with the BusinessContactMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactMethod

`func (o *User) SetBusinessContactMethod(v string)`

SetBusinessContactMethod sets BusinessContactMethod field to given value.

### HasBusinessContactMethod

`func (o *User) HasBusinessContactMethod() bool`

HasBusinessContactMethod returns a boolean if a field has been set.

### GetBusinessCategoryName

`func (o *User) GetBusinessCategoryName() string`

GetBusinessCategoryName returns the BusinessCategoryName field if non-nil, zero value otherwise.

### GetBusinessCategoryNameOk

`func (o *User) GetBusinessCategoryNameOk() (*string, bool)`

GetBusinessCategoryNameOk returns a tuple with the BusinessCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCategoryName

`func (o *User) SetBusinessCategoryName(v string)`

SetBusinessCategoryName sets BusinessCategoryName field to given value.

### HasBusinessCategoryName

`func (o *User) HasBusinessCategoryName() bool`

HasBusinessCategoryName returns a boolean if a field has been set.

### GetCategoryName

`func (o *User) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *User) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *User) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *User) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


