const resource = {
  // API
  // Login
  SIGNIN: "/sign-in",
  SIGNUP: "/sign-up",
  VERIFY_TOKEN: "/auth/verify-token",
  // Member role
  MEMBER_ROLES: "/member-roles",
  // Member
  MEMBERS: "/members",
  NUMBER_OF_MEMBERS: "/members/total",
  MEMBERS_BY_PAGE: "/members/page/:page/size/:size",
  ADD_MEMBER: "/member",
  EDIT_MEMBER: "/member/:id",
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
  // Color
  COLOR_CATEGORIES: "/color-categories",
  COLORS: "/colors",
  // Tag
  TAGS: "/tags",
  NUMBER_OF_TAGS: "/tags/total",
  TAGS_BY_PAGE: "/tags/page/:page/size/:size",
  ADD_TAG: "/tag",
  EDIT_TAG: "/tag/:id",
  // Document
  DOCUMENTS: "/documents",
  NUMBER_OF_DOCUMENTS: "/documents/total",
  DOCUMENTS_BY_PAGE: "/documents/page/:page/size/:size",
  ADD_DOCUMENT: "/document",
  EDIT_DOCUMENT: "/document/:id",
}

export default {
  resource,
}