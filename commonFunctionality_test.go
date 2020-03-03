package rctlr

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gookit/color"
)

func TestGenerateRandomNumbers(t *testing.T) {
	color.Info.Println(`
	TestGenerateRandomNumbers: Go utiliza math/rand.Intn(x) para generar numeros aleatorios
	desde 0 hasta x. Estos numeros se repetiran frecuentemente a menos que se establesca una semilla
	antes de generar los numeors. 
	`)
	rand.Seed(time.Now().Unix())
	numbers := make([]int, 20)
	for i := 0; i < 20; i++ {
		random := rand.Intn(1000)
		if find(numbers, random) > 0 {
			color.Error.Println("\tNumeros generados no son suficientemente aleatorios")
			t.FailNow()
		}
		numbers[i] = random
	}
}

func TestFindDuplicates(t *testing.T) {
	color.Error.Println("TestFindDuplicates: implementar esto")
	t.Fail()
}

func find(numbers []int, target int) int {
	for i, num := range numbers {
		if num == target {
			return i
		}
	}
	return -1
}