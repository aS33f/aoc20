package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("20.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	codes := make(map[int64]byte)
	image := make(map[int][]byte)
	c := 0
	for scanner.Scan() {
		c++
		line := scanner.Bytes()
		if c == 1 {
			for i := range line {
				codes[int64(i)] = line[i]
			}
		}
		if c >= 3 {
			image[c-2] = []byte{46, 46}
			image[c-2] = append(image[c-2], line...)
			image[c-2] = append(image[c-2], []byte{46, 46}...)
		}
	}
	frame := []byte{}
	for _, el := range image[1] {
		el = 46
		frame = append(frame, el)
	}
	image[len(image)+1] = frame
	image[len(image)+1] = frame
	image[-1] = frame
	image[0] = frame
	result := image

	for n := 1; n <= 50; n++ {
		result = enhance(result, codes, n)
	}

	count := 0
	for v := range result {
		for w := range result[v] {
			if result[v][w] == 35 {
				count = count + 1
			}
		}
	}
	fmt.Println("pixels lit: ", count)
}
func enhance(image map[int][]byte, codes map[int64]byte, n int) map[int][]byte {
	temp := make(map[int][]byte)
	for i := 1 - n; i <= len(image)-(2+n); i++ {
		for p := 1; p <= len(image[1])-2; p++ {
			bin := []byte{}
			bin = append(bin, image[i-1][p-1])
			bin = append(bin, image[i-1][p])
			bin = append(bin, image[i-1][p+1])
			bin = append(bin, image[i][p-1])
			bin = append(bin, image[i][p])
			bin = append(bin, image[i][p+1])
			bin = append(bin, image[i+1][p-1])
			bin = append(bin, image[i+1][p])
			bin = append(bin, image[i+1][p+1])
			for i := range bin {
				if bin[i] == 46 {
					bin[i] = 48
				} else {
					bin[i] = 49
				}
			}
			dec, _ := strconv.ParseInt(string(bin), 2, 16)
			temp[i] = append(temp[i], codes[dec])
		}
	}
	if n%2 != 0 {
		for i := range temp {
			temp[i] = append([]byte{35, 35}, temp[i]...)
			temp[i] = append(temp[i], []byte{35, 35}...)
		}
		frame := []byte{}
		for _, el := range temp[1] {
			el = 35
			frame = append(frame, el)
		}
		temp[len(temp)-(n-1)] = frame
		temp[len(temp)-(n-1)] = frame
		temp[-1*(n+1)] = frame
		temp[-1*(n)] = frame
		return temp
	}
	if n%2 == 0 {
		for i := range temp {
			temp[i] = append([]byte{46, 46}, temp[i]...)
			temp[i] = append(temp[i], []byte{46, 46}...)
		}
		frame := []byte{}
		for _, el := range temp[1] {
			el = 46
			frame = append(frame, el)
		}
		temp[len(temp)-(n-1)] = frame
		temp[len(temp)-(n-1)] = frame
		temp[-1*(n+1)] = frame
		temp[-1*(n)] = frame
		return temp
	}
	return temp
}
