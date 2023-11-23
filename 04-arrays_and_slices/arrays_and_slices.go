package arrays_and_slices

import "fmt"

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		var tail []int
		if len(numbers) > 0 {
			tail = numbers[1:]
		}
		sums = append(sums, Sum(tail))
	}

	return
}

// AffectingSlicedArrayExample shows an example of how changing the slice affects the original array;
// but a "copy" of the slice will not affect the original array
func AffectingSlicedArrayExample() {
	x := [3]string{"Лайка", "Белка", "Стрелка"}

	y := x[:] // slice "y" points to the underlying array "x"
	y = append(y, "Шарик")

	z := make([]string, len(x))
	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

	y[1] = "Belka" // the value at index 1 is now "Belka" for both "y" and "x"

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}

// CopyingAfterSlicing shows an example of why it's a good idea
// to make a copy of a slice after slicing a very large slice
func CopyingAfterSlicing() {
	a := make([]int, 1e6) // slice "a" with len = 1 million
	b := a[:2]            // even though "b" len = 2, it points to the same the underlying array "a" points to

	c := make([]int, len(b)) // create a copy of the slice so "a" can be garbage collected
	copy(c, b)
	fmt.Println(c)
}
