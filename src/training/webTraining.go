package main

import (
	"fmt"
	"reflect"
)

func main() {
	outputTitle("05: 標準出力")
	section5()
	outputTitle("06: 変数")
	section6()
	outputTitle("07: データ型")
	section7()
	outputTitle("08: 配列")
	section8()
	outputTitle("09: 演算子")
	section9()
	outputTitle("10: 条件分岐")
	section10()
	outputTitle("11: 繰り返し")
	section11()
	outputTitle("12: 関数")
	section12()
	outputTitle("13: 構造体")
	section13()
	outputTitle("14: メソッド")
	section14()
}

func outputTitle(title string) {
	fmt.Printf("=== %s ===\n", title)
}

// 05: 標準出力
func section5() {
	fmt.Println("Good morning")
	fmt.Println("Good afternoon")
	fmt.Println("Good evening")
}

// 06: 標準出力
func section6() {
	var number int
	number = 1
	fmt.Println(number)

	var num = 5
	fmt.Println(num)

	num07 := 7
	fmt.Println(num07)
}

// 07: データ型
func section7() {
	var num8 int8 = 8
	var num16 int16 = 16
	var num32 int32 = 32
	var num64 int64 = 64
	fmt.Println(reflect.TypeOf(num8))
	fmt.Println(reflect.TypeOf(num16))
	fmt.Println(reflect.TypeOf(num32))
	fmt.Println(reflect.TypeOf(num64))

	var f32 float32 = 32
	var f64 float64 = 64
	fmt.Println(reflect.TypeOf(f32))
	fmt.Println(reflect.TypeOf(f64))

	var s string = "string"
	fmt.Println(reflect.TypeOf(s))

	var b bool = false
	fmt.Println(reflect.TypeOf(b))
}

// 08: 配列
func section8() {
	a := [3]string{"hoge", "fuga", "bizz"}
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])

	a[0] = "piyo"
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])

	b := [...]string{"hoge", "fuga", "bizz"}
	fmt.Println(b)
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])

	// c := [2][2]string{{"hoge", "fuga"}, {"bizz", "piyo"}}
	c := [...][2]string{{"hoge", "fuga"}, {"bizz", "piyo"}}
	fmt.Println(c)
	// NOTE: 配列数が違う場合はどのように表現する？
}

// 09: 演算子
func section9() {
	x := 10
	y := 2
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	fmt.Println(x > y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x == y && x != y)
	fmt.Println(x == y || x != y)

	fmt.Println(y)
	y++
	fmt.Println(y)
	y--
	fmt.Println(y)
	// NOTE: Cなどで使用する下記は不可
	// --y
	// ++y

	a := 10
	fmt.Println(a)
	// NOTE: インクリメントしながらの代入は不可
	// b := a++
	// c := ++a
}

// 10: 条件分岐
func section10() {
	// age := 30
	// if age > 20 {
	// 	fmt.Println("Adult")
	// } else if age == 0 {
	// 	fmt.Println("Baby")
	// } else {
	// 	fmt.Println("Child")
	// }

	// NOTE: 上記の簡易文
	if age := 30; age > 20 {
		fmt.Println("Adult")
	} else if age == 0 {
		fmt.Println("Baby")
	} else {
		fmt.Println("Child")
	}

	age2 := 20
	if age2 = 30; age2 > 20 {
		fmt.Printf("Age %s in if\n", age2)
	}
	fmt.Printf("Age %s out if\n", age2)
}

// 11: 繰り返し
// NOTE: whireはない
func section11() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("i = %d\n", i)
		if i > 4 {
			break
		}
		for j := 0; j < 3; j++ {
			fmt.Printf("j = %d\n", j)
		}
	}
	fmt.Println()

	z := 0
	for z < 4 {
		fmt.Printf("z = %d\n", z)
		z++
	}

	// TODO: 配列をforEach的なことはできる？
	// a := {1,2,3}
}

// 12: 関数
func section12() {
	// NOTE: function in functionはできない
	// func demo() {
	// 	fmt.Println("demo")
	// }
	fmt.Println(inclement(10))

	// NOTE: 無名関数可能
	func(s string) {
		fmt.Println(s)
	}("Hello")

	hello := func(s string) {
		fmt.Println(s)
	}
	hello("Hello")
}

func inclement(a int) int {
	return a + 1
}

// NOTE: カリー型は定義不可
// func add(a int)(b int) int {
// 	return a + b
// }

// 13: 構造体
func section13() {
	// NOTE: 構造体の初期値は指定できない
	// type Example struct {
	// 	type string = "global"
	// }
	type RectLocal struct {
		name string
		x, y int
	}

	var rect RectLocal
	rect.name = "name"
	rect.x = 80
	rect.y = 70

	fmt.Println(rect)

	rect2 := RectLocal{name: "name", x: 60, y: 40}
	fmt.Println(rect2)

	rect3 := RectLocal{x: 60, y: 40, name: "name"}
	fmt.Println(rect3)

	// NOTE: 改行はできない
	// rect3 := RectLocal{
	// 	name: "name",
	// 	x: 60,
	// 	y: 40
	// }
	// fmt.Println(rect3)
}

// 14: メソッド
// Goにはclassがない
func section14() {
	rect1 := RectGlobal{x: 60, y: 40, name: "name"}
	fmt.Println(rect1.getX())
	fmt.Println(rect1.sumX(10))
	// 無限ループは実行時エラー
	// fmt.Println(rect1.getA())
}

type RectGlobal struct {
	name string
	x, y int
}

func (rect RectGlobal) getX() int {
	return rect.x
}

// func (rect RectGlobal) getA() int {
// 	return rect.getB()
// }

// func (rect RectGlobal) getB() int {
// 	return rect.getA()
// }

func (rect RectGlobal) sumX(i int) int {
	return rect.getX() + i
}
