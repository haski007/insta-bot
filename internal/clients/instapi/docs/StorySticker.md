# StorySticker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "gif"]
**X** | **float32** |  | 
**Y** | **float32** |  | 
**Z** | Pointer to **int32** |  | [optional] [default to 1000005]
**Width** | **float32** |  | 
**Height** | **float32** |  | 
**Rotation** | Pointer to **float32** |  | [optional] [default to 0.0]
**Extra** | **map[string]interface{}** |  | 

## Methods

### NewStorySticker

`func NewStorySticker(x float32, y float32, width float32, height float32, extra map[string]interface{}, ) *StorySticker`

NewStorySticker instantiates a new StorySticker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryStickerWithDefaults

`func NewStoryStickerWithDefaults() *StorySticker`

NewStoryStickerWithDefaults instantiates a new StorySticker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorySticker) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorySticker) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorySticker) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StorySticker) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *StorySticker) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorySticker) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorySticker) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorySticker) HasType() bool`

HasType returns a boolean if a field has been set.

### GetX

`func (o *StorySticker) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *StorySticker) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *StorySticker) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *StorySticker) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *StorySticker) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *StorySticker) SetY(v float32)`

SetY sets Y field to given value.


### GetZ

`func (o *StorySticker) GetZ() int32`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *StorySticker) GetZOk() (*int32, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *StorySticker) SetZ(v int32)`

SetZ sets Z field to given value.

### HasZ

`func (o *StorySticker) HasZ() bool`

HasZ returns a boolean if a field has been set.

### GetWidth

`func (o *StorySticker) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *StorySticker) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *StorySticker) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *StorySticker) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *StorySticker) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *StorySticker) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetRotation

`func (o *StorySticker) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *StorySticker) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *StorySticker) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *StorySticker) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetExtra

`func (o *StorySticker) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *StorySticker) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *StorySticker) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


