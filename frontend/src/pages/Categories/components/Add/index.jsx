import { useState } from "react";

// Const
import apiConfig from "@/const/config/api";

// Component
import Input from "@/components/Input";
import Button from "@/components/Button";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  category: "Category is required.",
  duplicated: "Category is duplicated.",
};

const Add = (props) => {
  const { onInitialization } = props;

  // State
  const [category, setCategory] = useState("");

  // Method
  const addCategory = async () => {
    if (category === "") {
      messageUtil.showErrorMessage(errorMessage.category);
      return;
    }

    try {
      await apiHandler.post(apiConfig.resource.ADD_CATEGORY, {
        name: category,
      });
      setCategory("");
      messageUtil.showSuccessMessage(commonMessage.success);
      onInitialization();
    } catch (error) {
      if (error.response.data.error.code === 500) {
        messageUtil.showErrorMessage(errorMessage.duplicated);
        return;
      }
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  return (
    <div className='section'>
      <div className='input-group'>
        <Input
          id='category'
          type='text'
          name='category'
          autoComplete='off'
          placeholder='New category'
          value={category}
          onChange={(e) => setCategory(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              addCategory();
            }
          }}
          required
        />
      </div>
      <div className='space-t-2'></div>
      <div className='text-right'>
        <Button onClick={addCategory}>
          <i className='fa-solid fa-plus'></i>
          <span className='space-l-1'>Add</span>
        </Button>
      </div>
    </div>
  );
};

export default Add;
