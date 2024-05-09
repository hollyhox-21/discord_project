package api

import "net/http"

type errorResp struct {
	Error string
}

// DownloadFile
//
//	@Summary	Скачивание файла
//	@Id			DownloadFile
//	@Tags		media_service
//	@Router		/v1/download/{file_id} [get]
//	@Accept		json
//	@Produce	mpfd
//	@Param		file_id	path		int	true	"file id"
//	@Success	200		{array}		byte
//	@Failure	400		{object}	errorResp
//	@Failure	500		{object}	errorResp
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DownloadFile"))
}

// UploadFile
//
//	@Summary	Загрузка файла
//	@Id			UploadFile
//	@Tags		media_service
//	@Router		/v1/upload/{user_id} [post]
//	@Accept		mpfd
//	@Produce	json
//	@Param		user_id	path		int		true	"user id"
//	@Param		file	formData	file	true	"query param"
//	@Success	200		{object}	UploadFileResponse
//	@Failure	400		{object}	errorResp
//	@Failure	500		{object}	errorResp
func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UploadFile"))
}

type UploadFileResponse struct {
	Result string
	Link   string
}
