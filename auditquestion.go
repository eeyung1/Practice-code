package main

import (
  "fmt"
   "time"
   )
  
  func doSomething() string {
    fmt.Println("Here's the current time")
    
    now := time.Now()
    
    fmt.Println("Date included", now)
    
    
    return ""
  }
  
  func main(){
    fmt.Println("About to do something that does not exist")
    doSomething()
  }
