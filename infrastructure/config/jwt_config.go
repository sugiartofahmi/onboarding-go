package config

var (
	JWTSecret    = GetRequired("JWT_SECRET")
	JWTExpiredIn = Get("JWT_EXPIRED_IN", "24h")
)
