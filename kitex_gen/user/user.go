/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/west2-online/DomTok/kitex_gen/model"
)

type RegisterRequest struct {
	Username string `thrift:"username,1,required" frugal:"1,required,string" json:"username"`
	Password string `thrift:"password,2,required" frugal:"2,required,string" json:"password"`
	Email    string `thrift:"email,3,required" frugal:"3,required,string" json:"email"`
	Phone    string `thrift:"phone,4,required" frugal:"4,required,string" json:"phone"`
}

func NewRegisterRequest() *RegisterRequest {
	return &RegisterRequest{}
}

func (p *RegisterRequest) InitDefault() {
}

func (p *RegisterRequest) GetUsername() (v string) {
	return p.Username
}

func (p *RegisterRequest) GetPassword() (v string) {
	return p.Password
}

func (p *RegisterRequest) GetEmail() (v string) {
	return p.Email
}

func (p *RegisterRequest) GetPhone() (v string) {
	return p.Phone
}
func (p *RegisterRequest) SetUsername(val string) {
	p.Username = val
}
func (p *RegisterRequest) SetPassword(val string) {
	p.Password = val
}
func (p *RegisterRequest) SetEmail(val string) {
	p.Email = val
}
func (p *RegisterRequest) SetPhone(val string) {
	p.Phone = val
}

func (p *RegisterRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RegisterRequest(%+v)", *p)
}

func (p *RegisterRequest) DeepEqual(ano *RegisterRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Username) {
		return false
	}
	if !p.Field2DeepEqual(ano.Password) {
		return false
	}
	if !p.Field3DeepEqual(ano.Email) {
		return false
	}
	if !p.Field4DeepEqual(ano.Phone) {
		return false
	}
	return true
}

