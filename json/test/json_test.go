package json

import (
	"io/ioutil"
	"testing"

	"github.com/andewx/dieselvk/json"
)

func TestMarshal(t *testing.T) {
	var root json.Vlk
	contents, err := ioutil.ReadFile("../vlk_example.json")
	if err != nil {
		t.Errorf("Error opening json file\n")
	}

	json_err := root.UnmarshalJSON(contents)

	if json_err != nil {
		t.Errorf("Error unmarshalling json file\n%s", json_err.Error())
	}

}
