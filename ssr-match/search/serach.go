package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exits := matchers[feed.Type]
		if !exits {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()

		close(results)

	}()

	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exits := matchers[feedType]; exits {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
