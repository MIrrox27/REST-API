package gorutines

import (
	"fmt"
	"math/rand"
	"sync"
)

func g1(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print(num)

}

func Task1() { // Запусти 10 горутин, каждая выводит своё число от 1 до 10. Используй sync.WaitGroup для ожидания завершения.
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go g1(i, &wg)
	}
	wg.Wait() // ждём завершения всех горутин

}

// --------------

func reeder(n int) {
	fmt.Println(n)

}

func speecer(ch chan<- int) {
	n := rand.Intn(100) + 1 // Intn(100) возвращает 0..99, +1 → 1..100
	ch <- n
}

func Task2() { // Реализуй производитель‑потребитель: один горутина генерирует случайные числа (0‑100) и отправляет их в канал, другая читает и выводит только чётные. Останови после 20 чисел.

	//var wg sync.WaitGroup
	ch := make(chan int)

	for i := 20; i <= 3; i++ {
		go speecer(ch)
		n := <-ch
		reeder(n)
	}
}
