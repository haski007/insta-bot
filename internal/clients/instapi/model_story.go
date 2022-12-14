/*
instagrapi-rest

RESTful API Service for instagrapi

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instapi

import (
	"encoding/json"
	"time"
)

// Story struct for Story
type Story struct {
	Pk string `json:"pk"`
	Id string `json:"id"`
	Code string `json:"code"`
	TakenAt time.Time `json:"taken_at"`
	MediaType int32 `json:"media_type"`
	ProductType *string `json:"product_type,omitempty"`
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	User UserShort `json:"user"`
	VideoUrl *string `json:"video_url,omitempty"`
	VideoDuration *float32 `json:"video_duration,omitempty"`
	Mentions []StoryMention `json:"mentions"`
	Links []StoryLink `json:"links"`
	Hashtags []StoryHashtag `json:"hashtags"`
	Locations []StoryLocation `json:"locations"`
	Stickers []StorySticker `json:"stickers"`
	Medias []StoryMedia `json:"medias,omitempty"`
}

// NewStory instantiates a new Story object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStory(pk string, id string, code string, takenAt time.Time, mediaType int32, user UserShort, mentions []StoryMention, links []StoryLink, hashtags []StoryHashtag, locations []StoryLocation, stickers []StorySticker) *Story {
	this := Story{}
	this.Pk = pk
	this.Id = id
	this.Code = code
	this.TakenAt = takenAt
	this.MediaType = mediaType
	var productType string = ""
	this.ProductType = &productType
	this.User = user
	var videoDuration float32 = 0.0
	this.VideoDuration = &videoDuration
	this.Mentions = mentions
	this.Links = links
	this.Hashtags = hashtags
	this.Locations = locations
	this.Stickers = stickers
	return &this
}

// NewStoryWithDefaults instantiates a new Story object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryWithDefaults() *Story {
	this := Story{}
	var productType string = ""
	this.ProductType = &productType
	var videoDuration float32 = 0.0
	this.VideoDuration = &videoDuration
	return &this
}

// GetPk returns the Pk field value
func (o *Story) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Story) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Story) SetPk(v string) {
	o.Pk = v
}

// GetId returns the Id field value
func (o *Story) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Story) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Story) SetId(v string) {
	o.Id = v
}

// GetCode returns the Code field value
func (o *Story) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Story) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Story) SetCode(v string) {
	o.Code = v
}

// GetTakenAt returns the TakenAt field value
func (o *Story) GetTakenAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TakenAt
}

// GetTakenAtOk returns a tuple with the TakenAt field value
// and a boolean to check if the value has been set.
func (o *Story) GetTakenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TakenAt, true
}

// SetTakenAt sets field value
func (o *Story) SetTakenAt(v time.Time) {
	o.TakenAt = v
}

// GetMediaType returns the MediaType field value
func (o *Story) GetMediaType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value
// and a boolean to check if the value has been set.
func (o *Story) GetMediaTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaType, true
}

// SetMediaType sets field value
func (o *Story) SetMediaType(v int32) {
	o.MediaType = v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *Story) GetProductType() string {
	if o == nil || o.ProductType == nil {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Story) GetProductTypeOk() (*string, bool) {
	if o == nil || o.ProductType == nil {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *Story) HasProductType() bool {
	if o != nil && o.ProductType != nil {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *Story) SetProductType(v string) {
	o.ProductType = &v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *Story) GetThumbnailUrl() string {
	if o == nil || o.ThumbnailUrl == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Story) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || o.ThumbnailUrl == nil {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *Story) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl != nil {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *Story) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}

// GetUser returns the User field value
func (o *Story) GetUser() UserShort {
	if o == nil {
		var ret UserShort
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Story) GetUserOk() (*UserShort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Story) SetUser(v UserShort) {
	o.User = v
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise.
func (o *Story) GetVideoUrl() string {
	if o == nil || o.VideoUrl == nil {
		var ret string
		return ret
	}
	return *o.VideoUrl
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Story) GetVideoUrlOk() (*string, bool) {
	if o == nil || o.VideoUrl == nil {
		return nil, false
	}
	return o.VideoUrl, true
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *Story) HasVideoUrl() bool {
	if o != nil && o.VideoUrl != nil {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given string and assigns it to the VideoUrl field.
func (o *Story) SetVideoUrl(v string) {
	o.VideoUrl = &v
}

// GetVideoDuration returns the VideoDuration field value if set, zero value otherwise.
func (o *Story) GetVideoDuration() float32 {
	if o == nil || o.VideoDuration == nil {
		var ret float32
		return ret
	}
	return *o.VideoDuration
}

// GetVideoDurationOk returns a tuple with the VideoDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Story) GetVideoDurationOk() (*float32, bool) {
	if o == nil || o.VideoDuration == nil {
		return nil, false
	}
	return o.VideoDuration, true
}

// HasVideoDuration returns a boolean if a field has been set.
func (o *Story) HasVideoDuration() bool {
	if o != nil && o.VideoDuration != nil {
		return true
	}

	return false
}

// SetVideoDuration gets a reference to the given float32 and assigns it to the VideoDuration field.
func (o *Story) SetVideoDuration(v float32) {
	o.VideoDuration = &v
}

// GetMentions returns the Mentions field value
func (o *Story) GetMentions() []StoryMention {
	if o == nil {
		var ret []StoryMention
		return ret
	}

	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value
// and a boolean to check if the value has been set.
func (o *Story) GetMentionsOk() ([]StoryMention, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mentions, true
}

// SetMentions sets field value
func (o *Story) SetMentions(v []StoryMention) {
	o.Mentions = v
}

// GetLinks returns the Links field value
func (o *Story) GetLinks() []StoryLink {
	if o == nil {
		var ret []StoryLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *Story) GetLinksOk() ([]StoryLink, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *Story) SetLinks(v []StoryLink) {
	o.Links = v
}

// GetHashtags returns the Hashtags field value
func (o *Story) GetHashtags() []StoryHashtag {
	if o == nil {
		var ret []StoryHashtag
		return ret
	}

	return o.Hashtags
}

// GetHashtagsOk returns a tuple with the Hashtags field value
// and a boolean to check if the value has been set.
func (o *Story) GetHashtagsOk() ([]StoryHashtag, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hashtags, true
}

// SetHashtags sets field value
func (o *Story) SetHashtags(v []StoryHashtag) {
	o.Hashtags = v
}

// GetLocations returns the Locations field value
func (o *Story) GetLocations() []StoryLocation {
	if o == nil {
		var ret []StoryLocation
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *Story) GetLocationsOk() ([]StoryLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locations, true
}

// SetLocations sets field value
func (o *Story) SetLocations(v []StoryLocation) {
	o.Locations = v
}

// GetStickers returns the Stickers field value
func (o *Story) GetStickers() []StorySticker {
	if o == nil {
		var ret []StorySticker
		return ret
	}

	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value
// and a boolean to check if the value has been set.
func (o *Story) GetStickersOk() ([]StorySticker, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stickers, true
}

// SetStickers sets field value
func (o *Story) SetStickers(v []StorySticker) {
	o.Stickers = v
}

// GetMedias returns the Medias field value if set, zero value otherwise.
func (o *Story) GetMedias() []StoryMedia {
	if o == nil || o.Medias == nil {
		var ret []StoryMedia
		return ret
	}
	return o.Medias
}

// GetMediasOk returns a tuple with the Medias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Story) GetMediasOk() ([]StoryMedia, bool) {
	if o == nil || o.Medias == nil {
		return nil, false
	}
	return o.Medias, true
}

// HasMedias returns a boolean if a field has been set.
func (o *Story) HasMedias() bool {
	if o != nil && o.Medias != nil {
		return true
	}

	return false
}

// SetMedias gets a reference to the given []StoryMedia and assigns it to the Medias field.
func (o *Story) SetMedias(v []StoryMedia) {
	o.Medias = v
}

func (o Story) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["taken_at"] = o.TakenAt
	}
	if true {
		toSerialize["media_type"] = o.MediaType
	}
	if o.ProductType != nil {
		toSerialize["product_type"] = o.ProductType
	}
	if o.ThumbnailUrl != nil {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.VideoUrl != nil {
		toSerialize["video_url"] = o.VideoUrl
	}
	if o.VideoDuration != nil {
		toSerialize["video_duration"] = o.VideoDuration
	}
	if true {
		toSerialize["mentions"] = o.Mentions
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if true {
		toSerialize["hashtags"] = o.Hashtags
	}
	if true {
		toSerialize["locations"] = o.Locations
	}
	if true {
		toSerialize["stickers"] = o.Stickers
	}
	if o.Medias != nil {
		toSerialize["medias"] = o.Medias
	}
	return json.Marshal(toSerialize)
}

type NullableStory struct {
	value *Story
	isSet bool
}

func (v NullableStory) Get() *Story {
	return v.value
}

func (v *NullableStory) Set(val *Story) {
	v.value = val
	v.isSet = true
}

func (v NullableStory) IsSet() bool {
	return v.isSet
}

func (v *NullableStory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStory(val *Story) *NullableStory {
	return &NullableStory{value: val, isSet: true}
}

func (v NullableStory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


