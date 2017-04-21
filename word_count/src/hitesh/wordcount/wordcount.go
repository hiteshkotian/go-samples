package wordcount

import (
   "fmt"
   "strings"
)

/**
 * Function returning a map of words and it's count in the given
 * string
 */ 
func WordCount( s string ) map[string] int {
   fmt.Println ( "Processing string ", s )
   
   var count_map = make(map[string] int)

   string_slices := getStringSlice (s)
   fmt.Println ( "The slice for the string is :", string_slices )
   
   // Iterate through the array thus created and add add it to the 
   // map
   for _, word := range string_slices {
      count, _ := count_map[word]
      count_map[word] = count + 1
   }
   return count_map
}

/**
 * Returns all the words present in the given string.
 */
func getStringSlice ( s string ) []string {
   // Split the string based on the spaces used in the string
   string_slices := strings.Fields (s)
   return string_slices
}
