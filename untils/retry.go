package untils

import (
	"time"
)

func Retry(attempts int, sleep time.Duration, fn func() error) error {
	err := fn()
	if err == nil {
		return nil
	}

	if attempts--; attempts > 0 {
		time.Sleep(sleep)
		return fn()
	}

	return err
}
