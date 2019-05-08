package mytime

import (
	"log"
	"time"
)

const (
	LocationName = "Asia/Tokyo"
)

var (
	loc     *time.Location
	fixTime *time.Time
)

func init() {
	var err error
	loc, err = time.LoadLocation(LocationName)
	if err != nil {
		log.Panicln(err)
	}
	time.Local = loc
}

func Now() time.Time {
	if fixTime != nil {
		return *fixTime
	}
	return time.Now().Truncate(time.Second)
}
