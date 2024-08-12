package handler

import (
	"net/http"

	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (Ctrl *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	var u model.UserInterface = &model.User{}
	var sp model.SignUpParams

	err := util.DecodeJSONBody(r, &sp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(sp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid email, username, password, or confirm password",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	if sp.Email == "" || sp.Name == "" || sp.Password == "" || sp.ConfirmPassword == "" {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Email, username, password, and confirm password are required",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	if sp.Password != sp.ConfirmPassword {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Password and confirm password do not match",
		}
		
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

  id := u.GetId(sp)
  if id != 0 {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Email is already registered",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	sp.Password, err = util.HashPassword(sp.Password)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to hash password",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	err = u.Create(sp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to insert user",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"message": "User created successfully",
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	var u model.UserInterface = &model.User{}
	var si model.SignInParams

	err := util.DecodeJSONBody(r, &si)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(si)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid email or password",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	if si.Email == "" || si.Password == "" {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Email and password are required",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	err = u.Query(si)
	if err != nil {
		resErr := map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"message": "Invalid email or password",
		}

		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	user := u.(*model.User)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get token",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	tokenString, err := util.CreateToken(user.Id, user.Name)
	if err != nil {
		resErr := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to create token",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	var t model.TokenInterface = &model.Token{}
	err = t.Create(user.Id, tokenString, util.GetNow(), util.GetNow().AddDate(0, 0, 1)) // 1 day

	if err != nil {
		resErr := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to insert token",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id":    user.Id,
		"email": user.Email,
		"name":  user.Name,
		"token": tokenString,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}