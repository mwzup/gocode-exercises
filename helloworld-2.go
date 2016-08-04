package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    //ORIG CODE | ReadString() requires an argument
    /*reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    str, _ := reader.ReadString('\n') //adds new line to string
    fmt.Printf("\nHello %v, welcome to my world!",str)*/

//-------------------------------------------------------

   fmt.Print("\nEnter your name: ")
   scan := bufio.NewScanner(os.Stdin)
   scan.Scan()   //fmt.Println(scan.Text())
   
   str := scan.Text()
   fmt.Printf("Hello %v, welcome to my world!\n",str)
}