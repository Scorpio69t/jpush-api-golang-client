## ADDED Requirements

### Requirement: Context-first Client APIs
SDK SHALL expose a unified `Client` with context-aware methods for Push/Schedule/CID/SMS and support dependency injection (custom http.Client/RoundTripper) plus per-request timeout configuration.

#### Scenario: Uses context-aware push
- **WHEN** a caller invokes `Client.Push` with a context and payload
- **THEN** the request is issued using that context and respects cancellation/timeout while authenticating via Basic Auth

#### Scenario: Injects custom transport
- **WHEN** a caller constructs `Client` with an injected `http.Client` or `RoundTripper`
- **THEN** all outbound requests use the provided transport, enabling tracing, retries, or mocking

### Requirement: Structured Responses and Errors
SDK SHALL return typed response objects and structured errors that distinguish validation, HTTP/network, and JPush API errors, while preserving HTTP status and API error codes/messages.

#### Scenario: Validation failure
- **WHEN** required payload fields (e.g., platform or audience) are missing
- **THEN** `Client.Push` returns a validation error without issuing an HTTP request

#### Scenario: API error mapping
- **WHEN** JPush returns a non-2xx response with an error payload
- **THEN** the SDK surfaces an error containing the HTTP status and parsed API error code/message

### Requirement: Wire-format Compatibility and Deprecation Shims
SDK SHALL preserve existing JSON wire format for Push/Schedule/CID/SMS requests and keep legacy string-returning methods as deprecated shims that delegate to the new APIs.

#### Scenario: Wire-compatible payload
- **WHEN** the same logical payload is marshaled via the new builders
- **THEN** the resulting JSON matches the prior wire format for equivalent fields

#### Scenario: Deprecated wrapper works
- **WHEN** a caller uses an old string-returning method
- **THEN** it internally calls the new typed API and returns the stringified response while emitting deprecation metadata in code docs
