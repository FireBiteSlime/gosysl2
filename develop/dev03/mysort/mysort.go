package mysort

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type args struct {
	key       string
	number    bool
	reverse   bool
	unique    bool
	ignorelb  bool
	check     bool
	sortcheck bool
}

type Sort struct {
	args
	files  []string
	result []string
}

func CreateSort(key string, number bool, reverse bool, unique bool, ignorelb bool, check bool, files []string) *Sort {
	return &Sort{
		args: args{
			key:       key,
			number:    number,
			reverse:   reverse,
			unique:    unique,
			ignorelb:  ignorelb,
			check:     check,
			sortcheck: false,
		},
		files: files,
	}
}

func ReadFile(filepath string, ignrbl bool) ([]string, error) {
	result := make([]string, 0)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		line = line[:len(line)-1]
		if ignrbl {
			result = append(result, strings.TrimSpace(string(line)))
		} else {
			result = append(result, string(line))
		}

	}
	return result, nil
}

func isUnique(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return false
		}
	}
	return true
}

func makeUnique(file []string) []string {
	result := make([]string, 0)
	for _, v := range file {
		if isUnique(result, v) {
			result = append(result, v)
		}
	}
	return result
}

func makeReverse(file []string) []string {
	result := make([]string, len(file), cap(file))
	for i, v := range file {
		result[len(file)-i-1] = v
	}
	return result
}

// func ReadFiles(fls []string) []string {
// 	files := make([]string, 0)
// 	for _, v := range fls {
// 		file, err := readFile(v, false)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		files = append(files, file...)
// 	}
// 	return files
// }

func (s *Sort) Run() error {
	files := make([]string, 0)

	for _, v := range s.files {
		file, err := ReadFile(v, s.ignorelb)
		if err != nil {
			return err
		}
		files = append(files, file...)
	}

	if s.args.unique {
		files = makeUnique(files)
	}

	key, err := strconv.Atoi(s.key)
	if err != nil {
		key = 0
	}

	sort.Slice(files, func(i, j int) bool {
		aVals := strings.Split(files[i], " ")
		bVals := strings.Split(files[j], " ")
		if len(aVals) <= key || len(bVals) <= key {
			return false
		}
		if s.args.number {
			aInt, errA := strconv.Atoi(aVals[key])
			bInt, errB := strconv.Atoi(bVals[key])
			if errA == nil && errB == nil {
				if s.args.check && bInt > aInt {
					s.args.sortcheck = true
					return false
				}
				return aInt < bInt
			} else if errA != nil && errB != nil {
				if s.args.check && bVals[key] > aVals[key] {
					s.args.sortcheck = true
					return false
				}
				return aVals[key] < bVals[key]
			} else if errA != nil {
				return true
			} else if errB != nil {
				return false
			}
		} else {
			if s.args.check && bVals[key] > aVals[key] {
				s.args.sortcheck = true
				return false
			}
			return aVals[key] < bVals[key]
		}
		return false
	})

	if s.args.reverse {
		files = makeReverse(files)
	}
	s.result = files
	return nil
}

func (s *Sort) Output() error {
	if s.args.check && !s.args.sortcheck {
		return nil
	} else if s.args.check {
		return fmt.Errorf("Файл не отсортирован")
	}
	_, err := fmt.Fprintln(os.Stdout, strings.Join(s.result, "\n"))
	return err
}
