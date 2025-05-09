# Go Playtomic API Client

A Go client library for interacting with the [Playtomic](https://playtomic.io) API - the sports facility booking system.

## Project Status

**Pre-1.0 Software**: This library is under active development and has not reached a major version release yet. As per semantic versioning practices, minor version releases (0.x.y) may include backward incompatible changes until we reach version 1.0.0.

## Features

- Coverage of some of the Playtomic API endpoints (WIP)
- Simple and intuitive API
- Strongly typed request and response models with context-aware request handling
- Customizable request timeouts and retries
- Error handling with detailed API error information

## Installation

```bash
go get github.com/rafa-garcia/go-playtomic-api
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rafa-garcia/go-playtomic-api/client"
	"github.com/rafa-garcia/go-playtomic-api/models"
)

func main() {
	// Create a client with custom options
	c := client.NewClient(
		client.WithTimeout(10 * time.Second),
		client.WithRetries(3),
	)

	// Set up search parameters
	params := &models.SearchClassesParams{
		Sort:          "start_date,ASC",
		Status:        "PENDING,IN_PROGRESS",
		TenantIDs:     []string{"tenant-id-1", "tenant-id-2"},
		FromStartDate: time.Now().Format("2006-01-02") + "T00:00:00",
	}

	// Fetch classes
	ctx := context.Background()
	classes, err := c.GetClasses(ctx, params)
	if err != nil {
		log.Fatalf("Error fetching classes: %v", err)
	}

	// Display results
	for _, class := range classes {
		fmt.Printf("Class: %s at %s (%s)\n", 
			class.CourseSummary.Name,
			class.Tenant.TenantName,
			class.StartDate)
	}
}
```

## Client Configuration

The client can be customized with various options:

```go
client := client.NewClient(
    // Set a custom base URL (useful for testing)
    client.WithBaseURL("https://api.playtomic.io/v1"),
    
    // Set HTTP client timeout
    client.WithTimeout(15 * time.Second),
    
    // Configure request retries
    client.WithRetries(3),
    
    // Enable debug logging
    client.WithDebug(true),
    
    // Set custom User-Agent
    client.WithUserAgent("MyApp/1.0"),
    
    // Use a custom HTTP client
    client.WithHTTPClient(customHTTPClient),
)
```

## API Documentation

For detailed information about API endpoints, parameters, and examples, see:

- [Endpoint Documentation](./docs/endpoints.md) - Complete details on all supported API endpoints
- [Examples](./examples) - Code examples showing usage patterns

## Error Handling

The client provides detailed error handling:

```go
classes, err := client.GetClasses(ctx, params)
if err != nil {
    // Check if it's an API error
    if apiErr, ok := err.(*client.APIError); ok {
        fmt.Printf("API Error: %s (Status: %d)\n", apiErr.Message, apiErr.StatusCode)
        
        // Access details from the error response
        if details, ok := apiErr.Details["more_info"]; ok {
            fmt.Printf("Additional info: %v\n", details)
        }
    } else {
        // Handle network/client errors
        fmt.Printf("Request error: %v\n", err)
    }
}
```

## Examples

See the [examples](./examples) directory for more usage examples.

## Contributing

Contributions to improve go-playtomic-api are welcome! Here's how you can help:

1. **Report Issues**: File bugs or feature requests on the issue tracker
2. **Suggest Improvements**: Propose changes through pull requests
3. **Add Endpoints**: Implement support for new Playtomic API endpoints
4. **Improve Documentation**: Help keep docs clear, accurate and up-to-date

To contribute code:

```bash
# Clone the repository
git clone https://github.com/rafa-garcia/go-playtomic-api.git

# Create a feature branch
git checkout -b my-new-feature

# Make your changes and commit
git commit -am 'Add new feature'

# Push to your fork
git push origin my-new-feature

# Create a Pull Request
```

Please include tests and documentation with your changes!

## License

MIT License - see [LICENSE](./LICENSE) file for details.
