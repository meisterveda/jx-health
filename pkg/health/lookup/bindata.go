// Code generated for package lookup by go-bindata DO NOT EDIT. (@generated)
// sources:
// pkg/health/lookup/static_data/info.yaml
package lookup

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

// Mode return file modify time
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

var _pkgHealthLookupStatic_dataInfoYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xd4\x4d\x4e\xf3\x30\x10\xc6\xf1\x7d\x4f\xd1\x03\xbc\xae\x5f\x3e\x56\xdd\x21\xe8\x92\x0d\x37\xb0\x9d\xa1\x76\x6d\xcf\x44\x9e\xb1\x9a\xde\x1e\xb5\xa4\x1f\x02\x82\x1a\xa5\x5c\xe0\xff\x7b\x94\x64\xd2\x18\xc8\x84\x0c\xb2\x9c\xcd\xe7\x5e\xa4\xe5\xa5\xd6\xeb\x20\xbe\xda\x85\xa3\xac\x9f\x29\x3b\xc3\xa2\x63\xb5\x50\x3c\x98\x24\x7e\xa7\x6d\x22\xab\xef\x1f\xfe\xbb\xc7\xf7\x3b\xed\x72\xa3\x4f\x15\xe5\x3c\xb8\xa8\xdf\x56\x4f\x2f\xaf\xab\x45\x6e\x66\x0d\xb4\x89\x76\x19\x70\x22\x70\xca\x7c\x17\x90\x15\x8b\x91\xca\x2a\xa0\x40\x41\x93\xa6\x51\xc8\xaa\x00\x53\xaa\x12\x08\x7f\xe3\xa0\xfb\x7b\x6e\xd3\x29\x4b\xa2\x84\x22\xe0\x80\xb3\x01\x8c\x01\x59\x75\xaa\x4d\x75\x1d\x90\xf5\xa6\x53\xd1\xf7\xa9\x03\x98\x0d\x0b\x94\x83\x77\x19\x3c\x3b\xff\x8e\xdd\x53\x6c\x11\x48\x3b\xca\xb9\x62\x90\xdd\x7e\x86\xa3\x02\xb7\x5a\xb0\x6f\x8d\xc2\x03\xb2\x98\x34\xf4\xa4\x47\xfb\x7d\x6e\xd4\x04\x06\x57\x40\xf8\x56\x13\xfa\xdc\xa8\x09\x5b\xb0\x9e\x28\xde\x6c\xc3\xb1\x77\xfd\x08\x04\xd9\x52\x89\xca\x11\x22\xb8\xf3\x17\x3b\xe9\x04\x86\xa2\x17\x67\xd0\x52\xb3\xbf\x12\x31\x65\xf0\x0d\x5c\x67\x5d\x86\x7e\x54\x3e\x6f\x7b\xb2\xd1\xff\x22\xbe\x08\x1f\x01\x00\x00\xff\xff\xf6\x39\xd9\x2e\x6e\x05\x00\x00")

func pkgHealthLookupStatic_dataInfoYamlBytes() ([]byte, error) {
	return bindataRead(
		_pkgHealthLookupStatic_dataInfoYaml,
		"pkg/health/lookup/static_data/info.yaml",
	)
}

func pkgHealthLookupStatic_dataInfoYaml() (*asset, error) {
	bytes, err := pkgHealthLookupStatic_dataInfoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/health/lookup/static_data/info.yaml", size: 1390, mode: os.FileMode(420), modTime: time.Unix(1602173142, 0)}
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
	"pkg/health/lookup/static_data/info.yaml": pkgHealthLookupStatic_dataInfoYaml,
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
	"pkg": &bintree{nil, map[string]*bintree{
		"health": &bintree{nil, map[string]*bintree{
			"lookup": &bintree{nil, map[string]*bintree{
				"static_data": &bintree{nil, map[string]*bintree{
					"info.yaml": &bintree{pkgHealthLookupStatic_dataInfoYaml, map[string]*bintree{}},
				}},
			}},
		}},
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
