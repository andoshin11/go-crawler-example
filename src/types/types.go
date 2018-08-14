package types

// FetcherResult type
type FetcherResult struct {
	URL string
}

// Channels type
type Channels struct {
	FetcherResult chan FetcherResult
	FetcherDone   chan int
	UploaderDone  chan int
}

// NewChannels returns new ref
func NewChannels() *Channels {
	return &Channels{
		FetcherResult: make(chan FetcherResult, 10),
		FetcherDone:   make(chan int, 10),
		UploaderDone:  make(chan int, 10),
	}
}

// DetailFetcherResult type
type DetailFetcherResult struct {
	ID   string
	Item *Museum
}

// DetailChannels type
type DetailChannels struct {
	FetcherResult chan DetailFetcherResult
	FetcherDone   chan int
	UploaderDone  chan int
}

// NewDetailChannels returns new ref
func NewDetailChannels() *DetailChannels {
	return &DetailChannels{
		FetcherResult: make(chan DetailFetcherResult, 10),
		FetcherDone:   make(chan int, 10),
		UploaderDone:  make(chan int, 10),
	}
}
