import "../module.scss";
import Logo from "@/assets/img/brand.png";

import { useState, useRef } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";
import ReCAPTCHA from "react-google-recaptcha";
import _ from "lodash";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Slice
import { userActions } from "@/store/slices/user";

// Util
import apiHandler from "@/util/api.util";
import messageUtil from "@/util/message.util";

// Components
import Input from "@/components/Input";
import Button from "@/components/Button";

const SignIn = ({ successMessage, errorMessage, stageType, onChangeStage }) => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  // State
  const [form, setForm] = useState({
    email: "",
    password: "",
  });
  const [isTextTypeP, setIsTextTypeP] = useState(true);
  // const [recaptchaToken, setRecaptchaToken] = useState(null);
  // const recaptchaRef = useRef();

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
      password: "",
    });
  };

  const handleSignIn = async () => {
    // if (!form.email || !form.password || !recaptchaToken) {
    //   messageUtil.showErrorMessage(errorMessage.fieldsNotFill);
    //   return;
    // }
    if (!form.email || !form.password) {
      messageUtil.showErrorMessage(errorMessage.fieldsNotFill);
      return;
    }

    const payload = {
      email: form.email,
      password: form.password,
    };
    // const url = `${apiConfig.resource.SIGNIN}?recaptchaToken=${recaptchaToken}`;
    const url = `${apiConfig.resource.SIGNIN}`;

    // Call API
    try {
      const response = await apiHandler.post(url, payload);

      apiHandler.setToken(response.data.data.token);
      messageUtil.showSuccessMessage(successMessage.signin);
      dispatch(userActions.setUser(response.data.data));
      localStorage.setItem("user", JSON.stringify(response.data.data));
      navigate(routerConfig.routes.DASHBOARD);
    } catch (error) {
      // setRecaptchaToken(null);
      // recaptchaRef.current.reset();
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
              value={form.email}
              onChange={(e) => {
                changeForm("email")(e);
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
          <div className='space-b-4'></div>
          {/* <div className='recaptcah-group'>
            <ReCAPTCHA
              sitekey={import.meta.env.VITE_RECAPTCHA_SITE_KEY}
              onChange={(token) => setRecaptchaToken(token)}
              ref={recaptchaRef}
            />
          </div>
          <div className='space-b-3'></div> */}
          <div className='button-container'>
            <Button id='submitBtn' onClick={handleSignIn}>
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
          <div className='button-container button-container--right'>
            <div className='me-2'>Do you have not an account?</div>
            <Link
              className='link-button'
              onClick={onChangeStage(stageType.REGISTER)}
            >
              Sign up
            </Link>
          </div>
          <div className='button-container button-container--right'>
            <div className='me-2'>Did you forget your password?</div>
            <Link
              className='link-button'
              onClick={onChangeStage(stageType.FORGOT_PASSWORD)}
            >
              Forget password
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SignIn;
