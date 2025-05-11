package main
import {
  "fmt"
}

var count int = 0;

func incCount() {
  count++;  
}

func doCount() {
  for i :=0; i< 1000000; i++ {
    incCount()
  }

}


func main() {

count  = 0';
doCount();
}
