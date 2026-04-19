package main

import (
	"fmt"
)

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	for b != 0 {
		remainder := a % b

		a = b
		b = remainder
	}

	return a
}

func Lcm(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	//higher numbers to lower overflow risk
	// g := Gcd(a, b)
	// return (a / g) * b

	x := (a*b)/Gcd(a,b)

	return x
}

func AreCoprime(a, b uint) bool {
	if a == 0 || b == 0 {
		return false
	}

	return Gcd(a, b) == 1
}

func SimplifyFraction(num, den uint) (uint, uint) {
	if num == 0 || den == 0 {
		return 0, 0
	}

	commonDivisor := Gcd(num, den)

	return (num/commonDivisor), (den/commonDivisor)
}

func IsPrime(n uint) bool {

	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n % 2 == 0 {
		return false
	}

	for i := uint(3); i * i <= n; i+= 2 {
		if n % i == 0 {
			return false
		}

	}

	return true

}

func main() {
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(7))
	fmt.Println(IsPrime(9))
	fmt.Println(IsPrime(1))
	fmt.Println(SimplifyFraction(42, 10))
	fmt.Println(SimplifyFraction(12, 18))
	fmt.Println(SimplifyFraction(0, 5))
	fmt.Println(Lcm(12, 18))
	fmt.Println(Lcm(7, 5))
	fmt.Println(Lcm(0, 5))
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
	fmt.Println(Gcd(0, 10))
	fmt.Println(AreCoprime(17, 3))
	fmt.Println(AreCoprime(12, 18))
}



