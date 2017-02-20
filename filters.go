package gothumbor

import (
	"fmt"
	"net/url"
	"strings"
)

func DistributedCollageFilter(orientation, alignment string, urls []string) string {
	return fmt.Sprintf(
		"distributed_collage(%s,%s,%s)",
		orientation,
		alignment,
		strings.Join(urls, url.QueryEscape("|")),
	)
}

func WatermarkFilter(imageURL string, x, y, alpha int) string {
	return fmt.Sprintf("watermark(%s,%d,%d,%d)", imageURL, x, y, alpha)
}
