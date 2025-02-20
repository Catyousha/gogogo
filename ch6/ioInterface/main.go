package main

import (
	"bufio"
	"fmt"
	"io"
)

type S1 struct {
	F1 int
	F2 string
}

// read bytes from terminal
func (s *S1) Read(p []byte) (n int, err error) {
	fmt.Print("Give me your name:")
	fmt.Scanln(&p)
	s.F2 = string(p)
	return len(p), nil
}

// write bytes to terminal
func (s *S1) Write(p []byte) (n int, err error) {
	if s.F1 < 0 {
		return -1, nil
	}

	for i := 0; i < s.F1; i++ {
		fmt.Printf("%s ", p)
	}

	fmt.Println()

	return s.F1, nil
}

type S2 struct {
	F1 S1
	text []byte
}

func (s S2) eof() bool {
	return len(s.text) == 0;
}

func (s *S2) readByte() byte {
	// assume eof check was done before
	temp := s.text[0]
	s.text = s.text[1:]
	return temp
}

// read from terminal in traditional way
func (s *S2) Read(p []byte) (n int, err error) {
	if s.eof() {
		err = io.EOF
		return
	}

	l := len(p)
	if l > 0 {
		for n < l {
			p[n] = s.readByte()
			n++
			if s.eof() {
				s.text = s.text[0:0]
				break
			}
		}
	}
	return
}

func main()  {
	// S1: {4 Hello}
	s1 := S1{4, "Hello"}
	fmt.Println("S1:", s1)

	// Give me your name:awanama
	// S1.F1 != read result: 4 7
	// S1.F2: awanama
	buf := make([]byte, 2)
	readResult, err := s1.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("S1.F1 !== read result:", s1.F1, readResult)
	fmt.Println("S1.F2:", s1.F2)

	// repeate s1.F1 (4) times
	// Hello there! Hello there! Hello there! Hello there!
	_, _ = s1.Write([]byte("Hello there!"))

	// s2: {{4 awanama} [72 101 108 108 111 32 119 111 114 108 100 33 33]}
	s2 := S2{F1: s1, text: []byte("Hello world!!")}
	fmt.Println("s2:", s2)

	// f2.text chunked to len(buf) (2)
	// ** 2 He
	// ** 2 ll
	// ** 2 o 
	// ** 2 wo
	// ** 2 rl
	// ** 2 d!
	// ** 1 !
	r := bufio.NewReader(&s2)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("*", err)
			break
		}
		fmt.Println("**", n, string(buf[:n]))
	}
}