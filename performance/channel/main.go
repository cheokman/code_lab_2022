package main

import "fmt"

func main() {
	buffered_channel := make(chan string, 2)
	buffered_channel <- "foo"
	buffered_channel <- "bar"

	fmt.Printf("Channel Length %d and Cap %d After Add: ", len(buffered_channel), cap(buffered_channel))

	fmt.Println(<-buffered_channel)
	fmt.Println(<-buffered_channel)

	fmt.Printf("Channel Length %d and Cap %d After Add: ", len(buffered_channel), cap(buffered_channel))

	buffered_channel <- "baz"

	out := <-buffered_channel

	fmt.Println(out)
	close(buffered_channel)
}
