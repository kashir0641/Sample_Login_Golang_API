package common


/*
Database and server details.
Always fetch this details from properties or yml file.
*/
const (
	MongoUrl      = "127.0.0.1"
	Database      = "golangDB"
	UserCollecion = "user"
	ServerAddress = "127.0.0.1:8080"
)

/*
Unique key for JWT Token.
Do not store credentials inside projects.
*/
var JWT_KEY = []byte("pwd@123")
