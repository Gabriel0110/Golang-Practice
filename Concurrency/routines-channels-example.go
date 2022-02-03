// The one thing we don't have here is channels communicating with other channels...

package concurreny

import (
  "fmt"
  "time"
)

const CONCURRENCY_LEVEL = 3
var requestIds = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

func main () {
  // Create channel
  queue := make(chan bool, CONCURRENCY_LEVEL)
  
  for _, _ID := range requestIds {
    queue <- true
    go func(ID int) {
      defer func() { <- queue }()
      makeRequest(ID)
    }(_ID)
  }
  
  // Flush the remaining queue
  for i := 0; i < CONCURRENCY_LEVEL; i++ {
    queue <- true 
  }
}

func makeRequest(ID int) {
  time.Sleep(time.Second)
  fmt.Println(ID)
}
