package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBuildURL test with 1 not matching placeholder
func TestBuildURL(t *testing.T) {
	var originalURL = "http://sample_domain_endpoint.com/data?title={mascot}&image={location}&foo={bar}"
	// var data []map[string]string
	var data = []map[string]string{{
		"mascot":   "Gopher",
		"location": "https://blog.golang.org/gopher/gopher.png",
	}}

	var actual = BuildURL(originalURL, data)
	var expected = "http://sample_domain_endpoint.com/data?title=Gopher&image=https%3A%2F%2Fblog.golang.org%2Fgopher%2Fgopher.png&foo="
	assert.Equal(t,
		expected,
		actual)
}

// TestBuildURL2 test with more than 1 not matching placeholders
func TestBuildURL2(t *testing.T) {
	var originalURL = "http://sample_domain_endpoint.com/data?zoo={car}&title={mascot}&image={location}&foo={bar}"
	// var data []map[string]string
	var data = []map[string]string{{
		"mascot":   "Gopher",
		"location": "https://blog.golang.org/gopher/gopher.png",
	}}

	var actual = BuildURL(originalURL, data)
	var expected = "http://sample_domain_endpoint.com/data?zoo=&title=Gopher&image=https%3A%2F%2Fblog.golang.org%2Fgopher%2Fgopher.png&foo="
	assert.Equal(t,
		expected,
		actual)
}
