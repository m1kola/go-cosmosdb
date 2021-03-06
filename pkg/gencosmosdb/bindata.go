// Package gencosmosdb Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// cosmosdb/collection.go
// cosmosdb/cosmosdb.go
// cosmosdb/database.go
// cosmosdb/document.go
// cosmosdb/template.go
// cosmosdb/trigger.go
package gencosmosdb

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _collectionGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\x5f\x6f\xdb\x38\x12\x7f\x96\x3e\x05\x57\x0f\x45\x9c\xb8\xce\xcb\xe1\x1e\x0a\xf8\xa1\x48\x76\xb3\x41\x9b\x36\x68\x73\xb7\x07\x04\x41\x97\x91\xc6\x36\x2f\x32\xa9\x25\xa9\x6d\x72\x41\xbe\xfb\x81\xa4\x28\xf1\xaf\xed\x74\x0f\xb8\x3e\xa4\x31\xe7\x37\xc3\xf9\x3f\x63\xa6\xc3\xf5\x03\x5e\x03\xaa\x99\xd8\x32\xd1\xdc\x97\x25\xd9\x76\x8c\x4b\x74\x54\x16\x55\xcd\xa8\x84\x47\x59\x95\x45\x45\x41\x9e\x6e\xa4\xec\xaa\x72\x56\x96\xa7\xa7\xe8\x8c\xb5\x2d\xd4\x92\x30\x8a\x38\x74\x1c\x04\x50\x29\x10\x46\xf5\x78\x5e\xca\xa7\x0e\x5c\x9c\x90\xbc\xaf\x25\x7a\x2e\x8b\xcb\x73\x94\xfe\x27\x24\x27\x74\x9d\xa2\xfc\xfe\x6f\xc1\xe8\xbb\x8a\x34\x73\xb6\x25\x12\xb6\x9d\x7c\xaa\x7e\x2f\x8b\x2f\x20\x58\xcf\x6b\x88\x24\xee\x95\xf4\x8d\x87\xb2\x6e\xc8\x16\x84\xc4\xdb\x2e\xe4\x20\x54\x66\xf4\xb5\xb2\xa4\xf0\x45\x7d\x85\x76\xf5\x63\x06\x7e\x13\xd0\xae\x7c\x61\x3f\xdf\xe0\x24\xc7\x01\xc2\x40\xe2\xb5\x2f\xec\x9c\xd5\xfd\x56\xc7\xea\xd5\xc2\x1a\x56\x87\x66\x4a\xc6\xa1\xb9\xe6\xac\x86\xa6\xe7\x20\x5e\x63\x66\xc7\x23\x71\x37\x9c\xac\xd7\xc0\x23\xd5\x0e\x10\x27\x07\x56\x5f\xe0\x3f\x04\xf0\x73\x58\x11\x0a\xcd\x2f\x3d\xd5\x59\x28\x0e\x14\xd8\x37\xab\x40\xd8\x19\xa3\xab\x96\xd4\x3f\xe2\xb9\xda\xb2\xfa\x12\x2f\x69\x03\x8f\x84\xae\xaf\x59\x4b\xea\x27\x87\xef\x38\x4b\xb1\x65\xe0\xd1\x7d\xa9\xd7\x98\x4b\xa2\x4c\xfd\x00\x4f\x9e\x36\xc7\x59\xca\x20\xb5\x73\xe8\x81\x23\x29\xf9\xa3\x87\x0f\xf0\x14\x2a\x74\x9c\xa5\x0c\x32\x7b\x9f\x9e\x76\xa9\xaa\xe2\xb6\x57\x37\x0f\x52\x8e\xb3\x94\x41\x6c\x9d\xa1\xfb\xf2\x2f\x80\x89\x0e\x4b\x82\x5b\x25\x8f\x4c\x31\x3a\xce\x52\x06\xf9\xeb\x80\xee\xc9\x7d\xd1\x1d\x30\x88\x91\xdb\x05\x29\xb2\x01\x42\x9d\xa6\x9a\x5e\x18\x70\x4c\xfd\xf0\x7d\x2f\xd9\x16\x4b\x52\x5b\x2d\xee\x19\x6b\xd3\xb9\x84\x2d\x34\x9d\x4a\x57\xac\x01\x0d\xf7\xef\xd2\xc7\x41\xee\xa8\xb3\x50\x48\xdd\xf6\x0d\x34\xd7\x58\x6e\x74\x96\xdf\xde\xb9\x47\x7e\x02\x3a\xd0\xa0\x5d\x3d\x1e\x2c\x05\x1e\xb3\x52\xce\xd8\xb6\x63\x82\x48\xd0\x86\x80\x40\xb7\x77\xfe\xd1\x24\xa5\x0e\xa0\x7b\x63\xa5\xbd\xb1\x33\x5e\x68\xcb\x1a\x48\x05\x4d\xb3\x9a\x82\xcf\x09\xae\x19\x15\x12\x53\x29\x4a\xfd\x9b\x1a\xa4\x31\xea\x8c\x51\x41\x84\x04\x2a\x53\x22\x96\xa8\x9a\x00\x55\x8a\xff\x23\xfe\x8f\xad\xb5\x34\xbf\x02\xd8\x59\xed\xb9\x3f\xb4\xdb\x90\x50\x87\xe5\xc6\x1a\xec\xa0\xa7\x1c\xb5\xb1\xb3\xdd\x6e\xec\x19\x72\x93\x48\x45\x1d\x30\x13\x27\x37\xed\xf2\xc1\x49\xc5\xc3\x09\x80\xa3\xc8\x39\x96\xf8\x46\x11\x0c\x65\xfc\x38\xdc\xd3\x0c\x9f\x7d\xa5\x3e\x10\xda\x38\xee\x1a\x3f\x0e\x4c\x0f\x84\x06\xcb\xc0\x35\x87\x9a\x08\xb5\xb6\xf8\xd3\xdf\x9a\x6d\xc9\x39\x7b\x46\xb5\x12\x76\x21\xa5\x23\x52\xc6\x39\x16\x8e\x0c\x61\x76\x8d\x84\x5c\x62\x59\xc0\xd7\x69\x0e\xf9\x8c\x4b\x54\x19\x5a\x15\x70\x7c\xea\xb7\xf7\xc0\xd3\x1c\x86\x16\x72\x5c\xb3\xd1\x1d\x11\x87\xa6\xc5\x0c\xed\xd3\x9a\xd1\x0c\x83\xa6\x85\x2c\x1f\x09\xb5\xb6\x44\x2c\x13\x6d\x4a\x6e\x1b\xcf\x94\xa7\x55\x60\x1d\x27\x6b\x5c\xe8\x60\x7d\x98\x73\xae\x22\xfe\x8a\xc5\xc6\x4f\x9d\x25\xaa\xd4\x61\xe5\x80\xbe\x60\xba\x86\x10\xa4\x0f\x5d\xd4\x57\x33\x54\x7c\xd4\x70\x68\x2d\x72\x1b\x68\x60\x94\x6d\x98\x4e\xb9\x7a\xe8\xa0\x5c\x87\x5a\xcd\x95\xea\xcb\xb0\xc9\x7b\x6d\x35\xd8\xe6\x07\x9a\x5b\x90\x01\xc3\xed\x5d\xd8\x24\xf6\x5c\x5b\x7c\xe6\x0d\x70\x64\x7e\x5a\x14\x53\x9f\x12\xda\x19\x94\xef\x05\x8d\x35\xba\x18\xb2\x13\x51\x73\x90\x88\xa6\x26\xbc\x17\x35\xd0\x46\x57\x89\x01\x2e\x51\x85\xed\x59\x35\x80\xce\x61\x44\x8d\xa0\x06\x26\x94\x89\x92\xb7\x4c\x79\x4e\x1b\xd7\x28\xf4\x00\xc3\xe4\xf7\xc0\xbe\xbb\x84\x9e\x90\xc1\x0e\xe9\xb8\x4e\xa4\x1b\x99\x2b\x51\x9f\x65\x3b\xd9\x3f\x81\x27\xfa\xd8\x74\xc9\x9f\x86\x9e\xf0\x7d\x74\x47\xde\x4c\xa7\xd0\x22\x2e\x27\x3a\x11\x2d\x11\xa8\x10\xa3\xab\x2f\x62\x1c\x2b\xd0\x44\x23\x5c\x43\x3d\x4d\xcd\x0e\xaa\xd5\x74\xd7\xb1\x90\x67\x8a\xcb\x48\x51\x53\x6c\xfc\x10\xad\xb4\xa9\x61\x36\xa1\x33\x2a\x04\x77\x47\xd9\x30\xe6\x42\x36\x07\x6c\xd9\x66\x36\xe4\xa0\x80\x0d\x0a\xf1\x11\xe6\xf9\x20\x2b\x65\x52\x6b\x5c\x29\x13\xff\x72\xec\xee\xbe\xb9\x8d\xf6\xcc\x04\x97\x5d\x0c\x77\x7c\x99\x72\x17\xbe\x14\xff\xde\x3b\xec\xd7\xd4\x1f\xbd\xc3\xf2\xbf\x22\x18\xd1\xa6\xb9\x23\x20\xce\xce\xb9\x53\x9c\x53\x4e\x3b\x71\x89\xd2\xda\x85\xff\x88\x85\xfc\x8d\x13\x09\xfc\x37\x42\xc5\x6e\xd1\x7a\xbf\x74\xe1\xd5\x6e\xd9\x67\xbd\x90\x6c\x7b\x58\xd6\xa8\xdd\x57\xc3\x6d\x69\x47\x5f\xd5\x3c\x7f\x4e\x5f\xd4\xb4\x6b\xc9\xda\x78\x30\x62\x9a\xf2\xf9\x26\x45\x77\x37\x48\x19\x6e\x8f\x2f\x49\x3d\xc2\xed\x2e\xd2\x44\x6f\x78\x22\xad\x4f\xb8\xe9\x25\xe9\x89\x08\xa6\x70\x17\xc0\xd6\x1c\x77\x9b\xa7\xb4\x94\x25\xaa\x46\x44\xfc\x7a\x27\x5c\x13\xa6\xc7\x3b\x11\xbe\xde\x09\xc7\x83\x67\xac\x1f\x67\x49\x72\x3b\xfe\x56\x2b\x44\xfe\xa5\xce\x2f\xc0\xec\x83\x9c\x7b\xfb\xed\xdd\xb1\xf3\x94\x38\xb0\xd8\xc7\x2c\x07\xb8\x67\x8c\xe9\x85\xcc\xb3\xd9\x1f\x63\x5c\xd3\xe3\x41\x36\xf0\xa5\x7d\x90\x74\xc5\xab\x1c\x82\xf2\xcd\x2f\xeb\x9d\x84\x7a\xb7\x77\xd1\xa1\xe5\x8f\xd1\x87\x78\x6a\xd7\xc4\xd7\xae\xca\x78\x6a\xf7\x5b\x6f\xae\xff\x1e\xfe\xc6\xbb\x47\xc2\x41\x6f\xbb\xb9\x57\xdd\x57\xbc\xe9\xee\x53\xe3\xa0\xb7\xdc\x7d\x42\xe2\x37\xdc\x2b\xfc\xa8\xb7\x7e\x41\xfe\x84\xc3\x84\x6c\x1d\x8e\x40\x16\xa1\xfa\x0b\xff\x2b\x64\x39\x1c\xb9\x40\x5d\x73\x58\x91\xc7\xfd\x4e\xe6\xa4\x31\xd0\x20\x5a\x1b\xce\xfa\xf5\xa6\xeb\xe5\x2f\x1c\x9b\x8a\x5f\xb5\x0c\xcb\xbf\xff\x2d\x27\x48\x46\x1c\xe1\x4b\x35\x96\xbd\xff\x72\x1b\x25\xed\x80\x19\x24\x0a\xfd\x29\x2a\xb9\xf0\xe9\x3c\x5a\xdd\x7d\xbd\x3a\xc3\x71\x48\xad\x0d\xd7\xef\xab\x38\x64\x14\xcb\x14\xde\x20\x24\xb3\x72\xbb\x90\x3d\x9b\xb7\x03\xfd\x4c\x5b\x42\xf3\x77\x2d\x51\xc5\x34\x42\x0f\x17\x69\x06\x97\xed\xc7\x67\x2d\x01\x2a\x9d\x76\x70\xdc\x60\x89\xef\xb1\x00\x43\x29\x8b\x6e\xfa\xa2\x3a\x2e\x53\x01\x37\xf1\xff\xb0\x84\x6a\xc3\x1a\x4c\x28\x0b\xa6\x12\xf8\x0a\xd7\xa0\xbb\x34\x07\x2c\xe1\x68\xf8\x3b\xd6\xe2\xcc\xfc\x3f\x47\xce\x2c\x99\xa1\x23\xe7\xd3\x1c\x01\xe7\x8c\xcf\xca\xe2\x23\x11\xf2\x68\xe6\xc8\xbf\x94\xc0\xb1\x64\xdc\x90\xde\xb7\x6d\x28\xd6\x97\x24\x26\x51\x17\x20\x63\x15\x8c\xc9\xb9\xdb\xcf\xa1\x85\xbd\x8a\x6b\xb0\x2a\xbb\xae\xc5\xf5\x8f\x59\x19\x8f\x85\x5d\x9a\x26\x86\x88\x95\xf4\x12\x85\x5e\x79\xc9\xfa\xcc\x4d\x80\x30\x39\xca\x42\xdd\x48\x68\x8f\xed\x5f\x09\x55\x2a\x14\x0d\xa3\xd3\xd7\x8e\x7b\xc6\xda\x28\x3b\x46\xe1\x61\x7e\x10\x1b\xa9\x20\x43\x26\x06\x37\x47\x3e\xc1\x63\x14\x9e\x5c\x28\x8d\x0a\x9f\xe0\x7b\x94\x76\x1c\x64\xcf\xa9\x52\x84\xc2\xf7\x44\xb2\xae\x7a\x5a\xa7\x18\x8f\x6a\x74\xee\x55\xc4\x1c\x35\xf7\xa4\x19\x9d\x1e\x5d\xf4\x5c\x16\xe6\x2e\xf4\x26\xf4\xe4\x73\x59\x14\x7e\x79\xbd\x43\xf5\xe2\x28\x28\xb9\xd9\xbc\x2c\x74\xd5\xbd\x73\xda\x55\xd5\xdc\x8b\xd3\x0a\x9d\xe8\xcb\xe7\x65\xf1\xa2\x4c\xd5\x3a\x1f\xd5\x28\x0a\xd9\x0c\x61\x95\xfe\xf2\x11\x45\xc9\x42\x12\xfe\xce\x78\x53\x99\x82\xdb\x56\x09\x17\xe8\xdd\x12\xbd\x71\x30\xcf\x2f\x65\x59\xac\x18\x57\x98\x42\x23\x34\x97\x82\x91\x85\x89\x98\x7c\x9c\x95\x45\x41\x56\xfa\xfc\xa7\x25\xa2\xa4\xd5\x68\xeb\x1e\x4a\x5a\xcd\x53\x16\xc5\x8b\x01\x9a\x9b\x96\x0e\xf4\x9e\x03\x7e\xd0\x80\xb2\x18\x75\x59\x98\xdd\xee\x64\x89\x9c\x8f\x2e\xdd\x59\x48\x2c\x66\x3a\xf2\x05\x4d\x0b\xec\x12\xe1\xae\x03\xda\x1c\xa5\xa8\x73\x14\x1d\x2d\x16\x8b\x59\xa9\x15\x1b\xec\xb1\x7c\x73\xa5\xfe\x9e\xf0\xd8\xbe\x97\x8a\x10\x85\xef\x8a\x21\xe8\x0e\xe1\x91\x71\xf7\x14\x28\xf5\x69\x89\xea\x45\xc3\x94\xd4\x39\xda\x48\xd9\x2d\xae\x40\x6e\x58\x73\xcd\x84\x9c\xa3\x7a\xa1\x72\xea\xa4\x3a\xd5\x4a\x56\x73\x54\xd9\x5f\x0c\x65\x60\x31\x43\xc3\xe8\xd7\xcc\xd1\x9b\x41\x9b\xb9\xc9\x66\x6d\xdb\xcc\x9a\xbc\xc7\xc8\x6c\x87\x4e\xd7\x88\xdb\x8f\x9e\x43\x69\xef\x50\xbd\x2f\xe5\xc7\xae\x1f\x3b\x75\x47\x82\x0f\x7a\xd4\x8b\xa1\x62\x94\x3b\x8c\xe2\xb3\x3d\xf7\xe9\xa9\x91\x0a\xa0\x82\x3a\x1d\xe2\xaf\x84\xee\x02\xa2\xc8\x9d\x56\x27\xe6\x82\x28\x82\x09\x84\x13\xd2\xcf\x1f\xe6\xa6\xe4\x5e\x1f\x48\x3b\xec\x72\xc6\x26\xa6\x9e\xb2\x6b\xa8\xe8\x85\xde\xa9\x97\x4b\x54\x55\xba\xa8\x07\x8f\xff\xcc\xb9\x22\x7c\x81\x3f\x7a\xc2\xa1\x51\xa5\x54\x6c\x00\x37\xc0\x75\xb3\xd1\x9a\xff\xaa\x3f\x3f\x4f\x94\xc5\x57\x90\x47\xd5\xe5\xea\xed\x15\x96\xf5\xa6\x9a\x4f\x17\xcc\x9c\x50\x26\x1c\x69\x4c\xc8\x78\x6a\x71\x79\xbe\xcf\x99\x1a\xe2\x78\xf3\x13\xd3\x1e\x50\xe3\x40\x3b\x55\xff\x18\x94\xdc\x97\x38\xe3\x3e\xf0\x7f\xab\xfe\xd3\xea\x64\xb8\x67\x8f\xe9\x2e\xea\x7f\xdd\x1e\x52\xab\xcd\x21\xe5\xd4\x3d\x70\x81\x72\xab\xce\x5f\x2f\xab\x93\xea\xb4\x7b\x30\xcf\x0c\xaa\x49\x3a\xbf\xbf\xae\xca\x94\x9a\x39\x7f\x10\xd7\x1f\x6e\xdf\x9b\x21\x3b\x3b\x13\x2d\xcc\x8c\xc7\xa8\x91\x39\x26\x93\x15\x22\x0b\xbd\x98\x4d\x85\x66\x66\x54\xbe\xb2\x34\x8f\xb7\xe1\xfd\x34\x96\xaa\x57\x75\xff\x7a\x7b\x25\xde\x9e\x39\xc0\x6a\x1e\x70\x0e\xf3\xd0\x38\x9e\x64\x1d\x4f\xb2\x93\x88\xc4\x93\xc8\x6f\x5b\xc2\x29\xb2\xc4\x6a\xe1\x5a\x1c\x18\xb5\xb4\x8c\x8b\x8b\xb4\x31\x4a\xa0\xf1\xdd\x32\x74\x88\xee\x5d\xa5\x13\xc6\xff\x06\x00\x00\xff\xff\xa0\xd2\x0a\xa5\x95\x27\x00\x00")

