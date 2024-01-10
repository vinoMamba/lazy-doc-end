package id

import (
	"fmt"
	"testing"
)

func TestGenSnowflakeID(t *testing.T) {
	Init("2024-01-01 00:00:00", 1)
	id := GenSnowflakeID()
	fmt.Println(id)
	id2 := GenSnowflakeID()
	fmt.Println(id2)
	id3 := GenSnowflakeID()
	fmt.Println(id3)
}
