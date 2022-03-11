package main

import (
	"fmt"
	"net/http"

	"github.com/vrnvu/go-viper/internal/models"
	"github.com/vrnvu/go-viper/internal/naive"
)

var configPath = "./configs/"

func read(w http.ResponseWriter, req *http.Request) {
	v := naive.NewViper(configPath, "full", "yaml")
	naive.ReadInConfig(v)

	var c models.ConfigViper
	naive.Unmarshall(v, &c)

	fmt.Fprintf(w, "%v", c)
}

func main() {
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8090", nil)
}
