package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shirou/gopsutil/mem"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	// First END POINT
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		m, _ := mem.VirtualMemory()
		free := fmt.Sprintf("%v", m.Available)
		w.Write([]byte(free))
	})

	//	Second END POINT
	http.HandleFunc("/fact", func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("n")
		n, _ := strconv.Atoi(s)
		res := fact(n)

		fmt.Fprintf(w, "result : %d", res)
	})

	http.ListenAndServe("0.0.0.0:8000", nil)
}
