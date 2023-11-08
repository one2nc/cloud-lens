package internal

// ContextKey represents context key.
type ContextKey string

// A collection of context keys.
const (
	KeySelectedCloud     ContextKey = "seleted_cloud"
	KeyFactory           ContextKey = "factory"
	KeyApp               ContextKey = "app"
	KeyActiveProfile     ContextKey = "active-profile"
	KeyActiveRegion      ContextKey = "active-region"
	KeyActiveProject     ContextKey = "active-project"
	KeySession           ContextKey = "session"
	BucketName           ContextKey = "bucket_name"
	StorageBucketName    ContextKey = "storage_bucket_name"
	ObjectName           ContextKey = "object_name"
	FolderName           ContextKey = "folder_name"
	KeyAliases           ContextKey = "aliases"
	UserName             ContextKey = "user_name"
	GroupName            ContextKey = "group_name"
	RoleName             ContextKey = "role_name"
	VpcId                ContextKey = "vpc_id"
	LowercaseY           string     = "y"
	UppercaseY           string     = "Y"
	LowercaseYes         string     = "yes"
	UppercaseYes         string     = "YES"
	LowercaseN           string     = "n"
	UppercaseN           string     = "N"
	LowercaseNo          string     = "no"
	UppercaseNo          string     = "NO"
	LowercaseEc2         string     = "ec2"
	UppercaseEc2         string     = "Ec2"
	LowercaseS3          string     = "s3"
	UppercaseS3          string     = "S3"
	LowercaseEBS         string     = "ebs"
	UppercaseEBS         string     = "EBS"
	LowercaseSg          string     = "sg"
	UppercaseSg          string     = "SG"
	LowercaseIamUser     string     = "iam:u"
	UppercaseIamUser     string     = "IAM:U"
	LowercaseIam         string     = "iam"
	UppercaseIam         string     = "IAM"
	LowercaseIamGroup    string     = "iam:g"
	UppercaseIamGroup    string     = "IAM:g"
	LowercaseIamRole     string     = "iam:r"
	UppercaseIamRole     string     = "IAM:R"
	LowercaseEc2Snapshot string     = "ec2:S"
	UppercaseEc2Snapshot string     = "Ec2:S"
	LowercaseEc2Image    string     = "ec2:i"
	UppercaseEc2Image    string     = "Ec2:I"
	LowercaseSQS         string     = "sqs"
	UppercaseSQS         string     = "SQS"
	LowercaseVPC         string     = "vpc"
	UppercaseVPC         string     = "VPC"
	LowercaseSubnet      string     = "subnet"
	UppercaseSubnet      string     = "SUBNET"
	LowercaseLamda       string     = "lambda"
	UppercaseLamda       string     = "LAMBDA"
	LowercaseStorage     string     = "storage"
	UppercaseStorage     string     = "STORAGE"
	Help                 string     = "help"
	LowercaseH           string     = "h"
	QuestionMark         string     = "?"
	Quit                 string     = "quit"
	LowercaseQ           string     = "q"
	UppercaseQ           string     = "Q"
	QFactorial           string     = "q!"
	Aliases              string     = "aliases"
	Alias                string     = "alias"
	LowercaseA           string     = "a"
	Object               string     = "OBJ"
	StorageObject        string     = "STORAGE_OBJ"
	UserPolicy           string     = "User Policy"
	UserGroupPolicy      string     = "User Group Policy"
	RolePolicy           string     = "Role Policy"
	GroupUsers           string     = "Group Users"
)

const (
	AWS_SCREEN    string = "AWS"
	GCP_SCREEN    string = "GCP"
	MAIN_SCREEN   string = "MAIN"
	SPLASH_SCREEN string = "SPLASH"
)
const (
	AWS string = "AWS"
	GCP string = "GCP"
)

const (
	AWS_PROFILE        string = "AWS_PROFILE"
	AWS_DEFAULT_REGION string = "AWS_DEFAULT_REGION"
)

const (
	FOLDER_TYPE string = "Folder"
	FILE_TYPE   string = "File"
	NONE        string = "-"
)

const (
	LOCALSTACK_PORT                string = "LOCALSTACK_PORT"
	GOOGLE_APPLICATION_CREDENTIALS string = "GOOGLE_APPLICATION_CREDENTIALS"
)
