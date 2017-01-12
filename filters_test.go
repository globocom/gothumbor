package gothumbor

import (
	"strings"
	"testing"
)

var auxURLs = []string{
	"http://my.server.com/some/path/to/image-01.jpg",
	"http://my.server.com/some/path/to/image-02.jpg",
	"http://my.server.com/some/path/to/image-03.jpg",
	"http://my.server.com/some/path/to/image-04.jpg",
}

func TestDistributedCollageFilter(t *testing.T) {
	orientation := "horizontal"
	alignment := "smart"
	filter := DistributedCollageFilter(orientation, alignment, auxURLs)

	index := strings.Index(filter, "distributed_collage")
	if index < 0 {
		t.Errorf("%s should contain “distributed_collage”", filter)
	}

	index = strings.Index(filter, "horizontal")
	if index < 0 {
		t.Errorf("%s should contain “horizontal”", filter)
	}

	index = strings.Index(filter, "smart")
	if index < 0 {
		t.Errorf("%s should contain “smart”", filter)
	}

	for _, url := range auxURLs {
		index = strings.Index(filter, url)
		if index < 0 {
			t.Errorf("%s should contain “%s”", filter, url)
		}
	}

}
