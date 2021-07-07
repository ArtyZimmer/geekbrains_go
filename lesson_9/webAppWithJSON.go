package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Name  string
	Port  string
	Count int
}

func runWebServer(config Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ""
		for i := 0; i < config.Count; i++ {
			data += fmt.Sprintf("Hello, %s\n", config.Name)
		}
		_, err := fmt.Fprint(w, data)
		if err != nil {
			return
		}
	})
	http.ListenAndServe("localhost:"+config.Port, nil)
}

func main() {
	var config Config

	jsonData, err := os.ReadFile("C:\\Users\\User\\GolandProjects\\geekbrains_go\\lesson_9\\" +
		"config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		panic(err)
	}

	runWebServer(config)

}
