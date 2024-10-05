const overlayStyle = {
  position: "fixed",
  top: 0,
  left: 0,
  width: "100%",
  height: "100%",
  backgroundColor: "rgba(0, 0, 0, 0.5)",
  zIndex: 1000,
  justifyContent: "center",
  alignItems: "center",
  display: "none",
};

const spinnerStyle = {
  width: "50px",
  height: "50px",
  border: "8px solid rgba(255, 255, 255, 0.3)",
  borderTop: "8px solid #fff",
  borderRadius: "50%",
  animation: "spin 1s linear infinite",
};

const styles = document.createElement("style");
styles.innerHTML = `
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}`;
document.head.appendChild(styles);

const Spinner = () => {
  return (
    <div id='spinner' style={overlayStyle}>
      <div style={spinnerStyle}></div>
    </div>
  );
};

export default Spinner;
