// Package texttmpl provides convenience utilities for using text templates in
// an embedded filesystem.
package texttmpl

import (
	"path/filepath"
	"text/template"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/xio"
	"github.com/richardwilkes/toolbox/xio/fs/embedded"
)

// Load the templates found at the path, omitting any that the filter function
// returns true for. The filter function may be nil, in which case all files
// are loaded. The filter function is not called for the initial path.
func Load(tmpl *template.Template, fs embedded.FileSystem, path string, filter func(path string, isDir bool) bool) (*template.Template, error) {
	dir, err := fs.Open(path)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer xio.CloseIgnoringErrors(dir)
	fi, err := dir.Stat()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	if fi.IsDir() {
		fis, derr := dir.Readdir(-1)
		if derr != nil {
			return nil, errs.Wrap(derr)
		}
		for _, fi := range fis {
			onePath := filepath.Join(path, fi.Name())
			isDir := fi.IsDir()
			if filter == nil || !filter(onePath, isDir) {
				if isDir {
					if _, err = Load(tmpl, fs, onePath, filter); err != nil {
						return nil, err
					}
				} else {
					if err = load(tmpl, fs, onePath); err != nil {
						return nil, err
					}
				}
			}
		}
	} else if err = load(tmpl, fs, path); err != nil {
		return nil, err
	}
	return tmpl, nil
}

func load(tmpl *template.Template, fs embedded.FileSystem, path string) error {
	str, ok := fs.ContentAsString(path)
	if !ok {
		return errs.New("Unable to read " + path)
	}
	if _, err := tmpl.New(path).Parse(str); err != nil {
		return errs.NewWithCause("Unable to parse "+path, err)
	}
	return nil
}
