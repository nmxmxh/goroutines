package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type response struct {
	url     string
	code    int
	message string
}

type summarizedResponse struct {
	code       int
	occurences int
}

type appStatistics struct {
	successResponse atomic.Int64
	failureResponse atomic.Int64
	idleTime        atomic.Int64
	requestTime     atomic.Int64
	appTotalTime    atomic.Int64
}

var waitGroup sync.WaitGroup
var successChannel chan response
var failureChannel chan response
var summaryOfResponse []*summarizedResponse
var mutex sync.Mutex
var appStats appStatistics

func main() {
	startTime := time.Now()
	successChannel = make(chan response)
	failureChannel = make(chan response)
	appStats = appStatistics{}

	urls := []string{
		"https://google.com",
		"https://openai.com",
		"https://google.com/not-found",
		"https://none-existent-url.com/not-found",
	}

	waitGroup.Add(len(urls))

	for _, url := range urls {
		go fetchURL(url)
	}

	go func() {
		for {
			startTime := time.Now()
			select {
			case success := <-successChannel:
				logSummaryOfResponse(success)
				fmt.Println(success)
				appStats.successResponse.Add(1)
				waitGroup.Done()
			case failure := <-failureChannel:
				logSummaryOfResponse(failure)
				fmt.Println(failure)
				appStats.failureResponse.Add(1)
				waitGroup.Done()
			default:
				idleTime := time.Since(startTime).Microseconds()
				appStats.idleTime.Add(idleTime)
				// implement idle functionality
			}
		}
	}()

	waitGroup.Wait()

	appTime := time.Since(startTime).Microseconds()
	appStats.appTotalTime.Add(appTime)

	fmt.Println("-- summary of responses --")
	for _, sr := range summaryOfResponse {
		fmt.Printf("result: %v\n", sr)
	}

	fmt.Println("-- statistics --")
	fmt.Printf("successes: %v\n", appStats.successResponse.Load())
	fmt.Printf("failures: %v\n", appStats.successResponse.Load())
	fmt.Printf("idle time (sec): %v\n", float64(appStats.idleTime.Load())/1000/1000)
	fmt.Printf("request time (sec): %v\n", float64(appStats.requestTime.Load())/1000/1000)
	fmt.Printf("app time (sec): %v\n", float64(appStats.appTotalTime.Load())/1000/1000)
}

func fetchURL(url string) {
	startTime := time.Now()
	defer func() {
		requestTime := time.Since(startTime).Microseconds()
		appStats.requestTime.Add(requestTime)
	}()
	resp, err := http.Get(url)
	if err != nil {
		failureChannel <- response{url: url, message: fmt.Sprintf("error occurred on %s: %v ", url, err.Error())}
		return
	}

	if resp.StatusCode != http.StatusOK {
		failureChannel <- response{url: url, code: resp.StatusCode, message: fmt.Sprintf("error occurred on %s: %v", url, resp.StatusCode)}
		return
	}

	defer resp.Body.Close()
	successChannel <- response{url: url, code: resp.StatusCode, message: fmt.Sprintf("request successful on %s %v", url, resp.StatusCode)}
}

func logSummaryOfResponse(resp response) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, r := range summaryOfResponse {
		if r.code == resp.code {
			r.occurences++
			return
		}
	}
	summaryOfResponse = append(summaryOfResponse, &summarizedResponse{code: resp.code, occurences: 1})
}

// EXERCISE 3. Added Mutex for Shared Resource

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// type response struct {
// 	url     string
// 	code    int
// 	message string
// }

// type summarizedResponse struct {
// 	code       int
// 	occurences int
// }

// var waitGroup sync.WaitGroup
// var successChannel chan response
// var failureChannel chan response
// var summaryOfResponse []*summarizedResponse
// var mutex sync.Mutex

// func main() {
// 	successChannel = make(chan response)
// 	failureChannel = make(chan response)

