package databaseUtility

//Retrieve the DSN from parameters passed
//parameters["DB_USER"] => Database User
//parameters["DB_PASSWORD"] => Database Password
//parameters["DB_HOST"] => Database Host
//parameters["DB_NAME"] => Database Name
func RetrieveDSN(parameters map[string]string) string {
	return parameters["DB_USER"] + ":" + parameters["DB_PASSWORD"] + "@tcp" + "(" + parameters["DB_HOST"] + ")" + "/" + parameters["DB_NAME"]
}
