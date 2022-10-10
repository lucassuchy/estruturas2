package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkPrimeNumbers(b *testing.B) {

	rand.Seed(time.Now().Unix())
	array := rand.Perm(10)
	bucketSort(array)

}
