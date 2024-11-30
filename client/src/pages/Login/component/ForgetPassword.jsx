import "../module.css";
import logo from "@/assets/img/brand.png";

import { useState } from "react";

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

const ForgotPassword = ({
  successMessage,
  errorMessage,
  stageType,
  onChangeStage,
}) => {
  // State
  const [form, setForm] = useState({
    email: "",
  });
  const [isLoading, setIsLoading] = useState(false);

  // Method
  const changeForm = (key) => (value) => {
    setForm({ ...form, [key]: value });
  };

  const emptyForm = () => {
    setForm({
      email: "",
    });
  };

  const handleForgotPassword = async () => {
    if (!form.email) {
      messageUtil.showErrorMessage(errorMessage.fieldsNotFill);
      return;
    }

    const payload = {
      email: form.email,
    };

    // Call API
    try {
      setIsLoading(true);
      await apiHandler.post(apiConfig.resource.CREATE_FORGOT_PASSWORD, payload);
      emptyForm();
      messageUtil.showSuccessMessage(successMessage.forgotPassword);
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
              <MainTitle extraClasses={["text-general"]}>
                Forgot Password
              </MainTitle>
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
          </div>
          <div className='login-block login-button'>
            <ButtonBox onClick={handleForgotPassword} isSave={true}>
              Reset Password
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

export default ForgotPassword;
