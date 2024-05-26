package api

import "net/http"

type errorResp struct {
	Error string
}

// RequestFriendship
//
//	@Summary	Запрос на дружбу
//	@Id			RequestFriendship
//	@Tags		friendship_service
//	@Router		/v1/friend/request  [post]
//	@Accept		json
//	@Produce	json
//	@Param		request	body	RequestFriendshipRequest	true	"query params"
//	@Success	200
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func RequestFriendship(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RequestFriendship"))
}

type RequestFriendshipRequest struct {
	FromUserID string `validate:"required"`
	ToUserID   string `validate:"required"`
}

// DeleteFriendship
//
//	@Summary	Удалить пользователя из друзей
//	@Id			DeleteFriendship
//	@Tags		friendship_service
//	@Router		/v1/friend/request  [delete]
//	@Accept		json
//	@Produce	json
//	@Param		request	body	DeleteFriendshipRequest	true	"query params"
//	@Success	200
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func DeleteFriendship(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RequestFriendship"))
}

type DeleteFriendshipRequest struct {
	UserID       string `validate:"required"`
	DeleteUserID string `validate:"required"`
}

// AcceptFriendship
//
//	@Summary	Подтвердить запрос дружбы
//	@Id			AcceptFriendship
//	@Tags		friendship_service
//	@Router		/v1/friend/accept  [put]
//	@Accept		json
//	@Produce	json
//	@Param		request	body	AcceptFriendshipRequest	true	"query params"
//	@Success	200
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func AcceptFriendship(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RequestFriendship"))
}

type AcceptFriendshipRequest struct {
	FromUserID string `validate:"required"`
	ToUserID   string `validate:"required"`
}

// DeclineFriendship
//
//	@Summary	Отклонить запрос дружбы
//	@Id			DeclineFriendship
//	@Tags		friendship_service
//	@Router		/v1/friend/decline  [put]
//	@Accept		json
//	@Produce	json
//	@Param		request	body	DeclineFriendshipRequest	true	"query params"
//	@Success	200
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func DeclineFriendship(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RequestFriendship"))
}

type DeclineFriendshipRequest struct {
	FromUserID string `validate:"required"`
	ToUserID   string `validate:"required"`
}

// ListFriendship
//
//	@Summary	Получить список друзей
//	@Id			ListFriendship
//	@Tags		friendship_service
//	@Router		/v1/friend/list/{user_id}  [get]
//	@Produce	json
//	@Param		user_id	path		int	true	"user id"
//	@Success	200		{array}		ListFriendshipResponse
//	@Failure	400		{object}	errorResp
//	@Failure	500		{object}	errorResp
func ListFriendship(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RequestFriendship"))
}

type ListFriendshipResponse struct {
	ID             int
	Nickname       string
	Email          string
	AvatarPhotoUrl string
}
