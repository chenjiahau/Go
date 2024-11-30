import "../module.css";
import logo from "@/assets/img/brand.png";

import { useState } from "react";
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
import LinkButton from "@/components/LinkButton";
import LoadingBox from "@/components/LoadingBox";

const SignUp = ({ successMessage, errorMessage, stageType, onChangeStage }) => {
  // State
  const [form, setForm] = useState({
    email: "",
    username: "",
    password: "",
    confirmPassword: "",
  });
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);
  const [isConfirmPasswordVisible, setIsConfirmPasswordVisible] =
    useState(false);
  const [isLoading, setIsLoading] = useState(false);

  // Method
  const changeForm = (key) => (value) => {
    setForm({ ...form, [key]: value });
  };

  const emptyForm = () => {
    setForm({
      email: "",
      username: "",
      password: "",
      confirmPassword: "",
    });
  };

  const handleSignUp = async () => {
    if (
      !form.email ||
      !form.username ||
      !form.password ||
      !form.confirmPassword
    ) {
      messageUtil.showErrorMessage(errorMessage.fieldsNotFill);
      return;
    }

    if (form.password !== form.confirmPassword) {
      messageUtil.showErrorMessage(errorMessage.passwordNotMatch);
      return;
    }

    const payload = {
      email: form.email,
      username: form.username,
      password: form.password,
      confirmPassword: form.confirmPassword,
    };

    // Call API
    try {
      setIsLoading(true);
      await apiHandler.post(apiConfig.resource.SIGNUP, payload);
      emptyForm();
      messageUtil.showSuccessMessage(successMessage.signup);
    } catch (error) {
      messageUtil.showErrorMessage(apiHandler.extractErrorMessage(error));
    } finally {
      setIsLoading(false);
    }
  };

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
              <MainTitle extraClasses={["text-general"]}>Sign Up</MainTitle>
            </FormGroup>
            <FormGroup>
              <FormLabel forName='email'>E-mail</FormLabel>
              <InputBox
                type='text'
                id='email'
                name='email'
                placeholder='E-mail'
                value={form.email}
                onChange={(value) => {
                  changeForm("email")(value);
                }}
              />
            </FormGroup>
            <FormGroup>
              <FormLabel forName='username'>Username</FormLabel>
              <InputBox
                type='text'
                id='username'
                name='username'
                placeholder='Username'
                value={form.username}
                onChange={(value) => {
                  changeForm("username")(value);
                }}
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
                      setIsConfirmPasswordVisible(!isPasswordVisible)
                    }
                  />
                }
              </InputBox>
            </FormGroup>
          </div>
          <div className='login-block login-button'>
            <ButtonBox onClick={handleSignUp} isSave={true}>
              Sign Up
            </ButtonBox>
            <ButtonBox
              isClose={true}
              extraClasses={["cancel-shadow"]}
              onClick={emptyForm}
            >
              Reset
            </ButtonBox>
          </div>
          <div className='login-block login-link'>
            <div>Do you have an account?</div>
            <LinkButton
              to={routerConfig.routes.LOGIN}
              onClick={onChangeStage(stageType.LOGIN)}
              title='Back'
            />
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

export default SignUp;
