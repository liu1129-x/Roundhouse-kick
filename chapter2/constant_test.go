package chapter2

import "testing"

const (
	Monday  = 1 + iota //初始为0，+1后初始为1
	Tuesday            //2...以此类推
	Wednesday
)

const (
	Readable = 1 << iota //向左位移1位，初始001
	Writable
	Executable
)

func TestConstant1(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstant2(t *testing.T) {
	a := 1 //001 读写 写 读
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
