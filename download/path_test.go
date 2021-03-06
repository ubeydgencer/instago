package igdl

import (
	"testing"
)

func TestGetPostFilePath(t *testing.T) {
	path := getPostFilePath("instagram", "25025320", "Bh7kySfDYq8", "https://instagram.fkhh1-2.fna.fbcdn.net/vp/893534d61bdc5ea6911593d3ee0a1922/5B6363AB/t51.2885-19/s320x320/14719833_310540259320655_1605122788543168512_a.jpg?abc=1", 1520056661)
	if path != "Instagram/instagram/posts/instagram-25025320-post-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.jpg" {
		t.Error(path)
		return
	}
}

func TestGetStoryFilePath(t *testing.T) {
	path := getStoryFilePath("instagram", "25025320", "Bh7kySfDYq8", "123.mp4", 1520056661)
	if path != "Instagram/instagram/stories/instagram-25025320-story-2018-03-03T13:57:41+08:00-Bh7kySfDYq8-1520056661.mp4" {
		t.Error(path)
		return
	}
}

func TestGetPostLiveFilePath(t *testing.T) {
	path := getPostLiveFilePath("instagram", "25025320", "123.mp4", "video", 1520056661)
	if path != "Instagram/instagram/postlives/instagram-25025320-postlive-video-2018-03-03T13:57:41+08:00-1520056661.mp4" {
		t.Error(path)
		return
	}
}

func TestGetPostLiveMergedFilePath(t *testing.T) {
	vpath := "Instagram/instagram/postlives/instagram-25025320-postlive-video-2018-03-03T13:57:41+08:00-1520056661.mp4"
	apath := "Instagram/instagram/postlives/instagram-25025320-postlive-audio-2018-03-03T13:57:41+08:00-1520056661.mp4"
	path := getPostLiveMergedFilePath(vpath, apath)
	if path != "Instagram/instagram/postlives/instagram-25025320-postlive-2018-03-03T13:57:41+08:00-1520056661.mp4" {
		t.Error(path)
		return
	}
}

func TestAppendIndexToFilename(t *testing.T) {
	nf := appendIndexToFilename("instagram-25025320-post-Bh7kySfDYq8-2018-03-03T13:57:41+08:00-1520056661.mp4", 0)
	if nf != "instagram-25025320-post-Bh7kySfDYq8-2018-03-03T13:57:41+08:00-1520056661-0.mp4" {
		t.Error(nf)
	}
	nf = appendIndexToFilename("instagram-25025320-post-Bh7kySfDYq8-2018-03-03T13:57:41+08:00-1520056661.mp4", 1)
	if nf != "instagram-25025320-post-Bh7kySfDYq8-2018-03-03T13:57:41+08:00-1520056661-1.mp4" {
		t.Error(nf)
	}
}

func TestGetUserProfilPicFilePath(t *testing.T) {
	path := getUserProfilPicFilePath("instagram", "25025320", "https://instagram.fkhh1-2.fna.fbcdn.net/vp/893534d61bdc5ea6911593d3ee0a1922/5B6363AB/t51.2885-19/s320x320/14719833_310540259320655_1605122788543168512_a.jpg", 1520056661)
	if path != "Instagram/instagram/instagram-25025320-profile_pic-1520056661.jpg" {
		t.Error(path)
		return
	}
}
