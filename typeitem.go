package instago

import (
	"errors"
	"strconv"
	"strings"
)

// item struct shared by *reels tray* and *timeline* feed
// The JSON strcut of *reels tray* and *timeline* are slightly different

type IGItem struct {
	TakenAt         int64  `json:"taken_at"`
	Pk              int64  `json:"pk"`
	Id              string `json:"id"`
	DeviceTimestamp int64  `json:"device_timestamp"` // not reliable value
	MediaType       int64  `json:"media_type"`
	Code            string `json:"code"`
	ClientCacheKey  string `json:"client_cache_key"`
	FilterType      int64  `json:"filter_type"`

	// timeline only
	CarouselMedia []struct {
		Id               string             `json:"id"`
		MediaType        int64              `json:"media_type"`
		ImageVersions2   ItemImageVersion2  `json:"image_versions2"`
		OriginalWidth    int64              `json:"original_width"`
		OriginalHeight   int64              `json:"original_height"`
		VideoVersions    []ItemVideoVersion `json:"video_versions"`
		Pk               int64              `json:"pk"`
		CarouselParentId string             `json:"carousel_parent_id"`
	} `json:"carousel_media"`

	// timeline only
	Location struct {
		Pk               int64   `json:"pk"`
		Name             string  `json:"name"`
		Address          string  `json:"address"`
		City             string  `json:"city"`
		ShortName        string  `json:"short_name"`
		Lng              float64 `json:"lng"`
		Lat              float64 `json:"lat"`
		ExternalSource   string  `json:"external_source"`
		FacebookPlacesId int64   `json:"facebook_places_id"`
	} `json:"location"`

	ImageVersions2  ItemImageVersion2 `json:"image_versions2"`
	OriginalWidth   int64             `json:"original_width"`
	OriginalHeight  int64             `json:"original_height"`
	CaptionPosition float64           `json:"caption_position"`
	IsReelMedia     bool              `json:"is_reel_media"`

	VideoVersions []ItemVideoVersion `json:"video_versions"`
	HasAudio      bool               `json:"has_audio"`
	VideoDuration float64            `json:"video_duration"`

	User IGUser `json:"user"`

	//Caption              string `json:"caption"`	// not string type

	CaptionIsEdited      bool   `json:"caption_is_edited"`
	PhotoOfYou           bool   `json:"photo_of_you"`
	CanViewerSave        bool   `json:"can_viewer_save"`
	OrganicTrackingToken string `json:"organic_tracking_token"`
	ExpiringAt           int64  `json:"expiring_at"`

	IsDashEligible int64 `json:"is_dash_eligible"`

	//"reel_mentions"
	//"story_locations"
	//"story_events"
	//"story_hashtags"
	//"story_polls"
	//"story_feed_media"
	//"story_sound_on"

	CanReshare            bool `json:"can_reshare"`
	SupportsReelReactions bool `json:"supports_reel_reactions"`
}

// Used to decode JSON in item.
type ItemVideoVersion struct {
	Type   int64  `json:"type"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	Url    string `json:"url"`
	Id     string `json:"id"`
}

// Used to decode JSON in item.
type ItemImageVersion2 struct {
	Candidates []struct {
		Width  int64  `json:"width"`
		Height int64  `json:"height"`
		Url    string `json:"url"`
	} `json:"candidates"`
}

// media type:
//   0: ???
//   1: single photo
//   2: single video
//   8: multiple photos/videos
func (i *IGItem) IsRegularMedia() bool {
	// remove ads
	if i.User.FriendshipStatus.Following == false {
		return false
	}

	switch i.MediaType {
	case 1:
		return true
	case 2:
		return true
	case 8:
		return true
	default:
		return false
	}
}

func (i *IGItem) GetUsername() string {
	return i.User.Username
}

func (i *IGItem) GetUserId() string {
	return strconv.FormatInt(i.User.Pk, 10)
}

func (i *IGItem) GetPostUrl() string {
	return strings.Replace(
		"https://www.instagram.com/p/{{CODE}}/",
		"{{CODE}}",
		i.Code,
		1)
}

func (i *IGItem) GetPostCode() string {
	return i.Code
}

func (i *IGItem) GetTimestamp() int64 {
	return i.TakenAt
}

// Return best resolution photo/video URL(s) in item
func (i *IGItem) GetMediaUrls() (urls []string, err error) {
	switch i.MediaType {
	case 1:
		// single photo
		urls = append(urls, i.ImageVersions2.Candidates[0].Url)
	case 2:
		// single video
		urls = append(urls, i.VideoVersions[0].Url)
	case 8:
		// multiple photos/videos
		for _, media := range i.CarouselMedia {
			switch media.MediaType {
			case 1:
				urls = append(urls, media.ImageVersions2.Candidates[0].Url)
			case 2:
				urls = append(urls, media.VideoVersions[0].Url)
			default:
				err = errors.New("Cannot be multiple photos/videos in carousel_media")
				return
			}
		}
	default:
		err = errors.New("Not Regular Media Type")
		return
	}

	/*
		for i, url := range urls {
			urls[i], err = StripQueryString(url)
			if err != nil {
				return
			}
		}
	*/
	return
}
