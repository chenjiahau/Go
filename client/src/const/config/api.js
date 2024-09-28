const resource = {
  // API
  // Login
  SIGNIN: "/sign-in",
  ACTIVE_ACCOUNT: "/activate-account/:token",
  SIGNUP: "/sign-up",
  CREATE_FORGOT_PASSWORD: "/create-reset-password",
  CHECK_RESET_PASSWORD_TOKEN: "/check-reset-password-token/:email/:token",
  RESET_PASSWORD: "/reset-password",
  VERIFY_TOKEN: "/auth/verify-token",
  // Member role
  MEMBER_ROLES: "/auth/member-roles",
  // Member
  MEMBERS: "/auth/members",
  NUMBER_OF_MEMBERS: "/auth/members/total",
  MEMBERS_BY_PAGE: "/auth/members/page/:page/size/:size",
  ADD_MEMBER: "/auth/member",
  EDIT_MEMBER: "/auth/member/:id",
  // Statistic
  STATISTIC_MOST_PUBLISHERS: "/auth/statistic/most-publishers",
  STATISTIC_MOST_COMMENTS: "/auth/statistic/most-comments",
  // Category
  CATEGORIES: "/auth/categories",
  NUMBER_OF_CATEGORIES: "/auth/categories/total",
  CATEGORIES_BY_PAGE: "/auth/categories/page/:page/size/:size",
  ADD_CATEGORY: "/auth/category",
  EDIT_CATEGORY: "/auth/category/:id",
  // Subcategory
  SUBCATEGORIES: "/auth/category/:id/subcategories",
  NUMBER_OF_SUBCATEGORIES: "/auth/category/:id/subcategories/total",
  SUBCATEGORIES_BY_PAGE: "/auth/category/:id/subcategories/page/:page/size/:size",
  ADD_SUBCATEGORY: "/auth/category/:id/subcategory",
  EDIT_SUBCATEGORY: "/auth/category/:id/subcategory/:subId",
  // Color
  COLOR_CATEGORIES: "/auth/color-categories",
  COLORS: "/auth/colors",
  // Tag
  TAGS: "/auth/tags",
  NUMBER_OF_TAGS: "/auth/tags/total",
  TAGS_BY_PAGE: "/auth/tags/page/:page/size/:size",
  ADD_TAG: "/auth/tag",
  EDIT_TAG: "/auth/tag/:id",
  // Document
  DOCUMENTS: "/auth/documents",
  NUMBER_OF_DOCUMENTS: "/auth/documents/total",
  DOCUMENTS_BY_PAGE: "/auth/documents/page/:page/size/:size",
  ADD_DOCUMENT: "/auth/document",
  EDIT_DOCUMENT: "/auth/document/:id",
  DOCUMENT_COMMENTS: "/auth/document/:id/comments",
  ADD_DOCUMENT_COMMENT: "/auth/document/:id/comment",
  EDIT_DOCUMENT_COMMENT: "/auth/document/:id/comment/:commentId",
  SEARCH_DOCUMENTS: "/auth/documents/search",
}

export default {
  resource,
}