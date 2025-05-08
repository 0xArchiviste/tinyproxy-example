package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"os"
)



type ProxyConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

func readProxyConfig(filename string) (ProxyConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return ProxyConfig{}, fmt.Errorf("failed to open proxy config file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Read header
	_, err = reader.Read()
	if err != nil {
		return ProxyConfig{}, fmt.Errorf("failed to read CSV header: %w", err)
	}
	
	// Read first proxy entry
	record, err := reader.Read()
	if err != nil {
		return ProxyConfig{}, fmt.Errorf("failed to read proxy config: %w", err)
	}
	
	if len(record) < 4 {
		return ProxyConfig{}, fmt.Errorf("invalid proxy config format")
	}
	
	return ProxyConfig{
		Username: record[0],
		Password: record[1],
		Host:     record[2],
		Port:     record[3],
	}, nil
}

func main() {
	// Read proxy configuration
	proxyConfig, err := readProxyConfig("proxies.csv")
	if err != nil {
		log.Fatalf("Error reading proxy configuration: %v", err)
	}
	
	// Create proxy URL
	proxyURL := fmt.Sprintf("http://%s:%s@%s:%s", 
		proxyConfig.Username, 
		proxyConfig.Password, 
		proxyConfig.Host, 
		proxyConfig.Port)
	
	fmt.Printf("Using proxy: %s\n", strings.Replace(proxyURL, proxyConfig.Password, "****", 1))
	
	// Create HTTP client with proxy
	proxyURLParsed, err := url.Parse(proxyURL)
	if err != nil {
		log.Fatalf("Error parsing proxy URL: %v", err)
	}
	
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURLParsed),
		},
	}
	

	serverURL := "http://SERVER:8080/ipcheck"
	fmt.Printf("Making request to: %s\n", serverURL)
	
	resp, err := client.Get(serverURL)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()
	
	// Read and print response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	
	fmt.Println("Server response:")
	fmt.Println(string(body))
}

