package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  inputUser()
}

func inputUser(){
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Enter value between 1 & 9")
  scanner.Scan()

  UserInput, err := strconv.Atoi(scanner.Text())

  fmt.Printf("UserInput : %d , err : %d", UserInput ,err )
  if err != nil {
    fmt.Println("This is not a interger")
    os.Exit(2)
  }else {
    fmt.Println(UserInput)
  }
}
