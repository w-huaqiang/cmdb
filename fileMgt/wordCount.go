package word

import (
	"bufio"
	"io"
	"os"
	"sort"
)

// ValStruct is a struct
type ValStruct struct {
	Keys   []string
	Values []int
}

//TextCount is function
func TextCount(file string) (ValStruct, error) {

	var words map[string]int

	ret := ValStruct{
		Keys:   nil,
		Values: nil,
	}
	words = make(map[string]int)

	f, err := os.Open(file)
	if err != nil {
		return ret, err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return ret, err
		}

		for _, x := range line {
			if x == 10 {
				continue
			}
			if num, ok := words[string(x)]; ok {
				num++
				words[string(x)] = num
			} else {
				words[string(x)] = 1
			}
		}

	}

	wordStruct := toStruct(words)
	sort.Sort(wordStruct)

	return *wordStruct, nil

}

func toStruct(m map[string]int) *ValStruct {
	length := len(m)
	vs := &ValStruct{
		Keys:   make([]string, 0, length),
		Values: make([]int, 0, length),
	}
	for k, v := range m {
		vs.Keys = append(vs.Keys, k)
		vs.Values = append(vs.Values, v)
	}

	return vs
}

func (a *ValStruct) Len() int { return len(a.Keys) }
func (a *ValStruct) Swap(i, j int) {
	a.Keys[i], a.Keys[j] = a.Keys[j], a.Keys[i]
	a.Values[i], a.Values[j] = a.Values[j], a.Values[i]
}
func (a *ValStruct) Less(i, j int) bool { return a.Values[i] > a.Values[j] }
