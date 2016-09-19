package benchmethodset

import (
	"io"
	"testing"
)

const N = 100000000

var s MyStruct

func foo(s MyStruct) {

}

func bar(r io.Reader) {

}

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < N; i++ {
		foo(s)
	}
}

func BenchmarkInteface(b *testing.B) {
	for i := 0; i < N; i++ {
		bar(s)
	}
}

type MyStruct struct{}

func (t MyStruct) Method1() {

}

//
//func (t MyStruct) Method2() {
//
//}
//
//func (t MyStruct) Method3() {
//
//}
//
//func (t MyStruct) Method4() {
//
//}
//
//func (t MyStruct) Method5() {
//
//}
//
//func (t MyStruct) Method6() {
//
//}
//
//func (t MyStruct) Method7() {
//
//}
//
//func (t MyStruct) Method8() {
//
//}
//
//func (t MyStruct) Method9() {
//
//}
//
//func (t MyStruct) Method10() {
//
//}
//
//func (t MyStruct) Method11() {
//
//}
//
//func (t MyStruct) Method12() {
//
//}
//
//func (t MyStruct) Method13() {
//
//}
//
//func (t MyStruct) Method14() {
//
//}
//
//func (t MyStruct) Method15() {
//
//}
//
//func (t MyStruct) Method16() {
//
//}
//
//func (t MyStruct) Method17() {
//
//}
//
//func (t MyStruct) Method18() {
//
//}
//
//func (t MyStruct) Method19() {
//
//}
//
//func (t MyStruct) Method20() {
//
//}

func (t MyStruct) Read(b []byte) (int, error) {
	return 0, nil
}
