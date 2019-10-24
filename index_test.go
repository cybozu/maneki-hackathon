package main

import "testing"

func TestCreateIndex(t *testing.T) {
	result := createIndex("testIndex")
	if result != "success" {
		t.Errorf("Create Index Failed, expected %v, got %v", "success", result)
	}
}
