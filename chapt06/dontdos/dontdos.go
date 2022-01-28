package main

func main() {

}

// Don't
// func MakeFoo(f *Foo) error {
// 	f.Field1 = "val"
// 	f.Field2 = 20
// 	return nil
//   }

// Do:
//   func MakeFoo() (Foo, error) {
// 	f := Foo{
// 	  Field1: "val",
// 	  Field2: 20,
// 	}
// 	return f, nil
//   }

// Resonable use of pointers

// f := struct {
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// }
// err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)
