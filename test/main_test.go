// go test ./...
package main

import (
	"encoding/json"
	"github.com/allochi/jsone"
	"io/ioutil"
	"testing"
)

var content []byte
var root map[string]interface{}

func init() {
	var err error
	content, err = ioutil.ReadFile("sample.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &root)
	if err != nil {
		panic(err)
	}
}

func TestArray(t *testing.T) {
	skills, err := jsone.Dive(root, []interface{}{"skills"})
	if err != nil {
		t.Fatalf("Failed to retrieve the array: %s", err)
	}

	got, ok := skills.([]interface{})
	if !ok {
		t.Fatalf("Failed to retrieve the array: %s", err)
	}

	expected := 3
	if len(got) != expected {
		t.Fatalf("Wrong length of array. got %d expected %d", len(got), expected)
	}
}

func TestIterate(t *testing.T) {
	expectedCategories := []string{"JavaScript", "CouchDB", "Node.js"}
	node, err := jsone.Dive(root, []interface{}{"skills"})
	if err != nil {
		t.Fatalf("Failed to retrieve the array: %s", err)
	}

	for i, v := range node.([]interface{}) {
		c, err := jsone.Dive(v, "category")
		if err != nil {
			t.Fatalf(`Failed to get node @ "category"`)
		}
		i = i
		if c.(string) != expectedCategories[i] {
			t.Fatalf(`Got the wrong value. got "%s" expected "%s"`, c, expectedCategories[i])
		}
	}
}

func TestStringPath(t *testing.T) {
	path := "personal/favorites/color"
	expected := "Blue"
	color, err := jsone.Dive(root, path)
	if err != nil {
		t.Fatalf(`Failed to retrieve node @ "%s"`, path)
	}
	if color.(string) != expected {
		t.Fatalf(`Got the wrong value. got "%s" expected "%s"`, color, expected)
	}

	path = "personal/interests/1"
	expected = "Mountain Biking"
	interest, err := jsone.Dive(root, path)
	if err != nil {
		t.Fatalf(`Failed to retrieve node @ "%s"`, path)
	}
	if interest.(string) != expected {
		t.Fatalf(`Got the wrong value. got "%s" expected "%s"`, interest, expected)
	}

}

func TestInterfaceType(t *testing.T) {
	something := 4
	var fake interface{} = something
	path := "personal/interests/1"
	_, err := jsone.Dive(fake, path)
	if err == nil {
		t.Fatalf(`Expected an error when the wrong type of node is passed`)
	}
}

func TestGetANumber(t *testing.T) {
	path := "age"
	expected := 39.0
	age, err := jsone.Dive(root, path)
	if err != nil {
		t.Fatalf(`Failed to retrieve node @ "%s"`, path)
	}
	if age.(float64) != expected {
		t.Fatalf(`Got the wrong value. got "%f" expected "%f"`, age, expected)
	}
}
