package main

import "fmt"
import "github.com/wkoszek/go-texttestdata"

func main() {

	// To print out a big body of text from each of the public
	// domain books, you can use:

	fmt.Println(string(gotestdata.MobyDick));
	fmt.Println(string(gotestdata.ATaleOfTwoCities));
	fmt.Println(string(gotestdata.Dracula));
	fmt.Println(string(gotestdata.GrimsFairyTales));
	fmt.Println(string(gotestdata.Metamorphosis));
	fmt.Println(string(gotestdata.PictureOfDorianGrey));
}
