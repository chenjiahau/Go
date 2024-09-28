// import "bootstrap/dist/css/bootstrap.css";
import "react-toastify/dist/ReactToastify.css";
import "./App.css";

import { HashRouter, Routes, Route } from "react-router-dom";
import { ToastContainer } from "react-toastify";

import routerConfig from "@/const/config/router";

import Layout from "@/ui/Layout";
import Login from "@/pages/Login";
import ActiveAccount from "./pages/ActiveAccount";
import ResetPassword from "./pages/ResetPassword";
import Dashboard from "@/pages/Dashboard";
import Members from "@/pages/Members";
import Categories from "@/pages/Categories";
import Category from "@/pages/Category";
import Tags from "@/pages/Tags";
import Documents from "./pages/Documents";
import Document from "./pages/Document";
import AddDocument from "./pages/Documents/components/AddDocument";
import EditDocument from "./pages/Document/components/EditDocument";
import AddDocumentComment from "./pages/Document/components/AddDocumentComment";
import EditDocumentComment from "./pages/Document/components/EditDocumentComment";

function App() {
  return (
    <HashRouter>
      <Routes>
        <Route path={routerConfig.routes.LOGIN} element={<Login />} />
        <Route
          path={routerConfig.routes.ACTIVE_ACCOUNT}
          element={<ActiveAccount />}
        />
        <Route
          path={routerConfig.routes.RESET_PASSWORD}
          element={<ResetPassword />}
        />
        <Route path={routerConfig.routes.HOME} element={<Layout />}>
          <Route path={routerConfig.routes.DASHBOARD} element={<Dashboard />} />
          <Route path={routerConfig.routes.MEMBERS} element={<Members />} />
          <Route
            path={routerConfig.routes.CATEGORIES}
            element={<Categories />}
          />
          <Route path={routerConfig.routes.CATEGORY} element={<Category />} />
          <Route path={routerConfig.routes.TAGS} element={<Tags />} />
          <Route path={routerConfig.routes.DOCUMENTS} element={<Documents />} />
          <Route path={routerConfig.routes.DOCUMENT} element={<Document />} />
          <Route
            path={routerConfig.routes.ADD_DOCUMENT}
            element={<AddDocument />}
          />
          <Route
            path={routerConfig.routes.EDIT_DOCUMENT}
            element={<EditDocument />}
          />
          <Route
            path={routerConfig.routes.ADD_DOCUMENT_COMMENT}
            element={<AddDocumentComment />}
          />
          <Route
            path={routerConfig.routes.EDIT_DOCUMENT_COMMENT}
            element={<EditDocumentComment />}
          />
        </Route>
      </Routes>
      <ToastContainer />
    </HashRouter>
  );
}

export default App;
