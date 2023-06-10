package test

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestNew(t *testing.T) {
	v, ok := myMap.Get("1")
	then.AssertThat(t, ok, is.True())
	then.AssertThat(t, v.GetValue(), is.EqualTo(1))
}

func TestHas(t *testing.T) {
	then.AssertThat(t, myMap.Has("1"), is.True())
}

func TestIterBuffered(t *testing.T) {
	counter := 0
	for range myMap.IterBuffered() {
		counter += 1
	}
	then.AssertThat(t, counter, is.EqualTo(100))
}
