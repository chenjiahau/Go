import { useEffect, useState } from "react";
import "./App.css";

const serverUrl = "http://localhost:8080";

function App() {
  const [appName, setAppName] = useState("");
  const [version, setVersion] = useState("");
  const [message, setMessage] = useState("");

  useEffect(() => {
    fetch(`${serverUrl}/api`)
      .then((res) => res.json())
      .then((res) => {
        setAppName(res.data.appName);
        setVersion(res.data.version);
        setMessage(res.data.message);
      });
  }, []);

  return (
    <>
      <h1>App Name: {appName}</h1>
      <h1>Version: {version}</h1>
      <h1>Message: {message}</h1>
    </>
  );
}

export default App;
