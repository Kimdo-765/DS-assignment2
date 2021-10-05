package ds_hw_0

import (
	"bufio"
	"io"
	"strconv"
    "os"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
    sum := 0
    for i := range nums {
        sum = sum + i
    }
    out <- sum
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers
    file, err := os.OpenFile(fileName, os.O_RDWR, os.FileMode(0644),)
    if err != nil {
        return 0
    }

    r := bufio.NewReader(file)
    numList, _ := readInts(r)
    bufLen := len(numList) / num

    var inputs []chan int
    for i := 0; i < num; i++ {
        inputs = append(inputs, make(chan int, bufLen))
    }
    output := make(chan int, num)

    for i := 0; i < num; i++ {
        for j := i*bufLen; j < (i+1)*bufLen; j++ {
            inputs[i] <- numList[j]
        }
        close(inputs[i])
        go sumWorker(inputs[i], output)
    }

    sumResult := 0

    for i := 0; i < num; i++{
        sumResult = sumResult + <-output
    }

	return sumResult
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
