// Code generaTed by fileb0x at "2019-08-03 22:30:16.809668786 +0800 CST m=+0.260926315" from config file "b0x.yml" DO NOT EDIT.
// modified(2018-11-16 23:20:23 +0800 CST)
// original path: dist/images/expand.gif

package swaggerFiles

import (
	"os"
)

// FileImagesExpandGif is "/images/expand.gif"
var FileImagesExpandGif = []byte("\x47\x49\x46\x38\x39\x61\x0b\x00\x0b\x00\x91\x00\x00\x4a\x4d\x4a\x7b\x7d\x7b\xff\xff\xff\x00\x00\x00\x21\xf9\x04\x00\x00\x00\x00\x00\x2c\x00\x00\x00\x00\x0b\x00\x0b\x00\x00\x02\x1a\x8c\x8f\x26\xcb\xac\x0c\xda\x5b\xd1\x85\x56\x97\x02\xbc\xc7\x29\x64\x02\x28\x92\xd2\xd5\x9c\x49\x52\x00\x00\x3b")

func init() {

	f, err := FS.OpenFile(CTX, "/images/expand.gif", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileImagesExpandGif)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
