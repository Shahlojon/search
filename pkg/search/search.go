package search

import (
	"context"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

//Result ..
type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}


//Any 
func Any(ctx context.Context, phrase string, files []string) <-chan Result {

	ch := make(chan Result)
	wg := sync.WaitGroup{}
	result := Result{}
	for i := 0; i < len(files); i++ {

		data, err := ioutil.ReadFile(files[i])
		if err != nil {
			log.Println("error not opened file err => ", err)
		}
		filetext := string(data)

		if strings.Contains(filetext, phrase) {
			res := FindAnyMatchTextInFile(phrase, filetext)
			if (Result{}) != res {
				result = res
				break
			}
		}

	}
	log.Println(result)

	wg.Add(1)
	go func(ctx context.Context, ch chan<- Result) {
		defer wg.Done()
		if (Result{}) != result {
			ch <- result
		} 
	}(ctx, ch)

	go func() {
		defer close(ch)
		wg.Wait()
	}()
	return ch
}

//FindAnyMatchTextInFile 
func FindAnyMatchTextInFile(phrase, filetext string) (res Result) {

	//ch := make(chan Result)

	temp := strings.Split(filetext, "\n")

	for i, line := range temp {
		//fmt.Println("[", i+1, "]\t", line)
		if strings.Contains(line, phrase) {

			return Result{
				Phrase:  phrase,
				Line:    line,
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(line, phrase)) + 1,
			}

		}
	}

	return res
}