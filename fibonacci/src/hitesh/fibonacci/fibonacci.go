package fibonacci

func Fibonacci() func() int {
   prev1, prev2 := 0, 1
   return func () int {
      to_return := prev1
      prev1, prev2 = prev2, prev1 + prev2
      return to_return
   }
}
