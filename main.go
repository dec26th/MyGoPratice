package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

type T struct {
	Name string
}

func (t *T) Print() {
	fmt.Println(t.Name)
	wg.Done()
}

func main() {

	data := []T{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}}

	//{
	//	for i := 0; i < len(data); i++ {
	//		wg.Add(1)
	//		go data[i].Print()
	//	}
	//}
	{
		for _, v := range data {
			wg.Add(1)
			go v.Print()
		}
	}

	wg.Wait()
}
