package main

import (
	"fmt"
	"time"

	"github.com/mix3/buildtag-sample/mytime"
)

func main() {
	fmt.Println(mytime.Now())
	mytime.FixTime(time.Time{})
	fmt.Println(mytime.Now())
	mytime.Reset()
	fmt.Println(mytime.Now())
}
