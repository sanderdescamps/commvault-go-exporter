package client_test

import (
	"os"
	"strconv"
	"testing"

	"gitlab.global.ingenico.com/sdescamps/commvault-go-exporter/client"
)

func testAccPreChecks(t *testing.T) {

	target := os.Getenv("COMMVAULT_TARGET")
	username := os.Getenv("COMMVAULT_USERNAME")
	password := os.Getenv("COMMVAULT_PASSWORD")
	if target == "" {
		t.Fatalf("COMMVAULT_TARGET must be set for acceptance tests")
	}

	if (username == "") || (password == "") {
		t.Fatalf("COMMVAULT_USERNAME and COMMVAULT_PASSWORD must be set for acceptance tests")
	}
}

func testAccGenerateClient(t *testing.T) *client.CommvaultClient {

	username := os.Getenv("COMMVAULT_USERNAME")
	password := os.Getenv("COMMVAULT_PASSWORD")
	target := os.Getenv("COMMVAULT_TARGET")
	insecure := false
	if i := os.Getenv("COMMVAULT_INSECURE"); i != "" {
		insecure, _ = strconv.ParseBool(i)
	}

	c, err := client.NewClient(target, username, password, insecure)
	if err != nil {
		t.Fatalf("error setting up client: %s", err)
	}
	return c
}

func TestAccClient(t *testing.T) {
	testAccPreChecks(t)
	testAccGenerateClient(t)
}

// Test that a NewClient call with no authentication returns an error
func TestNewClientNoAuth(t *testing.T) {
	insecure := false
	if i := os.Getenv("COMMVAULT_INSECURE"); i != "" {
		insecure, _ = strconv.ParseBool(i)
	}
	_, err := client.NewClient("target", "", "", insecure)
	if err == nil {
		t.Errorf("An Error was NOT raised when no authentication methods provided")
	}
}
