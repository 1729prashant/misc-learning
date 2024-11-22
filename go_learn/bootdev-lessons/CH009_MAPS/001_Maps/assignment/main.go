package main

import (
    "errors"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
    if len(names) != len(phoneNumbers) {
        return nil, errors.New("invalid sizes")
    }

    usr := make(map[string]user) // initialize the map

    for i := range names {
        usr[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
    }

    return usr, nil
}

type user struct {
	name        string
	phoneNumber int
}
