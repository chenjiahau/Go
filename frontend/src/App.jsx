// import "bootstrap/dist/css/bootstrap.css";
import "react-toastify/dist/ReactToastify.css";
import "./App.css";

import { HashRouter, Routes, Route } from "react-router-dom";
import { ToastContainer } from "react-toastify";

import routerConfig from "@/const/config/router";

import Layout from "@/ui/Layout";
import Login from "@/pages/Login";
import Dashboard from "@/pages/Dashboard";
import Categories from "@/pages/Categories";
import Category from "@/pages/Category";
import Tags from "@/pages/Tags";

function App() {
  return (
    <HashRouter>
      <Routes>
        <Route path={routerConfig.routes.LOGIN} element={<Login />} />
        <Route path={routerConfig.routes.HOME} element={<Layout />}>
          <Route path={routerConfig.routes.DASHBOARD} element={<Dashboard />} />
          <Route
            path={routerConfig.routes.CATEGORIES}
            element={<Categories />}
          />
          <Route path={routerConfig.routes.CATEGORY} element={<Category />} />
          <Route path={routerConfig.routes.TAGS} element={<Tags />} />
        </Route>
      </Routes>
      <ToastContainer />
    </HashRouter>
  );
}

export default App;
