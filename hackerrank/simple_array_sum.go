package hackerrank

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
	var total int32 = 0
	for _, e := range ar {
		total += e
	}
	return total
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
//github.com/spf13/cobra
func SimpleArraySum() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for arItr := 0; arItr < int(arCount); arItr++ {
		arItemTemp, err := strconv.ParseInt(arTemp[arItr], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := simpleArraySum(ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}