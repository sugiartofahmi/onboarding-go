package config

var (
	Storage                = Get("STORAGE", "local")
	FileUploadMaxSize      = Get("FILE_UPLOAD_MAX_SIZE", "10") // in MB
	FileUploadAllowedTypes = Get("FILE_UPLOAD_ALLOWED_TYPES", "image/jpeg,image/png")
	AWSS3Endpoint          = Get("AWS_S3_ENDPOINT", "http://localhost:9000")
	AWSS3AccessKey         = Get("AWS_S3_ACCESS_KEY", "admin")
	AWSS3SecretKey         = Get("AWS_S3_SECRET_KEY", "admin")
	AWSS3Region            = Get("AWS_S3_REGION", "us-east-1")
	AWSS3Bucket            = Get("AWS_S3_BUCKET", "event-backend")
)
