package contentdir

import (
	// "fmt"
	"io/ioutil"
	"os"
	// "reflect"
)

// DirectoryReadWrite is used to implement the ReaderWriter interface
type DirectoryReadWrite struct {
}

// DefaultReaderWriter creates the reader/ writer
func DefaultReaderWriter() *DirectoryReadWrite {
	return &DirectoryReadWrite{}
}

// Stat is proxied  to os.stat
func (rw *DirectoryReadWrite) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }

// ReadDir is proxied to ioutil.ReadDir
func (rw *DirectoryReadWrite) ReadDir(dirname string) (fi []os.FileInfo, err error) {
	// fmt.Printf("!!!!!!!!!!!!!!! %+v\n", dirname)
	return ioutil.ReadDir(dirname)
}

// Mkdir is proxied to os.Mkdir
func (rw *DirectoryReadWrite) Mkdir(path string, perm os.FileMode) error {
	return os.Mkdir(path, perm)
}

// Rmdir is proxied to os.Rmdir
func (rw *DirectoryReadWrite) Rmdir(path string) error {
	return os.Remove(path)
}
