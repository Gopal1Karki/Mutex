package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	money int
)

func init() {
	money = 2500
}

func Ram(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Ram is giving me %d from  %d  money i had\n", value, money)
	money += value
	mutex.Unlock()
	wg.Done()
}

func Shyam(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Shyam is taking from me %d: %d\n", value, money)
	money -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Mutex Example:")
	var wg sync.WaitGroup
	wg.Add(2)
	go Ram(50, &wg)
	go Shyam(800, &wg)
	wg.Wait()

	fmt.Println("The total money i have now is ", money)
}
