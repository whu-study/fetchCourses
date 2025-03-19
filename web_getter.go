package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func doWebGetterWithRetry(client *http.Client, req *http.Request) OneJson {
	count := 0
	var resp *http.Response
	var err error
	for {
		resp, err = client.Do(req)
		if err != nil && count < 30 {
			count++
			log.Println(err)
			log.Println("Retrying...")
			time.Sleep(time.Second)
			continue
		}
		break
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	var oneJson []OneJson

	if err := json.Unmarshal(bodyText, &oneJson); err != nil {
		log.Fatal(err)
	}

	return oneJson[0]
}