func collectionGoBytes() ([]byte, error) {
	return bindataRead(
		_collectionGo,
		"collection.go",
	)
}

func collectionGo() (*asset, error) {
	bytes, err := collectionGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "collection.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cosmosdbGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\xdf\x6f\xdc\x36\x12\x7e\x96\xfe\x8a\x89\x00\x1b\x92\x2d\x73\xd7\xbe\x26\x38\x38\xdd\x87\xd4\x76\x5a\x5f\x62\xd7\x67\x6f\xd1\x87\xb4\x68\x68\x69\xb4\xcb\x5a\x22\x15\x92\x72\xbc\x31\xf6\x7f\x3f\x0c\xc9\x5d\x69\xbd\x4e\x10\xdc\x53\x17\x08\x42\xf1\xc7\xf0\x9b\xf9\x66\xbe\xa1\x5b\x5e\xdc\xf1\x19\x42\xa1\x4c\xa3\x4c\x79\x1b\xc7\xa2\x69\x95\xb6\x90\xc6\x51\x72\xbb\xb0\x68\x92\x38\x4a\x0a\x25\x2d\x3e\x58\x37\xd4\x8b\xd6\xaa\xd1\xbc\xe1\xc5\xe0\xd3\xcc\xf9\xd1\xcb\x57\x34\x81\xb2\x50\xa5\x90\xb3\xd1\x2d\x37\xf8\xea\x07\x9a\xaa\x1a\x77\x52\xa8\x91\x50\x9d\x15\x35\x7d\x48\xb4\xa3\xb9\xb5\xed\x6a\x4c\xe6\x5b\xad\xac\x5a\x4d\x74\xda\xed\x33\x56\x17\x4a\xde\x87\xa1\x90\x33\x87\xc7\x8a\x06\x93\x38\x8e\x92\x99\xb0\xf3\xee\x96\x15\xaa\x19\x75\x33\xa5\xff\x16\xa3\x99\x1a\x15\xaa\xc4\x22\x89\xb3\x38\x1e\x8d\xe0\xd7\xd6\x0a\x25\x0d\x68\x6c\x35\x1a\x94\xd6\xc0\x9b\xab\x73\x50\x7e\x3a\xb6\x8b\x16\xd7\x7b\x8c\xd5\x5d\x61\xe1\x31\x8e\x2e\xd5\xd9\x94\xcf\x60\xe3\x77\xab\x54\x1d\x47\x57\x1a\xa7\x5a\xcc\x66\xa8\xcd\x7a\xe5\xc3\x9f\x1e\x5b\x1c\x5d\x29\x63\x9f\x2e\x0f\x56\xb9\xb6\x82\xae\x7a\x87\x8b\x6b\x2e\x67\x78\x7e\x0a\x61\x6d\xe9\xc0\x9e\x69\xad\xf4\x10\x2a\x97\x80\x34\xe7\x71\xfa\xe5\x1e\xe5\x8d\xe5\xb6\x33\x27\xaa\x44\x10\xd2\xc6\x91\x1b\xf9\x9f\x37\x0b\x1f\xff\x36\x4a\x1e\x27\x14\x91\xe4\x63\x1c\x5d\xa0\x31\xc4\xf6\xd6\x86\xc6\x2f\x24\x1f\x09\x48\xd5\xc9\x02\x52\x84\x3d\x77\x5f\xe6\xaf\x4d\xb3\xd5\x91\xc7\x38\xd2\x68\x3b\x2d\xa1\x6a\x2c\xbb\x69\xb5\x90\xb6\x4a\x93\x9d\x12\x76\xcc\x31\xec\x98\x24\x07\x64\x3d\x34\xfa\x5a\xfd\x1f\xee\xcf\x82\xbb\xe7\xc6\x99\x1e\xb8\xe1\x0d\x1b\xb0\xba\x43\x10\x15\x39\x0f\xc2\x80\xaa\x60\x10\x00\x2e\x4b\x10\xd6\x40\x7f\x8e\x8c\x35\xdc\x16\x73\x24\x12\xd7\xb3\xce\x91\xad\x4b\x52\x32\xea\xa2\x9a\x0f\x36\x53\x04\x33\xc7\x31\x39\xe8\xaf\xce\x41\xdd\xc1\xf1\x84\x86\x2c\x0d\xd1\x78\x4d\x73\x8f\x71\xb4\x8a\x01\xad\x0d\x1c\x98\x4c\x86\x00\xa2\x65\x1f\x2b\x5e\x1b\xec\x69\xa6\xf4\xba\xc6\x4f\x9d\xd0\x58\x92\x87\x76\x8e\x1e\x53\x08\x01\xcd\x56\x6e\xd6\x25\x62\x25\xb0\x76\xfb\xa4\xb2\xd0\xaa\xb6\xab\xb9\x45\x50\x12\x38\xd9\xbb\xfa\x6d\x0a\x4a\xc3\xe9\xd9\xfb\xb3\xe9\x19\xa8\x16\x35\xa7\x2c\x8b\xef\xb9\xde\xba\x6b\xe2\x58\x73\xbe\x54\x69\xe2\x8c\x0b\xaa\x0e\xbf\x9c\xf8\xa2\xb9\x46\xab\x17\xbf\xca\x2b\x8d\x85\x92\xa5\x4b\xd9\xb7\x5c\xd4\x58\x12\x3a\x2d\xd0\x00\x07\x0a\x2e\x2d\x10\x4e\x61\xa1\xe2\xa2\x36\x50\x76\x08\x56\x39\x4c\x5b\x67\x3d\x1d\x5f\x35\x9d\x56\xce\x64\x9a\xf9\x38\x64\xd0\xf3\x94\x51\xc0\x2b\xa5\x41\x10\x1b\xe3\xd7\x20\xe0\x47\x78\xf9\x1a\xc4\xfe\xbe\xa3\x82\x36\x4e\xa0\x4a\xb3\x38\x22\xe6\x5e\x3c\x4b\x79\x0e\xa4\x37\x81\xab\xed\xeb\xdd\x1d\x2b\x56\xe3\x88\x98\x8b\x48\x66\xd8\x4d\x8d\xd8\xa6\x6e\x78\xda\xf9\xc0\xa6\x87\xe3\xf1\x9e\xc8\x60\x0f\xdc\xf4\x85\xa8\x6b\x61\x9c\xc1\x6c\x40\x79\x5f\x4a\x05\xec\x95\xdc\x72\x52\xc4\x93\x5a\x20\x25\x1a\xef\xec\x5c\x69\xf1\x05\x89\x18\x34\x36\xd5\xf8\x09\xf6\x1c\xc2\x30\x93\x83\x46\xa3\x3a\x5d\xe0\x74\xd1\x62\xff\xf5\x5e\xc8\xbb\x50\x8a\x0e\x73\x49\x99\x70\x3c\xf1\x50\x2e\xd5\xe7\x34\x63\xbf\x4d\x4f\xd2\x8c\xbd\x55\xba\xe1\x36\x4d\x2e\x94\xcc\x61\x7c\x04\xff\xe1\x12\x8e\xc6\xe3\x57\x70\xf8\xf2\x78\xfc\xc3\xf1\xf8\x25\xfc\x7c\x31\x25\xc2\xa3\x39\x9d\x27\x45\x67\x97\xf8\x39\xf5\x5a\x4e\xc3\x1c\x0a\xd6\x70\x63\x51\xbf\xc3\x45\x16\x47\x94\x39\x6f\x43\xbd\xcf\x73\x48\x76\xcc\x1f\x72\xf8\xef\x0f\x99\xe4\x01\x9a\x61\x53\xf5\x5e\x7d\x46\x4d\x8e\xb1\x0b\xb4\x73\x55\x66\xdf\x72\x69\xfb\x20\x39\x96\x11\x3c\xb2\xf0\x0b\xf2\x12\x35\xbb\x41\x9b\x26\x6f\x42\xec\x1c\x17\x49\x0e\x9d\xae\xd9\x7f\x3b\xd4\x8b\x33\x53\xf0\x16\xd3\x0d\x59\x22\xdd\x98\x78\x27\x76\xef\x51\x4f\x0e\xd9\x78\xd7\x88\xd9\xc4\xe9\x94\xef\x51\xec\xc6\x96\x67\xa1\x6d\x31\x37\xc0\xa9\xba\x71\x68\xd2\x39\xbb\xe9\x9a\x54\x8a\x3a\xa3\xdf\x36\x96\x87\x83\xc6\x1c\x10\xd2\x24\x07\x07\xf8\x9b\xac\x97\x2a\x2d\xec\x03\x84\x76\xca\x4e\xfc\xff\x39\x34\x2e\x3e\x39\xb4\xdc\xce\xbf\x83\xf7\x1c\xf0\xa1\xc5\xc2\x62\xb9\xd9\x01\x72\x10\x32\x07\xd5\x59\xfa\x40\x5d\xf1\x02\x1f\x97\x39\xcc\x1d\x5e\xe3\xf3\xdf\x83\x0f\x45\x46\xf9\x43\x22\xa1\xd1\xb4\xeb\xec\x33\xad\x92\x06\xfd\xc2\xba\x04\x63\x5f\x80\x54\xfe\x8b\x50\x84\x7e\xfc\xa3\xcb\x91\x87\x6b\xaf\x0b\x61\x36\xd4\x25\x99\xcd\xc1\x57\x67\xc1\xfe\xf2\xde\x7f\xbf\xb7\xcf\xb9\xb9\x76\x71\xed\xd6\xf7\x17\xfd\x54\xa9\x0b\x2e\x17\xa1\xbe\x4c\xa8\xf8\x5b\x8d\xfc\xce\x15\x7c\x1c\x45\x05\xab\xd5\x8c\xfd\xce\xb5\xa4\x8e\x66\x5c\x47\xe3\xd6\x62\xd3\x5a\xd8\x29\x43\x77\x7b\x8a\xdf\xea\x85\xf3\x92\x72\x35\x6a\x8c\x1b\x1f\x51\x90\xc2\xe3\x85\x5d\x71\x6d\xf0\x5c\x52\x91\x9b\x15\x03\xec\xe7\x75\xfa\x38\x0b\x07\xbc\xb2\xa8\x0f\x1a\x93\x64\x39\x1c\x8e\x73\x18\x07\xc7\x9c\xb1\x17\x13\x90\xa2\x1e\x4a\x94\x9b\x5f\xc1\xfe\xaa\x50\x35\xe6\xab\x2a\xe5\x9a\x9c\x23\x3e\x18\xdf\xdd\x5d\x67\xca\xe0\x3a\x62\xdd\x35\x41\x4d\x0f\x96\xf5\x0e\x87\xa4\xc4\x1a\x2d\xa6\x61\x2e\x87\xbb\x2c\x08\xa7\x3b\x94\xc3\x7d\x7f\x6e\xe0\xb9\x3f\x1b\x0e\x7d\xb8\xfb\x13\x26\x70\xef\xcf\x2d\xe3\x78\xe0\xdd\x37\x2b\xe9\xaf\x7f\x50\x29\xa5\x9b\x95\x93\x0f\x9a\x96\xc6\x4f\xbe\x00\x48\x60\x69\xd3\x25\x7e\x0e\xf9\xf7\xbb\xb0\xf3\x80\x7a\xb3\x2c\x12\xda\x68\x8e\x47\xa3\x64\xbf\x60\x2b\xbf\xdf\x14\x85\xea\xa4\xdd\x4f\x58\xa9\x8a\xae\xa1\xc7\x21\xe3\x5f\x3a\x8d\xee\xed\x9b\xec\x7b\x77\x49\xa7\x56\x8f\x97\x21\x89\x21\xa8\x52\xd4\x0e\xcd\x9a\x7e\x21\x87\xbb\x6e\xbb\x8a\x70\xee\xba\x67\x3f\xfb\xa9\xab\x2a\xd4\x8f\xcb\xd0\x5f\x8f\x27\xe0\xde\xd5\xe4\x81\xd7\x48\x9d\xde\x76\x15\xb5\x08\x7a\x42\xfe\xc2\x65\x59\x63\x16\xe4\x33\x15\xb2\xcf\xde\xe7\x92\xb7\x47\xe2\xf2\x85\x54\xf5\x27\x55\x2e\x60\x02\xfe\x4f\x04\x76\xa9\xda\x93\x5a\x19\x7f\x49\x16\xb6\x0c\x85\xd7\x85\x4e\xda\x03\x62\x36\xc9\x21\xe1\x6d\x5b\x8b\xc2\x65\xfd\x88\x00\x25\x21\xc9\xb7\x53\x71\x98\xc2\xbd\xd5\x0f\xeb\xbf\x40\xd8\x09\x97\x4a\x8a\x82\xd7\x17\xe7\x17\x67\x7e\xf5\x1d\x2e\xd2\xbb\x2c\xa4\xea\xf2\x99\x9e\xe4\x0a\xf9\x1e\xb5\xf1\x2d\x29\x39\x1a\x1f\xfe\xfb\xe0\xf0\xe8\xe0\x5f\x87\xae\xc3\x16\xec\xb9\x86\xff\xad\xfc\xf4\x8d\x6f\xa5\x9f\x14\x7d\x36\x2f\xd8\xa9\xa2\x73\xdf\xcb\x71\x54\x62\x85\x7a\xf5\xac\x5a\x29\xb2\x8b\x34\xbb\x46\x5e\xa6\x3e\x5f\x06\xb3\x2e\xe6\xf4\x8e\x5a\xa6\x04\xa0\xdc\xa0\xfd\x14\x3d\xed\xeb\xed\x4f\xc8\xef\x25\x65\xf8\x22\x7e\x31\x79\xae\xc4\x1e\xfb\xb4\xda\x75\xba\xed\x12\x6d\x75\x7c\x28\x91\x1b\x44\x67\xf4\xc0\xde\xe6\x3a\xc8\x11\xf3\x08\xd3\x5d\x27\xc6\x3e\xb5\x9e\xbe\xcf\x9f\xe2\xeb\x23\xb7\x8e\xf6\xba\x3c\xa8\xfc\x7b\x71\xfc\xff\x91\x6d\x5c\xd0\xa3\x54\x9d\xcd\x36\x44\xcf\x6f\x90\xa2\x8e\x97\xf1\xff\x02\x00\x00\xff\xff\xd0\x10\x8c\x8a\x99\x0f\x00\x00")

