package main

import (
  "fmt"
)

func stringSort(str string) string {
  return str;
}

func main() {
  // length: from 1 to 10^7
  // occurences: 'a' - 'z' characters
  var raw = "afduhldoskfnfjdkfkjaanvkjdksklaasdrhjlqonhsvgasbhfjdkanancjkndjkkslaweognjdfjplndjbqbajsbdpaznaajdsjoqkojfskjdbvfhrlsak"
	
  fmt.Println(raw)
	
  var sorted = stringSort(raw)	
  fmt.Println(sorted)
}

