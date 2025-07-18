# 🧩 Functional Options Pattern in Go

Functional options are a clean way to avoid bloated constructors with many parameters.

---

## Why Use It?

- Makes code flexible and readable
- Adds default values easily
- Keeps your constructor minimal

---


## Anti-patterns to Avoid
### 1. Using nils in options : Always check and reject nil values
```go
func WithLogger(l Logger) Option {
	return func(s *Service) {
		if l == nil {
			panic("logger cannot be nil") // or log/fallback
		}
		s.logger = l
	}
}
```
### 2. Making reuired fields optional : Use constructor for reuired fields
```go
func NewService(url string, opts ...Option) *Service {
	if url == "" {
		panic("url is required")
	}
	s := &Service{url: url}
	// apply options...
	return s
}
```
### 3. Interdependent options : Avoid order dependencies. If needed, validate after all options are applied
```go
func NewService(opts ...Option) (*Service, error) {
	s := &Service{...}
	for _, opt := range opts {
		opt(s)
	}

	if s.logger == nil {
		return nil, errors.New("logger is required")
	}
	return s, nil
}
```
### 4. Option logic doing too much : Options should just assign passed-in values. Keep side effects outside.
```go
func WithAutoConfigureDB() Option {
	return func(s *Service) {
		conn := connectToDB() // ❌ side effect
		s.db = conn
	}
}

///better
conn := connectToDB()
svc := NewService(WithDB(conn))
```
### 5. Too many options : Group into config bundles
```go
svc := NewService(
	WithLogger(...),
	WithMonitoringAndTracing(...),
)
```
