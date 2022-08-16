# UserShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] [default to ""]
**ProfilePicUrl** | Pointer to **string** |  | [optional] 
**ProfilePicUrlHd** | Pointer to **string** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**Stories** | Pointer to **[]interface{}** |  | [optional] [default to []]

## Methods

### NewUserShort

`func NewUserShort(pk string, ) *UserShort`

NewUserShort instantiates a new UserShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserShortWithDefaults

`func NewUserShortWithDefaults() *UserShort`

NewUserShortWithDefaults instantiates a new UserShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserShort) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserShort) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserShort) SetPk(v string)`

SetPk sets Pk field to given value.


### GetUsername

`func (o *UserShort) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserShort) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserShort) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserShort) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFullName

`func (o *UserShort) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserShort) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserShort) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UserShort) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetProfilePicUrl

`func (o *UserShort) GetProfilePicUrl() string`

GetProfilePicUrl returns the ProfilePicUrl field if non-nil, zero value otherwise.

### GetProfilePicUrlOk

`func (o *UserShort) GetProfilePicUrlOk() (*string, bool)`

GetProfilePicUrlOk returns a tuple with the ProfilePicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicUrl

`func (o *UserShort) SetProfilePicUrl(v string)`

SetProfilePicUrl sets ProfilePicUrl field to given value.

### HasProfilePicUrl

`func (o *UserShort) HasProfilePicUrl() bool`

HasProfilePicUrl returns a boolean if a field has been set.

### GetProfilePicUrlHd

`func (o *UserShort) GetProfilePicUrlHd() string`

GetProfilePicUrlHd returns the ProfilePicUrlHd field if non-nil, zero value otherwise.

### GetProfilePicUrlHdOk

`func (o *UserShort) GetProfilePicUrlHdOk() (*string, bool)`

GetProfilePicUrlHdOk returns a tuple with the ProfilePicUrlHd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicUrlHd

`func (o *UserShort) SetProfilePicUrlHd(v string)`

SetProfilePicUrlHd sets ProfilePicUrlHd field to given value.

### HasProfilePicUrlHd

`func (o *UserShort) HasProfilePicUrlHd() bool`

HasProfilePicUrlHd returns a boolean if a field has been set.

### GetIsPrivate

`func (o *UserShort) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *UserShort) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *UserShort) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *UserShort) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetStories

`func (o *UserShort) GetStories() []interface{}`

GetStories returns the Stories field if non-nil, zero value otherwise.

### GetStoriesOk

`func (o *UserShort) GetStoriesOk() (*[]interface{}, bool)`

GetStoriesOk returns a tuple with the Stories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStories

`func (o *UserShort) SetStories(v []interface{})`

SetStories sets Stories field to given value.

### HasStories

`func (o *UserShort) HasStories() bool`

HasStories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


