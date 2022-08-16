# Media

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | 
**Id** | **string** |  | 
**Code** | **string** |  | 
**TakenAt** | **time.Time** |  | 
**MediaType** | **int32** |  | 
**ProductType** | Pointer to **string** |  | [optional] [default to ""]
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**User** | [**UserShort**](UserShort.md) |  | 
**CommentCount** | Pointer to **int32** |  | [optional] [default to 0]
**LikeCount** | **int32** |  | 
**HasLiked** | Pointer to **bool** |  | [optional] 
**CaptionText** | **string** |  | 
**Usertags** | [**[]Usertag**](Usertag.md) |  | 
**VideoUrl** | Pointer to **string** |  | [optional] 
**ViewCount** | Pointer to **int32** |  | [optional] [default to 0]
**VideoDuration** | Pointer to **float32** |  | [optional] [default to 0.0]
**Title** | Pointer to **string** |  | [optional] [default to ""]
**Resources** | Pointer to [**[]Resource**](Resource.md) |  | [optional] [default to []]
**ClipsMetadata** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]

## Methods

### NewMedia

`func NewMedia(pk string, id string, code string, takenAt time.Time, mediaType int32, user UserShort, likeCount int32, captionText string, usertags []Usertag, ) *Media`

NewMedia instantiates a new Media object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaWithDefaults

`func NewMediaWithDefaults() *Media`

NewMediaWithDefaults instantiates a new Media object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Media) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Media) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Media) SetPk(v string)`

SetPk sets Pk field to given value.


### GetId

`func (o *Media) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Media) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Media) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *Media) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Media) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Media) SetCode(v string)`

SetCode sets Code field to given value.


### GetTakenAt

`func (o *Media) GetTakenAt() time.Time`

GetTakenAt returns the TakenAt field if non-nil, zero value otherwise.

### GetTakenAtOk

`func (o *Media) GetTakenAtOk() (*time.Time, bool)`

GetTakenAtOk returns a tuple with the TakenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakenAt

`func (o *Media) SetTakenAt(v time.Time)`

SetTakenAt sets TakenAt field to given value.


### GetMediaType

`func (o *Media) GetMediaType() int32`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *Media) GetMediaTypeOk() (*int32, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *Media) SetMediaType(v int32)`

SetMediaType sets MediaType field to given value.


### GetProductType

`func (o *Media) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Media) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Media) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *Media) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *Media) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *Media) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *Media) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *Media) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetLocation

`func (o *Media) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Media) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Media) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Media) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUser

`func (o *Media) GetUser() UserShort`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Media) GetUserOk() (*UserShort, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Media) SetUser(v UserShort)`

SetUser sets User field to given value.


### GetCommentCount

`func (o *Media) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *Media) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *Media) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *Media) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *Media) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *Media) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *Media) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.


### GetHasLiked

`func (o *Media) GetHasLiked() bool`

GetHasLiked returns the HasLiked field if non-nil, zero value otherwise.

### GetHasLikedOk

`func (o *Media) GetHasLikedOk() (*bool, bool)`

GetHasLikedOk returns a tuple with the HasLiked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLiked

`func (o *Media) SetHasLiked(v bool)`

SetHasLiked sets HasLiked field to given value.

### HasHasLiked

`func (o *Media) HasHasLiked() bool`

HasHasLiked returns a boolean if a field has been set.

### GetCaptionText

`func (o *Media) GetCaptionText() string`

GetCaptionText returns the CaptionText field if non-nil, zero value otherwise.

### GetCaptionTextOk

`func (o *Media) GetCaptionTextOk() (*string, bool)`

GetCaptionTextOk returns a tuple with the CaptionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionText

`func (o *Media) SetCaptionText(v string)`

SetCaptionText sets CaptionText field to given value.


### GetUsertags

`func (o *Media) GetUsertags() []Usertag`

GetUsertags returns the Usertags field if non-nil, zero value otherwise.

### GetUsertagsOk

`func (o *Media) GetUsertagsOk() (*[]Usertag, bool)`

GetUsertagsOk returns a tuple with the Usertags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsertags

`func (o *Media) SetUsertags(v []Usertag)`

SetUsertags sets Usertags field to given value.


### GetVideoUrl

`func (o *Media) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *Media) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *Media) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *Media) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetViewCount

`func (o *Media) GetViewCount() int32`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *Media) GetViewCountOk() (*int32, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *Media) SetViewCount(v int32)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *Media) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetVideoDuration

`func (o *Media) GetVideoDuration() float32`

GetVideoDuration returns the VideoDuration field if non-nil, zero value otherwise.

### GetVideoDurationOk

`func (o *Media) GetVideoDurationOk() (*float32, bool)`

GetVideoDurationOk returns a tuple with the VideoDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDuration

`func (o *Media) SetVideoDuration(v float32)`

SetVideoDuration sets VideoDuration field to given value.

### HasVideoDuration

`func (o *Media) HasVideoDuration() bool`

HasVideoDuration returns a boolean if a field has been set.

### GetTitle

`func (o *Media) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Media) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Media) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Media) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetResources

`func (o *Media) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Media) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Media) SetResources(v []Resource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Media) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetClipsMetadata

`func (o *Media) GetClipsMetadata() map[string]interface{}`

GetClipsMetadata returns the ClipsMetadata field if non-nil, zero value otherwise.

### GetClipsMetadataOk

`func (o *Media) GetClipsMetadataOk() (*map[string]interface{}, bool)`

GetClipsMetadataOk returns a tuple with the ClipsMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipsMetadata

`func (o *Media) SetClipsMetadata(v map[string]interface{})`

SetClipsMetadata sets ClipsMetadata field to given value.

### HasClipsMetadata

`func (o *Media) HasClipsMetadata() bool`

HasClipsMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


