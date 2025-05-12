package main
import (
  "fmt"
  "sync"
)

var mu sync.Mutex
var wg sync.WaitGroup

var count int = 0;

func incCount() {
  mu.Lock();
  count++;
  mu.Unlock();
  wg.Done()  
}

func doCount() {
  for i :=0; i< 1000000; i++ {
    wg.Add(1)
    incCount()
  }
}


func main() {
  count  = 0; 
  go doCount();
  wg.Wait();
}
