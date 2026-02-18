package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/wesyu/gopl/ch1/lissajous"
)

func Server1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Server2() {
	var mu sync.Mutex
	var count int
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count %d\n", count)
		mu.Unlock()
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Server3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Exercise_1_12() {
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		var opts []lissajous.Option
		q := r.URL.Query()
		if cycles, err := strconv.Atoi(q.Get("cycles")); err == nil {
			opts = append(opts, lissajous.WithCycles(cycles))
		}
		if res, err := strconv.ParseFloat(q.Get("resolution"), 64); err == nil {
			opts = append(opts, lissajous.WithResolution(res))
		}
		if size, err := strconv.Atoi(q.Get("size")); err == nil {
			opts = append(opts, lissajous.WithSize(size))
		}
		if nframes, err := strconv.Atoi(q.Get("nframes")); err == nil {
			opts = append(opts, lissajous.WithNFrames(nframes))
		}
		if delay, err := strconv.Atoi(q.Get("delay")); err == nil {
			opts = append(opts, lissajous.WithDelay(delay))
		}
		lissajous.Lissajous(w, opts...)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
