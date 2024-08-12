import { toast } from "react-toastify";

export const commonMessage = {
  success: "Success",
  error: "Something went wrong, please try again",
};

const toastConfig = {
  position: "top-center",
  autoClose: 3000,
  hideProgressBar: true,
  closeOnClick: true,
  pauseOnHover: true,
  draggable: true,
};

const dismissAll = () => {
  toast.dismiss();
}

const showSuccessMessage = (message) => {
  dismissAll();
  toast.success(message, toastConfig);
}

const showErrorMessage = (message) => {
  dismissAll();
  toast.error(message, toastConfig);
}

const showInfoMessage = (message) => {
  dismissAll();
  toast.info(message, toastConfig);
}

export default {
  showSuccessMessage,
  showErrorMessage,
  showInfoMessage,
};