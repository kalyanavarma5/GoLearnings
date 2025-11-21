package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}

// Simulate image download
func downloadImage(id int) string {
	time.Sleep(time.Second) // simulate delay
	return fmt.Sprintf("Image %d downloaded", id)
}

func sensor(id int, ch chan<- string) {
	time.Sleep(time.Second * time.Duration(id))
	ch <- fmt.Sprintf("Sensor %d data", id)
}

// Generate numbers
func generate(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()
	return ch
}

// Square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {

	// Fan-out: Multiple Goroutines Processing Jobs Concurrently

	//var wg sync.WaitGroup
	//images := []int{1, 2, 3, 4, 5}
	//
	//for _, id := range images {
	//	wg.Add(1)
	//	go func(imageID int) {
	//		defer wg.Done()
	//		result := downloadImage(imageID)
	//		fmt.Println(result)
	//	}(id)
	//}
	//wg.Wait()
	//fmt.Println("All images downloaded.")

	//Fan-in: Merging Results from Multiple Sources

	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(sensorID int) {
			defer wg.Done()
			sensor(sensorID, ch)
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		fmt.Println("Received:", data)
	}
	fmt.Println("All sensor data received.")

	//Pipeline: Multi-Stage Processing

	gen := generate(1, 2, 3, 4)
	sq := square(gen)
	for n := range sq {
		fmt.Println(n)
	}
}

	// Pipeline: Multi-Stage Processing

	// Fan-in: Merging Results from Multiple Sources

	//ctx, cancel := context.WithCancel(context.Background())
	//go worker(ctx)
	//
	//time.Sleep(2 * time.Second)
	//cancel() // gracefully stop worker
	//time.Sleep(500 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://api.example.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}
