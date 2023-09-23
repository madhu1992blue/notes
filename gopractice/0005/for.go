package main

import "fmt"

func main() {
  // Simulating a while loop
  i := 1
  for i <= 3 {  // Loop with condition only.
    fmt.Println(i)
    i = i + 1
  }


  // classic for loop : initialization, condition, after
  for j:=1; j <= 10; j++ { //Scope of j is within this block only.
    fmt.Println(j)
  }
  
  // Infinite for loop: no condition specified (condition evaluates to true. 'break' or 'return' is required to exit this)
  for  {
    fmt.Println("loop")
    break // Break out of for loop.
  }
  
  // We can continue to next iteration with 'continue'
  for n := 0; n <= 5; n++ { //Scope of n is only within this block.
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}
