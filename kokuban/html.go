package main

func html(body string) AppResponse {
	return AppResponse{
		Body:        body,
		ContentType: "text/html",
	}
}
