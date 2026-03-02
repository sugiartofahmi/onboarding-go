package config

import "event-backend/infrastructure/utils"

var (
	DBHost                      = GetRequired("DB_HOST")
	DBPort                      = Get("DB_PORT", "5432")
	DBName                      = GetRequired("DB_NAME")
	DBUser                      = GetRequired("DB_USER")
	DBPassword                  = GetRequired("DB_PASSWORD")
	DBTimezone                  = Get("DB_TIMEZONE", "Asia/Jakarta")
	DBMaxIdleConns              = utils.StringToInt(Get("DB_MAX_IDLE_CONNS", "5"))
	DBMaxOpenConns              = utils.StringToInt(Get("DB_MAX_OPEN_CONNS", "10"))
	DBMaxIdleConnsInMinutes     = utils.StringToInt(Get("DB_MAX_IDLE_CONNS_IN_MINUTES", "10"))
	DBMaxLifetimeConnsInMinutes = utils.StringToInt(Get("DB_MAX_LIFETIME_CONNS_IN_MINUTES", "60"))
)
