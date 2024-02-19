package environment

import (
	"github.com/spf13/pflag"
)

type Config interface {
	ReadFromEnv() error
	ReadFromFile() error
	AddFlags(set *pflag.FlagSet)
}

//var Config map[string]string = map[string]string{
//	"DB_HOST":                          getEnvDefault("DB_HOST", ""),
//	"DB_PORT":                          getEnvDefault("DB_PORT", ""),
//	"DB_USER":                          getEnvDefault("DB_USER", ""),
//	"DB_PASSWORD":                      getEnvDefault("DB_PASSWORD", ""),
//	"DB_DBNAME":                        getEnvDefault("DB_DBNAME", ""),
//	"JWT_SECRET":                       getEnvDefault("JWT_SECRET", ""),
//	"ENV":                              getEnvDefault("ENV", ""),
//	"PORT":                             getEnvDefault("PORT", ""),
//	"TWILIO_ACCOUNT_SID":               getEnvDefault("TWILIO_ACCOUNT_SID", ""),
//	"TWILIO_AUTH_TOKEN":                getEnvDefault("TWILIO_AUTH_TOKEN", ""),
//	"TWILIO_SENDER_CONTACT":            getEnvDefault("TWILIO_SENDER_CONTACT", ""),
//	"AWS_ACCESS_KEY":                   getEnvDefault("AWS_ACCESS_KEY", ""),
//	"AWS_SECRET_ACCESS_KEY":            getEnvDefault("AWS_SECRET_ACCESS_KEY", ""),
//	"AWS_REGION":                       getEnvDefault("AWS_REGION", ""),
//	"AWS_BUCKET_NAME":                  getEnvDefault("AWS_BUCKET_NAME", ""),
//	"APPLE_DEVELOPER_BUNDLE_ID":        getEnvDefault("APPLE_DEVELOPER_BUNDLE_ID", ""),
//	"APPLE_DEVELOPER_CERT_PATH":        getEnvDefault("APPLE_DEVELOPER_CERT_PATH", ""),
//	"APPLE_DEVELOPER_CERT_PASSPHASE":   getEnvDefault("APPLE_DEVELOPER_CERT_PASSPHASE", ""),
//	"APPLE_DEVELOPER_TEAM_ID":          getEnvDefault("APPLE_DEVELOPER_TEAM_ID", ""),
//	"APPLE_DEVELOPER_KEY_ID":           getEnvDefault("APPLE_DEVELOPER_KEY_ID", ""),
//	"PROFILE_PIC_PLACEHOLDER_URL":      getEnvDefault("PROFILE_PIC_PLACEHOLDER_URL", "gibberish.com"),
//	"PACT_DISPLAY_PIC_PLACEHOLDER_URL": getEnvDefault("PACT_DISPLAY_PIC_PLACEHOLDER_URL", "gibberish-pact.com"),
//	"GETSTREAM_API_KEY":                getEnvDefault("GETSTREAM_API_KEY", ""),
//	"GETSTREAM_API_SECRET":             getEnvDefault("GETSTREAM_API_SECRET", ""),
//}