func (p *RegisterRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Username, src) != 0 {
		return false
	}
	return true
}
func (p *RegisterRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Password, src) != 0 {
		return false
	}
	return true
}
func (p *RegisterRequest) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Email, src) != 0 {
		return false
	}
	return true
}
func (p *RegisterRequest) Field4DeepEqual(src string) bool {

	if strings.Compare(p.Phone, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_RegisterRequest = map[int16]string{
	1: "username",
	2: "password",
	3: "email",
	4: "phone",
}

type RegisterResponse struct {
	Base   *model.BaseResp `thrift:"base,1,required" frugal:"1,required,model.BaseResp" json:"base"`
	UserID int64           `thrift:"userID,2,required" frugal:"2,required,i64" json:"userID"`
}

func NewRegisterResponse() *RegisterResponse {
	return &RegisterResponse{}
}

func (p *RegisterResponse) InitDefault() {
}

var RegisterResponse_Base_DEFAULT *model.BaseResp

func (p *RegisterResponse) GetBase() (v *model.BaseResp) {
	if !p.IsSetBase() {
		return RegisterResponse_Base_DEFAULT
	}
	return p.Base
}

func (p *RegisterResponse) GetUserID() (v int64) {
	return p.UserID
}
func (p *RegisterResponse) SetBase(val *model.BaseResp) {
	p.Base = val
}
func (p *RegisterResponse) SetUserID(val int64) {
	p.UserID = val
}

func (p *RegisterResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *RegisterResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RegisterResponse(%+v)", *p)
}

func (p *RegisterResponse) DeepEqual(ano *RegisterResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	if !p.Field2DeepEqual(ano.UserID) {
		return false
	}
	return true
}

func (p *RegisterResponse) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}
func (p *RegisterResponse) Field2DeepEqual(src int64) bool {

	if p.UserID != src {
		return false
	}
	return true
}

var fieldIDToName_RegisterResponse = map[int16]string{
	1: "base",
	2: "userID",
}

type LoginRequest struct {
	Username        string `thrift:"username,1" frugal:"1,default,string" json:"username"`
	Password        string `thrift:"password,2" frugal:"2,default,string" json:"password"`
	ConfirmPassword string `thrift:"confirm_password,3" frugal:"3,default,string" json:"confirm_password"`
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func (p *LoginRequest) InitDefault() {
}

func (p *LoginRequest) GetUsername() (v string) {
	return p.Username
}

func (p *LoginRequest) GetPassword() (v string) {
	return p.Password
}

func (p *LoginRequest) GetConfirmPassword() (v string) {
	return p.ConfirmPassword
}
func (p *LoginRequest) SetUsername(val string) {
	p.Username = val
}
func (p *LoginRequest) SetPassword(val string) {
	p.Password = val
}
func (p *LoginRequest) SetConfirmPassword(val string) {
	p.ConfirmPassword = val
}

func (p *LoginRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LoginRequest(%+v)", *p)
}

func (p *LoginRequest) DeepEqual(ano *LoginRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Username) {
		return false
	}
	if !p.Field2DeepEqual(ano.Password) {
		return false
	}
	if !p.Field3DeepEqual(ano.ConfirmPassword) {
		return false
	}
	return true
}

func (p *LoginRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Username, src) != 0 {
		return false
	}
	return true
}
func (p *LoginRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Password, src) != 0 {
		return false
	}
	return true
}
func (p *LoginRequest) Field3DeepEqual(src string) bool {

	if strings.Compare(p.ConfirmPassword, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_LoginRequest = map[int16]string{
	1: "username",
	2: "password",
	3: "confirm_password",
}

type LoginResponse struct {
	Base *model.BaseResp `thrift:"base,1" frugal:"1,default,model.BaseResp" json:"base"`
	User *model.UserInfo `thrift:"user,2" frugal:"2,default,model.UserInfo" json:"user"`
}

func NewLoginResponse() *LoginResponse {
	return &LoginResponse{}
}

func (p *LoginResponse) InitDefault() {
}

var LoginResponse_Base_DEFAULT *model.BaseResp

func (p *LoginResponse) GetBase() (v *model.BaseResp) {
	if !p.IsSetBase() {
		return LoginResponse_Base_DEFAULT
	}
	return p.Base
}

var LoginResponse_User_DEFAULT *model.UserInfo

func (p *LoginResponse) GetUser() (v *model.UserInfo) {
	if !p.IsSetUser() {
		return LoginResponse_User_DEFAULT
	}
	return p.User
}
func (p *LoginResponse) SetBase(val *model.BaseResp) {
	p.Base = val
}
func (p *LoginResponse) SetUser(val *model.UserInfo) {
	p.User = val
}

func (p *LoginResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *LoginResponse) IsSetUser() bool {
	return p.User != nil
}

func (p *LoginResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LoginResponse(%+v)", *p)
}

func (p *LoginResponse) DeepEqual(ano *LoginResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	if !p.Field2DeepEqual(ano.User) {
		return false
	}
	return true
}

func (p *LoginResponse) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}
func (p *LoginResponse) Field2DeepEqual(src *model.UserInfo) bool {

	if !p.User.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LoginResponse = map[int16]string{
	1: "base",
	2: "user",
}

type UserService interface {
	Register(ctx context.Context, req *RegisterRequest) (r *RegisterResponse, err error)

	Login(ctx context.Context, req *LoginRequest) (r *LoginResponse, err error)
}

type UserServiceRegisterArgs struct {
	Req *RegisterRequest `thrift:"req,1" frugal:"1,default,RegisterRequest" json:"req"`
}

func NewUserServiceRegisterArgs() *UserServiceRegisterArgs {
	return &UserServiceRegisterArgs{}
}

func (p *UserServiceRegisterArgs) InitDefault() {
}

var UserServiceRegisterArgs_Req_DEFAULT *RegisterRequest

func (p *UserServiceRegisterArgs) GetReq() (v *RegisterRequest) {
	if !p.IsSetReq() {
		return UserServiceRegisterArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *UserServiceRegisterArgs) SetReq(val *RegisterRequest) {
	p.Req = val
}

func (p *UserServiceRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserServiceRegisterArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceRegisterArgs(%+v)", *p)
}

func (p *UserServiceRegisterArgs) DeepEqual(ano *UserServiceRegisterArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *UserServiceRegisterArgs) Field1DeepEqual(src *RegisterRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceRegisterArgs = map[int16]string{
	1: "req",
}

type UserServiceRegisterResult struct {
	Success *RegisterResponse `thrift:"success,0,optional" frugal:"0,optional,RegisterResponse" json:"success,omitempty"`
}

func NewUserServiceRegisterResult() *UserServiceRegisterResult {
	return &UserServiceRegisterResult{}
}

func (p *UserServiceRegisterResult) InitDefault() {
}

var UserServiceRegisterResult_Success_DEFAULT *RegisterResponse

func (p *UserServiceRegisterResult) GetSuccess() (v *RegisterResponse) {
	if !p.IsSetSuccess() {
		return UserServiceRegisterResult_Success_DEFAULT
	}
	return p.Success
}
func (p *UserServiceRegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*RegisterResponse)
}

func (p *UserServiceRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserServiceRegisterResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceRegisterResult(%+v)", *p)
}

func (p *UserServiceRegisterResult) DeepEqual(ano *UserServiceRegisterResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *UserServiceRegisterResult) Field0DeepEqual(src *RegisterResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceRegisterResult = map[int16]string{
	0: "success",
}

type UserServiceLoginArgs struct {
	Req *LoginRequest `thrift:"req,1" frugal:"1,default,LoginRequest" json:"req"`
}

func NewUserServiceLoginArgs() *UserServiceLoginArgs {
	return &UserServiceLoginArgs{}
}

func (p *UserServiceLoginArgs) InitDefault() {
}

var UserServiceLoginArgs_Req_DEFAULT *LoginRequest

func (p *UserServiceLoginArgs) GetReq() (v *LoginRequest) {
	if !p.IsSetReq() {
		return UserServiceLoginArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *UserServiceLoginArgs) SetReq(val *LoginRequest) {
	p.Req = val
}

func (p *UserServiceLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserServiceLoginArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceLoginArgs(%+v)", *p)
}

func (p *UserServiceLoginArgs) DeepEqual(ano *UserServiceLoginArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *UserServiceLoginArgs) Field1DeepEqual(src *LoginRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceLoginArgs = map[int16]string{
	1: "req",
}

type UserServiceLoginResult struct {
	Success *LoginResponse `thrift:"success,0,optional" frugal:"0,optional,LoginResponse" json:"success,omitempty"`
}

func NewUserServiceLoginResult() *UserServiceLoginResult {
	return &UserServiceLoginResult{}
}

func (p *UserServiceLoginResult) InitDefault() {
}

var UserServiceLoginResult_Success_DEFAULT *LoginResponse

func (p *UserServiceLoginResult) GetSuccess() (v *LoginResponse) {
	if !p.IsSetSuccess() {
		return UserServiceLoginResult_Success_DEFAULT
	}
	return p.Success
}
func (p *UserServiceLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*LoginResponse)
}

func (p *UserServiceLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserServiceLoginResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceLoginResult(%+v)", *p)
}

func (p *UserServiceLoginResult) DeepEqual(ano *UserServiceLoginResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *UserServiceLoginResult) Field0DeepEqual(src *LoginResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceLoginResult = map[int16]string{
	0: "success",
}
