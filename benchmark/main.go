package benchmark

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/getoutreach/benchmarker/lib/benchmarker"
)

var counter uint32

func BenchmarkerMain(f *benchmarker.Options) {
	res, err := http.DefaultClient.Get(f.Addr)
	if err != nil {
		fmt.Println("error", err)
	}
	atomic.AddUint32(&counter, 1)

	fmt.Printf("success:\n\tcount: %d\n\tcode: %d\n---\n", counter, res.StatusCode)
}
