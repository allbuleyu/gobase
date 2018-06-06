package ch7

import (
	"fmt"
	"flag"
	"io"
	"bytes"
	"os"
)

type Celsius float64

type Fahrenheit float64

type celsiusFlag struct {
	Celsius
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5+32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32)*5/9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":

		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperture %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{Celsius:value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}


type Tss struct {
	I int
}

// CelsiusFlag的测试方法
func ttt() {
	fmt.Println(os.Args[1:])
	var temp = CelsiusFlag("temp", 20.0, "the temperature")
	flag.Parse()

	fmt.Println(temp)
}

// 接口的陷阱
func xianjing() {
	debug := false
	//var buf io.Writer			// 不会出错
	var buf *bytes.Buffer		// 会出错, panic   因为把*bytes.Buffer 转为io.Writer接口, io.Writer接口的动态类型是*bytes.Buffer,所以w != nil 为 true,调用w.Write时候panic
	if debug {
		buf = new(bytes.Buffer)
	}

	f := func(w io.Writer) {
		if w != nil {
			w.Write([]byte("hello"))
			fmt.Println("not nil")
		}else {
			fmt.Println("nil")
		}
	}

	f(buf)
}
