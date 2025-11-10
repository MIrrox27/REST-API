package stats

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// AvgAge возвращает средний возраст (avg) и ошибку, если срез пуст.
func AvgAge(people []Person) (float64, error) {
	// TODO: реализовать
}
