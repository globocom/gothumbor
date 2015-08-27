package gothumbor

import (
	"strings"
	"testing"
)

const imageURL = "my.server.com/some/path/to/image.jpg"
const width = 300
const height = 200
const encryptedURL = "8ammJH8D-7tXy6kU3lTvoXlhu4o=/300x200/my.server.com/some/path/to/image.jpg"
const unsafeURL = "/300x200/my.server.com/some/path/to/image.jpg"

func TestGetUrlPartialWithWidthAndHeight(t *testing.T) {
	thumborOptions := ThumborOptions{Width: 1, Height: 1, Smart: false}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}

	urlE := strings.Join([]string{"1x1", imageURL}, "/")
	if url != urlE {
		t.Errorf("Got an unxpected partial url: %s != %s", url, urlE)
	}
}

func TestGetUrlPartialWithSmart(t *testing.T) {
	thumborOptions := ThumborOptions{Width: 1, Height: 1, Smart: true}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}

	urlE := strings.Join([]string{"1x1", "smart", imageURL}, "/")
	if url != urlE {
		t.Errorf("Got an unxpected partial url: %s != %s", url, urlE)
	}
}

func TestGetUrlPartialOnlyWithWidthAndHeight(t *testing.T) {
	thumborOptions := ThumborOptions{Width: width, Height: height}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to generate the thumbor url", err)
	}
	if url != strings.Join([]string{"300x200", imageURL}, "/") {
		t.Error("Got an unxpected partial path:", url)
	}
}

func TestEscapeURLByRFC3986(t *testing.T) {
	thumborOptions := ThumborOptions{}
	url, err := getURLParts("/a-path with spaces.jpg", thumborOptions)
	if err != nil {
		t.Error("Got an error when tried to process a path with spaces", err)
	}
	if url != "/a-path%20with%20spaces.jpg" {
		t.Error("Got an unxpected partial path:", url)
	}
}
func TestFitInParameter(t *testing.T) {
	thumborOptions := ThumborOptions{FitIn: true}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if !strings.Contains(url, "fit-in") {
		t.Errorf("url doesn't have a fit-in parameter")
	}

}

func TestOneFilterParameter(t *testing.T) {
	filter := "max-age(360000)"
	filters := []string{filter}
	thumborOptions := ThumborOptions{Width: 200, Height: 300, Filters: filters}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if !strings.Contains(url, "filters:"+filter) {
		t.Errorf("url doesn't have a filters parameter")
	}

}

func TestTwoFiltersParameter(t *testing.T) {
	firstFilter := "max-age(360000)"
	secondFilter := "grayscale()"
	filters := []string{firstFilter, secondFilter}
	thumborOptions := ThumborOptions{Width: 200, Height: 300, Filters: filters}

	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if !strings.Contains(url, "filters:"+firstFilter+":"+secondFilter) {
		t.Errorf("url doesn't have a first filter parameter")
	}
}


func TestFlipWithoutAnyOtherParameter(t *testing.T) {
	thumborOptions := ThumborOptions{Flip: true}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if url != strings.Join([]string{"-0x0", imageURL}, "/") {
		t.Errorf("url is not fliped", url)
	}
}


func TestFlopWithoutAnyOtherParameter(t *testing.T) {
	thumborOptions := ThumborOptions{Flop: true}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if url != strings.Join([]string{"0x-0", imageURL}, "/") {
		t.Errorf("url is not floped", url)
	}
}


func TestFlipFlopWithoutAnyOtherParameter(t *testing.T) {
	thumborOptions := ThumborOptions{Flop: true, Flip: true}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if url != strings.Join([]string{"-0x-0", imageURL}, "/") {
		t.Errorf("url is not flipfloped", url)
	}
}


func TestFlipFlopWithWidthAndHeigh(t *testing.T) {
	thumborOptions := ThumborOptions{Flop: true, Flip: true, Height: 500, Width: 400}
	url, err := getURLParts(imageURL, thumborOptions)
	if err != nil || url == "" {
		t.Errorf("Got an error when tried to generate the thumbor url")
	}

	if url != strings.Join([]string{"-400x-500", imageURL}, "/") {
		t.Errorf("url is not flipfloped", url)
	}
}