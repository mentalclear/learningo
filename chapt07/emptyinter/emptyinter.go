package main

import (
	"encoding/json"
	"io/ioutil"
)

func main() {
	// var i interface{}

	// i = 20
	// i = "hello"
	// i = struct {
	// 	FirstName string
	// 	LastName  string
	// }{"Fred", "Fredson"}

	// one set of braces for the interface{} type,
	// the other to instantiate an instance of the map
	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile("testdata/sample.json")
	if err != nil {
		return
	}
	//defer contents.Close()
	json.Unmarshal(contents, &data)
	// the contents are now in the data map
}

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}
