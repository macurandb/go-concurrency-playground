package main

import (
	"fmt"
	"github.com/macurandb/go-concurrency-playground/examples"
)

func main() {
	const separator = "--------"

	fmt.Println(separator)
	examples.MainGoroutines()
	fmt.Println(separator)
	examples.MainChannels1()
	fmt.Println(separator)
	examples.MainChannels3()
	fmt.Println(separator)
	examples.MainSelect1()
	fmt.Println(separator)
	examples.MainSelect2()
	fmt.Println(separator)
	examples.MainFindBind()
	fmt.Println(separator)
	examples.MainDownload()
	fmt.Println(separator)
	examples.MainDonwloadCurrency()
}
