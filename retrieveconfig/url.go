package retrieveconfig

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// FetchURLConfigFile will retrieve a remote configuration file and write it to local disk
func FetchURLConfigFile(configPath string) ([]byte, error) {
	resp, err := http.Get(configPath)

	if err != nil {
		return []byte(""), errors.Wrapf(err, "Could not retrieve remote configuration file from url %s.", configPath)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
