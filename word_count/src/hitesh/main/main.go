package main

import (
   "fmt"
   "os"
   "hitesh/wordcount"
   "strings"
)

func main () {
   if os.Args == nil || len(os.Args) < 2 {
      printUsage ()
      os.Exit (1)
   }
   
   str := strings.Join (os.Args[1:], " ")
   fmt.Println ("String passed in is ", str)
   count := wordcount.WordCount(str)
   fmt.Println (count)
}

func printUsage () {
   fmt.Println ("main [sentence to be inspected by word count]")
}
