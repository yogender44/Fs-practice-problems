package main

import (
    "fmt"
    "sync"
    "math/rand"
)


type ConcurrentQueue struct {
    queue []int32
    mu sync.Mutex
}

func(q *ConcurrentQueue) Enqueue(item int32) {
    q.mu.Lock()
    defer q.mu.Unlock() // whatever instruction is executed is done after the function exits.
    q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Dequeue() int32 {
    q.mu.Lock()
    defer q.mu.Unlock()
    if len(q.queue) == 0 {
        panic("Cannot dequeue from  an empty queue")
    }
    
    item := q.queue[0]
    q.queue = q.queue[1:]
    return item
}

func (q *ConcurrentQueue) Size() int {
    q.mu.Lock()
    defer q.mu.Unlock()
    return len(q.queue)
}

var wgE sync.WaitGroup
var wgD sync.WaitGroup

func main() {
    q1 := ConcurrentQueue{
        queue : make([] int32, 0),
    }

    for i:=0; i<1000000; i++ {
        wgE.Add(1)
        go func() { // anonymout function
            q1.Enqueue(rand.Int31()) // generate random values
            wgE.Done()
        }()
    }    



    for i:=0; i<1000000; i++ {
        wgD.Add(1)
        go func() { // anonymout function
            q1.Dequeue() // generate random values
            wgD.Done()
        }()
    }  

        wgE.Wait()
    wgD.Wait()

        fmt.Println(q1.Size())


    // fmt.Println(q1.Dequeue())
    //     fmt.Println(q1.Dequeue())
    // fmt.Println(q1.Dequeue())
    // fmt.Println(q1.Dequeue())

}
