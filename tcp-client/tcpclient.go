package main

import (
  "encoding/gob"
  "fmt"
  "net"
)


func client(){
  // connect to the server
  c,err :=net.Dial("tcp","127.0.0.1:9999")
  if err !=nil{
    fmt.Println(err)
    return
  }

  //send the message
  msg := "Hello,World!"
  fmt.Println("Sending",msg)
  err = gob.NewEncoder(c).Encode(msg)
  if err != nil{
    fmt.Println(err)
  }

  c.Close()

}

func main(){
  go client()

  fmt.Println("-----finish go client() 1------")

  var input string
  fmt.Scanln(&input)

  fmt.Println("-----finish go client() 2------")

  
}
