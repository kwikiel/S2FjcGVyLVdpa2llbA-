package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Work struct {
	ID       int64
	URL      string
	INTERVAL int64
}

type Download struct {
	Response   string
	Duration   time.Duration
	Created_at time.Time
}

var database = make(map[int64]Work)
var downloads = make(map[int64][]Download)

func fetcherhistory(w http.ResponseWriter, r *http.Request) {
	djson, _ := json.Marshal(downloads)
	w.WriteHeader(200)
	w.Write(djson)

}
func helloworld(t time.Time, foo int64) {
	url := "http://www.google.com"
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	tnow := time.Now()
	d1 := Download{Response: string(responseData), Duration: time.Since(start), Created_at: tnow}
	downloads[1] = append(downloads[1], d1)

}

func fetchurl(url string, id int64) string {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	tnow := time.Now()
	d1 := Download{Response: string(responseData), Duration: time.Since(start), Created_at: tnow}
	downloads[id] = append(downloads[id], d1)
	fmt.Println("Download completed")
	return "ok"
}

func polling(seconds int64, url string, id int64) {
	for {
		<-time.After(time.Duration(seconds))
		go fetchurl(url, id)
	}
}

func worker(w http.ResponseWriter, r *http.Request) {

	//doEvery(1*time.Second, helloworld)
	// Time 10**9 in nanoseconds interval
	// Url
	for k, v := range database {
		fmt.Printf("key[%v] value[%v]\n", k, v)
		go polling(v.INTERVAL*1000000000, v.URL, k)
	}
	//go polling(1000000000, "https://httpbin.org/range/15", 1 )

}

func fetcher(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json, _ := json.Marshal(database)
		w.Write(json)
	}

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var t Work
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		database[t.ID] = t
		id, _ := json.Marshal(t)
		w.WriteHeader(200)
		w.Write(id)

	}

	if r.Method == "DELETE" {
		decoder := json.NewDecoder(r.Body)
		var t Work
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		deltedid := t.ID
		delete(database, t.ID)
		w.WriteHeader(200)

		deltedjson, _ := json.Marshal(deltedid)

		w.Write(deltedjson)

	}

}

func greeter(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "hello\n")

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/", greeter)
	http.HandleFunc("/api/fetcher", fetcher)
	http.HandleFunc("/api/fetcher/history", fetcherhistory)
	http.HandleFunc("/worker", worker)

	server.ListenAndServe()
}
