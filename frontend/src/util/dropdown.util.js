// Dropdown
const openDropdown = (className = 'dropdown') => {
  const optionMenuList = document.querySelectorAll(`.${className}`);

  for (let i = 0; i < optionMenuList.length; i++) {
    const optionMenu = optionMenuList[i];
    const selectBtn = optionMenu.querySelector(".item");
    const options = optionMenu.querySelectorAll(".option");
    const text = optionMenu.querySelector(".text");

    selectBtn.addEventListener("click", () => {
      const arrowUp = selectBtn.querySelector('.arrow-up');
      const arrowDown = selectBtn.querySelector('.arrow-down');

      arrowUp.classList.toggle("hidden");
      arrowDown.classList.toggle("hidden");

      optionMenu.classList.toggle("active");
    });

    options.forEach(option => {
      option.addEventListener("click", () => {
        options.forEach(option => option.classList.remove("is-active"));
        option.classList.add("is-active");
        const selectedOption = option.innerText;
        text.innerText = selectedOption;
        optionMenu.classList.remove("active");
        optionMenu.querySelector('.arrow-up').classList.add("hidden");
        optionMenu.querySelector('.arrow-down').classList.remove("hidden");
      });
    });
  }
}

const closeDropdown = (className = 'dropdown') => {
  const body = document.querySelector("body");
  const dropdowns = document.querySelectorAll(`.${className}`);

  body.addEventListener("click", (e) => {
    dropdowns.forEach(dropdown => {
      if (!dropdown.contains(e.target)) {
        dropdown.classList.remove("active");
        const arrowUp = dropdown.querySelector('.arrow-up');
        const arrowDown = dropdown.querySelector('.arrow-down');
        arrowUp.classList.add("hidden");
        arrowDown.classList.remove("hidden");
      }
    });
  });
}

export default { openDropdown, closeDropdown };