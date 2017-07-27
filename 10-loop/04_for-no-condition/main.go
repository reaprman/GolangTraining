package main

import "fmt"

func main() {
	i := 0
	// infinite loop
	for {
		fmt.Println("i - ", i)
		i++
	}
}

/*
Run Result:
...appendi -  86724
i -  86725
i -  86726
i -  86727
i -  86728
i -  86729
i -  86730
i -  86731
i -  86732
i -  86733
i -  86734
i -  86735
i -  86736
i -  86737
i -  86738
i -  86739
i -  86740
i -  86741
i -  86742
exit status 2

*/
