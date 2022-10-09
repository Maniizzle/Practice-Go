package main

import (
	"errors"
	"fmt"

	"github.com/maniizzle/webservice/models"
)

// const (
// 	first  = iota + 6
// 	second = iota
// )

//iota resets in a different constant block
const (
	first = iota + 6
	second
	third
)

const (
	forth = iota
)

func main() {

	loop()
	// controllers.RegisterControllers()
	// http.ListenAndServe(":3200", nil)

}

func startWebServer(port, retries int) error {
	fmt.Println("starting the server", port, retries)
	return errors.New("something happened")
}
func startWebServerWithMultipleReturn(port, retries int) (int, error) {
	fmt.Println("starting the server", port, retries)
	return port, errors.New("something happened")
}

func practise() {
	//fmt.Println("Hello Go")

	//declaring variables
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	//pointer datatype

	// var test *string
	// *test = "TestArthur"
	// fmt.Println(test)
	//error bcuz test is an uninitialized pointer

	//initializing a pointer
	var lastName *string = new(string)
	//dereferencing
	*lastName = "Arthur"
	fmt.Println(lastName) //returns memory address

	testName := "Bose"
	//& address operator
	ptr := &testName
	fmt.Println(ptr, *ptr)
	//result is 0xc000050270 Bose

	testName = "Tricia"
	fmt.Println(ptr, *ptr)
	//result is 0xc000050270 Tricia

	//Constants
	//values must be known at compile time
	//const pi= 3.1415
	//fmt.Println(pi)

	//iota in constants
	fmt.Println(first, second, third)
	fmt.Println(forth)

	//collections
	//Array - fixed size
	// slices
	//maps(key value relationship)
	//Structs

	//Arrays- fixed size collection of similar data types
	arrr := [3]int{1, 2, 3}
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arrr)

	//slice datatype- built on top of an array
	//creating a slice off an array
	slicearr := [3]int{1, 2, 3}
	slice := slicearr[:]
	fmt.Println(slicearr, slice)
	slice[1] = 42
	slice[2] = 43
	fmt.Println(slicearr, slice)

	//creating a slice
	sliceTwo := []int{1, 2, 3}
	fmt.Println(sliceTwo)
	sliceTwo = append(sliceTwo, 4, 42, 27)
	fmt.Println(sliceTwo)

	s3 := sliceTwo[1:]
	//starting from index 0 up to and not including index 2
	s4 := sliceTwo[:2]
	//starting from index 1 up to and not including index 2
	s5 := sliceTwo[1:2]
	fmt.Println(s3, s4, s5)

	//Maps- associate arbitrary keys with the values

	//declaring a map with key type being a string and value being an int
	m := map[string]int{"foo": 42, "fo": 34}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["fo"] = 30
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	//struct
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	fmt.Println(u)

	u.ID = 2
	u.FirstName = "sewa"
	u.LastName = "taiwo"
	fmt.Println(u)

	u2 := user{ID: 3,
		FirstName: "bose",
		LastName:  "good",
	}
	fmt.Println(u2)

	u3 := user{ID: 3, FirstName: "bola", LastName: "bad"}
	fmt.Println(u3)

	us := models.User{
		Id:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(us)

	err := startWebServer(3000, 2)
	fmt.Println(err)

	p, err2 := startWebServerWithMultipleReturn(3000, 2)
	fmt.Println(p, err2)

	_, err3 := startWebServerWithMultipleReturn(3000, 2)
	fmt.Println(err3)

}

func loop() {

	//four types of for loop
	//loop till condition
	// infinite loop
	//loop till condition with post clause
	//loop over collections

	//loop that terminates based on a condition
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			continue
		}
		println("continuing...")
	}
	var j int
	//loop with post clauses
	//for i:=0 ;j < 5; i++ {
	for ; j < 5; i++ {
		println(i)
		if i == 3 {
			continue
		}
		println("continuing...")
	}

	//infinite loop
	var k int
	//for ;;{
	for {
		println(i)
		if k == 3 {
			continue
		}
		k++
		println("continuing...")
	}
}

func loopOverCollections() {

	slice := []int{2, 3, 4}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	//i is the key and v is the value
	for i, v := range slice {
		println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v)
	}

	//print out just the keys
	for k := range wellKnownPorts {
		println(k)
	}

	//print out just the values
	for _, v := range wellKnownPorts {
		println(v)
	}
}

//similar to exceptions in other c type langiuages
func panic() {
	panic()
}

//similar to exceptions in other c type langiuages
func Conditions() {
	u1 := models.User{
		Id:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	u2 := models.User{
		Id:        2,
		FirstName: "Fred",
		LastName:  "Prefect",
	}
	if u1 == u2 {
		println("Same User")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user")
	} else {
		println("Different user")

	}

}

type HTTPRequest struct {
	Method string
}

func Switch() {
	r := HTTPRequest{Method: "GET"}
	switch r.Method {
	case "GET":
		println("Get request")
		fallthrough
	case "DELETE":
		println("delete request")
	case "POST":
		println("post request")
	case "PUT":
		println("put request")
	default:
		println("Unhandled method")

	}
}
