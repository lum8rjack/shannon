package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
)

func StringEntropy(s string) float64 {
	m := map[rune]float64{}
	for _, r := range s {
		m[r]++
	}
	var hm float64
	for _, c := range m {
		hm += c * math.Log2(c)
	}
	l := float64(len(s))
	r := math.Log2(l) - hm/l

	return r
}

func DataEntropy(d []byte) float64 {
	m := map[byte]float64{}
	for _, r := range d {
		m[r]++
	}
	var hm float64
	for _, c := range m {
		hm += c * math.Log2(c)
	}
	l := float64(len(d))
	r := math.Log2(l) - hm/l

	return r
}

func FileEntropy(filename string) (float64, error) {
	var ent float64
	m := map[byte]float64{}

	f, err := os.Open(filename)
	if err != nil {
		return ent, err
	}
	defer f.Close()

	b := bufio.NewReader(f)

	var buf [1]byte
	filelength := 0

	for {
		_, err := b.Read(buf[:])
		if err == io.EOF {
			break
		}

		m[buf[0]]++
		filelength++
	}

	l := float64(filelength)

	var hm float64
	for _, c := range m {
		hm += c * math.Log2(c)
	}

	ent = math.Log2(l) - hm/l

	return ent, nil
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func main() {
	stringdata := flag.String("string", "", "String to calculate the Shannon entropy of")
	filename := flag.String("file", "", "File to calculate the Shannon entropy of")
	flag.Parse()

	if *stringdata == "" && *filename == "" {
		flag.PrintDefaults()
		fmt.Println("\nYou must provide either a string of file to check")
		os.Exit(0)
	}

	if *stringdata != "" {
		entropy := StringEntropy(*stringdata)
		fmt.Println(entropy)
	} else {
		entropy, err := FileEntropy(*filename)
		must(err)
		fmt.Println(entropy)
	}
}
