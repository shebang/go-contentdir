package contentdir

import "sort"

// Filelist is used to maintain a sorted list of paths.
type Filelist struct {
	files []FileInfo
}

// NewFilelist creates a new path list
func NewFilelist(initSize ...int) *Filelist {
	size := 0
	if len(initSize) == 1 {
		size = initSize[0]
	}
	return &Filelist{
		files: make([]FileInfo, size),
	}
}

func (p *Filelist) appendRealloc(data []FileInfo) {
	l := len(p.files)
	if l+len(data) > cap(p.files) { // reallocate
		newSlice := make([]FileInfo, (l+len(data))*2)
		copy(newSlice, p.files)
		p.files = newSlice
	}
	p.files = p.files[0 : l+len(data)]
	copy(p.files[l:], data)
}

// SortByName sorts the file list by name
func (p *Filelist) SortByName() {
	sort.SliceStable(p.files, func(i, j int) bool { return p.files[i].Name() < p.files[j].Name() })
}

// Values return the values as slice
func (p *Filelist) Values() []FileInfo {
	return p.files
}

// Append adds a path to the path list.
func (p *Filelist) Append(file *FileInfo) {
	data := []FileInfo{*file}
	p.appendRealloc(data)
}

// Count returns the number of entries in the path list.
func (p *Filelist) Count() int {
	return len(p.files)
}
