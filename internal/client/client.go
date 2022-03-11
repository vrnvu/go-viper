package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func ReadConfig() {
	url := "http://localhost:8090/read"
	for i := 0; i < 100000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(b))
		time.Sleep(1 * time.Second)
	}
}
