package main

import ( 
	"fmt" 
	"sync"	
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var data = []string{"lesgo", "lesgoo", "lesgooo"}
var result = []string{}

func main(){
	t0 := time.Now()
	for i:= 0; i<len(data);i++{
		wg.Add(1)
		go dataCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal Execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are: %v", result)
}

func dataCall(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(data[i])
	log()
	wg.Done()
}

func save(results string){
	m.Lock()
	result = append(result, results)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("\nThe current results are: %v", result)
	m.RUnlock()
}