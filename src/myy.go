package main

import (
	"fmt"
	"math"

	"runtime"
	"strconv"
	"time"
	"unsafe"
)


func main() {
//	fmt.Println("hello world!"+Int64ToString(88));

	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)
	//fmt.Println(math.MaxInt64)
	var a uint64 = StringToUint64("284467440737095")
	//myAtoUi(str string)
	fmt.Println("StringToUint64", a, runtime.Version())

	 var b uint64 = myAtoUi("284467440737095")

	 fmt.Println("myAtoUi", b, runtime.Version())

	 var c string=Int64ToString(9223372036854775807)
	fmt.Println("Int64ToString", c, runtime.Version())
	 var d string=myInt64ToString(9223372036854775807)
    fmt.Println("myInt64ToString", d)
//	var c int64 = myAtoi("18446744073709551615")
	//myAtoUi(str string)
//	fmt.Println("myAtoi", c, runtime.Version())
	//测试test_func的执行时间
	start := time.Now().Unix()
	//test_func_StringToUint64()
	end := time.Now().Unix()
	fmt.Println("test_func_StringToUint64 执行消耗的时间为:%v秒", end - start)

	//测试test_func的执行时间
	start= time.Now().Unix()
	//test_func_myAtoi()
	end= time.Now().Unix()
	fmt.Println("test_func_myAtoi 执行消耗的时间为:%v秒", end - start)

	//测试test_func的执行时间
	start= time.Now().Unix()
//	test_func_myAtoUi()
	end= time.Now().Unix()
	fmt.Println("test_func_myAtoUi 执行消耗的时间为:%v秒", end - start)

}
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func StringToUint64(s string) uint64 {
	ret, _ := strconv.ParseUint(s, 10, 64)
	return ret
}

const digits = "0123456789"
const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"
func myInt64ToString(n int64) (string){
	//s := []byte("01234567890123456789")

	//	var i int=19
	var a [19+ 1]byte // +1 for sign of 64bit value in base 2
//	i := len(a)
	i := 20
//	fmt.Println("len", i)
	var neg bool=false
	if(n<0){
		n=-n;
		neg=true
	}
	/*
		for {
			 s[i]= byte(n%10+'0');
	 		i--
			n/=10
			if 0==n {
				break
			}
		}*/
	/*
		for n >= 10 {

			  			// Avoid using r = a%b in addition to q = a/b
			  			// since 64bit division and modulo operations
			 			// are calculated by runtime functions on 32bit machines.
			  			q :=n / 10
			 			s[i] = digits[uint(n-q*10)]
		  		     	n = q
			           i--
			  		}
		   		// u < base

		  		s[i] = digits[uint(n)]
		         i--*/
	us := uint(n)
	for us >= 100 {
		is := us % 100 * 2
		us /= 100
		i -= 2
		a[i+1] = smallsString[is+1]
		a[i+0] = smallsString[is+0]
	}

	// us < 100
	is := us * 2
	i--
	a[i] = smallsString[is+1]
	if us >= 10 {
		i--
		a[i] = smallsString[is]
	}
	if neg {
		i--
		a[i] = '-'
	}
	/*
		if neg{
			s[i]='-'
			i--
		}
		c:=s[i+1:]
		 return *(*string)(unsafe.Pointer(&c))*/
	//return c
	c:=a[i:]
	return *(*string)(unsafe.Pointer(&c))
	//return string(a[i:])


}
	func myAtoUi(str string) uint64 {
	var i int
	var num uint64
	var strl int
	strl= len(str)

	var flag bool
	flag=false
	for ; i < strl; i++ {
		if  str[i] < 48 || str[i] > 57 {
			num=0;
			flag=true
			break;
		}
		num=num*10+ uint64(str[i] - '0')
	}
	if strl>=20&&!flag&&str>"18446744073709551615"{//字符串中没有不合法字符且数字越界（超过最大值）
		return  math.MaxUint64;
	}
	return num
}