package otpl

import (
	"crypto/md5"
	"testing"
)

func TestRender(t *testing.T) {
	// ip := New()
	// ip.Render()

	data := []byte("123456")
	//fmt.Printf("%x", md5.Sum(data))

	t.Fatalf("%x", md5.Sum(data))
}
