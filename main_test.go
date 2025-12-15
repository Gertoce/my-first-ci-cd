package main

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestAddNegative(t *testing.T) {
	result := Add(-1, -1)
	expected := -2
	
	if result != expected {
		t.Errorf("Add(-1, -1) = %d; want %d", result, expected)
	}
}

func TestGetPlatformInfo(t *testing.T) {
	info := GetPlatformInfo()
	
	// Проверяем, что строка содержит ожидаемые части
	if !strings.Contains(info, "OS:") {
		t.Errorf("Platform info should contain 'OS:', got: %s", info)
	}
	
	if !strings.Contains(info, "Architecture:") {
		t.Errorf("Platform info should contain 'Architecture:', got: %s", info)
	}
	
	// Выводим информацию для отладки (будет видно в логах CI)
	t.Logf("Running test on: %s", info)
}
