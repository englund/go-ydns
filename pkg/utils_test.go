package pkg

import (
	"os"
	"testing"
)

func createTempFileWithIp(t *testing.T, ip string) *os.File {
	t.Helper()
	file, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	if _, err := file.WriteString(ip); err != nil {
		t.Fatalf("Failed to write IP address to file: %v", err)
	}
	return file
}

func TestReadIpFromFile(t *testing.T) {
	ip := "192.168.0.1"
	file := createTempFileWithIp(t, ip)
	defer os.Remove(file.Name())

	readIp, err := ReadIpFromFile(file.Name())
	if err != nil {
		t.Fatalf("Failed to read IP address from file: %v", err)
	}

	if readIp != ip {
		t.Errorf("Expected IP address %q, but got %q", ip, readIp)
	}
}

func TestReadIpFromFile_FileDoesNotExist(t *testing.T) {
	ip, err := ReadIpFromFile("does-not-exist")
	if err != nil {
		t.Fatalf("Failed to read IP address from file: %v", err)
	}

	if ip != "" {
		t.Errorf("Expected empty IP address, but got %q", ip)
	}
}

func TestWriteIpToFile(t *testing.T) {
	ip := "192.168.0.1"
	file := createTempFileWithIp(t, ip)
	defer os.Remove(file.Name())

	if err := WriteIpToFile(file.Name(), ip); err != nil {
		t.Fatalf("Failed to write IP address to file: %v", err)
	}

	readIp, err := ReadIpFromFile(file.Name())
	if err != nil {
		t.Fatalf("Failed to read IP address from file: %v", err)
	}

	if readIp != ip {
		t.Errorf("Expected IP address %q, but got %q", ip, readIp)
	}
}

func TestWriteIpToFile_FileDoesNotExist(t *testing.T) {
	ip := "192.168.0.1"
	err := WriteIpToFile("does-not-exist", ip)
	if err != nil {
		t.Fatalf("Failed to write IP address to file: %v", err)
	}
	defer os.Remove("does-not-exist")
}

func TestWriteIpToFile_FileAlreadyExists(t *testing.T) {
	ip := "192.168.0.1"
	file := createTempFileWithIp(t, ip)
	defer os.Remove(file.Name())

	newIp := "192.168.0.2"
	if err := WriteIpToFile(file.Name(), newIp); err != nil {
		t.Fatalf("Failed to write IP address to file: %v", err)
	}

	readIp, err := ReadIpFromFile(file.Name())
	if err != nil {
		t.Fatalf("Failed to read IP address from file: %v", err)
	}

	if readIp != newIp {
		t.Errorf("Expected IP address %q, but got %q", newIp, readIp)
	}
}
