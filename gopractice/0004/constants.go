package main 

// We can import multiple packages with shorthand syntax
import (
  "math"
  "fmt"
)

// s is a package scoped constant.
const s string = "constant"

func main() {
  // Go supports constants of character, string, boolean, and numeric values only.

  // const declares a constant value.


  // Numeric constant has no type until its given one such as by explicit conversion.

  // n is a numeric constant with no type specified.
  const n = 5000000

  // the constant expressions perform arithmetic with arbitrary precision.
  const d = 3e20 / n
  fmt.Println(d)
  
  // Convert the constant value to a speicific numeric type.
  fmt.Println(int64(d))

  // n will be used as float64 automatically since math.Sin expects a float64
  fmt.Println(math.Sin(n))

  
}