func cosmosdbGoBytes() ([]byte, error) {
	return bindataRead(
		_cosmosdbGo,
		"cosmosdb.go",
	)
}

func cosmosdbGo() (*asset, error) {
	bytes, err := cosmosdbGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cosmosdb.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _databaseGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x5f\x6f\xdb\xb6\x17\x7d\x26\x3f\xc5\xad\x1e\x02\xab\x55\xe5\xdf\x0f\x18\xf6\x60\xc0\x0f\x45\x12\xb4\x59\xd7\x6e\x48\x3a\x60\x40\x51\xb4\x34\x75\x6d\xb3\x95\x49\x8f\xa4\x96\x04\x41\xbe\xfb\x40\x52\x7f\x48\x49\x4e\xb0\xe5\xc1\x0e\x79\x2f\x2f\xcf\x3d\xe7\x90\xf4\x91\xf1\x1f\x6c\x87\xc0\x95\x39\x28\x53\x6d\x28\x15\x87\xa3\xd2\x16\x16\x94\x64\x5c\x49\x8b\x77\x36\xa3\x24\x43\xc9\x55\x25\xe4\x6e\xb9\x61\x06\x7f\xfe\xc9\x4d\x49\xb4\xcb\xbd\xb5\xc7\x8c\x52\x92\xed\x84\xdd\x37\x9b\x92\xab\xc3\xd2\x08\xdd\x1c\x0d\xca\x65\xad\x76\xba\x31\x59\x1a\x6d\x76\x4a\x7f\x17\xcb\x9d\x5a\x72\x55\x21\xcf\x68\x4e\xe9\x72\x09\x17\xcc\x32\x57\x1a\x34\x1e\x35\x1a\x94\xd6\x00\x83\xaa\x9d\xa5\xf6\xfe\x88\x43\x8e\xb1\xba\xe1\x16\x1e\x28\xb9\xba\x80\xfe\xcf\x58\x2d\xe4\x0e\xbe\x7d\x37\x4a\xae\x32\x51\x15\xea\x20\x2c\x1e\x8e\xf6\x3e\xfb\x46\xc9\x35\x1a\xd5\x68\x8e\x6e\x45\x9a\xf9\x55\x8f\x73\x3f\x89\x03\x1a\xcb\x0e\x47\x00\x10\xd2\xba\xe2\x5d\xae\x35\x69\xea\x0d\xd6\xdb\x79\x00\x5f\x0d\xd6\xdb\x34\xf9\xf2\x13\xdb\x9d\x48\x46\xcb\x76\x69\xf2\xb9\xaa\x6b\xe4\x56\x28\x69\xc6\xc9\x5c\xd5\xf5\x08\xc7\x1f\x06\xb5\x99\x2f\xdd\xb8\x50\x92\xfd\x98\x50\x6e\x62\xce\x3b\xc6\x4d\x4a\xb9\x89\x38\x3f\x57\x4d\xe0\xa4\x27\x07\x62\x86\xb8\x0b\x9f\xe4\xbe\xc5\x96\xac\x98\xf0\x3f\xec\x0a\x9f\xbf\xbc\xec\x65\x6f\xf3\xfb\xe8\xb8\x27\x8f\xb8\x6b\xe0\xbc\x16\x28\x6d\x04\xbb\x56\x3b\x88\xff\x5e\x06\x7b\x96\x97\xd2\xea\x7b\x4a\xf6\x1c\xd2\xb0\xb3\x76\x19\xaa\x50\xe2\x76\x7e\xc7\x64\x55\x63\x1b\xf5\xe6\x2d\x7f\xe9\xa7\x29\xe9\x36\x7e\xc3\x3d\x03\x6d\xa7\x94\x1c\x98\xb1\xa8\xdf\xe3\x7d\x5b\xf8\xf3\x97\xcd\xbd\x45\x37\x7f\x77\x8d\x56\x0b\x34\x3d\x97\x23\x61\xda\x16\x44\x7c\x12\x80\x07\x44\x89\x3a\x5d\xa2\xb4\xa8\xb7\x8c\xa3\x57\x49\x23\xb3\xb8\x68\x8f\x70\x79\x1e\xbe\x0b\xe8\xe9\xcc\x61\xd1\xff\x5f\x00\x6a\xad\x74\x4e\xc9\xaf\xc2\xd8\x45\xde\x57\xbe\xb2\xa8\x99\x55\x3a\x04\xde\xd4\xf5\xb8\x60\x5c\xc5\x0c\x65\xde\xa2\x9d\x6e\x1d\x18\x99\xdf\xf7\x02\x6b\x7c\x06\xae\x4f\x9d\xe8\xec\x70\x75\x28\x23\xb5\x5f\xa6\x3e\xa0\xc4\x55\x16\xb2\x61\xee\x40\xf5\xda\x54\x4a\x62\xaf\xf8\x46\xa9\x7a\xa4\x40\x5f\x38\xd5\x40\x74\xac\x24\x2a\x0c\xc9\xb1\x0e\x1f\xf1\x6e\x42\xc5\x3c\x69\x61\xeb\x8f\x78\x3b\x92\x55\xa3\x6d\xb4\x74\x00\x24\xde\x4e\x8c\xb0\x6d\x24\x9f\x2e\x5a\x38\xb7\x27\x0e\x2f\x60\xcf\x13\x53\x17\x10\x99\x7a\xe2\xe7\x02\x46\x7e\x2e\x60\x30\x72\x2f\x64\xba\x69\xd7\x89\x6b\xfb\x6f\xa6\xdd\xa8\x15\x8d\x12\x0e\xab\x35\x9c\xa5\xa2\x3c\x50\xe2\x4e\xe5\x2a\x39\x77\xb5\xda\x15\x94\x90\x3d\x4f\xa7\x61\xcf\xdd\xf4\x00\xb9\x0d\x0f\x13\x2e\x3c\xc2\xbc\x9a\x34\x41\x49\x74\xee\xda\x12\xff\xff\x5f\x41\xc9\xa3\xc3\x58\xf6\x2d\xfa\x56\x60\x0d\xe1\xb5\x2b\x6f\x6c\x75\xd9\x3e\x80\xe5\x05\x3a\xae\x6e\x3c\x05\x8b\x7e\x41\x4e\x89\xd8\xfa\x45\x2f\xd6\x20\x45\xed\x38\x20\x41\x39\x37\xf4\xf5\xc2\x2e\xed\x24\x2f\xdc\xbc\x13\xdd\x2b\xb8\xe0\x30\xb2\x6c\x0e\xcc\x1d\x37\x7b\x07\x93\x43\x21\x26\x9e\x9b\x75\x94\xc3\xc0\xea\xba\xda\x18\x4f\x7e\x1f\x7f\x70\x30\xb6\x4a\x7b\x8c\xd5\x26\xe4\xbb\x14\x51\x06\xb7\xda\xbb\x9c\x92\x99\x7e\xa6\x0d\x91\xc7\x90\xe8\xf6\x58\x47\x89\x1b\x8d\xec\x87\x0f\x53\xd2\x62\x28\xc3\xbb\xf1\x6a\x0d\xfd\x60\x88\x45\x2f\x44\x88\x0f\x13\x43\xd2\xf0\x2c\xac\x81\x1d\x8f\x28\xab\xc5\x38\x52\x40\x32\x2c\xcb\x32\x4f\x58\x0f\xf9\xcf\x53\xdf\xdd\x9e\x73\xec\x4b\xbc\xad\x36\xc9\x3d\x1a\x0f\x8b\xc1\xf6\x5e\x80\xe0\x23\x5e\x56\xca\x55\x2b\xc0\x9f\xc0\x0f\x68\xf7\xaa\xfa\x5d\x19\x5b\x40\x56\x6d\x4c\x36\x7c\x65\x6d\xca\x8d\x65\xb6\x31\x01\x47\x55\xc0\x99\xdf\xb5\x80\x33\xf7\x21\x45\x9d\x77\x2d\x3d\xd9\xc6\x89\xdb\xdc\xe1\x6a\x09\x39\x9b\xbb\x47\x1f\xd2\x3a\x2b\xe0\x8f\xcf\x6e\xf3\x66\xde\xac\x27\x8d\xd9\x1d\x83\xb2\x75\x79\x01\xbc\x0c\x70\xf3\x27\xf7\xf2\x2f\xcb\x9c\x2c\xd5\x46\x54\xc3\xcd\xf4\x1f\x25\x79\x8b\xad\x22\xcb\xec\x95\x2b\x98\xc8\xd3\xcf\x45\x02\xfd\xf6\xbe\x08\xc7\xe1\xdf\x09\xd3\x3d\x77\xf3\x8d\x4c\x5e\x3d\x87\xda\x9f\xb2\xd2\xff\x84\x5c\xaf\x21\xcb\xe2\x1b\xe6\x52\x6b\x17\xb8\xc6\xbf\x1a\xa1\xb1\x72\x96\x27\x7b\x64\x95\xfb\x4d\xb8\x5a\x07\xbc\xef\xfc\xf8\x61\x88\x94\x37\x68\x17\xd9\xd5\xf6\xf5\x07\x66\xf9\x3e\x2b\xba\xf2\x79\x24\xce\x0c\x45\x01\x7a\xc4\x48\x79\x75\x31\xa5\xc9\x4f\x46\x3c\x7d\x54\xbe\x3f\xf7\x48\x78\xba\xfc\x47\x0b\x24\x12\x5c\x0c\x54\xc5\x86\xcc\xa1\xbb\x98\x66\xfc\xe5\xee\x9e\x91\xc7\x22\xb1\xc5\x16\x44\xe9\xdf\xf9\x81\xaf\x70\x25\x9c\x26\xc8\xaf\x49\x7e\x30\xbc\xe8\x19\x4f\xc8\xfb\xf3\xf5\x07\xf3\xfa\x3c\x4a\xcc\x8a\xd1\xca\xf6\xfa\x09\x96\x13\x4f\x5a\xee\xe4\x25\x10\x7b\xcc\x44\xac\x9d\x7c\x71\xc2\x9e\xa3\x16\xd6\xdd\xc2\xf2\xed\x3c\x74\x57\x30\x30\xb5\x1e\xb7\xef\x0d\x47\x23\x67\xff\x13\x00\x00\xff\xff\x21\xcd\xa3\x1b\x37\x0e\x00\x00")

func databaseGoBytes() ([]byte, error) {
	return bindataRead(
		_databaseGo,
		"database.go",
	)
}

func databaseGo() (*asset, error) {
	bytes, err := databaseGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "database.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _documentGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\xcd\xaa\xc2\x40\x0c\x85\xd7\x9d\xa7\x08\x5d\x5f\xe8\xfe\x3e\x84\xe8\xc6\x8d\x08\x8d\x35\x94\xaa\x99\x19\x93\x54\x28\xe2\xbb\x4b\x7f\x9c\x32\x9d\xd5\x24\xe7\x7c\xe4\x8b\xd8\xdc\xb1\x25\x68\x82\x72\xd0\xeb\xc5\xb9\xaa\x82\x43\x4f\x32\x80\x50\x14\x52\xf2\xa6\x80\xf0\x1c\x57\xce\x86\x48\x4b\xaa\x26\x7d\x63\xf0\x76\xc5\x3c\x4f\x4f\x4d\x3a\xdf\xce\xff\xfa\xa6\xc1\xff\x97\x13\xf8\x17\xb8\x33\xe2\x68\x43\x59\xbb\x62\x8f\x82\x4c\x46\xa2\x70\x3a\xa7\xe1\x07\xc4\x94\x66\xd4\x67\x32\x5b\xdb\x99\x5d\x62\x66\xc3\xb5\xb5\x5a\xee\x90\x29\x09\x2e\xa7\x3c\x32\xe5\x6a\x47\x7c\xf4\xb4\x69\xbd\xc6\xdd\xc6\xe5\x1b\x00\x00\xff\xff\xa1\x09\x70\x18\x37\x01\x00\x00")

func documentGoBytes() ([]byte, error) {
	return bindataRead(
		_documentGo,
		"document.go",
	)
}

func documentGo() (*asset, error) {
	bytes, err := documentGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "document.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4d\x6f\xe3\x36\x13\x3e\x8b\xbf\x82\xab\x83\x61\xc5\x92\xfc\xbe\xd7\x00\x2e\xb0\x48\xd2\xad\xbb\x4d\x76\x9b\xcd\xa1\x40\x51\x20\xb4\x34\xb6\xb9\x91\x45\x2d\x49\x21\x09\x0c\xff\xf7\x82\xd4\x87\x29\xea\xc3\x76\x36\x68\xb3\xcd\x25\x89\x34\x9c\x79\x38\xf3\xcc\x97\x32\x12\x3d\x90\x15\xe0\x88\x89\x0d\x13\xf1\x02\x21\xba\xc9\x18\x97\x78\x8c\x1c\x37\x62\xa9\x84\x27\xe9\x22\xc7\x4d\x41\x4e\xd7\x52\x66\xea\x6f\x21\x39\x4d\x57\xc2\x45\xc8\xc9\x1e\x56\xd8\x5d\x51\xb9\xce\x17\x61\xc4\x36\xd3\xaf\x74\x13\x6c\x68\x2a\x81\x4f\x57\x2c\xa8\x94\x4e\xb3\x87\xd5\x74\x05\x69\xfd\x7f\xfd\x47\x9c\x6f\x36\xcf\x2e\xf2\x10\x92\xcf\x19\x60\x09\x9b\x2c\x21\x12\x2e\x12\x0a\xa9\xc4\x42\xf2\x3c\x92\x78\x8b\x9c\xb3\x98\x48\xb2\x20\xa2\x7c\x83\x9c\x8c\xc8\x35\x2e\x80\xa0\x1d\x42\xd3\x29\xbe\x6b\x9e\xa5\x02\x93\x5a\x1f\x8e\x8a\x63\xda\x88\x2d\xa8\xd0\x2e\x49\x04\xca\xce\x05\x07\x22\x61\x5c\x5e\x3c\xbc\x28\x7e\xfb\xa5\x25\x1f\x9f\x65\x0f\xab\xb0\x52\xe0\xe3\xb3\x4f\x99\xa4\x2c\x15\x1e\x1e\x5b\x6f\x80\x73\xc6\x3d\xe4\xfc\x46\x85\x1c\xef\xc5\x2a\x81\x5b\xf2\x38\x97\xc0\x89\x64\xbc\x90\x79\x9f\x24\x6d\xab\x7d\xea\xc5\x5e\xff\x07\x90\xfd\x68\x6b\xd4\x07\x61\xde\x42\x96\x90\xe8\x35\x2f\x7e\x09\x09\xbc\xc8\x93\x5a\x01\x72\x7e\xcf\x81\x3f\x8f\x6b\x69\xfd\xaf\x29\xd6\xe9\x49\x2d\xd5\xe9\xca\x7e\x45\xbd\xae\xbd\x58\x93\x74\x05\x3f\x03\xc4\x1d\x01\xac\x6d\xee\x6c\xea\xd6\xa7\x2a\x11\x93\xc6\x4d\x82\x23\x47\xe1\xa4\x69\x4e\x94\xf6\x8a\xce\x0e\x2b\x8c\x61\xfd\x53\x99\x6e\x19\x52\xac\x79\xb1\x89\x98\xa5\x80\xab\x9f\x05\x63\xc9\xb1\x56\xb5\xfb\x8e\x31\x9b\x11\x2e\xa9\xd2\xf1\x00\xcf\xb5\xd9\x6f\xea\x74\x65\xb6\x08\xc5\x2b\x00\x34\x92\xbf\x46\xd6\x4c\x7f\x5a\x45\xab\x51\x00\xf6\xc2\x66\x09\xb8\x81\xa7\x56\x4a\xf5\xb3\xa4\x69\xde\x20\xa3\x85\x80\x93\xc7\x1e\x14\x8d\x33\x26\x90\x16\xd5\x34\xb4\x5b\xf2\xd8\x66\x77\x7d\x6e\xbb\xab\x12\xa8\x00\x76\x03\x8f\x56\xb9\xe3\x20\x73\x9e\x2a\x6c\x29\x3c\xb6\x0a\xe4\x32\x4f\xa3\xf6\xa1\x71\xc4\x92\x24\xc2\x17\x2c\x49\x20\x52\x6e\x2f\x1e\xfb\x58\x3d\xa7\x71\x19\x34\xcf\xae\xac\x5b\xe4\x14\xd6\xf0\xa8\xc9\x8e\x2d\x72\x9c\x66\x41\x3f\xd7\xaa\xa2\x70\x7c\x16\x59\x46\xbc\xb0\x29\xe9\x23\x47\x17\xff\x73\xbc\xff\xe9\x3f\xab\xdb\xc4\x04\xbb\x53\xf5\x4a\x4c\x5d\x3c\x29\x31\xfb\xc8\xd9\x29\x1f\xe9\x0b\x8f\x23\x6c\xf1\xd7\xc3\x44\x55\x11\xf9\x84\xdb\xbe\x6e\xf1\xa7\x97\x1d\xca\x03\x24\x49\x2a\xd5\x02\x9f\xcf\xf0\xa8\x21\xb9\xdd\x21\xe4\x2c\x19\x57\x92\x8e\x6c\x9c\x57\xc2\x34\x2c\xd8\x28\x9f\x3c\xe4\x38\x74\xa9\x9f\xbf\x9b\xe1\x94\x26\xfa\x44\xe5\xdf\x94\x26\xfa\x0c\x72\x9c\x5d\x21\xb8\xb7\x39\x33\xc4\x17\x1c\xc8\x83\x16\x42\x4e\x03\x59\x78\xc1\xf2\x54\xe2\xc9\x0c\x5b\x8f\x6c\xb9\x5b\x10\x2c\xe7\x11\xcc\x2f\xb1\x29\xbb\x7f\x6c\x1f\xb8\xdb\x03\xc1\x24\xcb\x20\x8d\xc7\xdd\xef\x7d\xdc\xf1\x30\x0c\x43\x0f\x69\xb8\xe5\x4d\xcd\xb3\xbe\xba\xd8\x60\x10\xab\x76\xde\x15\xc7\x8e\x02\xe5\xab\x9c\xa8\x53\xc2\xea\x50\x55\xf9\x31\x3a\x47\x9f\xa8\x8a\xd2\x9e\x01\x6b\x20\x31\x70\x1d\x7c\x35\x41\x85\xbf\xe8\xff\xb7\xbb\xfa\x4d\xf8\x05\xe4\xd8\xfd\x23\xb8\x16\xc1\x25\x8b\xf2\x0d\xa4\x32\x5e\x04\x9f\x0d\x7c\xae\x8f\xef\xff\x74\xef\x27\x26\xe6\xc9\xbd\xfb\xd7\xbd\x87\x90\x8a\x76\x85\xcd\x88\x75\xfd\x08\x8f\x4a\xc0\xca\xe2\xae\x2e\xa3\xe1\x0d\xbb\xba\x23\x2b\x15\x45\x9e\x03\x42\x8e\x02\x3d\xc3\x51\x28\x40\x96\x07\xc6\xa5\x68\xc3\x2d\x3e\x2e\x51\x7b\xa8\x83\x90\x45\x94\x8a\x88\x55\x0a\x63\xa6\x02\xe0\x17\x97\xbf\x06\xb9\x66\xf1\x67\x26\x54\xf5\xd0\xe9\x39\x71\xa7\x31\x8b\x84\xeb\x63\xb7\xfc\x5d\x3c\x2f\x0f\x7c\x91\x44\xe6\xa2\x08\x64\xec\xe3\x51\x03\xca\xa8\x0b\x55\x89\x61\x88\x17\x7a\x2a\x6b\x07\xb4\xab\x24\x77\x14\x31\xb3\xf3\x6e\x9b\xaa\xcf\x71\x54\x33\xe5\xbc\xfa\x63\xb8\xd0\xd4\xd3\x5f\x17\x49\x3b\x48\xd7\x5f\x6a\x4a\x9c\x51\x58\xd6\x2e\xe5\x49\xf3\xa6\x9e\x37\x08\x44\x8f\x92\x87\x32\x65\x9f\xa5\x75\xdd\x7f\xc3\xa9\x31\xc8\x69\x55\x33\x5f\x81\xcb\x1f\xc0\xa6\xf2\xd4\x9d\xec\x9d\x64\xd3\xba\x5b\xc6\xa0\xfa\xa7\x8f\x25\xb6\x97\x92\xbb\x9e\xe5\xff\xd3\x55\xef\x1f\x2d\x57\x79\x47\x88\x0d\x23\xe1\xfc\xb2\x3f\xcc\xb6\x9c\x15\xea\xd7\x29\x68\xd5\xb6\x75\x74\xc8\x4f\x88\xf7\xdb\x89\xeb\xab\x06\xb5\x70\x59\x7f\x5a\x0e\x06\xb5\x2f\xa2\x37\x4c\xbb\x5d\xcd\xc6\x3a\x87\xad\x22\x73\x44\x24\x8b\xbd\xb7\x33\x66\xc5\xf6\x54\x6d\xb0\x2f\x6e\x5f\x8d\x15\xae\xab\x7f\x99\xc6\xcf\xad\xea\xaf\x31\x9c\x17\xbf\x4e\xed\x74\xfb\xed\xfc\x68\x9a\x1e\xba\xf2\x89\xdd\xb0\xed\x5c\xdf\xba\xca\x81\x1e\x69\x7c\x13\xe8\xf7\xff\x90\xf3\xdb\x9f\x07\xda\x11\x18\xf6\x62\x3b\x31\x6a\x08\xfd\x69\x5d\x65\xad\x91\xb2\xe5\xa2\xa8\x40\x76\x4f\x90\xfb\xdd\xa2\xc8\x25\x63\xab\xa8\xd2\x6d\x34\xc2\xef\xac\x79\x72\xdb\xdc\x3f\xc2\x62\xc8\x9c\x61\xd7\x6d\x6c\x2c\x57\x9c\xab\x37\xb7\xf0\x2d\xa7\x1c\xe2\x72\x71\x69\x94\x90\xf9\x32\xb8\x26\x32\x5a\xbb\x7e\x53\x9b\x5e\x08\x94\x8d\x04\xd2\xca\x05\xe1\x67\x0e\x77\x9c\xae\x56\x2a\xcf\xf0\x4f\xf8\x7f\xda\xd8\x70\x45\xe2\x10\x94\x47\x82\x79\x1a\x25\x79\x0c\x6e\xf5\x9d\x48\x84\xbf\x32\xda\xa9\xdc\xc7\xae\xef\x7a\xdd\x10\x98\x90\xa7\x62\x60\x42\x1e\x0d\xc2\x50\x3f\x84\xa2\x22\xf7\x47\x78\xbe\x55\x5c\x9b\x5f\x1e\x0b\xa6\x7d\xd2\xad\xf3\xa2\x53\x6f\x63\x33\x6b\xec\x62\xd4\xe0\x6d\x8b\xf1\x1e\xae\xb6\x5a\xbb\x0a\x18\xd3\x85\xc0\x1d\xa9\x7d\x7a\x1f\x7a\x1f\xcc\xaf\xd5\x6a\x31\x4f\x23\x0e\xea\xa2\x24\xc1\x4b\x80\xd8\x55\xbd\xa6\xed\x8e\x6b\xf2\x14\xcc\x25\x6c\x02\xbd\xfc\xaa\x83\xc1\xff\xdd\xa2\xcb\xd0\xb0\xf1\xa1\xea\x5d\x4d\x69\x9b\xb4\x37\x2c\x85\x9a\xb9\xcd\x53\x9e\xd9\x95\xa8\xd9\xe3\x68\xf8\x3d\x83\x29\xed\x1d\x4c\x69\xcf\x8e\x45\xdb\x3b\x56\x7b\xf0\x14\x16\x8e\xb9\xb8\x52\xfe\x2f\x57\x32\x16\x83\x9a\x0d\xac\xfe\x27\xaf\x59\x4c\x97\x14\x62\x1d\xa5\x12\x5f\x59\x43\x0e\xdc\xc5\xf2\xf0\xac\x32\x1e\xaa\xcd\xc4\xbd\x92\x64\xa5\xa3\x66\x37\x51\x83\x6b\xe6\x66\xf6\x2a\x2c\xab\xdc\x5b\x7f\x79\x53\x3e\xde\xfb\xa7\xa3\xa5\x0f\xa2\x29\x55\xb4\x9b\x9f\xfe\x3c\x68\x7e\xc4\xb3\xa6\x2e\xcd\x3f\xfd\x4d\xd4\x76\xda\x29\xf3\xd8\x77\xb3\x5b\x6b\xb9\x30\x04\x7f\x54\x86\x73\xf2\x78\x3c\x82\x61\x5e\xb6\x7d\xa2\x14\x16\xc1\x9a\xd9\x7e\xd5\x8d\x70\x98\xc3\x8d\xf1\xec\xdf\x27\x71\x07\x9c\x37\xca\xe2\xc1\xe6\x36\x17\x7a\xce\x53\xd2\x77\x3c\x87\x96\x7c\x39\xb6\x07\x77\xcf\x99\xea\xc0\x2e\xc9\xb2\x84\x46\x3a\x68\x53\x7d\x72\xf2\x55\x94\xa1\x5d\x6a\x6a\x19\xc3\xea\x50\xb2\x1c\x58\x81\x9a\x9a\xca\x25\xc8\xd9\x61\x48\x04\x1c\xd6\xa8\x63\x13\x5c\xa5\x64\x91\x40\xc4\x99\x10\xb5\x32\xf3\x9e\xbb\x1f\x37\xc1\x8b\xef\x84\x27\x66\xf8\x88\x86\xe5\x4c\xff\x76\xd2\xfc\xef\x00\x00\x00\xff\xff\x03\x7f\x59\xe7\x67\x1f\x00\x00")

func templateGoBytes() ([]byte, error) {
	return bindataRead(
		_templateGo,
		"template.go",
	)
}

func templateGo() (*asset, error) {
	bytes, err := templateGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _triggerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x4f\x6f\xdb\x36\x14\x3f\x8b\x9f\x82\xe5\xa1\xb0\x6a\x55\xbe\x07\xd0\x21\x73\x8a\x2c\xd8\x92\x16\x49\x0e\x03\x86\xa1\x65\xa8\x67\x87\xab\x2c\x6a\x24\x8d\x24\x30\xf2\xdd\x07\x92\xa2\x44\x4a\xb4\x97\xa2\x87\xf9\x90\xc8\x8f\xef\xef\xef\xf7\xde\xa3\xdc\x51\xf6\x9d\x6e\x01\x33\xa1\x76\x42\xd5\x0f\x08\xf1\x5d\x27\xa4\xc6\x0b\x94\x11\x26\x5a\x0d\xcf\x9a\xa0\x8c\xb4\xa0\x57\x8f\x5a\x77\x04\xe5\x08\xad\x56\xf8\x5e\xf2\xed\x16\x24\x96\xd0\x49\x50\xd0\x6a\x85\x29\xd6\x4e\x88\xf4\x4b\x07\x83\x86\xd2\x72\xcf\x34\x3e\xa0\xec\xea\x02\xc7\x1f\xa5\x25\x6f\xb7\x81\xe0\xdb\xdf\x4a\xb4\x67\x84\xd7\x85\xd8\x71\x0d\xbb\x4e\xbf\x90\x6f\x28\xbb\x05\x25\xf6\x92\xc1\xe0\xe0\x98\xe1\x57\x39\x35\xbd\xe7\x3b\x50\x9a\xee\x3a\xaf\xc9\x5b\x1d\x27\xe1\x4d\xb5\x8a\x2d\xef\xa0\xd9\xbc\x29\xdb\xaf\x0a\x9a\x4d\x6c\xfb\xe9\x9e\x6e\xdf\x66\x0b\x9a\x6e\x63\xdb\x5f\x44\xfd\xf2\x26\xdb\x07\x51\xbf\x4c\x8a\x75\x90\x7f\xee\x40\x52\xcd\x45\x8b\x67\x82\xde\x54\x4f\xe4\x49\x37\xf7\x86\x45\xfb\x99\x09\x62\x37\x46\x1e\x79\x78\x0d\x5b\x64\x0c\x9e\xea\x15\x2c\xfc\x71\xd4\x35\xa3\x91\xab\x3e\xed\x90\x89\x56\x69\xda\x6a\x85\xec\x93\xe9\xd9\xa9\xce\x79\xd3\x84\x25\x8c\xb6\x15\x26\xe7\x4d\x43\xe6\x16\x6b\x09\x54\x43\xda\xc2\x9d\x25\x8c\x6e\xa1\x6b\x28\x4b\xa4\x5f\x61\xd2\x9f\x25\xac\x2e\xa0\x81\x63\xa1\xdc\xd9\x64\xdc\x2c\x03\x49\x18\x0d\x78\x11\x82\x56\x75\x0e\x9e\x15\x1f\xc7\xcd\x1c\x7f\x91\x10\x53\x5e\x61\xf2\x45\x06\xf9\x5b\x25\xa1\xf4\x4c\x49\x28\x3d\x49\x58\x85\xd9\xf6\xb9\xaa\x28\x4f\x15\x2c\x88\xb5\xd8\xfb\xe9\x0c\xc6\xd4\x8f\x0a\x33\xa7\x47\x17\x43\x30\x25\xc7\x97\x81\x0f\x89\xf1\x9f\x7f\x7d\xf0\x0b\xaa\x57\xf7\x87\xd3\x4e\xb6\xc9\xf6\xa9\xaf\x1b\x0e\xad\x0e\x32\xfe\x50\x53\x4d\x1f\xa8\x02\x77\x82\xb2\x8e\xea\x47\x0f\x7b\x34\x06\xbd\x29\x0f\x49\x63\xce\x28\x84\xc3\xab\xb5\x1a\xe4\xc6\x74\x94\x81\xc5\x76\xdd\xa2\xdf\xc6\xe5\xda\xfd\x2f\xb0\xaf\x20\xc7\x0b\xff\x58\x60\x90\x52\xc8\x1c\x65\xbf\x73\xa5\x17\xb9\x77\x7b\xa5\x4d\x6b\x09\xe9\xe4\xe7\x4d\x33\xf5\x16\xf8\x50\xa3\x93\x4b\xd0\xf3\xb0\xae\xba\x64\x50\xd7\xb4\xa7\x32\xb5\x9a\x86\x38\x3b\x13\x3f\x56\xd3\x84\x0c\x53\x89\xaf\x2b\xa4\x24\xe2\x0a\x65\x26\x04\x6f\xf7\xd1\x36\xc9\x6a\xd1\xc2\xb0\x4f\x1f\x84\x68\x62\xb2\x06\xb7\x11\x5d\xdc\x83\x18\x12\x36\xaa\x86\x94\xdd\xc0\xf3\x0c\xb9\x24\xc4\x2e\xec\x0d\x3c\xc5\xfc\x4b\xd0\x7b\xd9\x9a\xe0\x2d\x3c\x4d\xfb\x65\xb3\x6f\xd9\xcc\x64\xc1\x44\xd3\x30\xbc\x16\x4d\x03\xcc\xee\x31\x2b\x2e\xb0\x91\xf3\x7a\x60\x2d\x8e\x73\x40\x99\x0b\x85\xdf\x47\xb0\x1d\x50\x96\xc5\xad\x7d\x66\x1d\xb1\x72\xf1\x81\x4d\x42\xe4\x65\xac\x59\xa0\xcc\x8e\xc1\x59\x70\x65\x1d\xb7\xb5\x03\xb3\xc4\x64\x65\x8e\xd4\x8a\xe0\x65\x9f\x71\x81\xb2\x57\x03\x8f\xad\x76\xc1\x70\xcc\x6b\x8e\xa9\x69\x62\xfd\x8c\x67\x2d\xc4\xa7\xd4\xa4\x80\x37\x95\xd3\xa6\xf1\x2b\x09\x9f\x55\xf8\xbd\xd7\x39\xbc\x22\x94\x6d\x84\x34\x3a\x99\x0e\x0d\x8d\x1a\x2f\x1d\xb9\xfa\x39\x47\x59\xc6\x37\x56\xfe\xae\xc2\x2d\x6f\xac\x81\x07\xb4\xe5\x8d\xb5\x41\x59\xf6\xea\x14\x87\x60\x55\xa0\xfd\x20\x81\x7e\xb7\x3a\x28\x0b\x33\x2a\xdd\x32\x5c\x56\x38\x96\x4c\xb4\x82\x0d\x18\x68\x8e\xd2\x89\xfa\xb0\x01\x2b\x4c\xbb\x0e\xda\x7a\x91\x3a\x2d\xf0\x4c\x54\x96\x65\x8e\x6c\x92\x7d\x79\x81\x5d\x61\xaa\x39\x45\x95\x5f\x60\x29\xb6\x5a\x78\xf2\xed\x1d\xcc\xfe\x54\xe4\xd0\x1f\xa9\x33\xdf\x2a\xcc\xca\x5a\x18\xaf\x05\x36\xaf\xa6\xe5\x35\xe8\x47\x51\x9b\x5b\xa8\xc0\xcc\x36\xd6\x92\xac\x7c\x92\xa4\xc0\x24\x78\x76\xe7\xbd\xe1\x9d\xa6\x7a\xaf\x5c\x96\x75\x81\xdf\x8f\x39\x15\xc3\x5c\xd8\x22\x73\x5f\xfe\xa9\x6a\xd3\x9b\x37\x31\x6a\xe1\x02\x3b\x44\x4e\xce\x30\x3b\xd9\xfb\xc3\x12\x9f\x23\x7a\xac\xd9\xfb\xd8\xac\xec\xe7\xc6\x40\xe0\x32\xcd\x4f\x45\xb2\x37\x40\x8a\xb7\x5e\x2f\xd8\x2c\x3f\x43\xda\x25\xa4\x38\x5b\x91\xe5\x10\x26\x45\xdf\x51\xd5\x80\xd5\xcf\xbf\x15\x6e\x14\x7f\x9c\x48\x7f\x9b\x9d\x28\x7f\x7a\xb3\x99\x32\xc7\x49\x2f\xed\xef\x80\xaa\xc2\x84\xd8\x61\xef\x39\xf8\x24\xa5\x39\xb8\x85\x7f\xf6\x5c\x42\x6d\xc6\x2a\x7b\x04\x5a\xf7\x7b\xc8\x26\xff\xab\xfd\x7e\x18\x4f\xca\x3b\xd0\x0b\x72\xb5\xf9\x78\x4d\x35\x7b\x24\x45\x14\x23\x0f\xf8\x4d\xa0\xeb\x0a\x39\x89\x5a\x79\x75\xf1\x66\x84\xad\x6e\x00\xf1\x8d\xb0\xb0\x98\xcb\xc6\x22\x6d\xff\xf4\x69\x9f\x6c\xae\xe1\x2d\xe0\xff\x5c\x0c\x2b\xb2\x1c\x43\xbd\x09\x86\x89\xfa\x4f\xaf\x10\x3e\x00\x13\x6e\x84\x1c\xfb\x7b\x26\x31\xe2\xc3\x55\x12\xcf\x7a\x00\x04\xdf\x60\x5e\xda\xd7\x9c\xb1\xf3\xdc\x02\x3f\xde\x6a\xd6\x26\x7a\x5f\x7a\x37\xf4\x6e\xd4\x86\x7f\x7c\xbc\x56\x1f\xd7\x81\x22\x29\x26\x96\xfd\x65\xe1\xe8\xe0\x47\x47\x9e\xff\xc7\x9a\xe6\xf3\x35\x3d\x1b\x68\x15\x34\x5b\xe2\x3e\x0e\x4b\x9f\x54\x57\x79\xc3\xf2\x32\x5d\x95\x71\xe8\x40\xac\xa6\xc8\xd8\xa9\x46\x01\x99\xff\x06\x00\x00\xff\xff\xd4\x0c\x52\x6d\x4a\x11\x00\x00")

func triggerGoBytes() ([]byte, error) {
	return bindataRead(
		_triggerGo,
		"trigger.go",
	)
}

func triggerGo() (*asset, error) {
	bytes, err := triggerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "trigger.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"collection.go": collectionGo,
	"cosmosdb.go":   cosmosdbGo,
	"database.go":   databaseGo,
	"document.go":   documentGo,
	"template.go":   templateGo,
	"trigger.go":    triggerGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"collection.go": &bintree{collectionGo, map[string]*bintree{}},
	"cosmosdb.go":   &bintree{cosmosdbGo, map[string]*bintree{}},
	"database.go":   &bintree{databaseGo, map[string]*bintree{}},
	"document.go":   &bintree{documentGo, map[string]*bintree{}},
	"template.go":   &bintree{templateGo, map[string]*bintree{}},
	"trigger.go":    &bintree{triggerGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
