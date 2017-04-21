package main

import (
   "hitesh/nsqrt"
   "fmt"
   "os"
   "strconv"
)

func main () {
   
   // Read the command line arguments.
   if os.Args == nil || len(os.Args) != 2 {
      printUsage()
      os.Exit (1)
   }
   
   numString := os.Args[1]  
   fmt.Println ( "Number added is ", numString )
   num, err := strconv.ParseFloat ( numString, 64 )
   if err != nil {
      fmt.Println (err)
      os.Exit (1)
   }

   val := nsqrt.Sqrt ( num )
   fmt.Println ( "sqrt (", num, ") = ", val )
}

func printUsage () {
   fmt.Println ( "Usage ./main [number whose sqrt is to be computed]" )
}
