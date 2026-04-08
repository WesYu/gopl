package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/wesyu/gopl/ch1/lissajous"
	"github.com/wesyu/gopl/ch3/surface"
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

// Exercise 1.12: Serves a lissajous gif.
func Lissajous() {
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
		lissajous.Generate(w, opts...)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Exercise 3.4: Serve a surface svg.
func Surface() {
	http.HandleFunc("/surface", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		var opts []surface.Option
		q := r.URL.Query()
		if width, err := strconv.Atoi(q.Get("width")); err == nil {
			opts = append(opts, surface.WithWidth(width))
		}
		if height, err := strconv.Atoi(q.Get("height")); err == nil {
			opts = append(opts, surface.WithHeight(height))
		}
		if cells, err := strconv.Atoi(q.Get("cells")); err == nil {
			opts = append(opts, surface.WithCells(cells))
		}
		if xyrange, err := strconv.ParseFloat(q.Get("xyrange"), 64); err == nil {
			opts = append(opts, surface.WithXYRange(xyrange))
		}
		if angle, err := strconv.ParseFloat(q.Get("angle"), 64); err == nil {
			opts = append(opts, surface.WithAngleDegree(angle))
		}
		if color := q.Get("color"); len(color) == 6 {
			ri, err1 := strconv.ParseUint(color[0:2], 16, 8)
			gi, err2 := strconv.ParseUint(color[2:4], 16, 8)
			bi, err3 := strconv.ParseUint(color[4:6], 16, 8)
			if err1 == nil && err2 == nil && err3 == nil {
				opts = append(opts, surface.WithRed(uint8(ri)))
				opts = append(opts, surface.WithGreen(uint8(gi)))
				opts = append(opts, surface.WithBlue(uint8(bi)))
			}
		}
		if fn := q.Get("function"); fn != "" {
			opts = append(opts, surface.WithFunction(fn))
		}
		cfg := surface.NewConfig(opts...)
		cfg.CreateSVG("f", w)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
