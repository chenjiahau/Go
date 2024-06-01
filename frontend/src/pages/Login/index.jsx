import "./module.scss";
import Logo from "@/assets/img/brand.png";

import { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";
import _ from "lodash";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Slice
import { userActions } from "@/store/slices/user";

// Component
import Footer from "@/ui/Footer";
import Input from "@/components/Input";
import Button from "@/components/Button";

// Util
import apiHandler from "@/util/api.util";
import messageUtil from "../../util/message.util";

const stageType = {
  LOGIN: 0,
  REGISTER: 1,
};

const successMessage = {
  SIGNIN: "Login successfully",
  SIGNUP: "Login successfully, please back to login",
};

const errorMessage = {
  FIELDS_NOT_FILL: "Please fill in all fields",
  PASSWORD_NOT_MATCH: "Password and confirm password are not matched",
};

const Home = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  // State
  const [stage, setStage] = useState(0);
  const [signInData, setSignInData] = useState({
    email: "",
    password: "",
  });
  const [signUpData, setSignUpData] = useState({
    email: "",
    username: "",
    password: "",
    confirmPassword: "",
  });

  // Method
  const changeStage = (stage) => () => {
    setStage(stage);
  };

  const changeSignInData = (key) => (e) => {
    let value = "";
    if (_.isObject(e)) {
      value = e.target.value;
    }

    setSignInData({ ...signInData, [key]: value });
  };

  const emptySignInData = () => {
    setSignInData({
      email: "",
      password: "",
    });
  };

  const changeSignUpData = (key) => (e) => {
    let value = "";
    if (_.isObject(e)) {
      value = e.target.value;
    }

    setSignUpData({ ...signUpData, [key]: value });
  };

  const emptySignUpData = () => {
    setSignUpData({
      email: "",
      username: "",
      password: "",
      confirmPassword: "",
    });
  };

  const handleSignIn = async () => {
    if (!signInData.email || !signInData.password) {
      messageUtil.showErrorMessage(errorMessage.FIELDS_NOT_FILL);
      return;
    }

    const payload = {
      email: signInData.email,
      password: signInData.password,
    };

    // Call API
    try {
      const response = await apiHandler.post(
        apiConfig.resource.SIGNIN,
        payload
      );
      apiHandler.setToken(response.data.data.token);
      messageUtil.showSuccessMessage(successMessage.SIGNIN);
      dispatch(userActions.setUser(response.data.data));
      localStorage.setItem("user", JSON.stringify(response.data.data));
      navigate(routerConfig.routes.DASHBOARD);
    } catch (error) {
      messageUtil.showErrorMessage(apiHandler.extractErrorMessage(error));
    }
  };

  const handleSignUp = async () => {
    if (
      !signUpData.email ||
      !signUpData.username ||
      !signUpData.password ||
      !signUpData.confirmPassword
    ) {
      messageUtil.showErrorMessage(errorMessage.FIELDS_NOT_FILL);
      return;
    }
    if (signUpData.password !== signUpData.confirmPassword) {
      messageUtil.showErrorMessage(errorMessage.PASSWORD_NOT_MATCH);
      return;
    }

    const payload = {
      email: signUpData.email,
      username: signUpData.username,
      password: signUpData.password,
      confirmPassword: signUpData.confirmPassword,
    };

    // Call API
    try {
      await apiHandler.post(apiConfig.resource.SIGNUP, payload);
      emptySignUpData();
      messageUtil.showSuccessMessage(successMessage.SIGNUP);
    } catch (error) {
      messageUtil.showErrorMessage(apiHandler.extractErrorMessage(error));
    }
  };

  // Side effect
  useEffect(() => {
    const savedUser = localStorage.getItem("user");

    if (savedUser) {
      const userSavedInLocalStorage = JSON.parse(savedUser);

      apiHandler.setToken(userSavedInLocalStorage.token);
      dispatch(userActions.setUser(savedUser));
      localStorage.setItem("user", JSON.stringify(userSavedInLocalStorage));
      navigate(routerConfig.routes.HOME);
    }
  }, []);

  let content = null;
  if (stage === stageType.LOGIN) {
    content = (
      <div className='login-section'>
        <div className='login-block'>
          <div className='login-block__logo'>
            <img src={Logo} alt='logo' />
          </div>
          <div className='header-title login-block__title'>
            <h2>Sign In</h2>
          </div>
          <div className='login-block__body'>
            <div className='input-group'>
              <label htmlFor='username'>E-mail</label>
              <Input
                id='email'
                type='text'
                name='email'
                placeholder='Your email'
                value={signInData.email}
                onChange={(e) => {
                  changeSignInData("email")(e);
                }}
              />
            </div>
            <div className='space-b-3'></div>
            <div className='input-group'>
              <label htmlFor='password'>Password</label>
              <Input
                id='password'
                type='password'
                name='password'
                placeholder='at least 8 characters'
                value={signInData.password}
                onChange={(e) => {
                  changeSignInData("password")(e);
                }}
              />
            </div>
            <div className='space-b-4'></div>
            <div className='button-container'>
              <Button id='submitBtn' onClick={handleSignIn}>
                Submit
              </Button>
              <Button
                id='cancelBtn'
                extraClasses={["cancel-button"]}
                onClick={emptySignInData}
              >
                Reset
              </Button>
            </div>
            <div className='button-container button-container--right'>
              <div className='me-2'>Do you have not an account?</div>
              <Link
                className='link-button'
                onClick={changeStage(stageType.REGISTER)}
              >
                Sign up
              </Link>
            </div>
          </div>
        </div>
      </div>
    );
  } else if (stage === stageType.REGISTER) {
    content = (
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
                value={signUpData.email}
                onChange={(e) => {
                  changeSignUpData("email")(e);
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
                value={signUpData.username}
                onChange={(e) => {
                  changeSignUpData("username")(e);
                }}
              />
            </div>
            <div className='space-b-3'></div>
            <div className='input-group'>
              <label htmlFor='password'>Password</label>
              <Input
                id='password'
                type='password'
                name='password'
                placeholder='at least 8 characters'
                value={signUpData.password}
                onChange={(e) => {
                  changeSignUpData("password")(e);
                }}
              />
            </div>
            <div className='space-b-3'></div>
            <div className='input-group space-b-3'>
              <label htmlFor='password'>Confirm Password</label>
              <Input
                id='confirm-password'
                type='password'
                name='confirm-password'
                placeholder='confirm your password'
                value={signUpData.confirmPassword}
                onChange={(e) => {
                  changeSignUpData("confirmPassword")(e);
                }}
              />
            </div>
            <div className='space-b-4'></div>
            <div className='button-container'>
              <Button id='submitBtn' onClick={handleSignUp}>
                Submit
              </Button>
              <Button
                extraClasses={["cancel-button"]}
                onClick={emptySignUpData}
              >
                Reset
              </Button>
            </div>
            <div className='button-container button-container--right'>
              <div className='me-2'>Do you have an account?</div>
              <Link
                className='link-button'
                onClick={changeStage(stageType.LOGIN)}
              >
                Sign in
              </Link>
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <>
      {content}
      <Footer />
    </>
  );
};

export default Home;
