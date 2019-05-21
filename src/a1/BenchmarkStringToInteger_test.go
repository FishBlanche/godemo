package a1

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
var testarr [2000] string
var testintarr [2000] int64
func init() {
	fmt.Println("init. package is pk1")
	for i := 0; i <2000; i++ {
		testarr[i]=GetRandomString(12);
		a:=RandInt(0, 9223372036854775807)
		if i%2==0 {
			testintarr[i]= a
		} else{
			testintarr[i]= -a
		}

	}
	fmt.Println("init. package is pk1  over/////")
}
func RandInt(min, max int) int64 {

	return int64(rand.Intn(max-min) + min)
}
func  GetRandomString(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func BenchmarkStringToUint64(t *testing.B) {
	// var st string
	for i := 0; i < t.N; i++ {
		// st=GetRandomString(8)
		//	fmt.Println(st)
		StringToUint64(testarr[i%2000])
		//	StringToUint64("12888888888888899")
	}
}

func BenchmarkMyAtoUi(t *testing.B) {

	for i := 0; i < t.N; i++ {
		 myAtoUi(testarr[i%2000])
	//	myAtoUi("12888888888888899")
	}

}
func BenchmarkInt64ToString(t *testing.B) {

	for i := 0; i < t.N; i++ {
		Int64ToString(testintarr[i%2000])
		//	myAtoUi("12888888888888899")
	}

}

func BenchmarkMyInt64ToString(t *testing.B) {

	for i := 0; i < t.N; i++ {
		myInt64ToString(testintarr[i%2000])
		//	myAtoUi("12888888888888899")
	}

}
/*
func TestMyAtoUi(t *testing.T) {
	area := myAtoUi("18446744073709551616")
	if area != 18446744073709551616 {
		t.Error("测试失败")
	}
}
func TestStringToUint64(t *testing.T) {
	area := StringToUint64("18446744073709551616")
	if area != 18446744073709551616 {
		t.Error("测试失败")
	}
}*/