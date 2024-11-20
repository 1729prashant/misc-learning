package main

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (authInfo authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + authInfo.username + ":" + authInfo.password
}

