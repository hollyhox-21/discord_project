package api

import (
	"net/http"
)

type errorResp struct {
	Error string
}

// CreateUser
//
//	@Summary	Создание нового пользователя
//	@Id			CreateUser
//	@Tags		user_profile_service
//	@Router		/v1/user  [post]
//	@Accept		json
//	@Produce	json
//	@Param		request	body		CreateUserRequest	true	"query params"
//	@Success	200		{object}	CreateUserResponse
//	@Failure	400		{object}	errorResp
//	@Failure	500		{object}	errorResp
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUser"))
}

type CreateUserRequest struct {
	Nickname       string `validate:"required"`
	Email          string `validate:"required"`
	Password       string `validate:"required"`
	AboutUser      string
	AvatarPhotoUrl string
}

type CreateUserResponse struct {
	ID int
}

// GetUser
//
//	@Summary	Получение пользователя
//	@Id			GetUser
//	@Tags		user_profile_service
//	@Router		/v1/user/{id}  [get]
//	@Produce	json
//	@Param		id	path		int	true	"user id"
//	@Success	200	{object}	GetUserResponse
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUser"))
}

type GetUserResponse struct {
	ID             int
	Nickname       string
	Email          string
	AboutUser      string
	AvatarPhotoUrl string
}

// UpdateUser
//
//	@Summary	Обновление пользователя
//	@Id			UpdateUser
//	@Tags		user_profile_service
//	@Router		/v1/user/{id}  [put]
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int					true	"user id"
//	@Param		request	body		UpdateUserRequest	true	"query params"
//	@Success	200		{object}	UpdateUserResponse
//	@Failure	400		{object}	errorResp
//	@Failure	500		{object}	errorResp
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

type UpdateUserRequest struct {
	Nickname       string
	Email          string
	Password       string
	AboutUser      string
	AvatarPhotoUrl string
}

type UpdateUserResponse struct {
	ID             int
	Nickname       string
	Email          string
	AboutUser      string
	AvatarPhotoUrl string
}

// DeleteUser
//
//	@Summary	Удаление пользователя
//	@Id			DeleteUser
//	@Tags		user_profile_service
//	@Router		/v1/user/{id}  [delete]
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"user id"
//	@Success	200	{object}	DeleteUserResponse
//	@Failure	400	{object}	errorResp
//	@Failure	500	{object}	errorResp
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}

type DeleteUserResponse struct {
	ID             int
	Nickname       string
	Email          string
	AvatarPhotoUrl string
}
