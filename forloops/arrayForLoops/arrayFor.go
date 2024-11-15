package arrayforloops

import "fmt"

// function that takes an array of integers and returns the maximum element
func MaxElement() {
	x := [5]int{1, 2, 6, 9, 5}
	z := 0
	for _, num := range x {
		fmt.Println(num)
		if z <= num {
			z = num
		}
	}
	fmt.Printf("Biggest number is %d\n", z)
}

// program to reverse the elements of an array in place
func ReverseArray() {
	y := [6]int{1, 2, 3, 4, 5, 6}
	z := [len(y)]int{}
	for i := 0; i < len(y); i++ {
		z[i] = y[len(y)-1-i]
	}
	fmt.Println("Reversed array is:", z)
}

// function that counts the occurrences of a given element in an array
func CountOccur() {
	y := [6]string{"delhi", "Mumbai", "delhi", "delhi", "Shimla", "Amritsar"}
	x := "Chennai"
	i := 0
	for _, wrd := range y {
		if x == wrd {
			i++
		}
	}
	fmt.Printf("%s word comes %d times.\n", x, i)
}

// function that calculates the average of numbers in an array
func Average() {
	a := [6]int{1, 8, 11, 7, 4, 6}
	b := 0
	for _, x := range a {
		b = b + x
	}
	c := b / len(a)
	fmt.Printf("average is %d.\n", c)

}

// find all unique elements (elements that appear only once).

func Unique() {
	c := [7]int{9, 6, 11, 12, 6, 9, 6}

	uniqMap := make(map[int]int)

	for _, num := range c {
		uniqMap[num]++
	}
	noUniq := false

	for num, count := range uniqMap {
		if count == 1 {
			fmt.Println(num)
			noUniq = true
		}
	}
	if !noUniq {
		fmt.Println("No unique number in an array")
	}
}
