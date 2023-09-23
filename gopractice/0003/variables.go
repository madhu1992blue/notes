package main

import "fmt"

func main() {
  // var keyword is used for declaration.
  // var <variableName> <type> = <Val>

  // a's type is declared as float64 and set to 40.
  var a float64 = 40.0

  // b's type is inferred as String by looking at value.
  // Hence type could be skipped.
  var b = "Hello"

  // c and d are declared as integers and set to 20 and 30 respectively.
  var c, d int = 20, 30

  // e is not initialized and hence will be set to zero value of its type, in this case, 0.
  var e int

  // f is declared and initialized. It's type is set to inferred type(In this case int) and value set to 10.
  // 'f:=10' is equivalent to 'var f int = 10'.
  // This syntax is valid only within functions.
  f:=10
  fmt.Println(a,b,c,d,e,f)
}
