package main

func fallback(request AppRequest) (AppResponse, error) {
	return html("<h1>Fallback</h1>"), nil
}
