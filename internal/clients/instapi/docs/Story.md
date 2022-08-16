# Story

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
**User** | [**UserShort**](UserShort.md) |  | 
**VideoUrl** | Pointer to **string** |  | [optional] 
**VideoDuration** | Pointer to **float32** |  | [optional] [default to 0.0]
**Mentions** | [**[]StoryMention**](StoryMention.md) |  | 
**Links** | [**[]StoryLink**](StoryLink.md) |  | 
**Hashtags** | [**[]StoryHashtag**](StoryHashtag.md) |  | 
**Locations** | [**[]StoryLocation**](StoryLocation.md) |  | 
**Stickers** | [**[]StorySticker**](StorySticker.md) |  | 
**Medias** | Pointer to [**[]StoryMedia**](StoryMedia.md) |  | [optional] [default to []]

## Methods

### NewStory

`func NewStory(pk string, id string, code string, takenAt time.Time, mediaType int32, user UserShort, mentions []StoryMention, links []StoryLink, hashtags []StoryHashtag, locations []StoryLocation, stickers []StorySticker, ) *Story`

NewStory instantiates a new Story object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryWithDefaults

`func NewStoryWithDefaults() *Story`

NewStoryWithDefaults instantiates a new Story object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Story) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Story) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Story) SetPk(v string)`

SetPk sets Pk field to given value.


### GetId

`func (o *Story) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Story) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Story) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *Story) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Story) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Story) SetCode(v string)`

SetCode sets Code field to given value.


### GetTakenAt

`func (o *Story) GetTakenAt() time.Time`

GetTakenAt returns the TakenAt field if non-nil, zero value otherwise.

### GetTakenAtOk

`func (o *Story) GetTakenAtOk() (*time.Time, bool)`

GetTakenAtOk returns a tuple with the TakenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakenAt

`func (o *Story) SetTakenAt(v time.Time)`

SetTakenAt sets TakenAt field to given value.


### GetMediaType

`func (o *Story) GetMediaType() int32`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *Story) GetMediaTypeOk() (*int32, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *Story) SetMediaType(v int32)`

SetMediaType sets MediaType field to given value.


### GetProductType

`func (o *Story) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Story) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Story) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *Story) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *Story) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *Story) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *Story) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *Story) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetUser

`func (o *Story) GetUser() UserShort`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Story) GetUserOk() (*UserShort, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Story) SetUser(v UserShort)`

SetUser sets User field to given value.


### GetVideoUrl

`func (o *Story) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *Story) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *Story) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *Story) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetVideoDuration

`func (o *Story) GetVideoDuration() float32`

GetVideoDuration returns the VideoDuration field if non-nil, zero value otherwise.

### GetVideoDurationOk

`func (o *Story) GetVideoDurationOk() (*float32, bool)`

GetVideoDurationOk returns a tuple with the VideoDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDuration

`func (o *Story) SetVideoDuration(v float32)`

SetVideoDuration sets VideoDuration field to given value.

### HasVideoDuration

`func (o *Story) HasVideoDuration() bool`

HasVideoDuration returns a boolean if a field has been set.

### GetMentions

`func (o *Story) GetMentions() []StoryMention`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *Story) GetMentionsOk() (*[]StoryMention, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *Story) SetMentions(v []StoryMention)`

SetMentions sets Mentions field to given value.


### GetLinks

`func (o *Story) GetLinks() []StoryLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Story) GetLinksOk() (*[]StoryLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Story) SetLinks(v []StoryLink)`

SetLinks sets Links field to given value.


### GetHashtags

`func (o *Story) GetHashtags() []StoryHashtag`

GetHashtags returns the Hashtags field if non-nil, zero value otherwise.

### GetHashtagsOk

`func (o *Story) GetHashtagsOk() (*[]StoryHashtag, bool)`

GetHashtagsOk returns a tuple with the Hashtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashtags

`func (o *Story) SetHashtags(v []StoryHashtag)`

SetHashtags sets Hashtags field to given value.


### GetLocations

`func (o *Story) GetLocations() []StoryLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Story) GetLocationsOk() (*[]StoryLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Story) SetLocations(v []StoryLocation)`

SetLocations sets Locations field to given value.


### GetStickers

`func (o *Story) GetStickers() []StorySticker`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *Story) GetStickersOk() (*[]StorySticker, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *Story) SetStickers(v []StorySticker)`

SetStickers sets Stickers field to given value.


### GetMedias

`func (o *Story) GetMedias() []StoryMedia`

GetMedias returns the Medias field if non-nil, zero value otherwise.

### GetMediasOk

`func (o *Story) GetMediasOk() (*[]StoryMedia, bool)`

GetMediasOk returns a tuple with the Medias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedias

`func (o *Story) SetMedias(v []StoryMedia)`

SetMedias sets Medias field to given value.

### HasMedias

`func (o *Story) HasMedias() bool`

HasMedias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


