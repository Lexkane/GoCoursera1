package main

import (
	"bytes"
	"testing"
)

type User struct {
	ID       int
	RealName string
	Login    string
	Flags    int
}
type Avatar struct {
	ID  int
	Url string
}

func BenchmarkCodegen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := &User{}
		u.Unpack(data)
	}
}
func BenchmarkReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := &User{}
		UnpackReflect(u, data)
	}
}

func (in *User) Unpack(data []byte) error {
	r := bytes.NewReader(data)
	//ID
	var IDRaw uint32
	binary.Read(r, binary.LittleEndian, &IDRaw)
	in.ID = int(IDRaw)

	//Login
	var LoginLenRaw uint32
	binary.Read(r, binary.LittleEndian, &LoginLenRaw)
	LoginRaw := make([]byte, LoginLenRaw)
	binary.Read(r, binary.LittleEndian, &LoginRaw)
	in.Login = string(LoginRaw)

	//Flags
	var FlagsRaw uint32
	binary.Read(r, binary.LittleEndian, &FlagsRaw)
	in.Flags = int(FlagsRaw)
	return nil
}

//go test -bench .unpack_test.go
//go test -bench . -benchmem -cpuprofile=cpu.put -memprofile=mem.out -memprofilerate=1 unpack_test.go
//go tool pprof main.test.exe mem.out
//list unpack
//web
//top
//alloc_space
//top
////go tool pprof main.test.exe  cpu.out
