package test

import (
	"fmt"
	"os"
	"testing"

	xzmap "github.com/averyyan/xz-map/map"
)

var myMap *xzmap.Map[string, int]

func TestMain(m *testing.M) {
	myMap = xzmap.New[int]()
	for i := 0; i < 100; i++ {
		myMap.Set(fmt.Sprintf("%d", i), i)
	}
	os.Exit(m.Run())
}
