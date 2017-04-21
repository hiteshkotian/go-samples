package main

import (
   "fmt"
   "sync"
)

/**
 * Create instance of wait group.
 */
var wg = sync.WaitGroup {}

func printString (str string, i int) {
   defer wg.Done ()
   // for i := 0; i < 5; i ++ {
      fmt.Println ( str, " : ", i )
   // }
}

func create_go_routine (str string) {
   // wg.Add (1)
   for i := 0; i < 5; i ++ {
      wg.Add (1)
      go printString (str, i)
   }
}

func main () {

   create_go_routine ("megadeth")
   create_go_routine ("metallica")
   create_go_routine ("meshuggah")
   create_go_routine ("tool")
   
   wg.Wait()
   fmt.Println ( "Finished!!" )
}
