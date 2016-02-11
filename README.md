gothumbor
=========

[![GoDoc](https://godoc.org/github.com/globocom/gothumbor?status.svg)](https://godoc.org/github.com/globocom/gothumbor) [![Build Status](https://travis-ci.org/globocom/gothumbor.svg?branch=master)](https://travis-ci.org/globocom/gothumbor) [![Coverage Status](https://coveralls.io/repos/globocom/gothumbor/badge.svg?branch=master&service=github)](https://coveralls.io/github/globocom/gothumbor?branch=master)

gothumbor allows easy usage of [Thumbor] in Go.


Using it
-------
```go
myKey := "my-very-secret-key"
myImageURL := "my-domain.com/static/images/fancy-image.png"
thumborOptions := gothumbor.ThumborOptions{Width: 540, Height: 480}
newURL, err := gothumbor.GetCryptedThumborPath(myKey, myImageURL, thumborOptions)
```

With these variables the values of:

* newURL: AGp4diIF89Cm2ugmDGjhycikYjY=/540x480/my-domain.com/static/images/fancy-image.png
* err: nil


License
-------

[MIT][mit] © [Globo.com][globocom]


Contributors
------------

Click [here][contributors] to see the list of contributors.

[Thumbor]:      https://github.com/thumbor/thumbor
[mit]:          https://github.com/globocom/gothumbor/blob/master/LICENSE
[globocom]:     https://github.com/globocom
[contributors]: https://github.com/globocom/gothumbor/graphs/contributors
