package config

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// ValidateConfig validates all configuration values and returns detailed errors
func ValidateConfig(cfg *Config) error {
	// Run all validation checks
	checks := []func(*Config) error{
		validateHTTPPort,
		validateGRPCPort,
		validatePortConflict,
		validateHTTPAddress,
		validateEnvironment,
		validateJWTSecret,
		validateJWTExpiry,
		validateDatabaseDSN,
		validateRedisDSN,
		validateRedisDB,
		validateRedisPrefix,
		validateRedisTTL,
	}

	for _, check := range checks {
		if err := check(cfg); err != nil {
			return err
		}
	}

	// Log warnings for non-critical issues
	validateWarnings(cfg)

	return nil
}

// validateHTTPPort validates HTTP port is in valid range
func validateHTTPPort(cfg *Config) error {
	if cfg.HTTPPort <= 0 || cfg.HTTPPort > 65535 {
		return fmt.Errorf(
			"invalid HTTP_PORT: %d. Expected value between 1 and 65535. "+
				"Set APP_HTTP_PORT environment variable",
			cfg.HTTPPort,
		)
	}
	return nil
}

// validateGRPCPort validates GRPC port is in valid range
func validateGRPCPort(cfg *Config) error {
	if cfg.GRPCPort <= 0 || cfg.GRPCPort > 65535 {
		return fmt.Errorf(
			"invalid GRPC_PORT: %d. Expected value between 1 and 65535. "+
				"Set APP_GRPC_PORT environment variable",
			cfg.GRPCPort,
		)
	}
	return nil
}

// validatePortConflict validates that HTTP and GRPC ports are different
func validatePortConflict(cfg *Config) error {
	if cfg.HTTPPort == cfg.GRPCPort {
		return fmt.Errorf(
			"port conflict: HTTP_PORT (%d) and GRPC_PORT (%d) cannot be the same. "+
				"Set different values for APP_HTTP_PORT and APP_GRPC_PORT",
			cfg.HTTPPort,
			cfg.GRPCPort,
		)
	}
	return nil
}

// validateHTTPAddress validates HTTP address is not empty and is valid
func validateHTTPAddress(cfg *Config) error {
	if cfg.HTTPAddress == "" {
		return fmt.Errorf(
			"HTTP_ADDRESS is empty. " +
				"Set APP_HTTP_ADDRESS environment variable (e.g., 127.0.0.1 or 0.0.0.0)",
		)
	}

	// Validate it's a valid IP address
	if ip := net.ParseIP(cfg.HTTPAddress); ip == nil && cfg.HTTPAddress != "localhost" {
		return fmt.Errorf(
			"invalid HTTP_ADDRESS: %q. Expected valid IP address or 'localhost'. "+
				"Set APP_HTTP_ADDRESS environment variable",
			cfg.HTTPAddress,
		)
	}

	return nil
}

// validateEnvironment validates ENV is set and has valid value
func validateEnvironment(cfg *Config) error {
	if cfg.ENV == "" {
		return fmt.Errorf(
			"ENV is empty. " +
				"Set APP_ENV environment variable to one of: development, test, production",
		)
	}

	validEnvs := []string{"development", "test", "production"}
	isValid := false
	for _, valid := range validEnvs {
		if cfg.ENV == valid {
			isValid = true
			break
		}
	}

	if !isValid {
		return fmt.Errorf(
			"invalid ENV: %q. Expected one of: development, test, production. "+
				"Set APP_ENV environment variable",
			cfg.ENV,
		)
	}

	return nil
}

// validateJWTSecret validates JWT secret is not empty
func validateJWTSecret(cfg *Config) error {
	if cfg.JWTSecret == "" {
		return fmt.Errorf(
			"JWT_SECRET is empty. " +
				"Set APP_JWT_SECRET environment variable (must be at least 32 characters)",
		)
	}

	// Warn if secret is too short (less than 32 characters)
	if len(cfg.JWTSecret) < 32 {
		log.Printf("⚠️  WARNING: JWT_SECRET is too short (%d chars). "+
			"Recommended length: 32+ characters\n", len(cfg.JWTSecret))
	}

	return nil
}

// validateJWTExpiry validates JWT expiry hours is greater than 0
func validateJWTExpiry(cfg *Config) error {
	if cfg.JWTExpiryHours <= 0 {
		return fmt.Errorf(
			"invalid JWT_EXPIRY_HOURS: %d. Expected value greater than 0. "+
				"Set APP_JWT_EXPIRY_HOURS environment variable",
			cfg.JWTExpiryHours,
		)
	}

	// Warn if expiry is very long
	if cfg.JWTExpiryHours > 720 { // 30 days
		log.Printf("⚠️  WARNING: JWT_EXPIRY_HOURS is very long (%d hours / %d days). "+
			"Recommended: 24-72 hours for security\n",
			cfg.JWTExpiryHours,
			cfg.JWTExpiryHours/24,
		)
	}

	return nil
}

