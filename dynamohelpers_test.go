package lambdadynamohelpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var table string;

func TestMain(m *testing.M) {
	f, err := ioutil.ReadFile("test_conf.json")
	if err == nil {
		var properties map[string]map[string]string
		json.Unmarshal(f, &properties)
		println(properties["region"]["value"])
		os.Setenv(REGION, properties["region"]["value"])
		table = properties["table_name"]["value"]
	} else {
		log.Println("Error reading conf file", err)
	}
	os.Exit(m.Run())
}

type WithKey struct {
	Key string `json:"key"`
}
func TestPutAndGet(t *testing.T) {
	if err := PutItem(&WithKey{"foo"}, table); err != nil {
		t.Error(err)
	}
	var withKey WithKey
	GetItemS("key", "foo", table, &withKey)
	if withKey.Key != "foo" {
		t.Error("Key returned was incorrect")
	}
}
