# Test text data for development

This package is simple, and it makes it easy for a Golang program
to get a large body of a (non-trivial) English text.

## Quickstart

Install the package:

	go get github.com/wkoszek/go-texttestdata

## Example

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

There is an example which you can build:

	cd ./examples/
	go build example.go

## Contributions

I was able to provide some additional `.txt` thanks to: [https://github.com/ssbc/ssb-testnet.git](https://github.com/ssbc/ssb-testnet.git
) and project Gutenberg.

# Author

- Wojciech Adam Koszek, [wojciech@koszek.com](mailto:wojciech@koszek.com)
- [http://www.koszek.com](http://www.koszek.com)
