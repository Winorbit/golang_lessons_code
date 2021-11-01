package unit_tests_3

import (
	"io/ioutil"
	"net/http"
)

type API struct {
	Client  *http.Client
	url string
}

func (api *API) DoStuff() ([]byte, error) {
	resp, err := api.Client.Get(api.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// handling error and doing stuff with body that needs to be unit tested
	return body, err
}


// func main(){

// api := Api{}

// }
// // http://hassansin.github.io/Unit-Testing-http-client-in-Go