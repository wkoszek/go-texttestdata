package main

import "fmt"
import "github.com/wkoszek/go-texttestdata"

func main() {

	// To print out a big body of text from each of the public
	// domain books, you can use:

	fmt.Println(string(gotexttestdata.MobyDick));
	fmt.Println(string(gotexttestdata.ATaleOfTwoCities));
	fmt.Println(string(gotexttestdata.Dracula));
	fmt.Println(string(gotexttestdata.GrimsFairyTales));
	fmt.Println(string(gotexttestdata.Metamorphosis));
	fmt.Println(string(gotexttestdata.PictureOfDorianGrey));
}
