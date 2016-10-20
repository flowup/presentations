package backend

import (
	"fmt"; "io/ioutil"
	"net/http"; "net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomeBackendHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.Equal(t, nil, err)

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.Equal(t, nil, err)
	assert.Equal(t, greeting, "Hello, client")
}
