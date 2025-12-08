package response

// WebResponse adalah format standar untuk semua response API
type WebResponse struct {
	Code    int         `json:"code"`           // HTTP Status Code (200, 400, 500)
	Message string      `json:"message"`        // Pesan (Success, Invalid Input, dll)
	Data    interface{} `json:"data,omitempty"` // Payload data (bisa null kalau error)
}

// Helper function untuk respon Sukses
func Success(code int, message string, data interface{}) WebResponse {
	return WebResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// Helper function untuk respon Error
func Error(code int, message string) WebResponse {
	return WebResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
