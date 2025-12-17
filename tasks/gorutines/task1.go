package gorutines

import (
	"fmt"
	"sync"
)

func g1(num int) {
	fmt.Println(num)
}

func Task1() { // Запусти 10 горутин, каждая выводит своё число от 1 до 10. Используй sync.WaitGroup для ожидания завершения.
	var wg sync.WaitGroup // созадем элемент WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) //
		go g1(i)

	}
	wg.Wait() // ждём завершения всех горутин

}
