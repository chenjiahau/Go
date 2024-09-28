import "../module.scss";
import Logo from "@/assets/img/brand.png";

import { useState } from "react";
import { Link } from "react-router-dom";
import _ from "lodash";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil from "@/util/message.util";

// Components
import Input from "@/components/Input";
import Button from "@/components/Button";

const SignUp = ({ successMessage, errorMessage, stageType, onChangeStage }) => {
  // State
  const [form, setForm] = useState({
    email: "",
    username: "",
    password: "",
    confirmPassword: "",
  });
  const [isTextTypeP, setIsTextTypeP] = useState(true);
  const [isTextTypeCP, setIsTextTypeCP] = useState(true);

  // Method
  const changeForm = (key) => (e) => {
    let value = "";
    if (_.isObject(e)) {
      value = e.target.value;
    }

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
      await apiHandler.post(apiConfig.resource.SIGNUP, payload);
      emptyForm();
      messageUtil.showSuccessMessage(successMessage.signup);
    } catch (error) {
      messageUtil.showErrorMessage(apiHandler.extractErrorMessage(error));
    }
  };

  return (
    <div className='login-section'>
      <div className='login-block'>
        <div className='login-block__logo'>
          <img src={Logo} alt='logo' />
        </div>
        <div className='header-title login-block__title'>
          <h2>Sign Up</h2>
        </div>
        <div className='login-block__body'>
          <div className='input-group'>
            <label htmlFor='username'>E-mail</label>
            <Input
              id='email'
              type='text'
              name='email'
              placeholder='Your email'
              value={form.email}
              onChange={(e) => {
                changeForm("email")(e);
              }}
            />
          </div>
          <div className='space-b-3'></div>
          <div className='input-group'>
            <label htmlFor='username'>Username</label>
            <Input
              id='username'
              type='text'
              name='username'
              placeholder='Your username'
              value={form.username}
              onChange={(e) => {
                changeForm("username")(e);
              }}
            />
          </div>
          <div className='space-b-3'></div>
          <div className='input-group'>
            <label htmlFor='password'>Password</label>
            <Input
              id='password'
              type={isTextTypeP ? "password" : "text"}
              name='password'
              placeholder='at least 8 characters'
              value={form.password}
              onChange={(e) => {
                changeForm("password")(e);
              }}
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
          <div className='input-group space-b-3'>
            <label htmlFor='password'>Confirm Password</label>
            <Input
              id='confirm-password'
              type={isTextTypeCP ? "password" : "text"}
              name='confirm-password'
              placeholder='confirm your password'
              value={form.confirmPassword}
              onChange={(e) => {
                changeForm("confirmPassword")(e);
              }}
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
            <Button id='submitBtn' onClick={handleSignUp}>
              Submit
            </Button>
            <Button extraClasses={["cancel-button"]} onClick={emptyForm}>
              Reset
            </Button>
          </div>
          <div className='button-container button-container--right space-t-3'>
            <div>Do you have an account?</div>
            <Link
              className='link-button'
              onClick={onChangeStage(stageType.LOGIN)}
            >
              Sign in
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
