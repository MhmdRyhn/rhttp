package rhttp

type RequestOptions struct {
	// Common part of the urls. For example, if you set the `BaseUrl` value as
	// `https://myapi.com`, then to send request to the API "https://myapi.com/api/path",
	// only the value of the path, i.e., "api/path" is needed to supply to the
	// corresponding request method. If it's value is empty string, then the full
	// URL must be provided to the corresponding request.
	BaseUrl string

	// API Key used by the corresponding API.
	APIKey *string

	// Function that performs the authentication using the username and password.
	// The function takes the username and password as input and returns the
	// auth token (usually the access token).
	BasicAuthFunc *func(string, string) string

	// Retry configuration
	Retry *RetryConfig

	// Number of nanoseconds the client waits for the reaponse.
	TimeoutNanoseconds *int
}

type RetryConfig struct {
	// Maximum number of times to retry when the API respond with error.
	MaxAttempt int

	// Number of nanoseconds to wait before successive retry.
	IntervalNanoseconds int

	// Multiplication factor used to increase the current retry interval from the
	// previous retry interval. If the the request needs to be configured as simple
	// request, it's value must be set as 1.0. If the value is less than or equal
	// to 0.0, then the request is considered as simple request.
	BackoffRate *float64
}
