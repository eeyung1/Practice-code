package main

import "fmt"

func Enigma(a ***int, b *int, c *******int, d ****int) {
	va := ***a
	vb := *b
	vc := *******c
	vd := ****d

	*******c = va // a -> c
	****d = vc    // c -> d
	*b = vd       // d -> b
	***a = vb     // b -> a
}


func main() {
	x := 7
	y := &x
	z := &y
	a := &z

	w := 6
	b := &w

	u := 3
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	k := 4
	l := &k
	m := &l
	n := &m
	d := &n

	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	Enigma(a, b, c, d)

	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

}
