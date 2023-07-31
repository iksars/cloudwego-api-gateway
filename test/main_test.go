package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func BenchmarkApiGateWayTest01(b *testing.B) {
	requestBody := []byte(`{"id": 102, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"],"gender":"male"}`)

	// Run the benchmark b.N times
	for i := 0; i < b.N; i++ {

		resp, err := http.Post("http://127.0.0.1:8888/agw/KitexServer/Register", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			b.Fatalf("Failed to send request: %s", err)
		}

		// Read and close the response body
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			b.Fatalf("Failed to read response: %s", err)
		}
		resp.Body.Close()
	}
}

func BenchmarkApiGateWayTest02(b *testing.B) {
	url := "http://127.0.0.1:8888/query?id=201"

	// Run the benchmark b.N times
	for i := 0; i < b.N; i++ {
		// Send the GET request to your server
		resp, err := http.Get(url)
		if err != nil {
			b.Fatalf("Failed to send request: %s", err)
		}

		// Close the response body
		resp.Body.Close()
	}
}

func BenchmarkApiGateWayTestWithParallel(b *testing.B) {

	b.SetParallelism(10)

	requestBody := []byte(`{"id": 102, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"],"gender":"male"}`)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < b.N; i++ {

				resp, err := http.Post("http://127.0.0.1:8888/agw/KitexServer/Register", "application/json", bytes.NewBuffer(requestBody))
				if err != nil {
					b.Fatalf("Failed to send request: %s", err)
				}

				// Read and close the response body
				_, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					b.Fatalf("Failed to read response: %s", err)
				}
				resp.Body.Close()
			}
		}
	})

}
