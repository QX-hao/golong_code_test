package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func repositoryFindUser(ctx context.Context, userID string, cost time.Duration) (string, error) {
	select {
	case <-time.After(cost):
		return "user " + userID, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func serviceGetUser(ctx context.Context, userID string, cost time.Duration) (string, error) {
	return repositoryFindUser(ctx, userID, cost)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	cost, err := time.ParseDuration(r.URL.Query().Get("delay"))
	if err != nil {
		http.Error(w, "bad delay", http.StatusBadRequest)
		return
	}

	user, err := serviceGetUser(ctx, "42", cost)
	if err != nil {
		http.Error(w, "query canceled: "+err.Error(), http.StatusGatewayTimeout)
		return
	}

	fmt.Fprintln(w, user)
}

func request(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("%s -> %s", resp.Status, body)
}

func main() {
	server := httptest.NewServer(http.HandlerFunc(userHandler))
	defer server.Close()

	request(server.URL + "/users/42?delay=80ms")
	request(server.URL + "/users/42?delay=500ms")
}
