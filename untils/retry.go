package untils

import (
	"log"
	"time"
)

func Retry(attempts int, sleep time.Duration, fn func() error) error {
	err := fn()
	if err == nil {
		return nil
	}


	if attempts--; attempts > 0 {
		log.Printf("调用失败，重新排队等待%d重试...", sleep)
		time.Sleep(sleep)
		return Retry(attempts, sleep, fn)
	}

	return err
}
