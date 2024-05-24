package altsrc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
)

var (
	ErrCannotReadFile = fmt.Errorf("cannot read file")
)

func readURI(uriString string) ([]byte, error) {
	u, err := url.Parse(uriString)
	if err != nil {
		return nil, err
	}

	if u.Host != "" { // i have a host, now do i support the scheme?
		switch u.Scheme {
		case "http", "https":
			res, err := http.Get(uriString)
			if err != nil {
				return nil, err
			}
			return io.ReadAll(res.Body)
		default:
			return nil, fmt.Errorf("%[1]w: scheme of %[2]q is unsupported", ErrCannotReadFile, uriString)
		}
	} else if u.Path != "" || (runtime.GOOS == "windows" && strings.Contains(u.String(), "\\")) {
		if _, notFoundFileErr := os.Stat(uriString); notFoundFileErr != nil {
			return nil, fmt.Errorf("%[1]w: cannot read from %[2]q because it does not exist", ErrCannotReadFile, uriString)
		}
		return os.ReadFile(uriString)
	}

	return nil, fmt.Errorf("%[1]w: unable to determine how to load from %[2]q", ErrCannotReadFile, uriString)
}
