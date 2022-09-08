package main

import (
	"fmt"
	"time"
	"reflect"
)

func main(){
	fmt.Printf("Sum: %d\n",sumNum(2,3))
	credential()
	sleep()
	fn(nil)
	var e *MyError
	check(e)
	check(nil)
}

func sumNum(num1 int, num2 int) int {
	return num1+num2;
}

//to test an issue
func credential() {
	username := "admin"
    var pwd = "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"
    fmt.Println("Doing something with: ", username, pwd)
}

//to test the client-side tool DeadCode
func subNum(num1 int, num2 int) int {
	return num2-num1;
}

//to test the client-side tool StaticCheck - SA1004
func sleep() {
    time.Sleep(8)
    fmt.Println("Sleep Over.....")
}

func fn(x *int) {
    fmt.Println(*x)

    // This nil check is equally important for the previous dereference
    if x != nil {
        fmt.Println("ERROR")
    }
}

type MyError struct {
	Err error
}

func (me MyError) Error() string {
	return me.Err.Error()
}

func check(err error) {
	fmt.Printf("type: %v value: %v\n", reflect.TypeOf(err), reflect.ValueOf(err))
}