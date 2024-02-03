package main

import "testing"

func TestSquence(t *testing.T) {
	for i := 0; i <= 10000000; i++ {
		result := squence(i)
		if result != (i * i) {
			// show all error test and continue to end test
			t.Errorf("excepted %d , but result %d", i*i, result)

			// show 1 error and stoped in that itration
			// t.Fatalf("excepted %d , but result %d",i*i,result)
		}
	}
}
