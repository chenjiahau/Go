const resource = {
  // API
  // Login
  SIGNIN: "/sign-in",
  SIGNUP: "/sign-up",
  VERIFY_TOKEN: "/auth/verify-token",
  // Category
  CATEGORIES: "/categories",
  ADD_CATEGORY: "/category",
  EDIT_CATEGORY: "/category/:id",
  // Subcategory
  SUBCATEGORIES: "/category/:id/subcategories",
  ADD_SUBCATEGORY: "/category/:id/subcategory",
  EDIT_SUBCATEGORY: "/category/:id/subcategory/:subId",
}

export default {
  resource,
}