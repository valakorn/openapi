// package function

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func readLines(name string) []string {
// 	f, err := os.Open(name)
// 	if os.IsNotExist(err) {
// 		return nil
// 	}
// 	check(err)
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	var lines []string
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	check(scanner.Err())

// 	return lines
// }

// func writeLines(lines []string, path string) error {
// 	// overwrite file if it exists
// 	//file, err := os.OpenFile("./DBFILE.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
// 	file, err := os.OpenFile(path, os.O_APPEND, os.ModeAppend)
// 	check(err)

// 	defer file.Close()
// 	w := bufio.NewWriter(file)
// 	_, err = fmt.Fprintln(file, lines)
// 	check(err)

// 	return w.Flush()
// }
