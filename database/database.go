package databaseUtility

import "errors"

//Retrieve the DSN from parameters passed
func RetrieveDSN(parameters ...string) (string, error) {
	if len(parameters) != 4 {
		return "", errors.New("Please enter at least 4 parameters")
	}
	return parameters[0] + ":" + parameters[1] + "@tcp" + "(" + parameters[2] + ")" + "/" + parameters[3], nil
}
