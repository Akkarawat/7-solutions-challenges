package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	overlappedbinarytreemaxsum "7-solutions-challenges/internal/overlapped_binary_tree_max_sum"
)

func main() {
    jsonURL := "https://raw.githubusercontent.com/7-solutions/backend-challenge/refs/heads/main/files/hard.json"

	// Fetch and parse JSON
	matrix, err := fetchJSON(jsonURL)
	if err != nil {
		fmt.Println("Error fetching JSON:", err)
		return
	}

	maxSum := overlappedbinarytreemaxsum.FindFastBTreeMaxSum(matrix)
	fmt.Println("Max Sum:", maxSum)
}

func fetchJSON(url string) ([][]int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JSON: %w", err)
	}
	defer resp.Body.Close() // Close response body to prevent leaks

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var data [][]int
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}