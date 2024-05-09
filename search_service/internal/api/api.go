package api

import "net/http"

type errorResp struct {
	Error string
}

// SearchUser
//
//	@Summary	Поиск пользователя по nickname
//	@Id			SearchUser
//	@Tags		search_service
//	@Router		/v1/search/user/{nickname}  [get]
//	@Produce	json
//	@Param		nickname	path		string	true	"user nickname"
//	@Success	200			{object}	SearchUserResponse
//	@Failure	400			{object}	errorResp
//	@Failure	500			{object}	errorResp
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SearchUser"))
}

type SearchUserResponse struct {
	ID             int
	Nickname       string
	Email          string
	AboutUser      string
	AvatarPhotoUrl string
}

// SearchServer
//
//	@Summary	Поиск сервера по server_name
//	@Id			SearchServer
//	@Tags		search_service
//	@Router		/v1/search/server/{server_name}  [get]
//	@Produce	json
//	@Param		server_name	path		string	true	"server name"
//	@Success	200			{object}	SearchServerResponse
//	@Failure	400			{object}	errorResp
//	@Failure	500			{object}	errorResp
func SearchServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SearchServer"))
}

type SearchServerResponse struct {
	ID          int
	ServerName  string
	AboutServer string
	Users       []string
}
