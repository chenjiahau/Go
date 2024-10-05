package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

const recaptchaAPIURL = "https://www.google.com/recaptcha/api/siteverify"

type RecaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

func (ctrl *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	// Validate request
	var sp model.SignUpParams
	err := util.DecodeJSONBody(r, &sp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(sp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1401)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	if sp.Email == "" || sp.Name == "" || sp.Password == "" || sp.ConfirmPassword == "" {
		resErr := util.GetReturnMessage(1402)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	if sp.Password != sp.ConfirmPassword {
		resErr := util.GetReturnMessage(1403)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if email is already registered
	u := model.NewUser()
  id := u.GetId(sp)
  if id != 0 {
		resErr := util.GetReturnMessage(1404)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Hash password
	sp.Password, err = util.HashPassword(sp.Password)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1405)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Insert user
	id, err = u.Create(sp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1406)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Insert user register
	ur := model.NewUserRegister()
  urp := model.UserRegisterParams{
		UserId: id,
		Token: util.CreateMD5Hash(id),
		ExpiredAt: util.GetNow().AddDate(0, 0, 1), // 1 day
	}
	err = ur.Create(urp)

	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1407)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Send confirmation email
	confirmUrl := ctrl.Config.PortalUrl + "/#/active-account?token=" + url.QueryEscape(urp.Token)
	err = util.SendEmail(
		ctrl.Config.EmailConf.Host,
		ctrl.Config.EmailConf.Port,
		ctrl.Config.EmailConf.User,
		ctrl.Config.EmailConf.Pass,
		ctrl.Config.EmailConf.User,
		sp.Email,
		"User Registration",
		"Click <a href=\"" + confirmUrl + "\">here</a> to confirm your registration.",
	)

	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1408)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(1201)
	resData["data"] = map[string]interface{}{
		"email": sp.Email,
		"name": sp.Name,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// Validate request
	token := chi.URLParam(r, "token")
	if token == "" {
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query user register
	ur := model.NewUserRegister()
	err := ur.Query(token)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1409)
		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	// Check token expiration
	expiredAt := ur.(*model.UserRegister).ExpiredAt
	now := util.GetNow()
	if now.After(expiredAt) {
		resErr := util.GetReturnMessage(1409)
		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	// Update user's status to active
	u := model.NewUser()
	err = u.Active(ur.(*model.UserRegister).UserId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1409)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete user register
	err = ur.Delete(ur.(*model.UserRegister).Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1409)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(1205)
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	// Validate recaptcha
	// recaptchaToken := r.URL.Query().Get("recaptchaToken")
	// if recaptchaToken == "" {
	// 	resErr := util.GetReturnMessage(400)
	// 	util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
	// 	return
	// }

	// if !verifyRecaptcha(ctrl, recaptchaToken) {
	// 	resErr := util.GetReturnMessage(1414)
	// 	util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
	// 	return
	// }

	// Validate request
	var si model.SignInParams
	err := util.DecodeJSONBody(r, &si)
	if err != nil {
		util.WriteErrorLog(err.Error())
    resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(si)
	if err != nil {
		util.WriteErrorLog(err.Error())
    resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query user
	u := model.NewUser()
	err = u.Query(si)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1411)
		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	// Create token
	user := u.(*model.User)
	tokenString, err := util.CreateToken(user.Id, user.Name)
	if err != nil {
		util.WriteErrorLog(err.Error())
    resErr := util.GetReturnMessage(1412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

  t := model.NewToken()
	err = t.Create(user.Id, tokenString, util.GetNow(), util.GetNow().AddDate(0, 0, 1)) // 1 day
	if err != nil {
		util.WriteErrorLog(err.Error())
    resErr := util.GetReturnMessage(1413)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(1202)
	resData["data"] = map[string]interface{}{
		"id": user.Id,
		"email": user.Email,
		"name": user.Name,
		"token": tokenString,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) CreateResetPassword(w http.ResponseWriter, r *http.Request) {
	// Validate request
	var cuep model.CheckUserExistParams
	err := util.DecodeJSONBody(r, &cuep)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if email has signed up, if not return success
	u := model.NewUser()
  existUser, err := u.CheckUserExist(cuep.Email)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1431)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	if existUser.Id == 0 {
		resData := util.GetReturnMessage(1206)
		resData["data"] = map[string]interface{}{
			"email": cuep.Email,
		}

		util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
		return
	}

	// Check if user is registered, if not return success
	if !existUser.IsRegistered {
		resData := util.GetReturnMessage(1206)
		resData["data"] = map[string]interface{}{
			"email": cuep.Email,
		}

		util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
		return
	}
	
	// Check the last reset password record is over 15 minutes, if not return success
	urp := model.NewUserResetPassword()
	existRecord, _ := urp.Query(cuep.Email)
	if existRecord.Id != 0 {
		overValidTimestamp := existRecord.ExpiredAt.Add(-60 * 24 * time.Minute + 15 * time.Minute).Unix()
		nowTimestamp := util.GetNow().Unix()

		if nowTimestamp < overValidTimestamp {
			resData := util.GetReturnMessage(1206)
			resData["data"] = map[string]interface{}{
				"email": cuep.Email,
			}

			util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
			return
	  }
	}

	// Invalidate the last reset password record
	err = urp.Invalidate(existUser.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1432)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil	, resErr))
		return
	}

	// Insert a new user reset password, expired in 1 day
	urpp := model.UserResetPasswordParams{
		UserId: existUser.Id,
		Token: util.CreateMD5Hash(existUser.Id),
		ExpiredAt: util.GetNow().AddDate(0, 0, 1),
	}
	err = urp.Create(urpp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1433)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Send confirmation email
	confirmUrl := ctrl.Config.PortalUrl + "/#/reset-password?email=" + url.QueryEscape(cuep.Email) + "&token=" + url.QueryEscape(urpp.Token)
	err = util.SendEmail(
		ctrl.Config.EmailConf.Host,
		ctrl.Config.EmailConf.Port,
		ctrl.Config.EmailConf.User,
		ctrl.Config.EmailConf.Pass,
		ctrl.Config.EmailConf.User,
		cuep.Email,
		"Reset Password",
		"Click <a href=\"" + confirmUrl + "\">here</a> to reset your password.",
	)

	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1434)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(1206)
	resData["data"] = map[string]interface{}{
		"email": cuep.Email,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) CheckRestPasswordToken(w http.ResponseWriter, r *http.Request) {
	// Validate request
	email := chi.URLParam(r, "email")
	token := chi.URLParam(r, "token")
	if email == "" || token == "" {
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check token is valid
	urp := model.NewUserResetPassword()
	existTicket, err := urp.Query(email)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1435)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	if existTicket.Id == 0 || existTicket.Token != token {
		resErr := util.GetReturnMessage(1435)
		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	// Check token expiration
	expiredAt := existTicket.ExpiredAt
	now := util.GetNow()
	if now.After(expiredAt) {
		resErr := util.GetReturnMessage(1435)
		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(1207)
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) ResetPassword(w http.ResponseWriter, r *http.Request) {
	// Validate request
	var rpp model.ResetPasswordParams
	err := util.DecodeJSONBody(r, &rpp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check token is valid
	urp := model.NewUserResetPassword()
	existTicket, _ := urp.Query(rpp.Email)
	if existTicket.Id == 0 {
		resErr := util.GetReturnMessage(1435)
		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	// Check token expiration
	expiredAt := existTicket.ExpiredAt
	now := util.GetNow()
	if now.After(expiredAt) {
		resErr := util.GetReturnMessage(1435)
		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	// Hash password
	rpp.Password, err = util.HashPassword(rpp.Password)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1436)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Change password
	u := model.NewUser()
	err = u.ResetPassword(existTicket.UserId, rpp.Password)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1437)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Invalidate user reset password
	err = urp.Invalidate(existTicket.UserId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1432)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(1208)
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func verifyRecaptcha(ctrl *Controller, recaptchaToken string) bool {
	secretKey := ctrl.Config.RecaptchaConf.SecretKey
	data := url.Values{
		"secret": {secretKey},
		"response": {recaptchaToken},
	}

	resp, err := http.PostForm(recaptchaAPIURL, data)
	if err != nil {
		log.Printf("Error sending recaptcha request: %v", err)
		return false
	}
	defer resp.Body.Close()

	var recaptchaResponse RecaptchaResponse
	err = json.NewDecoder(resp.Body).Decode(&recaptchaResponse)
	if err != nil {
		log.Printf("Error decoding recaptcha response: %v", err)
		return false
	}

	return recaptchaResponse.Success
}

func (ctrl *Controller) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Validate request
	var cpp model.ChangePasswordParams
	err := util.DecodeJSONBody(r, &cpp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(cpp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check new password and confirm password
	if cpp.NewPassword != cpp.ConfirmPassword {
		resErr := util.GetReturnMessage(1401)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check original password is correct
	u := model.NewUser()
	u.GetUserById(ctrl.User.Id)
	err = util.CheckPasswordHash(cpp.OriginalPassword, u.(*model.User).Password)
	if err != nil {
		resErr := util.GetReturnMessage(9402)
		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	// Hash new password
	hashPassword, err := util.HashPassword(cpp.NewPassword)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(9403)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Change password
	err = u.ResetPassword(ctrl.User.Id, hashPassword)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(9404)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(9201)
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}