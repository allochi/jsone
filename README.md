# Jsone

Jsone is a function that can be used to dive into `map[string]interface{}` that represent a json structure.
It's simple to use and straight forward. Not yet fully tested but it will be, and it worked so far.

## Getting started

1. Download

```
go get github.com/allochi/jsone
```

2. Import

```
import "github.com/allochi/jsone"
```

## Usage

Jsone currently has only one public function `jsone.Dive()`.

Basically: 

1. use `json.Unmarshal()` to decode a json into either `map[string]interface{}` or `[]interface{}`.
2. use a `string` or a `[]interface{}` as the path of the node to retrieve.

```go
var content []byte
var root map[string]interface{}
var err error

content, err = ioutil.ReadFile("sample.json")
err = json.Unmarshal(content, &root)

skills, err := jsone.Dive(root, []interface{}{"skills"})
// or
skills, err := jsone.Dive(root, "skills")

"personal/favorites/color"

color, err := jsone.Dive(root, []interface{}{"personal", "favorites", "color"})
// or
color, err := jsone.Dive(root, "personal/favorites/color")

interest, err := jsone.Dive(root, []interface{}{"personal", "interests", 1})
// or
interest, err := jsone.Dive(root, "personal/interests/1")

name, err := jsone.Dive(skills, "tests/0/name")
name, err := jsone.Dive(skills, "tests/1/name")

```