// 	urls := []string{
// 		"https://google.com",
// 		"https://openai.com",
// 		"https://google.com/not-found",
// 		"https://none-existent-url.com/not-found",
// 	}

// 	waitGroup.Add(len(urls))

// 	for _, url := range urls {
// 		go fetchURL(url)
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case success := <-successChannel:
// 				logSummaryOfResponse(success)
// 				fmt.Println(success)
// 				waitGroup.Done()
// 			case failure := <-failureChannel:
// 				logSummaryOfResponse(failure)
// 				fmt.Println(failure)
// 				waitGroup.Done()
// 			}
// 		}
// 	}()

// 	waitGroup.Wait()
// 	fmt.Println("-- summary of responses --")
// 	for _, sr := range summaryOfResponse {
// 		fmt.Printf("result: %v\n", sr)
// 	}
// }

// func fetchURL(url string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		failureChannel <- response{url: url, message: fmt.Sprintf("error occurred on %s: %v ", url, err.Error())}
// 		return
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		failureChannel <- response{url: url, code: resp.StatusCode, message: fmt.Sprintf("error occurred on %s: %v", url, resp.StatusCode)}
// 		return
// 	}

// 	defer resp.Body.Close()
// 	successChannel <- response{url: url, code: resp.StatusCode, message: fmt.Sprintf("request successful on %s %v", url, resp.StatusCode)}
// }

// func logSummaryOfResponse(resp response) {
// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	for _, r := range summaryOfResponse {
// 		if r.code == resp.code {
// 			r.occurences++
// 			return
// 		}
// 	}
// 	summaryOfResponse = append(summaryOfResponse, &summarizedResponse{code: resp.code, occurences: 1})
// }

// EXERCISE 2: Channels & Wait Groups

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// var waitGroup sync.WaitGroup
// var successChannel chan string
// var failureChannel chan string

// func main() {
// 	successChannel = make(chan string)
// 	failureChannel = make(chan string)

// 	urls := []string{
// 		"https://google.com",
// 		"https://openai.com",
// 		"https://google.com/not-found",
// 		"https://none-existent-url.com/not-found",
// 	}

// 	waitGroup.Add(len(urls))

// 	for _, url := range urls {
// 		go fetchURL(url)
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case success := <-successChannel:
// 				fmt.Println(success)
// 				waitGroup.Done()
// 			case failure := <-failureChannel:
// 				fmt.Println(failure)
// 				waitGroup.Done()
// 			}
// 		}
// 	}()

// 	waitGroup.Wait()
// }

// func fetchURL(url string) {
// 	response, err := http.Get(url)
// 	if err != nil {
// 		failureChannel <- fmt.Sprintf("error occurred on %s: %v \n", url, err.Error())
// 		return
// 	}

// 	if response.StatusCode != http.StatusOK {
// 		failureChannel <- fmt.Sprintf("error occurred on %s: %v \n", url, (response.StatusCode))
// 		return
// 	}

// 	defer response.Body.Close()
// 	successChannel <- fmt.Sprintf("request successful on %s %v \n", url, response.StatusCode)
// }

// EXAMPLE 1. GO ROUTINES AND WAIT GROUPS

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// var waitGroup sync.WaitGroup

// func main() {
// 	urls := []string{
// 		"https://google.com",
// 		"https://openai.com",
// 		"https://google.com/not-found",
// 		"https://none-existent-url.com/not-found",
// 	}

// 	waitGroup.Add(len(urls))

// 	for _, url := range urls {
// 		go fetchURL(url)
// 	}

// 	waitGroup.Wait()
// }

// func fetchURL(url string) {
// 	defer waitGroup.Done()
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("error occurred on %s: %v \n", url, err.Error())
// 		return
// 	}

// 	if response.StatusCode != http.StatusOK {
// 		fmt.Printf("error occurred on %s: %v \n", url, (response.StatusCode))
// 		return
// 	}

// 	defer response.Body.Close()
// 	fmt.Printf("request successful on %s %v \n", url, response.StatusCode)
// }
