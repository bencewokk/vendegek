package main

import "fmt"

type vendeg struct {
	erk int
	tav int
}

func bubbleSort(array []vendeg) []vendeg {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j].tav > array[j+1].tav {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func ujVendeg(erk int, tav int) vendeg {
	return vendeg{erk: erk, tav: tav}
}

func input() []int {
	var input1, input2 int

	fmt.Scanln(&input1, &input2)

	var output []int
	output = append(output, input1, input2)
	return output
}

func main() {

	var vendegek []vendeg

	vend := input()[0]
	for i := 0; i < vend; i++ {
		input := input()
		vendeg := ujVendeg(input[0], input[1])
		vendegek = append(vendegek, vendeg)
	}

	//fmt.Println(bubbleSort(vendegek))

	var photo []int

	for i := 0; i < len(vendegek)-1; i++ {
		//fmt.Println(i, vendegek[i])
		//fmt.Println(vendegek[i+1], "dddd")
		//fmt.Println(vendegek[i].erk,vendegek[i].tav)
		if vendegek[i].tav >= vendegek[i+1].erk {
			photo = append(photo, vendegek[i].tav)
		}

	}
	
	l:=len(vendegek)-1
	if vendegek[l].erk > vendegek[l-1].tav {
		photo = append(photo, vendegek[l].erk)
	}

	fmt.Println(photo)

}
