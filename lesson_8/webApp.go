package main

import (
	"flag"
	"fmt"
	"net/http"
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
	var name = flag.String("name", "User", "Name for Hello")
	var port = flag.String("port", "8080", "Port")
	var count = flag.Int("count", 1, "Number of repeats")
	flag.Parse()

	config := Config{
		Name:  *name,
		Port:  *port,
		Count: *count,
	}

	runWebServer(config)
}
