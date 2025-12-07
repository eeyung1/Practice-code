package piscine

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
