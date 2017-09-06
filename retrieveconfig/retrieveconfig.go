package retrieveconfig

import (
	"io/ioutil"
	"net/url"
	"os"

	"github.com/pkg/errors"
)

// RetrieveConfig will take a given path to config and attempt to load it locally, and write it to a place
// that it can be read by the rest of the silo agent.
func RetrieveConfig(path string) (composeBytes []byte, err error) {

	if isLocalFile(path) {
		return ioutil.ReadFile(path)
	}

	if isURL(path) {
		composeBytes, err = FetchURLConfigFile(path)

		if err != nil {
			return composeBytes, errors.Wrap(err, "Could not retrieve configuration file")
		}

		return
	}

	return composeBytes, nil
}

// Check if given path is a URL that we can retrieve config from
func isURL(path string) bool {
	// Check if this is a URL
	url, err := url.Parse(path)

	if err != nil {
		return false
	}

	if url.Scheme != "" && url.Host != "" {
		return true
	}

	return false
}

// isLocalFile Check if the given path is an OS path
func isLocalFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
