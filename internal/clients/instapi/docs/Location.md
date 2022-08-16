# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Phone** | Pointer to **string** |  | [optional] [default to ""]
**Website** | Pointer to **string** |  | [optional] [default to ""]
**Category** | Pointer to **string** |  | [optional] [default to ""]
**Hours** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Address** | Pointer to **string** |  | [optional] [default to ""]
**City** | Pointer to **string** |  | [optional] [default to ""]
**Zip** | Pointer to **string** |  | [optional] [default to ""]
**Lng** | Pointer to **float32** |  | [optional] 
**Lat** | Pointer to **float32** |  | [optional] 
**ExternalId** | Pointer to **int32** |  | [optional] 
**ExternalIdSource** | Pointer to **string** |  | [optional] 

## Methods

### NewLocation

`func NewLocation(name string, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Location) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Location) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Location) SetPk(v int32)`

SetPk sets Pk field to given value.

### HasPk

`func (o *Location) HasPk() bool`

HasPk returns a boolean if a field has been set.

### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *Location) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Location) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Location) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Location) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetWebsite

`func (o *Location) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Location) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Location) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Location) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetCategory

`func (o *Location) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Location) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Location) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Location) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetHours

`func (o *Location) GetHours() map[string]interface{}`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *Location) GetHoursOk() (*map[string]interface{}, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *Location) SetHours(v map[string]interface{})`

SetHours sets Hours field to given value.

### HasHours

`func (o *Location) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetAddress

`func (o *Location) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Location) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Location) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Location) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *Location) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Location) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Location) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Location) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZip

`func (o *Location) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *Location) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *Location) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *Location) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetLng

`func (o *Location) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *Location) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *Location) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *Location) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetLat

`func (o *Location) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *Location) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *Location) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *Location) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetExternalId

`func (o *Location) GetExternalId() int32`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Location) GetExternalIdOk() (*int32, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Location) SetExternalId(v int32)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Location) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalIdSource

`func (o *Location) GetExternalIdSource() string`

GetExternalIdSource returns the ExternalIdSource field if non-nil, zero value otherwise.

### GetExternalIdSourceOk

`func (o *Location) GetExternalIdSourceOk() (*string, bool)`

GetExternalIdSourceOk returns a tuple with the ExternalIdSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdSource

`func (o *Location) SetExternalIdSource(v string)`

SetExternalIdSource sets ExternalIdSource field to given value.

### HasExternalIdSource

`func (o *Location) HasExternalIdSource() bool`

HasExternalIdSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


