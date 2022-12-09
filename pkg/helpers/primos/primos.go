package primos

import (
	"math"
)

// Sieve3999 alteração da função Sieve do pacote https://github.com/fxtlabs/primes,
// a mesma retorna os 550 primos inferiores a 3999
// Sieve returns a list of the prime numbers less than or equal to n.
// If n is less than 2, it returns an empty list.
// The function uses the sieve of Eratosthenes algorithm
// (see https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)
// with the following optimizations:
//
// * The initial list of candidate primes includes odd numbers only.
//
// * Given a prime p, only multiples of p greater than or equal to p*p need to be marked off since smaller multiples of p have already been marked off by then.
//
// * The above also implies that the algorithm can terminate as soon as it finds  a prime p such that p*p is greater than n.
//
// Sieve takes O(n) memory and runs in O(n log log n) time.
func Sieve3999() []int {
	// ALTERAÇÃO  -- Define 3999 como o número de primos.
	n := 3999
	// ALTERAÇÃO  -- Adicionado essa variável para evitar o cálculo de caracteres para o make  Final
	primosValidos := 550

	// ALTERAÇÃO --  Removi as regras de n <2 vez que não se recebe mais n as mesmas sempre serão falsas.
	/*
		switch {
		case n < 2:
			return []int{}
		case n == 2:
			return []int{2}
		}
	*/

	// a[i] == false ==> p=2*i+3 is a candidate prime
	// p in [3,n] ==> i in [0,(n-3)/2]
	length := 1 + (n-3)/2
	a := make([]bool, length, length)
	// Start with number 3 and consider only odd numbers
	sqrtn := int(math.Sqrt(float64(n)))
	for i, p := 0, 3; p <= sqrtn; p += 2 {
		if !a[i] {
			// 2*i+1 is a prime number; mark off its multiples
			for j := (p*p - 3) / 2; j < length; j += p {
				a[j] = true
			}
		}
		i++
	}
	// ps will store the computed primes; its initial capacity is based
	// an estimate of the prime-counting function pi(n)
	// Alteração -- Recebe os primos válidos para evitar usar a função pi(n)
	ps := make([]int, 1, primosValidos)
	ps[0] = 2
	for i := 0; i < length; i++ {
		if !a[i] {
			ps = append(ps, 2*i+3)
		}
	}
	return ps
}
