package handler

import (
	"context"
	"micro/user/domain/model"
	"micro/user/domain/service"
	user "micro/user/proto"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, in *user.UserRegisterResquest, out *user.UserRegisterResponse) error {
	userRegister := &model.User{
		Username:     in.UserName,
		FirstName:    in.FirstName,
		HashPassword: in.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	out.Message = "Add user succeeded"
	return nil
}

func (u *User) Login(ctx context.Context, in *user.UserLoginRequest, out *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(in.UserName, in.Pwd)
	if err != nil {
		return err
	}
	out.IsSuccess = isOk
	return nil
}
func (u *User) GetUserInfo(ctx context.Context, in *user.UserInfoRequest, out *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(in.UserName)
	if err != nil {
		return err
	}
	out = UserInfoForResponse(userInfo)
	return nil
}

func UserInfoForResponse(userModel *model.User) *user.UserInfoResponse {
	resp := &user.UserInfoResponse{
		FirstName: userModel.FirstName,
		UserName:  userModel.Username,
		UserId:    userModel.ID,
	}
	return resp
}
