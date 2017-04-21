package main

import (
   "fmt"
   "os"
   "strconv"
   "hitesh/fibonacci"
)

func main () {
   if os.Args == nil || len (os.Args) != 2 {
      printUsage ()
      os.Exit (1)
   }

   num, error := strconv.Atoi(os.Args[1])
   if error != nil {
      fmt.Println (error)
      os.Exit (1)
   }

   fib_func := fibonacci.Fibonacci ()
   for i := 0; i < num; i ++ {
      fmt.Println (fib_func ())
   }
}

func printUsage () {
   fmt.Println ("Usage : main <maximum_number_of_the_series>")
}
