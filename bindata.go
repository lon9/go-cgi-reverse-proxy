// Code generated by go-bindata.
// sources:
// templates/.DS_Store
// templates/code_template.go.tmpl
// DO NOT EDIT!

package main

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesDs_store = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x98\x3b\x0e\xc2\x30\x10\x44\x77\x8c\x0b\x4b\x34\x2e\x29\xdd\x70\x00\x6e\x60\x45\xe1\x04\x5c\x80\x82\x2b\xd0\xfb\xe8\x24\xda\x11\xb2\x14\x52\x50\x25\x82\x79\x92\xf5\x56\x8a\x9d\x4f\xe3\xec\xd8\xcc\x30\x3c\x1f\x17\xb3\x3c\x95\xc9\xdc\x76\xb6\x8f\x24\x8e\x05\xa1\xab\xc1\x7b\x08\x21\x84\x10\x62\xdf\xc0\x95\x8e\xdb\xbe\x86\x10\x62\x87\xcc\xfb\x43\xa1\x2b\xdd\xdc\xe0\xf5\x40\xc7\x6e\x4d\xa6\x0b\x5d\xe9\xe6\x06\xe7\x05\x3a\xd2\x89\xce\x74\xa1\x2b\xdd\xdc\xdc\xb4\xc0\xf0\x01\x3e\x19\x4c\x28\x60\x0a\x41\xa1\xeb\x97\x1f\x2d\xc4\x9f\x70\x70\xe5\xf9\xff\x7f\xb5\xd5\xfc\x2f\x84\xf8\x61\x10\xc7\xdb\x38\xd8\x3b\x10\x2c\x27\x4c\xe3\xde\xd5\xcd\xd6\x9b\x80\xe0\x87\x85\xa7\x6e\x6d\xa1\x2b\xdd\xdc\x6a\x04\x84\xd8\x8a\x57\x00\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00")

func templatesDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_templatesDs_store,
		"templates/.DS_Store",
	)
}

func templatesDs_store() (*asset, error) {
	bytes, err := templatesDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1469352040, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCode_templateGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x92\xcf\x6e\xc2\x30\x0c\xc6\xcf\xcd\x53\x78\x3d\xa5\x53\x15\x76\xe2\xb0\x89\x03\xe2\x8f\x38\x4c\x50\x15\x24\x8e\x53\x05\x6e\x17\x2d\x24\x5d\x92\x8e\x4d\xac\xef\x3e\xa7\xd0\x89\x4d\xf4\xd0\x38\xb6\xeb\xef\xfb\x59\xad\x8b\xdd\x5b\x51\x21\x1c\x0a\xa9\x19\x93\x87\xda\x58\x0f\x9c\x45\x71\x79\xf0\x31\x1d\xca\x54\xe1\xd0\xe8\x07\xaf\xde\xd7\xd7\xf1\x60\x57\xc9\x3f\xf7\xf0\x6a\xbc\x54\x7d\xb2\xb1\x5d\x68\x5c\xcc\x12\xc6\xca\x46\xef\x3a\x19\x9e\xc0\x89\x31\x80\x32\x05\xb4\x16\x1e\x47\x60\x9c\x58\xd5\xa8\xe7\x52\x21\x8f\x4f\x27\x10\xcf\xa6\x0a\x17\x68\xdb\x38\xed\xaa\x2f\xf9\x74\x9b\x7f\x77\xd1\x24\x9f\x8d\x37\xb3\x73\x3c\xce\xb2\xd9\x72\x9a\xc2\xc3\x70\x38\x4c\x58\x24\xcb\x6e\xe2\xdd\x08\xb4\x54\x24\x12\x45\x64\x5f\xcc\x0b\x5f\x28\xa5\x39\x95\xa8\xa7\x65\xd1\x1e\x4b\xb4\x50\x8a\x89\x32\x0e\x39\xe5\x42\xd7\x1a\xfd\xaa\x09\xfe\x79\x49\x66\xfb\x51\x64\x8e\x28\xa9\x68\x3f\x90\x07\x3e\xb1\x28\xf4\x5e\xa1\x9d\x13\x0d\x0f\x48\xfc\x08\x5d\x3e\x47\x57\x1b\xed\x70\x6b\xa5\x47\x9b\x82\x85\xfb\x4b\xfe\xbd\x41\xe7\x03\x33\xd0\x43\x3b\xf9\xc5\xa6\x58\x64\x85\x75\x17\xe8\x4e\xc4\x06\x66\xb2\x74\x83\xe5\x06\x4c\xa0\x89\x6a\x6b\x3e\xbf\xc2\xbc\x7e\xff\x62\x89\xc7\xb5\xd4\x95\xc2\x85\x71\x3e\x47\x9a\xea\x30\x0b\x5d\x9c\x24\x93\xfe\x93\xb3\xe0\x62\xb3\xc9\xf8\x91\xfc\x86\xdd\x24\xc9\xd3\x7f\x55\xfa\x0f\x44\x66\xa5\xf6\x57\x1b\x6c\xd9\x4f\x00\x00\x00\xff\xff\xa8\x6e\x6d\x89\x37\x02\x00\x00")

func templatesCode_templateGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCode_templateGoTmpl,
		"templates/code_template.go.tmpl",
	)
}

func templatesCode_templateGoTmpl() (*asset, error) {
	bytes, err := templatesCode_templateGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/code_template.go.tmpl", size: 567, mode: os.FileMode(420), modTime: time.Unix(1469361376, 0)}
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
	"templates/.DS_Store": templatesDs_store,
	"templates/code_template.go.tmpl": templatesCode_templateGoTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		".DS_Store": &bintree{templatesDs_store, map[string]*bintree{}},
		"code_template.go.tmpl": &bintree{templatesCode_templateGoTmpl, map[string]*bintree{}},
	}},
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

