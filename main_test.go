package main

import (
	. "qiniu.com/mandycode/localone"
	"testing"
)

func TestLocalOne(t *testing.T) {
	LocalOne(1)
	LocalOne(2)
}
