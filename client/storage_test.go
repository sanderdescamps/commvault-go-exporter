package client_test

import (
	"testing"
)

func TestGetLibraries(t *testing.T) {

	c := testAccGenerateClient(t)

	libs, err := c.Storage.GetLibraries()
	if err != nil {
		t.Fatalf("unable to get libraries")
	}
	for _, l := range libs {
		t.Logf("library: %s (ID=%d)", l.Name, l.Id)
	}

}

func TestGetLibrariesDetails(t *testing.T) {

	c := testAccGenerateClient(t)

	libs, err := c.Storage.GetLibrariesDetails()
	if err != nil {
		t.Fatalf("unable to get libraries: \n %v", err)
	}
	for _, l := range libs {
		t.Logf("library: %s (ID=%d)", l.Library.LibraryName, l.Library.LibraryID)
	}

}
