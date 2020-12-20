package service

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/barrerajuanjose/item_deco/domain"
	"io/ioutil"
	"net/http"
)

type UserService interface {
	GetUserById(userId int32) *domain.User
}

type userService struct {

}

type userResponse struct {
	Id int32 `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

func NewUserService() UserService {
	return &userService{}
}

func (*userService) GetUserById(userId int32) *domain.User {
	response, err := http.Get(fmt.Sprintf("https://api.mercadolibre.com/users/%d", userId))
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	var userResponse userResponse
	_ = json.Unmarshal(respBody, &userResponse)

	return &domain.User{
		Id: userResponse.Id,
		Nickname: userResponse.Nickname,
	}
}
