package main

import (
   "fmt"
   "os"
   "hitesh/script_demo"
)

func main () {
   fmt.Println ("Executing git status command")
   
   output, err := script.ExecuteCommand ("bash", "dummy_script.sh")

   if err != nil {
      fmt.Println ("Error: ", err)
      os.Exit(1)
   }

   fmt.Println ("Output :", output)
}
