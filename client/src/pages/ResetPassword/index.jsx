import "@/pages/Login/module.scss";
import Logo from "@/assets/img/brand.png";

import { useState, useEffect, useCallback } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import _ from "lodash";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil from "@/util/message.util";

// Components
import Input from "@/components/Input";
import Button from "@/components/Button";

const message = {
  invalidRequest: "Invalid request",
  success: "Reset password successfully",
  fieldNotFill: "Please fill in all fields",
  passwordNotMatch: "New password and confirm new password are not matched",
  error: "Reset password failed",
};

const ResetPassword = () => {
  const navigate = useNavigate();
  const urlSearchParams = new URLSearchParams(useLocation().search);
  const [email, token] = urlSearchParams.values();
  const [isLoading, setIsLoading] = useState(true);

  // State
  const [form, setForm] = useState({
    email: "",
    token: "",
    password: "",
    confirmPassword: "",
  });
  const [isTextTypeP, setIsTextTypeP] = useState(true);
  const [isTextTypeCP, setIsTextTypeCP] = useState(true);

  // Method
  const handleInitialization = useCallback(async () => {
    if (!email || !token) {
      navigate(routerConfig.routes.LOGIN);
    }

    try {
      await apiHandler.get(
        apiConfig.resource.CHECK_RESET_PASSWORD_TOKEN.replace(
          ":email",
          email
        ).replace(":token", token)
      );

      setForm({
        email,
        token,
        password: "",
        confirmPassword: "",
      });

      setIsLoading(false);
    } catch (error) {
      messageUtil.showErrorMessage(message.invalidRequest);

      setTimeout(() => {
        navigate(routerConfig.routes.LOGIN);
      }, 3000);
    }
  }, [email, navigate, token]);

  const emptyForm = () => {
    setForm({
      email: "",
      token: "",
      password: "",
      confirmPassword: "",
    });
  };

  const changeForm = (key) => (e) => {
    let value = "";
    if (_.isObject(e)) {
      value = e.target.value;
    }

    setForm({ ...form, [key]: value });
  };

  const handleResetPassword = async () => {
    if (!form.password || !form.confirmPassword) {
      messageUtil.showErrorMessage(message.fieldsNotFill);
      return;
    }

    if (form.password !== form.confirmPassword) {
      messageUtil.showErrorMessage(message.passwordNotMatch);
      return;
    }

    const payload = {
      email: form.email,
      token: form.token,
      password: form.password,
    };

    // Call API
    try {
      await apiHandler.post(apiConfig.resource.RESET_PASSWORD, payload);

      messageUtil.showSuccessMessage(message.success);
      emptyForm();
      navigate(routerConfig.routes.LOGIN);
    } catch (error) {
      messageUtil.showErrorMessage(apiHandler.extractErrorMessage(error));
    }
  };

  useEffect(() => {
    handleInitialization();
  }, [handleInitialization, token]);

  if (isLoading) {
    return <div className='loader'>Loading...</div>;
  }

  return (
    <div className='login-section'>
      <div className='login-block'>
        <div className='login-block__logo'>
          <img src={Logo} alt='logo' />
        </div>
        <div className='header-title login-block__title'>
          <h2>Reset Password</h2>
        </div>
        <div className='login-block__body'>
          <div className='input-group'>
            <label htmlFor='username'>E-mail</label>
            <Input
              id='email'
              type='text'
              name='email'
              disabled={true}
              placeholder='Your email'
              extraClasses={["disabled"]}
              value={form.email}
            />
          </div>
          <div className='space-b-3'></div>
          <div className='input-group'>
            <label htmlFor='password'>New password</label>
            <Input
              id='password'
              type={isTextTypeP ? "password" : "text"}
              name='password'
              placeholder='at least 8 characters'
              value={form.password}
              onChange={(e) => changeForm("password")(e)}
            >
              {isTextTypeP ? (
                <i
                  className='fa-regular fa-eye'
                  title='show password'
                  onClick={() => setIsTextTypeP(!isTextTypeP)}
                ></i>
              ) : (
                <i
                  className='fa-regular fa-eye-slash'
                  title='hide password'
                  onClick={() => setIsTextTypeP(!isTextTypeP)}
                ></i>
              )}
            </Input>
          </div>
          <div className='space-b-3'></div>
          <div className='input-group'>
            <label htmlFor='password'>Confirm new password</label>
            <Input
              id='confirmPassword'
              type={isTextTypeCP ? "password" : "text"}
              name='password'
              placeholder='at least 8 characters'
              value={form.confirmPassword}
              onChange={(e) => changeForm("confirmPassword")(e)}
            >
              {isTextTypeCP ? (
                <i
                  className='fa-regular fa-eye'
                  title='show password'
                  onClick={() => setIsTextTypeCP(!isTextTypeCP)}
                ></i>
              ) : (
                <i
                  className='fa-regular fa-eye-slash'
                  title='hide password'
                  onClick={() => setIsTextTypeCP(!isTextTypeCP)}
                ></i>
              )}
            </Input>
          </div>
          <div className='space-b-4'></div>
          <div className='button-container'>
            <Button id='submitBtn' onClick={handleResetPassword}>
              Submit
            </Button>
            <Button
              id='cancelBtn'
              extraClasses={["cancel-button"]}
              onClick={emptyForm}
            >
              Reset
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ResetPassword;
