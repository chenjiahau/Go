import "./module.scss";
import Logo from "@/assets/img/brand.png";

import { useState, useEffect } from "react";
import PropTypes from "prop-types";
import { Link, useNavigate } from "react-router-dom";
import _ from "lodash";

import Footer from "@/ui/Footer";
import apiHandler from "@/util/api.util";
import messageUtil from "../../util/message.util";

const stageType = {
  LOGIN: 0,
  REGISTER: 1,
};

const successMessage = {
  SINGIN: "Login successfully",
  SINGUP: "Login successfully, please back to login",
};

const errorMessage = {
  FIELDS_NOT_FILL: "Please fill in all fields",
  PASSWORD_NOT_MATCH: "Password and confirm password are not matched",
};

const Home = (props) => {
  const { user, onSaveUser } = props;
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
      const response = await apiHandler.post("/sign-in", payload);
      apiHandler.setToken(response.data.data.token);
      localStorage.setItem("user", JSON.stringify(response.data.data));
      messageUtil.showSuccessMessage(successMessage.SINGIN);
      onSaveUser(response.data);
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
      await apiHandler.post("/sign-up", payload);
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
      onSaveUser(userSavedInLocalStorage);
      navigate("/dashboard");
    }
  }, [navigate, onSaveUser, user]);

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
            <form action='' method='post'>
              <div className='form-group mb-3'>
                <label htmlFor='username'>E-mail</label>
                <input
                  type='text'
                  name='email'
                  id='email'
                  className='form-control'
                  placeholder='Your email'
                  required
                  value={signInData.email}
                  onChange={(e) => {
                    changeSignInData("email")(e);
                  }}
                />
              </div>
              <div className='form-group mb-4'>
                <label htmlFor='password'>Password</label>
                <input
                  type='password'
                  name='password'
                  id='password'
                  className='form-control'
                  placeholder='at least 8 characters'
                  required
                  value={signInData.password}
                  onChange={(e) => {
                    changeSignInData("password")(e);
                  }}
                />
              </div>
              <div className='form-group d-flex justify-content-between'>
                <button
                  id='submitBtn'
                  type='button'
                  className='button mr-2'
                  onClick={handleSignIn}
                >
                  Submit
                </button>
                <button
                  type='button'
                  className='button cancel-button'
                  onClick={emptySignInData}
                >
                  Reset
                </button>
              </div>
              <div className='d-flex justify-content-end mt-4'>
                <div className='me-2'>Do you have not an account?</div>
                <Link
                  className='link-button'
                  onClick={changeStage(stageType.REGISTER)}
                >
                  Sign up
                </Link>
              </div>
            </form>
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
            <div className='form-group mb-3'>
              <label htmlFor='username'>E-mail</label>
              <input
                type='text'
                name='email'
                id='email'
                className='form-control'
                placeholder='Your email'
                required
                value={signUpData.email}
                onChange={(e) => {
                  changeSignUpData("email")(e);
                }}
              />
            </div>
            <div className='form-group mb-3'>
              <label htmlFor='username'>Username</label>
              <input
                type='text'
                name='username'
                id='username'
                className='form-control'
                placeholder='Your username'
                required
                value={signUpData.username}
                onChange={(e) => {
                  changeSignUpData("username")(e);
                }}
              />
            </div>
            <div className='form-group mb-3'>
              <label htmlFor='password'>Password</label>
              <input
                type='password'
                name='password'
                id='password'
                className='form-control'
                placeholder='at least 8 characters'
                required
                value={signUpData.password}
                onChange={(e) => {
                  changeSignUpData("password")(e);
                }}
              />
            </div>
            <div className='form-group mb-4'>
              <label htmlFor='password'>Confirm Password</label>
              <input
                type='password'
                name='confirm-password'
                id='confirm-password'
                className='form-control'
                placeholder='confirm your password'
                required
                value={signUpData.confirmPassword}
                onChange={(e) => {
                  changeSignUpData("confirmPassword")(e);
                }}
              />
            </div>
            <div className='form-group d-flex justify-content-between'>
              <button
                id='submitBtn'
                type='button'
                className='button mr-2'
                onClick={handleSignUp}
              >
                Submit
              </button>
              <button
                type='button'
                className='button cancel-button'
                onClick={emptySignUpData}
              >
                Reset
              </button>
            </div>
            <div className='d-flex justify-content-end mt-4'>
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

Home.propTypes = {
  user: PropTypes.object,
  onSaveUser: PropTypes.func.isRequired,
};

export default Home;
