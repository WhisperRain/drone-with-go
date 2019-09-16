package main

import (
	"fmt"
	"math"
	"os"
	"testing"
	"time"
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

func TestMaxTimeOfCI(t *testing.T) {
	t.Log("begin")

	for  i:=0;i<math.MaxInt32;i++ {
		time.Sleep(time.Second*1)
		t.Log("the ",i," second")
	}

	fmt.Println("hello world 2")
	if HelloWorld() != "hello world" {
		t.Errorf("got %s expected %s", HelloWorld(), "hello world")
	}
}
