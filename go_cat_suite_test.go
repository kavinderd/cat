package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGoCat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoCat Suite")
}
