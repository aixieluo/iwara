package untils

import (
	"log"
	"time"
)

func Fast(fn func())  {
	t := time.Now()
	fn()
	endTime:=time.Since(t)
	log.Println(endTime)
}
