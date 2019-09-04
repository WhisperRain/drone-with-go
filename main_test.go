package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "hello world" {
		t.Errorf("got %s expected %s", HelloWorld(), "hello world")
	}
}


func TestHelloWorld2(t *testing.T) {
	fmt.Println("hello world 2")
	if HelloWorld() != "hello world" {
		t.Errorf("got %s expected %s", HelloWorld(), "hello world")
	}
}
