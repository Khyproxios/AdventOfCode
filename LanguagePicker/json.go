package main

import "encoding/json"

func parseLanguages(data []byte) ([]Language, error) {
	var languages []Language

	error := json.Unmarshal(data, &languages)

	if error != nil {
		return nil, error
	}

	return languages, nil
}

func stringify(languages []Language) ([]byte, error) {
	data, error := json.Marshal(languages)

	if error != nil {
		return nil, error
	}

	return data, nil
}
