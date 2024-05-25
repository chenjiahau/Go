import { useState } from "react";

// Const
import apiConfig from "@/const/config/api";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  category: "Category is required.",
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
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  return (
    <div className='section'>
      <div className='mb-2'>
        <input
          type='text'
          name='category'
          id='category'
          className='form-control'
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
      <div className='text-right'>
        <button className='button w-100' onClick={addCategory}>
          <i className='fa-solid fa-plus'></i>
          <span className='ms-1'>Add</span>
        </button>
      </div>
    </div>
  );
};

export default Add;
