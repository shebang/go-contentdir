package contentdir

import (
	"os"
	"path"
)

// FilterCallback defines the function signature for filtering file names.
type FilterCallback func(os.FileInfo) bool

// FilterByExtension returns trur if a filename's extensions matches with an
// extension in extlist.
func FilterByExtension(fname string, extlist []string) bool {
	ext := path.Ext(fname)
	for _, v := range extlist {
		if v == ext {
			return true
		}
	}
	return false
}
