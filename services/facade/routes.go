/*
Facade routes specific endpoints to an external API. This supports a scenario
where the API is not fully implemented in Encore yet, but we want to expose
it to the public using Encore's API Gateway.
*/
package facade

import (
	"io"
	"net/http"

	"encore.dev/rlog"
)

const baseURL = "https://api.example.com"

var client = &http.Client{}

// Route all unmatched requests to another backend.
//
//encore:api public raw path=/!fallback
func Facade(w http.ResponseWriter, req *http.Request) {
	// Proxy the request to the provided baseURL.
	proxyURL := baseURL + req.URL.Path
	proxyReq, err := http.NewRequest(req.Method, proxyURL, req.Body)
	if err != nil {
		rlog.Error("Failed to create proxy request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Copy headers from the original request to the proxy request.
	proxyReq.Header = req.Header.Clone()

	// Create a new HTTP client and send the proxy request.
	proxyResp, err := client.Do(proxyReq)
	if err != nil {
		rlog.Error("Failed to send proxy request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer proxyResp.Body.Close()

	// Copy the response status code and headers from the proxy response to the original response.
	for key, values := range proxyResp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(proxyResp.StatusCode)

	// Copy the response body from the proxy response to the original response.
	_, err = io.Copy(w, proxyResp.Body)
	if err != nil {
		rlog.Error("Failed to copy response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
