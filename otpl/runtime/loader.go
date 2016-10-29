package runtime

import (
	"encoding/binary"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"

	"github.com/junhwong/mano/otpl/common"
	_ "github.com/junhwong/mano/otpl/opc"
)

var bufferLen = 256

var buffer = make([]byte, bufferLen)
var byteOrder binary.ByteOrder = binary.BigEndian

func Open(reader io.ReadCloser) (common.Loader, error) {
	//reader = bufio.NewReaderSize(reader, bufferLen)
	l := &loader{
		reader:  reader,
		codes:   make(map[common.Ptr]common.Opcode, 0),
		blocks:  make(map[string]common.Opcode, 0),
		bodyPtr: common.ZeroPtr,
	}
	err := l.readHeader()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func newErr(prefix, format string, v ...interface{}) error {
	return fmt.Errorf((prefix + ":" + format), v...)
}

type loader struct {
	reader     io.ReadCloser
	modified   int64
	template   string
	startPtr   common.Ptr
	codes      map[common.Ptr]common.Opcode
	blocks     map[string]common.Opcode
	bodyLoader common.Loader
	bodyPtr    common.Ptr
}

func (l *loader) readHeader() (err error) {

	buf, err := l.ReadBytes(12)
	if err != nil {
		return
	}

	//检查名称
	name := string(buf[0:4])
	if name != "OTIL" {
		//TODO: 错误
	}
	//检查版本
	if version := byteOrder.Uint16(buf[4:]); version != 0x02 {
		//TODO: 错误
		panic(version)
	}

	//buf[6] 编码(保留，默认统一UTF8)
	//buf[7:9] 保留3位

	l.startPtr = common.Ptr(byteOrder.Uint16(buf[10:]))

	l.modified, err = l.ReadLong()
	if err != nil {
		return
	}
	// panic("hhhhhhhhhhhhhhhhhh")
	l.template, err = l.ReadString()
	if err != nil {
		return
	}
	//panic(fmt.Sprintf("%v", l.modified))
	code, err := l.Load(l.startPtr)
	if err != nil {
		return
	}
	if code == nil {
		//TODO:未找到开始节点
	}
	return nil
}

func getCallMethod() string {
	method := ""
	pc, _, _, ok := runtime.Caller(4)
	if ok {
		name := runtime.FuncForPC(pc).Name()
		index := strings.LastIndex(name, ".")
		if index > 0 {
			method = "loader." + name[index+1:]
		} else {
			method = name
		}
	}
	return method
}

func (l *loader) read(data interface{}) error {

	err := binary.Read(l.reader, byteOrder, data)
	if err != nil {
		return newErr(getCallMethod(), "%v", err)
	}

	// _, err = l.reader.Seek(int64(binary.Size(data)), os.SEEK_CUR)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (l *loader) ReadBytes(count int) ([]byte, error) {

	buf := buffer
	if count > bufferLen {
		buf = make([]byte, count)
	}
	buf = buf[:count]
	//buf = make([]byte, count)
	//err := l.read(&buf)

	c, err := l.reader.Read(buf)
	if err == io.EOF {
		return nil, newErr("loader.ReadBytes", "Read data error:EOF")
	}
	if err != nil {
		return nil, err
	}
	// _, err = l.reader.Seek(int64(c), os.SEEK_CUR)
	// if err != nil {
	// 	return nil, err
	// }
	if c != count {
		return nil, newErr(getCallMethod(), "Read data error:%v/%d", c, count)
	}

	return buf, err
}

func (l *loader) ReadByte() (byte, error) {

	// buf, err := l.ReadBytes(1)
	// if err != nil {
	// 	return 0x00, nil
	// }
	// return buf[0], nil

	var r byte
	err := l.read(&r)
	return r, err
}

func (l *loader) ReadBool() (b bool, err error) {
	var x uint8
	err = l.read(&x)
	b = x != 0x00
	return
	// r, err := l.ReadByte()
	// if err != nil {
	// 	return false, err
	// }
	// return r == 0x00, err
}

func (l *loader) ReadShort() (int16, error) {
	var r int16
	err := l.read(&r)
	return int16(r), err
}

func (l *loader) ReadInt() (int, error) {
	var r int32
	err := l.read(&r)
	return int(r), err
}

func (l *loader) ReadPtr() (common.Ptr, error) {
	var r uint16
	err := l.read(&r)
	return common.Ptr(r), err
}

func (l *loader) ReadLong() (int64, error) {
	s, err := l.ReadString()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(s, 10, 64)
}

func (l *loader) ReadFloat() (float64, error) {
	s, err := l.ReadString()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(s, 64)
}

func (l *loader) ReadString() (string, error) {
	// strlen, err := l.ReadInt()
	// if err != nil {
	// 	return "", err
	// }
	var strlen uint16
	err := l.read(&strlen)
	if err != nil {
		return "", err
	}
	// panic(strlen)
	buf, err := l.ReadBytes(int(strlen))
	if err != nil {
		return "", err
	}
	return string(buf), err
}

func (l *loader) Load(ptr common.Ptr) (common.Opcode, error) {

	for {
		code, ok := l.codes[ptr]
		if ok {
			// panic(fmt.Sprintf("%v", ok))
			return code, nil
		}

		buf, err := l.ReadBytes(6)
		if err != nil {
			return nil, err
		}

		nextPtr := common.Ptr(byteOrder.Uint16(buf[0:]))

		typ := buf[2]
		line := common.LineNo(byteOrder.Uint16(buf[3:]))
		flag := buf[5]

		if nextPtr == common.ZeroPtr {
			//panic(newErr("loader.Load", "Invalid ptr;%d", nextPtr))
			return nil, newErr("loader.Load", "Invalid ptr;%d", len(l.codes))
		}

		code, err = common.NewOpcode(l, typ, nextPtr, line, flag)
		if err != nil {
			return nil, err
		}

		l.codes[nextPtr] = code

	}
}

func (l *loader) TemplateName() string {
	return l.template
}

func (l *loader) StartPtr() common.Ptr {
	return l.startPtr
}

func (l *loader) PutBlock(blockID string, block common.Opcode) {
	l.blocks[blockID] = block
}

func (l *loader) GetBlock(blockID string) (block common.Opcode, ok bool) {
	block, ok = l.blocks[blockID]
	return
}

func (l *loader) Blocks() map[string]common.Opcode {
	return l.blocks
}

func (l *loader) SetBody(body common.Loader, start common.Ptr) {
	l.bodyLoader = body
	l.bodyPtr = start

	bodyBlocks := body.Blocks()
	for id, block := range l.Blocks() {
		bodyBlocks[id] = block
	}
	l.blocks = bodyBlocks

}

func (l *loader) BodyLoader() (common.Loader, bool) {
	return l.bodyLoader, l.bodyLoader != nil
}

func (l *loader) BodyPtr() common.Ptr {
	return l.bodyPtr
}

func (l *loader) Close() (err error) {
	if l.reader != nil {
		err = l.reader.Close()
	}
	l.reader = nil
	l.blocks = nil
	l.codes = nil
	l.bodyLoader = nil

	return
}
