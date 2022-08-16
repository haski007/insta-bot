# StoryMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | Pointer to **float32** |  | [optional] [default to 0.5]
**Y** | Pointer to **float32** |  | [optional] [default to 0.4997396]
**Z** | Pointer to **float32** |  | [optional] [default to 0]
**Width** | Pointer to **float32** |  | [optional] [default to 0.8]
**Height** | Pointer to **float32** |  | [optional] [default to 0.60572916]
**Rotation** | Pointer to **float32** |  | [optional] [default to 0.0]
**IsPinned** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsSticker** | Pointer to **bool** |  | [optional] 
**IsFbSticker** | Pointer to **bool** |  | [optional] 
**MediaPk** | **int32** |  | 
**UserId** | Pointer to **int32** |  | [optional] 
**ProductType** | Pointer to **string** |  | [optional] 
**MediaCode** | Pointer to **string** |  | [optional] 

## Methods

### NewStoryMedia

`func NewStoryMedia(mediaPk int32, ) *StoryMedia`

NewStoryMedia instantiates a new StoryMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryMediaWithDefaults

`func NewStoryMediaWithDefaults() *StoryMedia`

NewStoryMediaWithDefaults instantiates a new StoryMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *StoryMedia) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *StoryMedia) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *StoryMedia) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *StoryMedia) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *StoryMedia) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *StoryMedia) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *StoryMedia) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *StoryMedia) HasY() bool`

HasY returns a boolean if a field has been set.

### GetZ

`func (o *StoryMedia) GetZ() float32`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *StoryMedia) GetZOk() (*float32, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *StoryMedia) SetZ(v float32)`

SetZ sets Z field to given value.

### HasZ

`func (o *StoryMedia) HasZ() bool`

HasZ returns a boolean if a field has been set.

### GetWidth

`func (o *StoryMedia) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *StoryMedia) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *StoryMedia) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *StoryMedia) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *StoryMedia) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *StoryMedia) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *StoryMedia) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *StoryMedia) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetRotation

`func (o *StoryMedia) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *StoryMedia) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *StoryMedia) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *StoryMedia) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetIsPinned

`func (o *StoryMedia) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *StoryMedia) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *StoryMedia) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *StoryMedia) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### GetIsHidden

`func (o *StoryMedia) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *StoryMedia) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *StoryMedia) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *StoryMedia) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsSticker

`func (o *StoryMedia) GetIsSticker() bool`

GetIsSticker returns the IsSticker field if non-nil, zero value otherwise.

### GetIsStickerOk

`func (o *StoryMedia) GetIsStickerOk() (*bool, bool)`

GetIsStickerOk returns a tuple with the IsSticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSticker

`func (o *StoryMedia) SetIsSticker(v bool)`

SetIsSticker sets IsSticker field to given value.

### HasIsSticker

`func (o *StoryMedia) HasIsSticker() bool`

HasIsSticker returns a boolean if a field has been set.

### GetIsFbSticker

`func (o *StoryMedia) GetIsFbSticker() bool`

GetIsFbSticker returns the IsFbSticker field if non-nil, zero value otherwise.

### GetIsFbStickerOk

`func (o *StoryMedia) GetIsFbStickerOk() (*bool, bool)`

GetIsFbStickerOk returns a tuple with the IsFbSticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFbSticker

`func (o *StoryMedia) SetIsFbSticker(v bool)`

SetIsFbSticker sets IsFbSticker field to given value.

### HasIsFbSticker

`func (o *StoryMedia) HasIsFbSticker() bool`

HasIsFbSticker returns a boolean if a field has been set.

### GetMediaPk

`func (o *StoryMedia) GetMediaPk() int32`

GetMediaPk returns the MediaPk field if non-nil, zero value otherwise.

### GetMediaPkOk

`func (o *StoryMedia) GetMediaPkOk() (*int32, bool)`

GetMediaPkOk returns a tuple with the MediaPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaPk

`func (o *StoryMedia) SetMediaPk(v int32)`

SetMediaPk sets MediaPk field to given value.


### GetUserId

`func (o *StoryMedia) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *StoryMedia) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *StoryMedia) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *StoryMedia) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetProductType

`func (o *StoryMedia) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *StoryMedia) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *StoryMedia) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *StoryMedia) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetMediaCode

`func (o *StoryMedia) GetMediaCode() string`

GetMediaCode returns the MediaCode field if non-nil, zero value otherwise.

### GetMediaCodeOk

`func (o *StoryMedia) GetMediaCodeOk() (*string, bool)`

GetMediaCodeOk returns a tuple with the MediaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaCode

`func (o *StoryMedia) SetMediaCode(v string)`

SetMediaCode sets MediaCode field to given value.

### HasMediaCode

`func (o *StoryMedia) HasMediaCode() bool`

HasMediaCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


