package api

import "net/http"

type errorResp struct {
	Error string
}

// MessageSend
//
//	@Summary	отправка сообщения
//	@Id			MessageSend
//	@Tags		chat_service
//	@Router		/v1/message/send  [post]
//	@Accept		json
//	@Param		typeMsg	query	string				true	"string enums"	Enums(unknown,server,user)
//	@Param		request	body	MessageSendRequest	true	"query params"
//	@Success	200
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func MessageSend(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MessageSend"))
}

type MessageSendRequest struct {
	ProducerId int
	ConsumerId int
	Content    string
}

// GetServerChatHistory
//
//	@Summary	Получение истории сообщений сервера
//	@Id			GetServerChatHistory
//	@Tags		chat_service
//	@Router		/v1/message/server/history/{server_id}  [get]
//	@Produce	json
//	@Param		server_id	path		int	true	"server id"
//	@Success	200			{object}	GetServerChatHistoryResponse
//	@Failure	400			{object}	errorResp
//	@Failure	500			{object}	errorResp
func GetServerChatHistory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetServerHistory"))
}

type GetServerChatHistoryResponse struct {
	Messages []string
}

// GetPrivetChatHistory
//
//	@Summary	Получение истории приватного чата
//	@Id			GetPrivetChatHistory
//	@Tags		chat_service
//	@Router		/v1/message/user/history/{user_id}  [get]
//	@Produce	json
//	@Param		recipient_user_id	query		string	true	" id user recipient"
//	@Param		user_id				path		int		true	"user id"
//	@Success	200					{object}	GetPrivetChatHistoryResponse
//	@Failure	400					{object}	errorResp
//	@Failure	500					{object}	errorResp
func GetPrivetChatHistory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetPrivetChatHistory"))
}

type GetPrivetChatHistoryResponse struct {
	Messages []string
}
