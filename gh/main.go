package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/google/go-github/v53/github"
)

type ObsidianReleaseStatus struct {
	Timestamp github.Timestamp `json:"timestamp"`
	Repo      string           `json:"repo"`
	User      string           `json:"user"`
}

func main() {
	s := &http.Server{
		Addr:           "127.0.0.1:55535",
		Handler:        ServerHandler{name: "test"},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	var oldObStat ObsidianReleaseStatus
	if _, err := os.Stat("time-cache.json"); err == nil {
		buffer, err := ioutil.ReadFile("time-cache.json")
		if err != nil {
			//return err
			fmt.Printf("%v\n", err)
		} else {
			json.Unmarshal(buffer, &oldObStat)
			fmt.Println("old timestamp --- ", oldObStat)
		}
	}
	// check repo main branch commit status per hour
	ticker := time.NewTicker(30 * time.Second)
	//ticker := time.NewTicker(1 * 60 * 60 * time.Second)
	go func() {
		for range ticker.C {
			ctx := context.Background()
			client := github.NewTokenClient(ctx, os.Getenv("GITHUB_AUTH_TOKEN"))

			branchCommits, _, err := client.Repositories.ListCommits(ctx, "obsidianmd", "obsidian-releases", nil)
			if err != nil {
				fmt.Printf("ListCommits: %v\n", err)
				continue
			}
			fmt.Println(*branchCommits[0].Commit.Author.Date)
			fmt.Printf("%v", branchCommits[0])
			latestTimestamp := *branchCommits[0].Commit.Author.Date
			latestObStat := ObsidianReleaseStatus{
				Timestamp: latestTimestamp,
				Repo:      "obsidian-releases",
				User:      "obsidianmd",
			}
			if latestTimestamp.After(oldObStat.Timestamp.Time) {
				fmt.Println("repo is updated")
				buff, err := json.Marshal(latestObStat)
				if err != nil {
					fmt.Println("marshal failed", err)
					continue
				}
				err = ioutil.WriteFile("time-cache.json", buff, 0600)
				if err != nil {
					fmt.Println("write file failed", err)
					continue
				}
			}
			fmt.Printf("old= %v\n", oldObStat)
			fmt.Printf("new= %v\n", latestTimestamp)
		}
	}()

	s.ListenAndServe()
}

type ServerHandler struct {
	name string
}

func (s ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()

	client := github.NewTokenClient(ctx, os.Getenv("GITHUB_AUTH_TOKEN"))

	branchCommits, _, err := client.Repositories.ListCommits(ctx, "google", "go-github", nil)
	if err != nil {
		fmt.Printf("ListBranchesHeadCommit: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for _, commit := range branchCommits {
		fmt.Fprintln(w, *commit.HTMLURL)
	}
	fmt.Printf("%v", branchCommits[0])
}
