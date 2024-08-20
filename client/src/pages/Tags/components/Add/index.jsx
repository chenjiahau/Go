import { useState, useCallback, useEffect } from "react";

// Const
import apiConfig from "@/const/config/api";

// Component
import Input from "@/components/Input";
import Button from "@/components/Button";
import Dropdown from "@/components/Dropdown";

// Util
import apiHandler from "@/util/api.util";
import messageUtil, { commonMessage } from "@/util/message.util";

const errorMessage = {
  tag: "Tag is required.",
  duplicated: "Tag is duplicated.",
};

const Add = (props) => {
  const { onInitialization, colorCategories, colors } = props;

  // State
  const [colorCategoryDropdown, setColorCategoryDropdown] = useState([]);
  const [colorDropdown, setColorDropdown] = useState([]);
  const [tag, setTag] = useState("");

  // Method
  const init = useCallback(() => {
    if (colorCategories.length === 0 || colors.length === 0) {
      return;
    }

    const updatedColorCategoryDropdown = colorCategories.map(
      (colorCategory) => {
        return {
          id: colorCategory.id,
          name: colorCategory.name,
          selected: false,
        };
      }
    );
    updatedColorCategoryDropdown[0].selected = true;

    const updatedColorDropdown = colors
      .filter(
        (color) => color.categoryId === updatedColorCategoryDropdown[0].id
      )
      .map((color) => {
        return {
          id: color.id,
          name: color.name,
          bgcolor: color.hexCode,
          selected: false,
        };
      });
    updatedColorDropdown[0].selected = true;

    setColorCategoryDropdown(updatedColorCategoryDropdown);
    setColorDropdown(updatedColorDropdown);
  }, [colorCategories, colors]);

  const addTag = async () => {
    if (tag === "") {
      messageUtil.showErrorMessage(errorMessage.tag);
      return;
    }

    try {
      await apiHandler.post(apiConfig.resource.ADD_TAG, {
        name: tag,
        colorId: colorDropdown.find((color) => color.selected).id,
      });
      setTag("");
      init();
      messageUtil.showSuccessMessage(commonMessage.success);
      onInitialization();
    } catch (error) {
      if (error.response.data.code === 2423) {
        messageUtil.showErrorMessage(errorMessage.duplicated);
        return;
      }
      messageUtil.showErrorMessage(commonMessage.error);
    }
  };

  // Side effect
  useEffect(() => {
    init();
  }, [init]);

  return (
    <div className='section'>
      <div className='input-group'>
        <Input
          id='tag'
          type='text'
          name='tag'
          autoComplete='off'
          placeholder='New tag'
          value={tag}
          onChange={(e) => setTag(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              addTag();
            }
          }}
          required
        />
      </div>
      <div className='space-t-2'></div>
      <div>
        <Dropdown
          zIndex={6}
          list={colorCategoryDropdown}
          selectedItem={colorCategoryDropdown.find((item) => item.selected)}
          onChange={(item) => {
            const updatedColorCategoryDropdown = colorCategoryDropdown.map(
              (colorCategory) => {
                return {
                  ...colorCategory,
                  selected: colorCategory.id === item.id,
                };
              }
            );
            setColorCategoryDropdown(updatedColorCategoryDropdown);

            const updatedColorDropdown = colors
              .filter((color) => color.categoryId === item.id)
              .map((color) => {
                return {
                  id: color.id,
                  name: color.name,
                  bgcolor: color.hexCode,
                  selected: false,
                };
              });
            updatedColorDropdown[0].selected = true;
            setColorDropdown(updatedColorDropdown);
          }}
        />
      </div>
      <div className='space-t-2'></div>
      <div>
        <Dropdown
          zIndex={5}
          hasBackground={true}
          list={colorDropdown}
          selectedItem={colorDropdown.find((item) => item.selected)}
          onChange={(item) => {
            const updatedColorDropdown = colorDropdown.map((color) => {
              return {
                ...color,
                selected: color.id === item.id,
              };
            });
            setColorDropdown(updatedColorDropdown);
          }}
        />
      </div>
      <div className='space-t-2'></div>
      <div className='text-right'>
        <Button onClick={addTag}>
          <i className='fa-solid fa-plus'></i>
          <span className='space-l-1'>Add</span>
        </Button>
      </div>
    </div>
  );
};

export default Add;
