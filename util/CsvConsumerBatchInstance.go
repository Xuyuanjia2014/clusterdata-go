package util

import (
	"fmt"
)

func ConvertCsv()  {
	fmt.Println("start convert")
	size :=0;
	for x := range Channel{
		fmt.Println(x)
		size++;
		if size >10 {
			break
		}
	}
}