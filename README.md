gothumbor
=========

[![GoDoc](https://godoc.org/github.com/globocom/gothumbor?status.svg)](https://godoc.org/github.com/globocom/gothumbor) [![Build Status](https://travis-ci.org/globocom/gothumbor.svg?branch=master)](https://travis-ci.org/globocom/gothumbor) [![Coverage Status](https://coveralls.io/repos/globocom/gothumbor/badge.svg?branch=master&service=github)](https://coveralls.io/github/globocom/gothumbor?branch=master)

gothumbor allows easy usage of thumbor in Go. **Requires Go 1.5**.


Using it
-------
    mykey := "my-very-secret-key"
    myImageURL := "my-domain.com/static/images/fancy-image.png"
    thumborOptions := gothumbor.ThumborOptions{Width: 540, Height: 480}
    newURL, err := gothumbor.GetCryptedThumborPath(myKey, myImageURL, thumborOptions)
    

Whith these variables the values of:

* newURL: AGp4diIF89Cm2ugmDGjhycikYjY=/540x480/my-domain.com/static/images/fancy-image.png
* err: nil

License
-------

MIT (see LICENSE file).

Contributors
------------

[Rômulo Jales](https://github.com/romulojales)
