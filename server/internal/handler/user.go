package handler

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

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
  urm := model.UserRegisterParams{
		UserId: id,
		Token: util.CreateMD5Hash(id),
		ExpiredAt: util.GetNow().AddDate(0, 0, 1), // 1 day
	}
	err = ur.Create(urm)

	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(1407)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Send confirmation email
	confirmUrl := os.Getenv("PORTAL_URL") + "/active-account?token=" + urm.Token
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