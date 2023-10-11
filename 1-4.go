// modify dup2 print names of files for all dupes
package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  counts := make(map[string][]string)
  files := os.Args[1:]
  if len(files) == 0 {
    countLines(os.Stdin, counts)
  } else {
    for _, arg := range files {
      f, err := os.Open(arg)
      if err != nil {
        fmt.Fprintf(os.Stderr, "modDup2: %v\n", err)
        continue
      }
      countLines(f, counts)
      f.Close()
    }
  }
  for line, filenames := range counts {
    if len(filenames) > 1 {
      for _, arg := range filenames {
        fmt.Printf("%s\t%s\n", arg, line)
      }
    }
  }
}
func countLines(f *os.File, counts map[string][]string) {
  input := bufio.NewScanner(f)
  for input.Scan() {
    counts[input.Text()] = append(counts[input.Text()], f.Name())
  }
  //TODO handle input.Err()
}
