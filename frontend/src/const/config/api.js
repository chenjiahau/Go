const resource = {
  // API
  // Login
  SIGNIN: "/sign-in",
  SIGNUP: "/sign-up",
  VERIFY_TOKEN: "/auth/verify-token",
  // Category
  CATEGORIES: "/categories",
  NUMBER_OF_CATEGORIES: "/categories/total",
  CATEGORIES_BY_PAGE: "/categories/page/:page/size/:size",
  ADD_CATEGORY: "/category",
  EDIT_CATEGORY: "/category/:id",
  // Subcategory
  SUBCATEGORIES: "/category/:id/subcategories",
  NUMBER_OF_SUBCATEGORIES: "/category/:id/subcategories/total",
  SUBCATEGORIES_BY_PAGE: "/category/:id/subcategories/page/:page/size/:size",
  ADD_SUBCATEGORY: "/category/:id/subcategory",
  EDIT_SUBCATEGORY: "/category/:id/subcategory/:subId",
}

export default {
  resource,
}