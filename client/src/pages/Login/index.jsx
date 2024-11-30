import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";

// Const
import routerConfig from "@/const/config/router";

// Slice
import { userActions } from "@/store/slices/user";

// Component
import SignIn from "./component/SignIn";
import SignUp from "./component/SignUp";
import ForgotPassword from "./component/ForgetPassword";

// Util
import apiHandler from "@/util/api.util";

const stageType = {
  LOGIN: 0,
  REGISTER: 1,
  FORGOT_PASSWORD: 2,
};

const successMessage = {
  signin: "Login successfully",
  signup: "Register successfully, please check your email to verify",
  forgotPassword: "Please check your email to reset password",
};

const errorMessage = {
  fieldsNotFill: "Please fill in all fields",
  passwordNotMatch: "Password and confirm password are not matched",
};

const Home = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  // State
  const [stage, setStage] = useState(0);

  // Method
  const changeStage = (stage) => () => {
    setStage(stage);
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
      <SignIn
        successMessage={successMessage}
        errorMessage={errorMessage}
        stageType={stageType}
        onChangeStage={changeStage}
      />
    );
  } else if (stage === stageType.REGISTER) {
    content = (
      <SignUp
        successMessage={successMessage}
        errorMessage={errorMessage}
        stageType={stageType}
        onChangeStage={changeStage}
      />
    );
  } else if (stage === stageType.FORGOT_PASSWORD) {
    content = (
      <ForgotPassword
        successMessage={successMessage}
        errorMessage={errorMessage}
        stageType={stageType}
        onChangeStage={changeStage}
      />
    );
  }

  return <>{content}</>;
};

export default Home;
