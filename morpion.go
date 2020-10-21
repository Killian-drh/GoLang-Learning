package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)


func main() {
  userInput := inputUser()
  grid(userInput)

}

func grid(value int) {

  var table = [3][3]int{ {1,2,3}, {4,5,6}, {7,8,9}}

  var i, j int

  /* output each array element's value */
  for  i = 0; i < 3; i++ {
     for j = 0; j < 3; j++ {
       fmt.Printf("%d ",table[i][j] )
       if j == 2{
        fmt.Printf("\n")
        }
     }
  }
}



// function input number in the grid
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
