package main

 import (
    "fmt"
    "bufio"
    "os"
    "strconv"
 )

 func isleapyear(year int64) bool {
     if year % 400 == 0 {
        return true
     } else if year % 100 == 0 {         
        return false
     } else if year % 4 == 0 {        
        return true
     } else {
        return false 
     }
 }

 func main() {

    fmt.Print("Enter a year: ")
    scan := bufio.NewScanner(os.Stdin)
    scan.Scan()
    year := scan.Text()
    intyear, _ := strconv.ParseInt(year, 0,64)
    yes := isleapyear(intyear)

    if yes {
        fmt.Printf("%v is a leap year.", year)
    } else {
        fmt.Printf("%v is not a leap year.", year)
    }
    //fmt.Println(year,yes)
 }