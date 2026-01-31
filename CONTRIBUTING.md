# Contributing to Marketstack Go Client

Thank you for your interest in contributing to the Marketstack Go Client! This document provides guidelines and instructions for contributing.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/marketstack-go.git
   cd marketstack-go
   ```
3. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git
- A Marketstack API key for testing (get one at [marketstack.com](https://marketstack.com/))

### Install Dependencies

```bash
go mod download
```

### Running Tests

Run all tests:
```bash
go test -v ./...
```

Run tests with coverage:
```bash
go test -v -cover ./...
```

Generate coverage report:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## Code Style

- Follow standard Go conventions and idioms
- Use `gofmt` to format your code:
  ```bash
  gofmt -w .
  ```
- Use `go vet` to check for common mistakes:
  ```bash
  go vet ./...
  ```
- Add comments to all exported types, functions, and methods
- Keep functions small and focused
- Write descriptive variable and function names

## Adding New Features

When adding a new feature:

1. **Write tests first** - Follow TDD principles
2. **Update documentation** - Add examples to README.md
3. **Maintain backwards compatibility** - Don't break existing APIs
4. **Add godoc comments** - Document all public APIs
5. **Follow existing patterns** - Look at existing code for consistency

### Adding a New Endpoint

If you're adding support for a new Marketstack API endpoint:

1. Create a new file (e.g., `newfeature.go`)
2. Define the options struct with `url` tags:
   ```go
   type NewFeatureOptions struct {
       Param1 string `url:"param1,omitempty"`
       Param2 int    `url:"param2,omitempty"`
   }
   ```
3. Define the response struct with `json` tags:
   ```go
   type NewFeatureResponse struct {
       Pagination Pagination `json:"pagination"`
       Data       []NewData  `json:"data"`
   }
   ```
4. Implement the client method:
   ```go
   func (c *Client) GetNewFeature(ctx context.Context, opts *NewFeatureOptions) (*NewFeatureResponse, error) {
       var result NewFeatureResponse
       if err := c.doRequest(ctx, "/newfeature", opts, &result); err != nil {
           return nil, err
       }
       return &result, nil
   }
   ```
5. Create comprehensive tests in `newfeature_test.go`
6. Update README.md with usage examples

## Testing Guidelines

- Write table-driven tests when appropriate
- Use `httptest` to mock API responses
- Test both success and error cases
- Verify query parameters are correctly encoded
- Check that responses are properly deserialized
- Test edge cases and boundary conditions

Example test structure:
```go
func TestNewFeature(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verify request
        if r.URL.Path != "/expected/path" {
            t.Errorf("unexpected path: %s", r.URL.Path)
        }
        
        // Return mock response
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(mockResponse)
    }))
    defer server.Close()
    
    client := NewClient("test-key", nil)
    client.SetBaseURL(server.URL)
    
    result, err := client.NewFeature(context.Background(), opts)
    
    // Assertions
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    // ... more assertions
}
```

## Documentation

- Update README.md with new features
- Add godoc comments to all exported symbols
- Include code examples in documentation
- Update CHANGELOG.md (if it exists)

## Submitting Changes

1. Ensure all tests pass:
   ```bash
   go test ./...
   ```

2. Format your code:
   ```bash
   gofmt -w .
   ```

3. Commit your changes with a descriptive message:
   ```bash
   git commit -m "Add support for new endpoint"
   ```

4. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

5. Create a Pull Request on GitHub

### Pull Request Guidelines

- Provide a clear description of the changes
- Reference any related issues
- Ensure CI tests pass
- Keep PRs focused on a single feature or fix
- Update documentation as needed

## Reporting Issues

When reporting issues, please include:

- Go version (`go version`)
- Operating system
- Clear description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Code samples if applicable

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers
- Focus on constructive feedback
- Help others learn and grow

## Questions?

If you have questions about contributing, feel free to:
- Open an issue on GitHub
- Contact the maintainer at sovletig@gmail.com

Thank you for contributing! 🎉
