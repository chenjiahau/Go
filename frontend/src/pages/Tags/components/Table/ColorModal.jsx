import { useState, useEffect, useCallback } from "react";
import Button from "@/components/Button";
import Dropdown from "@/components/Dropdown";

const ColorModal = (props) => {
  const {
    isOpen,
    onClose,
    onConfirm,
    colorCategories,
    colors,
    selectedTagColor,
  } = props;
  // State
  const [colorCategoryDropdown, setColorCategoryDropdown] = useState([]);
  const [colorDropdown, setColorDropdown] = useState([]);

  // Method
  const init = useCallback(() => {
    if (
      !selectedTagColor.id ||
      colorCategories.length === 0 ||
      colors.length === 0
    ) {
      return;
    }

    const updatedColorCategoryDropdown = colorCategories.map(
      (colorCategory) => {
        return {
          id: colorCategory.id,
          name: colorCategory.name,
          selected: colorCategory.id === selectedTagColor.colorCategoryId,
        };
      }
    );

    const activeColorCategory = updatedColorCategoryDropdown.find(
      (colorCategory) => colorCategory.selected
    );

    const updatedColorDropdown = colors
      .filter((color) => color.categoryId === activeColorCategory.id)
      .map((color) => {
        return {
          id: color.id,
          name: color.name,
          bgcolor: color.hexCode,
          selected: color.id === selectedTagColor.colorId,
        };
      });

    setColorCategoryDropdown(updatedColorCategoryDropdown);
    setColorDropdown(updatedColorDropdown);
  }, [colorCategories, colors, selectedTagColor]);

  const getSelectedColor = () => {
    return colorDropdown.find((color) => color.selected);
  };

  // Side effect
  useEffect(() => {
    init();
  }, [init]);

  return (
    <>
      <div className={`modal ${isOpen ? "display" : "hidden"}`}>
        <div className='modal-header'>
          <div className='modal-title'>Change color</div>
          <div className='modal-close' onClick={onClose}>
            &times;
          </div>
        </div>
        <div className='modal-body'>
          <div style={{ width: "100%" }}>
            <Dropdown
              zIndex={8}
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
          <div style={{ width: "100%" }}>
            <Dropdown
              zIndex={7}
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
        </div>
        <div className='modal-footer'>
          <Button extraClasses={["cancel-button"]} onClick={onClose}>
            Close
          </Button>
          <div className='space-r-2'></div>
          <Button onClick={() => onConfirm(getSelectedColor())}>Save</Button>
        </div>
      </div>
      <div className={`overlay ${isOpen ? "display" : "hidden"}`}></div>
    </>
  );
};

export default ColorModal;
