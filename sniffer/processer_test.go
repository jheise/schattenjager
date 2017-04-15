package main

import (
	"fmt"
	"testing"
)

func TestSplitQuery(t *testing.T) {
	subdomain, domain := splitDomain("localhost")
	if subdomain != "" {
		fmt.Printf("Expecting nothing, got back ->%s<-\n", subdomain)
		t.Error("single field subdomain not spliting")
	}
	if domain != "localhost" {
		t.Error("single field domain not spliting")
	}

	subdomain, domain = splitDomain("example.com")
	if subdomain != "" {
		t.Error("two field subdomain not spliting")
	}
	if domain != "example.com" {
		t.Error("two field domain not spliting")
	}

	subdomain, domain = splitDomain("www.example.com")
	if subdomain != "www" {
		t.Error("3 field subdomain not spliting")
	}
	if domain != "example.com" {
		t.Error("3 field domain not spliting")
	}

	subdomain, domain = splitDomain("1.www.example.com")
	if subdomain != "1.www" {
		t.Error("4+ field subdomain not spliting")
	}
	if domain != "example.com" {
		t.Error("4+ field domain not spliting")
	}
}
