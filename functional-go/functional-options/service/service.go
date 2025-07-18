package service

import (
	"errors"
	"log"
	"strconv"
	"time"
)

// Logger is a pluggable logging interface
type Logger interface {
	Info(msg string)
	Error(msg string)
}

type FeatureFlags struct {
	enableCache     bool
	enableAuditLog  bool
	enableDebugMode bool
}

// Service is our configurable target
type Service struct {
	logger   Logger
	timeout  time.Duration
	retries  int
	validate bool
	feature  FeatureFlags
}

// Option defines a functional config option
type Option func(*Service) error

// Defaults can be overridden via options
func NewService(opts ...Option) (*Service, error) {
	s := &Service{
		timeout: 10 * time.Second,
		retries: 3,
	}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	if s.logger == nil {
		return nil, errors.New("logger cannot be nil")
	}
	if s.validate && s.retries < 0 {
		return nil, errors.New("retries must be >= 0")
	}

	return s, nil
}

// --- Functional Options ---
func WithLogger(l Logger) Option {
	return func(s *Service) error {

		s.logger = l
		return nil
	}
}

func WithTimeout(t time.Duration) Option {
	return func(s *Service) error {
		s.timeout = t
		return nil
	}
}

func WithRetries(n int) Option {
	return func(s *Service) error {
		s.retries = n
		return nil
	}
}

func WithProductionDefaults() Option {
	return func(s *Service) error {
		s.timeout = 30 * time.Second
		s.retries = 5
		return nil
	}
}

func WithTestingDefaults() Option {
	return func(s *Service) error {
		s.timeout = 1 * time.Second
		s.retries = 0
		return nil
	}
}

func WithValidation(enable bool) Option {
	return func(s *Service) error {
		s.validate = enable
		return nil
	}
}

func WithFeature(name string, on bool) Option {
	return func(s *Service) error {
		switch name {
		case "cache":
			s.feature.enableCache = on
		case "audit":
			s.feature.enableAuditLog = on
		case "debug":
			s.feature.enableDebugMode = on
		}
		return nil
	}
}

// --- Example Logic ---
func (s *Service) Run() {
	s.logger.Info("Running with timeout: " + s.timeout.String())
	s.logger.Info("Max retries: " + strconv.Itoa(s.retries))

	if s.feature.enableDebugMode {
		s.logger.Info("Debug mode in ON")
	}

	if s.feature.enableCache {
		s.logger.Info("cache mode in ON")
	}

	if s.feature.enableAuditLog {
		s.logger.Info("audit mode in ON")
	}
}

// --- Default logger ---

type defaultLogger struct{}

func (d defaultLogger) Info(msg string)  { log.Println("[INFO]", msg) }
func (d defaultLogger) Error(msg string) { log.Println("[ERROR]", msg) }
