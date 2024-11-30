import "@/pages/Login/module.css";
import logo from "@/assets/img/brand.png";

import { useState, useEffect, useCallback } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil from "@/util/message.util";

// Components
import MainTitle from "@/components/MainTitle";
import FormGroup from "@/components/FormGroup";
import FormLabel from "@/components/FormLabel";
import InputBox from "@/components/InputBox";
import ButtonBox from "@/components/ButtonBox";
import LoadingBox from "@/components/LoadingBox";

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
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);
  const [isConfirmPasswordVisible, setIsConfirmPasswordVisible] =
    useState(false);

  // Method
  const handleInitialization = useCallback(async () => {
    if (!email || !token) {
      navigate(routerConfig.routes.LOGIN);
    }

    try {
      setIsLoading(true);
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
    } finally {
      setIsLoading(false);
    }
  }, [email, navigate, token]);

  const emptyForm = () => {
    setForm({
      email: form.email,
      token: "",
      password: "",
      confirmPassword: "",
    });
  };

  const changeForm = (key) => (value) => {
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

  return (
    <>
      <div className='login-container'>
        <div className='login-border light-primary-shadow'>
          <div className='login-block login-header light-primary-shadow'>
            <div className='icon'>
              <img src={logo} />
            </div>
            <MainTitle extraClasses={["!mb-0"]}>
              Management Information System
            </MainTitle>
          </div>
          <div className='login-block login-form light-primary-shadow'>
            <FormGroup>
              <MainTitle extraClasses={["text-general"]}>
                Reset Password
              </MainTitle>
            </FormGroup>
            <FormGroup>
              <FormLabel forName='email'>E-mail</FormLabel>
              <InputBox
                type='text'
                id='email'
                name='email'
                placeholder='E-mail'
                disabled={true}
                value={form.email}
                onChange={() => {}}
              />
            </FormGroup>
            <FormGroup>
              <FormLabel forName='password'>Password</FormLabel>
              <InputBox
                type={isPasswordVisible ? "text" : "password"}
                id='password'
                name='password'
                placeholder='Password'
                value={form.password}
                onChange={(value) => {
                  changeForm("password")(value);
                }}
              >
                {
                  <FontAwesomeIcon
                    icon={isPasswordVisible ? faEyeSlash : faEye}
                    onClick={() => setIsPasswordVisible(!isPasswordVisible)}
                  />
                }
              </InputBox>
            </FormGroup>
            <FormGroup>
              <FormLabel forName='password'>Confirm Password</FormLabel>
              <InputBox
                type={isConfirmPasswordVisible ? "text" : "password"}
                id='confirmPassword'
                name='confirm-password'
                placeholder='Confirm Password'
                value={form.confirmPassword}
                onChange={(value) => {
                  changeForm("confirmPassword")(value);
                }}
              >
                {
                  <FontAwesomeIcon
                    icon={isConfirmPasswordVisible ? faEyeSlash : faEye}
                    onClick={() =>
                      setIsConfirmPasswordVisible(!isConfirmPasswordVisible)
                    }
                  />
                }
              </InputBox>
            </FormGroup>
          </div>
          <div className='login-block login-button'>
            <ButtonBox onClick={handleResetPassword} isSave={true}>
              Submit
            </ButtonBox>
            <ButtonBox
              isClose={true}
              extraClasses={["cancel-shadow"]}
              onClick={emptyForm}
            >
              Reset
            </ButtonBox>
          </div>
          <div className='login-block login-footer light-primary-shadow'>
            © 2024 Ivan Solutions. All rights reserved.
          </div>
        </div>
      </div>
      <LoadingBox visible={isLoading} />
    </>
  );
};

export default ResetPassword;
