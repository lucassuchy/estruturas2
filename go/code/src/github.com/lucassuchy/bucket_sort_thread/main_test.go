package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkBucketSort(b *testing.B) {

	rand.Seed(time.Now().Unix())
	array := rand.Perm(50000000)
	bucket_sort(array)

}
