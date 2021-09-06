package main

import ( // 괄호로 import 그룹 => "factored" import
	"fmt"
	//"math"
	//"math/cmplx"
	//"runtime"
	//"time"
)

/*
// 패키지와 변수, 함수
var c, python, java bool //var문은 package나 함수 단에 존재할 수 있다

var (
	ToBe bool = false
	// unit64의 범위 : 0~18446744073709551615
	MaxInt uint64     = 1<<64 - 1 // 1을 빼주지 않으면 오버플로우가 발생한다.
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 인자가 없는 return문은 이름이 주어진 반환 값을 반환한다.
}
func swap(x, y string) (string, string) { // 반환값이 2개이기 때문에 이걸 받는 변수도 2개여야한다.
	return y, x
}

// func 함수명(매개변수) 반환형
func add(x, y int) int { //func add(x int, y int) int {
	return x + y
}

func main() {
	const Pi = 3.14
	fmt.Println("Happy", Pi, "Day")

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println(c, python, java)

	fmt.Println(split(17))

	a, b := swap("hello", "world") // 반환값이 2개라서
	fmt.Println(a, b)

	fmt.Println(add(42, 13))

	fmt.Println(math.Pi) // 대문자로 시작해야 접근할 수 있다.
}
*/

/*
// 흐름 제어 구문: for, if, else, switch 그리고 defer
func sqrt(x float64) string {
	if x < 0 { // {}괄호가 항상 필수
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x)) //Sprint()함수는 값을 문자열로 만든다.
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// main 함수 안에 있는 Println문보다 먼저 결과를 반환한다.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func deferStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// defer로 연기되어서 스택에 쌓였다. 그래서 후입선출 순서로 수행된다.
		// 9 8 7 6 5 4 3 2 1 0
		defer fmt.Printf("%d ", i)
	}

	fmt.Println("done")
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // {}괄호가 항상 필수
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 < 1000 { // Go의 "while"
		sum1 += sum1
	}
	fmt.Println(sum1)

	// for{} //무한 루프

	fmt.Printf("%T, %T \n", sqrt(2), sqrt(-4))
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(Sqrt(2))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // 운영체제 확인
	// 자동으로 break제공
	case "darwin":
		fmt.Println("Os X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}

	fmt.Print("When's Saturday?\t")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	t := time.Now()
	switch { //switch true{} 와 동일. (조건문 생략 가능)
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer문은 자신을 둘러싼 함수가 종료할 때까지 실행을 연기한다.
	defer fmt.Println("world") // main함수가 종료할 때까지 실행연기
	fmt.Print("hello ")

	deferStack()
}
*/

// 더 많은 타입들: struct와 slice, map
type Vertex struct { // 구조체
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}  // X:1 / Y:0
	v3 = Vertex{}      // X:0 / Y:0
	p  = &Vertex{1, 2} // type: *main.Vertex
)

func slicesAreLikeReferencesToArrays() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo", //Enter를 쳤기때문에, 마지막 원소 끝에 ,(쉼표)가 붙는다.
		//한 줄로 쓰면, 마지막 원소 끝에 ,(쉼표)가 붙지않는다.
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	names[0] = "Jeong" //기본 배열이 바뀌면 슬라이스도 바뀌게 된다.
	b[0] = "XXX"       //슬라이스의 요소를 변경하면 기본 배열의 해당 요소가 수정된다 => 기본배열의 한 영역을 나타내기 때문에
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s[1:4])
	fmt.Println(s[:2])
	fmt.Println(s[1:])
}

func sliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	// 여기서 s=s[:1]를 하고 s=s[2:]를 하게된다면 오류가 발생한다.
	//
	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	// 슬라이스의 길이(_length) : 슬라이스가 포함하는 요소의 개수
	// 슬라이스의 용량(_capcity) : 슬라이스의 첫 번째 요소부터 계산하는 기본 배열의 요소의 개수
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var p1 *int
	fmt.Println(p1 == nil)

	j := 2701
	p2 := &j // j 간접참조
	*p2 = *p2 / 37
	fmt.Println(j)
	fmt.Printf("%T\n", p2)

	v := Vertex{1, 2}
	p3 := &v //구조체 포인터를 통해서 구조체 필드에 접근 => 번거롭다.
	p3.X = 1e9
	fmt.Println(v)
	v.X = 4 //Vertex구조체의 X필드에 접근하기
	fmt.Println(v)

	fmt.Println(v1, v2, v3, p)
	fmt.Printf("%T\n", p)

	var a [2]string //배열 선언
	a[0], a[1] = "Hello", "World"
	primes := [6]int{2, 3, 5, 7, 11, 13} //배열 선언
	fmt.Println(a[0], a[1])
	fmt.Println(a, primes)

	var s []int = primes[1:4] //슬라이스 선언 : 배열과 달리 크기를 지정하지 않았다.
	fmt.Println(s)            // 첫번째 요소를 포함. 마지막 요소 제외

	slicesAreLikeReferencesToArrays()

	sliceDefaults()

	sliceLengthAndCapacity()
}
