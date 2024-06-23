package dto

import "email.v1/models"

type RequestToken struct {
	UsersID string `json:"users_id"`
	Token   string `json:"token"`
	Expire  int64  `json:"expire"`
}

type ResponseToken struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type RequestUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsVerify bool   `json:"isverify"`
}

// Mapping Model To Response
func ModelUserToResponse(data models.Users) ResponseUser {
	return ResponseUser{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		IsVerify: data.IsVerify,
	}
}

func ModelVerifTorResponse(data models.TokenVerif) ResponseToken {
	return ResponseToken{
		Token:  data.Token,
		Expire: data.Expire,
	}
}

// Mapping Request To Model
func RequestUserToModel(data RequestUser) models.Users {
	return models.Users{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}

func RequestToModel(data RequestToken) models.TokenVerif {
	return models.TokenVerif{
		UsersID: data.UsersID,
		Token:   data.Token,
		Expire:  data.Expire,
	}
}
