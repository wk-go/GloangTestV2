package wait_group

import (
	"sync"
	"testing"
)

// The error: WaitGroup is reused before previous Wait has returned
func TestWaitGroup01(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		wg := sync.WaitGroup{}
		for j := 0; j < 20; j++ {
			//wg.Add(1) // correct！！！
			go func(num int) {
				wg.Add(1) // The error: WaitGroup is reused before previous Wait has returned
				defer wg.Done()
				t.Logf("i=%04d,j=%04d;", i, num)
			}(j)
		}
		wg.Wait()
	}
}
