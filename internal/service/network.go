package service

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/MohammadTaghipour/flumint/internal/utils"
	"github.com/briandowns/spinner"
)

func RunNetworkCheck() error {
	fmt.Println()

	repos := []struct {
		Name string
		Host string
		Port string
	}{
		{"Pub.dev", "pub.dev", "443"},
		{"Flutter Releases", "storage.googleapis.com", "443"},
		{"Google Maven", "maven.google.com", "443"},
		{"Maven Central", "repo.maven.apache.org", "443"},
		{"CocoaPods CDN", "cdn.cocoapods.org", "443"},
	}

	s := spinner.New(spinner.CharSets[utils.SpinnerCharset], utils.SpinnerDuration)
	s.Suffix = " Checking network connectivity..."
	s.Color(utils.SpinnerColor)
	s.Start()
	defer s.Stop()

	type result struct {
		Name      string
		Reachable bool
		LatencyMS int64
		Err       error
	}

	results := make([]result, len(repos))
	var wg sync.WaitGroup

	for i, repo := range repos {
		wg.Add(1)
		go func(i int, rName, rHost, rPort string) {
			defer wg.Done()

			start := time.Now()
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(rHost, rPort), 2*time.Second)
			latency := time.Since(start).Milliseconds()

			if err != nil {
				results[i] = result{Name: rName, Reachable: false, LatencyMS: latency, Err: err}
				return
			}
			conn.Close()
			results[i] = result{Name: rName, Reachable: true, LatencyMS: latency, Err: nil}

		}(i, repo.Name, repo.Host, repo.Port)
	}

	wg.Wait()
	s.Stop()

	fmt.Println(utils.BrandWriter("Flumint Network Check"))
	fmt.Println("--------------------------------")

	hasFailure := false
	for _, r := range results {
		if r.Reachable {
			fmt.Println(utils.SuccessWriter(fmt.Sprintf("%-20s : %d ms", r.Name, r.LatencyMS)))
		} else {
			hasFailure = true
			fmt.Println(utils.ErrorWriter(fmt.Sprintf("%-20s : %v", r.Name, r.Err)))
		}
	}

	fmt.Println()
	if hasFailure {
		return fmt.Errorf("one or more repositories are unreachable")
	}

	fmt.Println(utils.SuccessWriter("All repositories are reachable."))
	fmt.Println()
	return nil
}
