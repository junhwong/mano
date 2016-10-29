package runtime

import (
	"os"
	"testing"

	"github.com/junhwong/mano/otpl/common"
	_ "github.com/junhwong/mano/otpl/opc"
)

func TestReads(t *testing.T) {
	file, err := os.Open("F:\\workspace\\clamp-projects\\exam-web\\bin\\otil\\4146ec82a0f0a638db9293a0c2039e6b.otil")
	if err != nil {
		t.Fatal(err)
	}

	l := &loader{
		reader: file,
		codes:  make(map[common.Ptr]common.Opcode, 0),
	}

	err = l.readHeader()
	if err != nil {
		t.Fatal(err)
	}
	//t.Fatal(l.startPtr)
	code, err := l.Load(l.startPtr)
	if err != nil {
		t.Fatal(err)
	}
	// buf, err := l.ReadBytes(8)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if buf != nil {

	// }

	// mt, err := l.ReadLong()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	t.Fatalf("结果:%v", code)
}

func xTestLoad(t *testing.T) {
	file, err := os.Open("F:\\workspace\\clamp-projects\\exam-web\\bin\\otil\\4e193ff69f4ef3daf559de82cb312ca7.otil")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	l, err := Open(file)

	if err != nil {
		t.Fatal(err)
	}

	str, err := l.ReadString()
	if err != nil {
		t.Fatal(err)
	}

	t.Fatalf("结果:%v", str)

}
