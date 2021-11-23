package utils

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

const (
	a = 0x00001
	b = 0x00010
	c = 0x00100
)

func TestArrayFlip(t *testing.T) {
	////flip := []uint64{1, 2, 3, 4}
	////out := ArrayFlip(flip)
	////t.Logf("out：%v\n", out)
	//
	//d := a & b & c
	//t.Log(d)

	list := []int{1, 2, 3}
	for _, v := range list {
		go func() {
			fmt.Printf("%d\n", v)
		}()
		time.Sleep(1 * time.Second)
	}
}

func TestArrayStrings(t *testing.T) {
	a := []uint64{1, 2, 3, 4, 5}
	t.Log(a)
	t.Log(reflect.TypeOf(a).String())
	b := ArrayStrings(a)
	t.Log(b)
	t.Log(reflect.TypeOf(b).String())
}

func TestArrayUnique(t *testing.T) {
	a := []uint64{1, 2, 3, 4, 5, 5, 5}
	b := ArrayUnique(a)
	t.Log(b)

	for _, v := range b {
		fmt.Printf(reflect.TypeOf(v).String())
	}
}
