#!/bin/sh

gofmt_check() {
	unformatted=$(gofmt -l .)
	if [ -n "$unformatted" ]; then
		echo "go fmt found problems with the following files:"
		echo "$unformatted"
		echo "Please check https://blog.golang.org/go-fmt-your-code :)"
		return 1
	fi
}
