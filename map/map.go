package main

import "fmt"

func main(){
  users := map[string]string{}

  users["Kurt"] = "kurt@gmail.com"
  users["Janis"] = "janis@gmail.com"

  fmt.Println("Map length: ",len(users))
}

