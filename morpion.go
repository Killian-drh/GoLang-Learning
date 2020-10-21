package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  userInput := inputUser()
  fmt.Println(userInput)
}

func inputUser()int{
  scanner := bufio.NewScanner(os.Stdin)
  var input int
  var err error

  for true {
    fmt.Println("Enter value between 1 & 9")
    scanner.Scan()
    input, err = strconv.Atoi(scanner.Text())

    if err != nil {
      fmt.Println("This is not a interger")
      continue
    }
    if input > 9 || input < 0 {
      fmt.Println("Your number must be between 0 & 9")
      continue
    }
    if input < 9 || input > 0{
      break
    }
  }
  return input
}
