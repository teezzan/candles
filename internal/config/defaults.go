package config

const (
	//defaultDBHost is the default database host
	defaultDBHost = "localhost"
	//defaultDBPort is the default database port
	defaultDBPort = 5432
	//defaultDBUsername is the default database username
	defaultDBUsername = "root"
	//defaultDBPassword is the default database password
	defaultDBPassword = "root"
	//defaultDBName is the default database name
	defaultDBName = "candles"

	//defaultServerHost is the default server host
	defaultServerHost = "localhost"
	//defaultServerPort is the default server port
	defaultServerPort = 8090

	//defaultDiscardInCompleteRow is the default value for discard incomplete row
	defaultDiscardInCompleteRow = false
	// defaultDataPointLimit is the default value for data point limit
	defaultDataPointLimit = 100

	//defaultS3Region is the default value for s3 region
	defaultS3Region = "eu-west-1"
	//defaultS3Bucket is the default value for s3 bucket
	defaultS3Bucket = "coiny-data-bucket"
	//defaultS3PresignURLExpiryTime is the default value for s3 presign url expiry time in hours
	defaultS3PresignURLExpiryTime = 2

	//defaultSQSRegion is the default value for sqs region
	defaultSQSRegion = "eu-west-1"
	//defaultSQSQueue is the default value for sqs queue
	defaultSQSQueue = "candle-files-notification-fifo"
	// defaultCronJobFrequencyInMinutes is the default value for cron job frequency in minutes
	defaultCronJobFrequencyInMinutes = 2
	// defaultCleanupCronJobFrequencyInDays is the default value for cleanup of stale data processing status in days
	defaultCleanupCronJobFrequencyInDays = 1
)
