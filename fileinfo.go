package contentdir

import (
	"os"
	"time"
)

// FileInfo represents a file in a directory. It is the same struct as
// os.FileInfo and implements the same interface.
type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	sys     interface{}
	isdir   bool

	DirInfo *FileInfo
}

// NewFile creates a new file. It is used to simplify the interface when only
// names are used.
func NewFile(name string, args ...interface{}) *FileInfo {
	return &FileInfo{
		name: name,
	}

}

// NewFromOs creates a new file from os.FileInfo.
func NewFromOs(fi os.FileInfo, parent ...*FileInfo) *FileInfo {
	var di *FileInfo = nil

	if len(parent) == 1 {
		di = parent[0]
	}
	return &FileInfo{
		name:    fi.Name(),
		size:    fi.Size(),
		mode:    fi.Mode(),
		modTime: fi.ModTime(),
		isdir:   fi.IsDir(),
		DirInfo: di,
	}
}

// Name returns the name of the file
func (fs *FileInfo) Name() string { return fs.name }

// Size returns the size of the file
func (fs *FileInfo) Size() int64 { return fs.size }

// Mode returns the FileMode of the file
func (fs *FileInfo) Mode() os.FileMode { return fs.mode }

// ModTime returns the modification time of the file
func (fs *FileInfo) ModTime() time.Time { return fs.modTime }

// Sys is not used but here to satisfy the interface
func (fs *FileInfo) Sys() interface{} { return nil }

// IsDir returns true if the file is a directory.
func (fs *FileInfo) IsDir() bool { return fs.isdir }
