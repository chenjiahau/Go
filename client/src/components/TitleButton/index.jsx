import "./module.scss";

const TitleButton = ({ title, onClick }) => {
  return (
    <button className='title-button' onClick={onClick}>
      {title}
    </button>
  );
};

export default TitleButton;
