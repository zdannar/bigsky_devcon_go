package main

import (
	log "github.com/zdannar/flogger"
	"net/http"
	"sync"
)

func main() {
	gopherCnt := 10
	var wg sync.WaitGroup
	for i := 0; i <= gopherCnt; i++ {
		wg.Add(1)
		go func() {
			resp, err := http.Get("https://devops.communicaretechnology.com/v3.0/admin/")
			defer resp.Body.Close()
			defer wg.Done()
			if err != nil {
				log.Fatal("Darn %s", err)
			}
			log.Infof("Response code : %d", resp.StatusCode)
		}()
	}
	wg.Wait()
}
