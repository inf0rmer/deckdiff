package loader

import (
	"bytes"
	"net/url"
	"path"
	"strings"

	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/go-fsimpl/filefs"
	"github.com/hairyhenderson/go-fsimpl/httpfs"
)

func Load(u *url.URL) (contents string, err error) {
	p := u.String()

	mux := fsimpl.NewMux()
	mux.Add(filefs.FS)
	mux.Add(httpfs.FS)

	if err != nil {
		return "", err
	}

	fsys, err := mux.Lookup(strings.TrimSuffix(p, path.Base(p)))

	if err != nil {
		return "", err
	}

	file, err := fsys.Open(path.Base(p))

	if err != nil {
		return "", err
	}

	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents = buf.String()

	return
}
