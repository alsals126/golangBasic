package main

import ( // 괄호로 import 그룹 => "factored" import
	"fmt"
	"math"
	//"math"
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
/*
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

func creatingASliceWithMake() {
	a := make([]int, 5) //동적 크기의 배열생성(길이가 5, 용량이 5)
	printSlice2("a", a)

	b := make([]int, 0, 5) // 길이가 0, 용량이 5
	printSlice2("b", b)

	c := b[:2] //길이가 2, 용량이 5
	printSlice2("c", c)

	d := c[2:5] //길이가 3, 용량이 3
	printSlice2("d", d)
}
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func appendingToASlice() {
	var s []int //길이 0, 용량 0
	printSlice(s)

	s = append(s, 0) //길이1, 길이1
	printSlice(s)

	s = append(s, 1) //길이2, 길이2
	printSlice(s)

	s = append(s, 2, 3, 4) //길이5, 용량
	printSlice(s)
}

func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42                       //map에 요소 추가하기
	fmt.Println("The value:", m["Answer"]) //요소 검색하기

	m["Answer"] = 48 //map에 요소 업데이트하기
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") //요소 제거하기
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] //키가 존재하는지 : v는 key인 "Answer"의 value고, ok는 "Answer"가 있는지 없는지
	fmt.Println("The value:", v, "Present?", ok)
}

//   함수명(매개변수명 매개변수타입             )  리턴타입
func compute(fn func(float64, float64) float64) float64 { // 함수도 값이기 때문에, 함수의 인수나 반환 값으로 사용될 수 있다.
	return fn(3, 4)
}
func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))    //함수가 저장된 hypot을 인수로 사용한다.
	fmt.Println(compute(math.Pow)) //compute 매개변수타입(함수의 타입?)이 같으면 어떤 함수든 상관없다.
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

	var s2 []int                      //슬라이스의 zero value는 nil이다.
	fmt.Println(s2, len(s2), cap(s2)) //nil 슬라이스의 길이와 용량은 0이며, 기본 배열을 가지고 있지 않다.
	if s2 == nil {
		fmt.Println("nil!")
	}

	creatingASliceWithMake()

	appendingToASlice()

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow { //i는 인덱스, v는 해당 인덱스 값의 복사본
		fmt.Printf("2**%d = %d\n", i, v)
	}

	type Vertex struct { // struct가 없으면 오류가 난다.
		Lat, Long float64
	}
	var m map[string]Vertex // nil맵
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.2134, -12.1234,
	}
	fmt.Println(m["Bell Labs"])

	var m2 = map[string]Vertex{
		"B": { //최상위 타입이 타입 이름일 경우, 리터럴의 요소에서 생략 가능
			40.234, -234.1234,
		},
		"G": Vertex{
			37.413, -24.124,
		},
	}
	fmt.Println(m2)

	functionValues()
}
*/

// Method와 Interface
type Vertex struct{ X, Y float64 }

// 리시버가 장착된 함수는 함수가 아니라 메소드가 된다.
func (v Vertex) Abs() float64 { //v라는 이름의 Vertex유형의 리시버가 있다.
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64 //구조체가 아닌 형식에 대해서도 메소드를 선언할 수 있다.

func (f MyFloat) Abs() float64 { //메소드와 동일한 패키지에 유형이 정의된 리시버가 있는 메소드만 선언할 수 있다.
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex2 struct{ X, Y float64 }

func (v Vertex2) Abs() float64 { //Vertext값을 복사해서 이용한다.
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex2) Scale(f float64) { // 포인터 리시버 : 리시버가 가리키는 값을 수정할 수 있다.
	v.X = v.X * f
	v.Y = v.Y * f
}

type Vertex3 struct{ X, Y float64 }

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex3, f float64) { //인수에 포인터 인수가 있다.
	v.X = v.X * f
	v.Y = v.Y * f
}

type I interface{ M() }
type T struct{ S string }

func (t T) M() {
	fmt.Println(t.S)
}

type I2 interface{ M2() }
type F float64

func (t *T) M2() {
	fmt.Println(t.S)
}
func (f F) M2() {
	fmt.Println(f)
}
func describe(i I2) {
	// 값과 타입 출력
	fmt.Printf("(%v, %T)\n", i, i)
}
func interfaceValues() {
	var i I2

	i = &T{"Hello"}
	describe(i) // (&{Hello}, *main.T)
	i.M2()      // Hello

	i = F(math.Pi)
	describe(i) // (3.14... , main.F)
	i.M2()      // 3.14...
}

type I3 interface{ M3() }

func (t *T) M3() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func describe3(i I3) {
	// 값과 타입 출력
	fmt.Printf("(%v, %T)\n", i, i)
}
func interfaceValuesWithNilUnderlyingValues() {
	var i I3

	var t *T
	i = t
	describe3(i) // (<nil>, *main.T)
	i.M3()       // <nil>

	i = &T{"hello"}
	describe3(i) // (&{hello}, *main.T)
	i.M3()       // hello
}

func describe4(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func theEmptyInterface() {
	var i interface{} // _empty_interface
	describe4(i)      // (<nil>, <nil>)

	i = 42
	describe4(i) // (42, int)

	i = "hello"
	describe4(i) // (hello, string)
}

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string) //string타입의 값이 있으므로 변수s에 할당한다.
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok) // hello, true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0, false

	//f=i.(float64) //panic
	fmt.Println(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2) // math.Sqrt2: 2의 제곱근 값
	fmt.Println(f.Abs())

	v2 := Vertex2{3, 4}
	v2.Scale(10)
	fmt.Println(v2.Abs())

	v3 := Vertex3{3, 4}
	v3.Scale(2)        //Scale메서드가 포인터 리시버를 가졌기 때문에, 편의상 v.Scale(5)를 (&v).Scale(5)로 해석한다.
	ScaleFunc(&v3, 10) //포인터 인수가 있어서 &연산자를 이용해 포인터를 생성한다.
	p := &Vertex3{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v, p)

	var i I = T{"hello"}
	i.M() //메소드를 실행함으로써 인터페이스를 구현한다.

	interfaceValues()

	interfaceValuesWithNilUnderlyingValues()

	theEmptyInterface()

	typeAssertions()
}
