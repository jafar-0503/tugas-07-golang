package main

import "fmt"
import "runtime"


func main(){
  runtime.GOMAXPROCS(2)


  var nilai = 10

  go tampil_pesan(5, nilai)
  tampil_pesan(5, nilai)


  var input int
  fmt.Scanln(&input)
}

func tampil_pesan(x int, pesan int){
  for i:=0;i<x;i++ {
    fmt.Println((i +1), pesan)
  }
  }
