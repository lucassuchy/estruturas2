package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Função que retorna lista array ordenado, pego em
// https://github.com/benoitvallon/go-programs/blob/master/insertion-sort.go
func insertion(array []int) []int {
	lenght := len(array)
	iterations := 0

	for i := 1; i < lenght; i++ {
		j := i
		for j > 0 && array[j-1] > array[j] {
			iterations++
			array[j-1], array[j] = array[j], array[j-1]
			j = j - 1
		}
	}
	return array
}

//Bucket sort com base em
//https://gostudent.github.io/Letsgo/implementation-of-bucket-sort-in-GO
func bucket_sort(array []int) []int {
	// Desenho o tamanho do array
	// Dados por bucket = total de digitos -2
	lenght := len(array)
	contador := 0
	for lenght != 0 {
		lenght /= 10
		contador += 1
	}
	total_buckets := contador - 2

	var max, min int

	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	var wg sync.WaitGroup
	nBuckets := int(max-min)/total_buckets + 1
	buckets := make([][]int, nBuckets)
	wg.Add(len(buckets))
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range array {
		idx := int(n-min) / total_buckets
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertion(bucket)
			defer wg.Done()
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func main() {
	// Gera uma lista aleatoria
	rand.Seed(time.Now().Unix())
	array := rand.Perm(1000)
	fmt.Println("Lista não sorteada:")
	fmt.Println(array)
	fmt.Println("Lista sorteada:")
	fmt.Println(bucket_sort(array))

}
