package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	filePath := flag.String("file", "", "the file whose checksum will be created")
	bitSize := flag.Int("n", 0, "bit size of the checksum that will be made (16, 32 or 64)")
	check := flag.Bool("check", false, "states if the checkChecksum functions will be run")
	checksumDefault := flag.String("checksum", "", "given checksum file path")

	flag.Parse()

	var result string
	var err error

	if *check {
		fmt.Println(checkChecksum(*filePath, *checksumDefault))
	} else {
		result, err = checksum(fmt.Sprintf("%v", *filePath), *bitSize)
		if err != nil {
			fmt.Println("checksum could not be calculated")
			os.Exit(0)
		}

		fmt.Println(result)
	}
}

func checksum(filename string, n int) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	sum := 0

	for _, ch := range file {
		sum += int(ch)
	}

	//fmt.Println(file)
	//fmt.Println(sum)

	var cSum string

	switch n {
	case 16:
		cSum = bit16(sum)

	case 32:
		cSum = bit32(sum)

	case 64:
		cSum = bit64(sum)

	default:
		fmt.Println("invalid bit size")
		os.Exit(0)
	}

	checksumFile, err := os.Create("checksum.txt")
	if err != nil {
		panic(err)
	}

	_, err = checksumFile.Write([]byte(cSum))
	return cSum, nil
}

func bit16(data int) string {
	binary := strconv.FormatInt(int64(data), 2)
	return fmt.Sprintf("%016s", binary)
}

func bit32(data int) string {
	binary := strconv.FormatInt(int64(data), 2)
	return fmt.Sprintf("%032s", binary)
}

func bit64(data int) string {
	binary := strconv.FormatInt(int64(data), 2)
	return fmt.Sprintf("%064s", binary)
}

func checkChecksum(filePath string, checksumPath string) bool {

	givenChecksumFile, err := os.ReadFile(checksumPath)
	if err != nil {
		fmt.Println("checksum file directory is not valid")
		os.Exit(0)
	}

	//fmt.Println(len(string(givenChecksumFile)))

	bitSize := len(string(givenChecksumFile))
	//fmt.Println(bitSize)

	checksumValue, err := checksum(filePath, bitSize)
	if err != nil {
		fmt.Println("there is not a valid file path")
		os.Exit(0)
	}

	givenChecksum, err := strconv.Atoi(string(givenChecksumFile))
	if err != nil {
		if string(givenChecksumFile) == "" {
			fmt.Println("there is not a valid file to check")
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	checkChecksum, _ := strconv.Atoi(checksumValue)

	if givenChecksum == checkChecksum {
		return true
	}

	return false
}
