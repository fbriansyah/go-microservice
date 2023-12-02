package echo

func generateSuccessResponse(data any) map[string]any {
	return map[string]any{
		"message":  "success created user",
		"data":     data,
		"is_error": false,
	}
}

func generateErrorResponse(message string) map[string]any {
	return map[string]any{
		"message":  message,
		"data":     nil,
		"is_error": true,
	}
}
