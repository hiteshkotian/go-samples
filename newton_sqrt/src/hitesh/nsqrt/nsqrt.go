package nsqrt

import "math"

/**
 * Function to find sqrt based on newton's square root approximation where
 * a number 'z' is a square root of 'x' if z = z - ( ( ( z ^ 2 ) - x ) / ( 2 * z ) )
 */
func Sqrt (x float64) float64 {
   
   // Find the whole number sqrt.
   wholeSqrt := newtonSqrtFunc ( x, 0.0, x / 2.0, 1.0 )

   // Find one decimal point sqrt.
   return newtonSqrtFunc ( x, wholeSqrt, wholeSqrt + 1, 0.01 )
}

func newtonSqrtFunc ( x, start, maxBound, increment float64 ) float64 {
   diff, min := math.MaxFloat64, 0.0
   
   for z := start; z <= maxBound; z += increment {
      // fmt.Println ( "----------------------------------------------" )
      // fmt.Println ( "Diff = ", diff, " and min ", min )
      value := computeVal (z, x)
      newDiff := math.Abs ( value -z )
      if diff > newDiff {
         diff = newDiff
         min = z
      }
      // fmt.Println ("Value for", z, "= ", value)
      // fmt.Println ( "----------------------------------------------" )
   }
   return min
}

func computeVal (z, x float64) float64 {
   if z == 0 {
      return math.MaxFloat64
   }
   value := z - ( ( math.Pow ( z, 2 ) - x ) / ( 2 * z ) )
   return value;
}
