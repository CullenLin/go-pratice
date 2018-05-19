package main

import "fmt"
import "encoding/json"

type Person struct {
  Name      string
  Age       int
  Class     string
  Gender    string
  Interests []string
}

func main() {
  // Testing json.Marshal function
  jacky := Person{"jacky", 20, "A", "girl", []string{"sport","music"} }
  jackyStr, _ := json.Marshal(jacky)
  fmt.Println(string(jackyStr))

  // Testing json.Unmarshal function
  jacky2 := Person {}
  json.Unmarshal(jackyStr, &jacky2)
  fmt.Println("name:", jacky2.Name, ", age:", jacky2.Age)
}
