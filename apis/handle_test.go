package apis

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRCToken(t *testing.T) {
	resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
