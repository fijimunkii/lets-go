// simple echo program that also prints name of command that invoked it
package main

import (
  "fmt"
  "os"
  "strings"
)
func main() {
  fmt.Println(os.Args[0])
  fmt.Println(strings.Join(os.Args[1:], " "))
}
