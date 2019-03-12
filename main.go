package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var encode, decode bool
	flag.BoolVar(&encode, "encode", false, "Encode the input")
	flag.BoolVar(&decode, "decode", false, "Decode the input")
	flag.Parse()

	if !encode && !decode {
		fmt.Fprintf(os.Stderr, "Shoud specify --encode or --decode\n")
		os.Exit(1)
		return
	}

	if encode && decode {
		fmt.Fprintf(os.Stderr, "You can only specify either --encode or --decode\n")
		os.Exit(1)
		return
	}

	var converter func(string) string
	if encode {
		converter = junkEncode
	} else {
		converter = junkDecode
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input := scanner.Text()
		output := converter(input)
		fmt.Println(output)
	}
}
