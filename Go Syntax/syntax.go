package main

// This is the code to be run in terminal

// import some packages to use
import (
	// "fmt" for print out to terminal

	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// create a fmt alias
var pl = fmt.Println

// For all functions use func() keyword
func main() {
	// pl("what is your name?")
	// reader := bufio.NewReader(os.Stdin)
	// //  store this input
	// name, err := reader.ReadString('\n')
	// // if don't want to store error
	// // name, _ := reader.ReadString('\n')
	// if err == nil {
	// 	pl("Hello", name)
	// } else {
	// 	log.Fatal(err)
	// }

	// // naming convention: Camal case
	// var vName string = "David Guo"
	// var v1, v2 = 1.2, 3.4
	// var v3 = "hello"
	// v4 := "hel"
	// v5 := 12
	// v6 := 5.4

	// DataTypes
	// 			int float64 boo string rune
	// default   0   0.0	false  ""

	// reflect package check the type
	pl(reflect.TypeOf(25))

	// Casting
	cV1 := 1.4
	cV2 := int(cV1)
	pl(cV1, cV2)

	// convert string to int
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	// convert int to string
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	pl(cV5, cV6, reflect.TypeOf(cV6))

	// convert string to float
	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	// convert float to string
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9, reflect.TypeOf(cV9))

	// conditional logic and ops
	iAge := 22
	if iAge >= 1 && iAge < 18 {
		pl("Important Birthday")
	} else if iAge == 21 || iAge == 50 {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("not an important birthday")
	}
	pl("!true = ", !true, "true = ", true)

	// String
	// arrays of bytes
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV1, replacer, sV2)
	pl("Length: ", len(sV2))
	pl("contains another: ", strings.Contains(sV2, "Another"))
	pl("o index:", strings.Index(sV2, "o"))
	// replace (the last arg is for occurance)
	pl("Replace: ", strings.Replace(sV2, "o", "0", 0))
	pl("Replace: ", strings.Replace(sV2, "o", "0", 1))
	pl("Replace: ", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome Words\n"
	pl(sV3)
	sV3 = strings.TrimSpace(sV3)
	pl(sV3)
	sV4 := strings.Split("a-b-c-d", "-")
	pl("Split: ", sV4, reflect.TypeOf(sV4))
	pl("Lower: ", strings.ToLower((sV2)))
	pl("Upper: ", strings.ToUpper(sV2))
	// return bool val
	pl("Prefix: ", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix: ", strings.HasSuffix("tacocat", "cat"))

	// Runes == Characters (unicode)
	rStr := "abcdefg"
	pl("Rune Count: ", utf8.RuneCount([]byte(rStr)))
	pl("Rune Count: ", utf8.RuneCountInString(rStr))

	// loop
	for i, runeVal := range rStr {
		// use printf have a new line
		// print unicode -> %#U
		// print int -> %d
		// print character -> %c
		fmt.Printf("%d: %#U : %c\n", i, runeVal, runeVal)
	}

	// How to track cur time
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	// Math func
	pl("5 + 4: ", 5+4)
	pl("5 - 4: ", 5-4)
	pl("5 * 4: ", 5*4)
	pl("5 / 4: ", 5/4)
	pl("5 % 4: ", 5%4)

	mInt := 1
	mInt += 1
	mInt++
	pl(mInt)

	// Understand precision when dealing with float
	pl("Float Precision: ", 0.111111111+0.22222222222)

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	// gen random val
	// Intn return non-negative pseudo random number
	randNum := rand.Intn(50) + 1
	pl("Random: ", randNum)

	pl("Abs(-10) = ", math.Abs(-10))
	pl("Pow(4, 2) = ", math.Pow(4, 2))

	// convert degree to radians
	r90 := 90 * math.Pi / 180
	// convert rad to degree
	d90 := r90 * (180 / math.Pi)
	pl(r90, d90)
	fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)

	// %d integer
	// %c character
	// %f float
	// %t bool
	// %s string
	// %o base8
	// %x base16
	// %v Guesses based on data type
	// %T type of supplied val
	fmt.Printf("%s, %d, %c, %f, %t, %o, %x", "Stuff", 1, 'c', 2.12, false, 1, 1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.14)
	fmt.Printf("%9f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)
	// sprintf format a string and return it
	sp1 := fmt.Sprintf("%9.f", 3.141592)
	pl(sp1)

	// for loop
	// for init; cond;
	// postStatement {body}

	// x is out of scope main, only exist inside for loop
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	for x := 5; x >= 1; x-- {
		pl(x)
	}
	//  use for to create while loop
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	// // while true loop to do a number guess
	// seedSecs1 := time.Now().Unix()
	// rand.Seed(seedSecs1)
	// randNum1 := rand.Intn(50) + 1

	// for true {
	// 	fmt.Print("Guess a num between 0 and 50: ")
	// 	pl("Rand num is: ", randNum1)
	// 	reader := bufio.NewReader(os.Stdin)
	// 	guess, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	guess = strings.TrimSpace(guess)
	// 	iGuess, err := strconv.Atoi(guess)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if iGuess > randNum {
	// 		pl("Pick a Lower val")
	// 	} else if iGuess < randNum {
	// 		pl("Pick a Higher val")
	// 	} else {
	// 		pl("You've got it!")
	// 		// call break to jump out of loop
	// 		break
	// 	}
	// }

	// cycle thru array using range
	aNums := []int{1, 2, 3}
	for i, num := range aNums {
		pl(i, ":", num)
	}

	// create empty array with default val 0
	var arr1 [5]int
	pl(arr1)
	arr1[0] = 1
	arr1[2] = 2
	pl(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("Index 0: ", arr2[0])
	pl("Arr Len: ", len(arr2))
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("%d: %d\n", i, v)
	}

	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune array: %d\n", v)
	}

	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	bStr1 := string(byteArr[0:2])
	bStr2 := string(byteArr[1:2])
	bStr3 := string(byteArr[1:3])
	pl("I am a string: ", bStr, bStr1, bStr2, bStr3)

	// create array of string -> array of char arrays
	sl1 := make([]string, 6)
	pl(sl1)

	// slicing
	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl(sl3)
	pl("1st 3: ", sArr[:3])
	pl("last 3: ", sArr[2:])

	sl3 = append(sl3, 13)
	pl("now sl3:", sl3)

	// functions
	// func funcName(parameter) returnType {BODY}
	// could return more return val
	pl(getTwo(5))
	pl(getQuotient(5, 4))
	pl(getSum2(1, 2, 3, 4, 5))
	vArr := []int{1, 2, 3, 4, 5, 6}
	pl("Array Sum: ", getArraySum(vArr))
	f3 := 5
	pl("f3 before function: ", f3)
	changeVal(f3)
	pl("f3 after function: ", f3)

	// pointers
	f6 := 10
	pl("f3 before function: ", f6)
	changeVal2(&f6)
	pl("f6 after function: ", f6)

	f4 := 10
	var f4Ptr *int = &f4
	pl("f4 addr: ", f4Ptr)
	pl("f4 val: ", *f4Ptr)
	*f4Ptr++
	pl("f4 val now:", *f4Ptr)

	// passing array pointers
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)

	// create avg
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average: %.3f\n", getAverage(iSlice...))

	// File IO
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// as soon as we leave main func, the file closed.
	defer f.Close()
	// create a list of prime numbers
	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	// cycle thru strings
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		// catch errors
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read from file and print it one by one
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime: ", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	_, err1 := os.Stat("data.txt")
	if errors.Is(err1, os.ErrNotExist) {
		pl("File Doesnt Exist")
	} else {
		f, err := os.OpenFile(
			"data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}

	// Map data structure
	// var myMap map[keyType]valueType
	var heroes map[string]string
	heroes = make(map[string]string)

	villians := make(map[string]string)

	heroes["batman"] = "bruce wayne"
	heroes["superman"] = "clark kent"
	heroes["the flash"] = "barry allen"
	villians["joker"] = "joker"
	villians["bad guy"] = "badguy"

	superPets := map[int]string{1: "krypto", 2: "Bat hound"}
	fmt.Printf("Batman is %v\n", heroes["batman"])
	pl("Chip: ", superPets[3])
	_, ok := superPets[3]
	pl("Is there a 3rd per: ", ok)

	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}
	for k, v := range villians {
		fmt.Printf("%s is %s\n", k, v)
	}

	delete(villians, "joker")
	delete(heroes, "batman")

	// Generics
	// for funcs that could work for multi-datatypes
	// generic parameter -> Capitalized

	pl("5 + 4: ", getSumGen(5, 4))
	pl("5.6 + 3.7: ", getSumGen(5.6, 3.7))

	// Struct
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 Main St"
	tS.bal = 234.56

	getCustInfo(tS)
	newCustAdd(&tS, "123 South St")

	pl("Now the Address: ", tS.address)

	// struct literal
	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name: ", sS.name)

	rect1 := rectangle{10.0, 15.0}
	pl("Rectangle area: ", rect1.Area())

	// nested struct
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}
	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}
	bus1.info()

	// Interfaces
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()
	kitty.HappySound()

	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	kitty2.AngrySound()
	kitty2.HappySound()
	pl("Cats Name: ", kitty2.Name())

	// Concurrency
	go printTo10()
	go printTo15()
	// You cannot trust what order you've run
	time.Sleep(2 * time.Second)

	// Channels? Could do concurrency in order
	// only could carry one type
	channel1 := make(chan int)
	channel2 := make(chan int)

	go nums1(channel1)
	go nums2(channel2)

	for i := 0; i < 3; i++ {
		pl(<-channel1)
	}
	for i := 0; i < 3; i++ {
		pl(<-channel2)
	}

	// Mutex/Lock Simulate
	var acct Account
	acct.balance = 100
	pl("Balance : ", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)

}

// Mutex/Lock
type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enough money in account")
	} else {
		fmt.Printf("%d withdrawn : Balance : %d\n",
			v, a.balance)
		a.balance -= v
	}
}

// channels
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

// Concurrency
func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("Func 1: ", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("Func 2: ", i)
	}
}

// Interfaces
type Animal interface {
	AngrySound()
	HappySound()
}
type Cat string

func (c Cat) Attack() {
	pl("Cat attacks its prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hisssssss")
}
func (c Cat) HappySound() {
	pl("Cat says purrrrrrr")
}

// struct3 -> nested important (Composition)
type contact struct {
	fName string
	lName string
	phone string
}
type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s",
		b.name,
		b.contact.fName,
		b.contact.lName)
}

// struct1
type customer struct {
	name    string
	address string
	bal     float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

// struct2
type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

// Generics
// for funcs that could work for multi-datatypes
// generic parameter -> Capitalized
type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	// receive either int or float and return same datatype
	return x + y
}

//
//
//

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You cannot divide by 0")
	} else {
		return x / y, nil
	}
}

// you could pass in multiple int number parameters
func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func changeVal2(myPtr *int) int {
	*myPtr = 12
	return *myPtr
}

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val

	}
	return (sum / numSize)
}
