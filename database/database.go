package databaseUtility

import "errors"

//Retrieve the DSN from parameters passed
//parameters["DB_USER"] => Database User
//parameters["DB_PASSWORD"] => Database Password
//parameters["DB_HOST"] => Database Host
//parameters["DB_NAME"] => Database Name
func RetrieveDSN(parameters ...string) (string, error) {
	if len(parameters) != 4{
		return "", errors.New("Please enter at least 4 parameters")
	}
	return parameters[0] + ":" + parameters[1] + "@tcp" + "(" + parameters[2] + ")" + "/" + parameters[3], nil
}
