package contentdir

import (
	"fmt"
	"os"
	"path"
)

// contentdir provides directory organization by a defined content type. For
// instance: If you want to store scheme files which define colors, you will
// be provided with the following directory structure:
// * schemes/default
// * schemes/material
// * schemes/solarized
//
// To organize files by using this structure you need to take care of three
// basic parts of a file path: content type (schemes), user directory (material)
// and the final directory entries (i.e: *.yaml files or whatever you store in
// those folders)
//

// DirectoryCreator defines an interface for creating directories.
type DirectoryCreator interface {
	Mkdir(path string, perm os.FileMode) error
}

// DirectoryDeleter defines an interface for deleting directories.
type DirectoryDeleter interface {
	Rmdir(path string) error
}

// FileStatReader defines an interface for testing the state of files and directories.
type FileStatReader interface {
	Stat(path string) (os.FileInfo, error)
}

// Reader defines an interface for reading a directory
type Reader interface {
	ReadDir(string) ([]os.FileInfo, error)
}

// ReaderWriter defines an interface for reading and writing a directory
type ReaderWriter interface {
	Reader
	FileStatReader
	DirectoryCreator
}

// Filter is used to filter files when querying a contentdir
type Filter struct {
}

// Options defines flags for modifying the behaviour of NewDirectory.
type Options struct {
	// readData enables reading and caching file data
	readData bool
}

// Directory defines the properties for a contentdir
type Directory struct {

	// PathInfo is used to store the root and content path. It also implements
	// the interface DirInfo.
	PathInfo PathInfo

	// contentTag defines a string for naming the content
	contentTag string

	// filter is used to filter files
	filter *Filter

	// rw is the ReaderWriter interface for reading and writing files and dirs
	rw ReaderWriter

	// perm defines the permission bits for creating dirs (before umask)
	perm os.FileMode

	// filterFunc is called to filter files when reading directories.
	filterFunc FilterCallback
}

// NewDirectory creates a new directory
func NewDirectory(rootPath string, contentTag string, args ...interface{}) (dir *Directory, err error) {
	var rw ReaderWriter = nil
	// var filter *Filter = nil
	var perm os.FileMode = 0700
	// var opts *Options = nil

	for _, v := range args {
		switch s := v.(type) {
		case Filter:
			// filter = s
		case ReaderWriter:
			rw = s
			break
		case Options:
			// opts = &s
			break
		}
	}

	if rw == nil {
		rw = DefaultReaderWriter()
	}

	if _, errPath := rw.Stat(rootPath); os.IsNotExist(errPath) {
		err = fmt.Errorf("root path '%s' does not exist (%+v)", rootPath, errPath)
		return
	}

	contentDir := path.Join(rootPath, contentTag)
	if _, errPath := rw.Stat(contentDir); os.IsNotExist(errPath) {
		err = rw.Mkdir(contentDir, perm)
		if err != nil {
			return
		}
	}

	dir = &Directory{
		PathInfo:   NewPathInfo(rootPath, contentTag),
		contentTag: contentTag,
		rw:         rw,
		perm:       perm,
		filterFunc: nil,
	}
	return
}

// ReadDir reads all file entries into memory
func (dir *Directory) ReadDir() (err error) {
	var files []os.FileInfo
	var userDir, userFile *FileInfo

	userdirs, err := dir.rw.ReadDir(dir.PathInfo.ContentPath())
	if err != nil {
		return
	}

	for _, d := range userdirs {
		userDir = NewFromOs(d)
		files, err = dir.rw.ReadDir(path.Join(dir.PathInfo.ContentPath(), userDir.Name()))
		for _, f := range files {
			userFile = NewFromOs(f, userDir)
			if dir.filterFunc != nil {
				if dir.filterFunc(f) {
					dir.PathInfo.AddFile(userDir, userFile)
				}
			} else {
				dir.PathInfo.AddFile(userDir, userFile)
			}
		}
	}
	return
}
