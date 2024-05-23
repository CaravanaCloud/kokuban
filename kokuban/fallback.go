package main

func fallback(data map[string]interface{}) (string, string, error) {
	return "text/html", "<h1>Default Response</h1>", nil
}
