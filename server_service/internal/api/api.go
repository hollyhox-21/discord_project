package api

import "net/http"

type errorResp struct {
	Error string
}

// CreateServer
//
//	@Summary	Создание нового ceрвера
//	@Id			CreateServer
//	@Tags		server_service
//	@Router		/v1/server  [post]
//	@Accept		json
//	@Produce	json
//	@Param		request	body		CreateServerRequest	true	"query params"
//	@Success	200		{object}	CreateServerResponse
//	@Failure	400		{object}	errorResp
//	@Failure	500		{object}	errorResp
func CreateServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateServer"))
}

type CreateServerRequest struct {
	Name           string `validate:"required"`
	AboutServer    string
	AvatarPhotoUrl string
}

type CreateServerResponse struct {
	ID int
}

// GetServer
//
//	@Summary	Получение информации о сервере
//	@Id			GetServer
//	@Tags		server_service
//	@Router		/v1/server/{server_id}  [get]
//	@Produce	json
//	@Param		server_id	path		int	true	"server id"
//	@Success	200			{object}	GetServerResponse
//	@Failure	400			{object}	errorResp
//	@Failure	500			{object}	errorResp
func GetServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetServer"))
}

type GetServerResponse struct {
	ID             int
	Name           string
	AboutServer    string
	AvatarPhotoUrl string
}

// SubscribeServer
//
//	@Summary	Подписаться на серверу
//	@Id			SubscribeServer
//	@Tags		server_service
//	@Router		/v1/server/subscribe/{server_id}  [post]
//	@Accept		json
//	@Param		server_id	path	int						true	"server id"
//	@Param		request		body	SubscribeServerRequest	true	"query params"
//	@Success	200
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func SubscribeServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SubscribeServer"))
}

type SubscribeServerRequest struct {
	UserId string `validate:"required"`
}

// UnsubscribeServer
//
//	@Summary	Отписаться от сервера
//	@Id			UnsubscribeServer
//	@Tags		server_service
//	@Router		/v1/server/unsubscribe/{server_id}  [post]
//	@Accept		json
//	@Param		server_id	path	int							true	"server id"
//	@Param		request		body	UnsubscribeServerRequest	true	"query params"
//	@Success	200
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func UnsubscribeServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UnsubscribeServer"))
}

type UnsubscribeServerRequest struct {
	UserId string `validate:"required"`
}

// ServerList
//
//	@Summary	Получить список серверов в которых состоит пользователь и созданные сервера
//	@Id			ServerList
//	@Tags		server_service
//	@Router		/v1/server/user/{user_id}  [get]
//	@Produce	json
//	@Param		user_id	path		int	true	"user id"
//	@Success	200		{object}	ServerListResponse
//	@Failure	400		{object}	errorResp
//	@Failure	500		{object}	errorResp
func ServerList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ServerList"))
}

type ServerListResponse struct {
	HostServers      []string
	SubscribeServers []string
}

// InviteToServer
//
//	@Summary	Пригласить пользователя на сервер
//	@Id			InviteToServer
//	@Tags		server_service
//	@Router		/v1/server/invite/{server_id}  [post]
//	@Accept		json
//	@Param		server_id	path	int						true	"server id"
//	@Param		request		body	InviteToServerRequest	true	"query params"
//	@Success	200
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func InviteToServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("InviteToServer"))
}

type InviteToServerRequest struct {
	UserId string `validate:"required"`
}
