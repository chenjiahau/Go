import "../module.css";
import logo from "@/assets/img/brand.png";

import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons";
// import ReCAPTCHA from "react-google-recaptcha";

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

// Slice
import { userActions } from "@/store/slices/user";

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

const SignIn = ({ successMessage, errorMessage, stageType, onChangeStage }) => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  // State
  const [form, setForm] = useState({
    email: "",
    password: "",
  });
  const [isPasswordType, setIsPasswordType] = useState(true);
  // const [recaptchaToken, setRecaptchaToken] = useState(null);
  // const recaptchaRef = useRef();
  const [isLoading, setIsLoading] = useState(false);

  // Method
  const changeForm = (key) => (value) => {
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
      setIsLoading(true);
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
              <MainTitle extraClasses={["text-general"]}>Sign In</MainTitle>
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
              <FormLabel forName='password'>Password</FormLabel>
              <InputBox
                type={isPasswordType ? "password" : "text"}
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
                    icon={isPasswordType ? faEyeSlash : faEye}
                    onClick={() => setIsPasswordType(!isPasswordType)}
                  />
                }
              </InputBox>
            </FormGroup>
          </div>
          <div className='login-block login-button'>
            <ButtonBox onClick={handleSignIn} isSave={true}>
              Sign In
            </ButtonBox>
            <ButtonBox
              onClick={emptyForm}
              isClose={true}
              extraClasses={["cancel-shadow"]}
            >
              Reset
            </ButtonBox>
          </div>
          <div className='login-block login-link'>
            <div>Did you forget your password?</div>
            <LinkButton
              to={routerConfig.routes.LOGIN}
              onClick={onChangeStage(stageType.FORGOT_PASSWORD)}
              title='Forget Password'
            />
          </div>
          <div className='login-block login-link'>
            <div>Do you have not an account?</div>
            <LinkButton
              to={routerConfig.routes.LOGIN}
              onClick={onChangeStage(stageType.REGISTER)}
              title='Sign Up'
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

export default SignIn;
