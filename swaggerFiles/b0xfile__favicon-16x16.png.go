// Code generaTed by fileb0x at "2017-06-25 01:10:56.979942041 +0800 CST" from config file "b0x.yml" DO NOT EDIT.

package swaggerFiles

import (
	"log"
	"os"
)

// FileFavicon16x16Png is "/favicon-16x16.png"
var FileFavicon16x16Png = []byte("\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x01\x84\x49\x44\x41\x54\x78\x01\x95\x53\x03\x4c\x75\x71\x1c\xfd\x8c\xf1\xc3\xec\x30\xa7\x29\xcd\x61\xb6\x6b\x36\xb2\x9b\xf9\xb2\x6b\xc8\x35\x2f\xdb\x8d\x71\x78\xc6\x94\x6d\xcc\x7b\xef\x7f\x4f\xff\xf3\x6c\xdc\xed\xf2\xe0\xfe\xf8\xc9\xff\x50\x14\x11\x2f\x14\x5b\xa3\x50\xc4\xa1\xbc\x3f\xf1\x74\x3e\x37\x12\x73\x13\x03\x85\xca\x37\x49\x52\x09\x61\xb5\x6a\x8f\xa7\x31\xbe\x5d\x88\xf6\xb9\x4c\xf0\x1c\x93\xcf\xda\xe3\x29\x10\x93\x66\x8d\xe4\x06\x13\xcf\xde\x3c\x9b\xd1\x34\x95\x8a\x92\x81\x4f\x41\xcf\x46\x89\xdd\x3c\x9b\x20\x4d\xe6\x7d\x4c\xe4\x07\x15\xc5\xf5\xe3\xff\x49\x0c\x7b\xd6\x8d\xff\x73\x99\x34\xba\x73\x66\x68\xae\x3f\xaf\x6b\x1a\x70\x72\x77\x10\x20\x3c\xb9\xdb\xc7\x86\xa6\xd1\x19\x49\x0a\xa8\xb1\xd7\x84\x79\x33\x67\x17\x31\x54\x24\xb5\x63\x7f\x71\xfb\x62\x71\xbf\x6b\x8e\x27\x1d\x51\xb0\xc2\x2c\x92\x0b\x78\x7c\x3b\x46\xe5\xf0\xef\x00\x83\xf2\xa1\x1f\x78\x7c\x3f\x71\xbd\xcb\xc2\x16\x80\x5a\x46\xf0\xc4\x4a\xf3\xe3\xe4\x6e\x31\xcc\x17\x6b\x60\x3a\x7d\xcb\x79\xe8\x98\xcb\x42\xc7\x7c\x36\x7a\x97\x72\xd1\x34\x9d\x06\xd3\xf9\x8a\xe4\x94\x90\x8b\xb6\xd9\x0c\x50\xeb\x63\x40\xd0\x7c\xbe\x2a\xc9\x34\xc8\xa7\x98\x27\xcd\x68\x00\xe3\xd9\x32\xa6\x76\x4b\x7d\x0c\x42\xa4\xf0\x2b\x44\x0a\xc7\x81\x29\xb0\x10\x9a\xe3\xa9\xd8\x8b\x78\xe4\x28\xa2\xbb\x8d\x6c\x0d\x01\xb6\x8a\x2d\xf3\x37\x38\xbe\xdd\xc7\xa6\xb6\xc9\xd9\xc6\x64\xd8\x5c\x6d\xf4\x0c\x92\x09\x75\x51\x0e\xd2\xf5\xb3\xd1\xf1\x77\xdf\x51\x16\xb3\x34\x61\x24\xa1\xc4\xc4\x28\x56\xbc\x46\xd9\xdf\xa4\x91\xe9\xb0\x26\x2c\x12\x2b\xcd\x93\xcf\x1c\x1c\x62\xdc\xca\x00\x71\x74\xeb\xcc\x2d\x14\x89\xfe\xfc\x0f\x6d\x32\x6a\x88\xec\xcc\x73\x18\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82")

func init() {

	f, err := FS.OpenFile(CTX, "/favicon-16x16.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(FileFavicon16x16Png)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}