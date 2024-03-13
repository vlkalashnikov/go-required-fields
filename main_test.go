package main

import (
	"testing"
)

func TestCheckRequiredFields(t *testing.T) {
	type User struct {
		ID   int    `json:"id" required:"true"`
		Name string `json:"name" required:"true"`
		Age  int    `json:"age" required:"true"`
	}

	type User2 struct {
		ID   int    `json:"id" required:"true"`
		Name string `json:"name" required:"true"`
		Age  int    `json:"age" required:"true"`
	}

	tests := []struct {
		name string
		obj  interface{}
	}{
		{
			name: "Test 1",
			obj:  &User{ID: 1, Name: "John", Age: 20},
		},
		{
			name: "Test 2",
			obj:  &User2{ID: 1, Name: "John", Age: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckRequiredFields(tt.obj); err != nil {
				t.Errorf("CheckRequiredFields() = %v", err)
			}
		})
	}
}
