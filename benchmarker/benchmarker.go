package benchmarker

import (
	"context"
	"errors"
	"net/http"

	"github.com/getoutreach/benchmarker/lib/benchmarker"
)

func BenchmarkerMain(_ context.Context, f *benchmarker.Options) error {
	res, err := http.DefaultClient.Get(f.Addr)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return errors.New(http.StatusText(res.StatusCode))
	}

	return nil
}
