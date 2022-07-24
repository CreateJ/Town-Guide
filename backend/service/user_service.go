package service

type UserService interface {
	GetUserInfo(openId string) Response
}

type userServiceImpl struct{}

func NewUserService() UserService {
	return userServiceImpl{}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (u userServiceImpl) GetUserInfo(openId string) Response {
	return Response{1, "ddd", nil}
}
