package messages

import (
	"testing"
)

// TestBuildURL test with 1 not matching placeholder
func TestBuildURL(t *testing.T) {
	message := Message{
		Endpoint{
			Method: "GET",
			URL:    "http://sample_domain_endpoint.com/data?title={mascot}&image={location}&foo={bar}",
		},
		[]map[string]string{{
			"mascot":   "Gopher",
			"location": "https://blog.golang.org/gopher/gopher.png",
		}},

		// Data{
		// 	{
		// 		"mascot":   "Gopher",
		// 		"location": "https://blog.golang.org/gopher/gopher.png",
		// 	},
		// },
	}

	ProcessMessage(message)
	// var expected = "http://sample_domain_endpoint.com/data?title=Gopher&image=https%3A%2F%2Fblog.golang.org%2Fgopher%2Fgopher.png&foo="
	// assert.Equal(t,
	// 	expected,
	// 	actual)
}