// validateDatabaseDSN validates database DSN is not empty and valid
func validateDatabaseDSN(cfg *Config) error {
	if cfg.Database.DSN == "" {
		return fmt.Errorf(
			"DATABASE_DSN is empty. " +
				"Set APP_DATABASE_DSN environment variable " +
				"(e.g., postgresql://user:pass@localhost:5432/database?sslmode=disable)",
		)
	}

	// Validate DSN format
	if !strings.HasPrefix(cfg.Database.DSN, "postgresql://") &&
		!strings.HasPrefix(cfg.Database.DSN, "postgres://") &&
		!strings.HasPrefix(cfg.Database.DSN, "mysql://") {
		return fmt.Errorf(
			"invalid DATABASE_DSN format: must start with postgresql://, postgres://, or mysql://. "+
				"Provided: %q",
			cfg.Database.DSN,
		)
	}

	return nil
}

// validateRedisDSN validates Redis DSN is not empty
func validateRedisDSN(cfg *Config) error {
	if cfg.Redis.DSN == "" {
		return fmt.Errorf(
			"REDIS_DSN is empty. " +
				"Set APP_REDIS_DSN environment variable (e.g., localhost:6379)",
		)
	}

	// Basic format check (host:port)
	parts := strings.Split(cfg.Redis.DSN, ":")
	if len(parts) < 2 {
		return fmt.Errorf(
			"invalid REDIS_DSN format: expected host:port. "+
				"Provided: %q",
			cfg.Redis.DSN,
		)
	}

	return nil
}

// validateRedisDB validates Redis DB is between 0 and 15
func validateRedisDB(cfg *Config) error {
	if cfg.Redis.DB < 0 || cfg.Redis.DB > 15 {
		return fmt.Errorf(
			"invalid REDIS_DB: %d. Expected value between 0 and 15. "+
				"Set APP_REDIS_DB environment variable",
			cfg.Redis.DB,
		)
	}
	return nil
}

// validateRedisPrefix validates Redis prefix is not empty
func validateRedisPrefix(cfg *Config) error {
	if cfg.Redis.Prefix == "" {
		return fmt.Errorf(
			"REDIS_PREFIX is empty. " +
				"Set APP_REDIS_PREFIX environment variable (e.g., go-simple)",
		)
	}

	// Validate prefix doesn't have invalid characters
	if strings.Contains(cfg.Redis.Prefix, " ") {
		return fmt.Errorf(
			"invalid REDIS_PREFIX: %q. Cannot contain spaces. "+
				"Set APP_REDIS_PREFIX environment variable",
			cfg.Redis.Prefix,
		)
	}

	return nil
}

// validateRedisTTL validates Redis default TTL is greater than 0
func validateRedisTTL(cfg *Config) error {
	if cfg.Redis.DefaultTTL <= 0 {
		return fmt.Errorf(
			"invalid REDIS_DEFAULT_TTL: %d. Expected value greater than 0 (in minutes). "+
				"Set APP_REDIS_DEFAULT_TTL environment variable",
			cfg.Redis.DefaultTTL,
		)
	}
	return nil
}

// validateWarnings logs non-critical warnings for configuration
func validateWarnings(cfg *Config) {
	// Warn about default JWT secret in production
	if cfg.IsProduction() {
		if cfg.JWTSecret == "this-is-a-secret-key" {
			log.Printf("⚠️  SECURITY WARNING: Using default JWT_SECRET in production! " +
				"Set APP_JWT_SECRET to a strong random secret (use: openssl rand -base64 32)\n")
		}

		// Warn about localhost address in production
		if cfg.HTTPAddress == "127.0.0.1" {
			log.Printf("⚠️  WARNING: HTTP_ADDRESS is 127.0.0.1 in production. " +
				"Set APP_HTTP_ADDRESS to 0.0.0.0 or use a reverse proxy\n")
		}

		// Warn about default database in production
		if strings.Contains(cfg.Database.DSN, "localhost") {
			log.Printf("⚠️  WARNING: DATABASE_DSN points to localhost in production. " +
				"Set APP_DATABASE_DSN to production database server\n")
		}

		// Warn about Redis on localhost in production
		if strings.Contains(cfg.Redis.DSN, "localhost") {
			log.Printf("⚠️  WARNING: REDIS_DSN points to localhost in production. " +
				"Set APP_REDIS_DSN to production Redis server\n")
		}
	}

	// Warn about short JWT expiry
	if cfg.JWTExpiryHours < 1 {
		log.Printf("⚠️  WARNING: JWT_EXPIRY_HOURS is very short (%d hours). "+
			"Recommended minimum: 1 hour\n", cfg.JWTExpiryHours)
	}
}
