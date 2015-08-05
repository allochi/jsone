package jsone

import (
	"fmt"
	"strconv"
	"strings"
)

// Dive accepts either `map[string]interface{}` or `[]interface{}` as the root object
// and uses a `string` or a `[]interface{}` as the path to the node.
func Dive(node interface{}, path interface{}) (interface{}, error) {
	keys, err := breakdownPath(path)
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		node, err = read(node, key)
		if err != nil {
			return nil, err
		}
	}
	return node, nil
}

func read(node interface{}, key interface{}) (interface{}, error) {
	switch node.(type) {
	case []interface{}:
		idx, ok := key.(int)
		if !ok {
			return nil, fmt.Errorf("Index is not an integer")
		}
		array := node.([]interface{})
		if idx >= len(array) {
			return nil, fmt.Errorf("Index out of bound")
		} else {
			node = array[idx]
		}
	case map[string]interface{}:
		key, ok := key.(string)
		if !ok {
			return nil, fmt.Errorf("Key is not a string")
		}
		node = node.(map[string]interface{})[key]
	default:
		return nil, fmt.Errorf("Node can only be of types map[string]interface{} or []interface{}")
	}

	if node == nil {
		return nil, fmt.Errorf("Couldn't find the node")
	}
	return node, nil
}

func breakdownPath(path interface{}) ([]interface{}, error) {
	var keys []interface{}
	switch path.(type) {
	case string:
		names := strings.Split(path.(string), "/")
		keys = make([]interface{}, len(names))
		for i, v := range names {
			if n, err := strconv.Atoi(v); err == nil {
				keys[i] = n
			} else {
				keys[i] = v
			}
		}
	case []interface{}:
		keys = path.([]interface{})
	default:
		return nil, fmt.Errorf("Path can only be of type string of []interface{}")
	}
	return keys, nil
}
