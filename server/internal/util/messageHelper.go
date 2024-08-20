package util

func GetReturnMessage(code int) map[string]interface{} {
	var resData map[string]interface{}
	var message string

	// Common
	if code >= 200 && code < 300 {
		message = CommonSuccessMessages[code]
	}

	if code >= 400 && code < 500 {
		message = CommonErrorMessages[code]
	}

	// User
	if code >= 1200 && code < 1300 {
		message = UserSuccessMessage[code]
	}

	if code >= 1400 && code < 1500 {
		message = UserErrorMessage[code]
	}

	// Member role
	if code >= 21200 && code < 21300 {
		message = MemberRoleSuccessMessage[code]
	}

	if code >= 21400 && code < 21500 {
		message = MemberRoleErrorMessage[code]
	}

	// Member
	if code >= 2200 && code < 2300 {
		message = MemberSuccessMessage[code]
	}

	if code >= 2400 && code < 2500 {
		message = MemberErrorMessage[code]
	}

	// Category
	if code >= 3200 && code < 3300 {
		message = CategorySuccessMessage[code]
	}

	if code >= 3400 && code < 3500 {
		message = CategoryErrorMessage[code]
	}

	// Subcategory
	if code >= 4200 && code < 4300 {
		message = SubcategorySuccessMessage[code]
	}

	if code >= 4400 && code < 4500 {
		message = SubcategoryErrorMessage[code]
	}

	// Tag
	if code >= 5200 && code < 5300 {
		message = TagSuccessMessage[code]
	}

	if code >= 5400 && code < 5500 {
		message = TagErrorMessage[code]
	}

	// Color
	if code >= 6200 && code < 6300 {
		message = ColorSuccessMessage[code]
	}

	if code >= 6400 && code < 6500 {
		message = ColorErrorMessage[code]
	}

	// Document
	if code >= 7200 && code < 7300 {
		message = DocumentSuccessMessage[code]
	}

	if code >= 7400 && code < 7500 {
		message = DocumentErrorMessage[code]
	}

	// Document comment
	if code >= 8200 && code < 8300 {
		message = DocumentCommentSuccessMessage[code]
	}

	if code >= 8400 && code < 8500 {
		message = DocumentCommentErrorMessage[code]
	}

	// Response
	resData = map[string]interface{} {
		"code": code,
		"message": message,
	}

	return resData
}

/**
 * Common
 */
var CommonSuccessMessages = map[int]string {
	200: "Success",
}

var CommonErrorMessages = map[int]string {
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	500: "Internal Server Error",
}

/**
 * User
 */
var UserSuccessMessage = map[int]string {
	// Sign up
	1201: "User created successfully",
	// Sign in
	1202: "User logged in successfully",
	// Verify token
	1203: "Success to verify token",
	// Log out
	1204: "User logged out successfully",
}

var	UserErrorMessage = map[int]string {
	// Sign up
	1401: "Invalid email, username, password, or confirm password",
	1402: "Email, username, password, and confirm password are required",
	1403: "Password and confirm password do not match",
	1404: "Email is already registered",
	1405: "Failed to hash password",
	1406: "Failed to create user",
	// Sign in
	1411: "Invalid email or password",
	1412: "Failed to create token",
	1413: "Failed to insert token",
	// Log out
	1421: "Failed to log out",
}

// Member role
var MemberRoleSuccessMessage = map[int]string {
	// Get all member roles
	21211: "Success to query all member roles",
}

var MemberRoleErrorMessage = map[int]string {
	// Get all member roles
	21411: "Failed to query all member roles",
}

// Member
var MemberSuccessMessage = map[int]string {
	// Create member
	2201: "Member created successfully",
	// Get all members
	2202: "Success to query all members",
	// Get member by page
	2203: "Success to query member by page",
	// Get member by id
	2204: "Success to query member by id",
	// Update member
	2205: "Member updated successfully",
	// Delete member
	2206: "Member deleted successfully",
}

var MemberErrorMessage = map[int]string {
	// Create member
	2401: "Invalid member name",
	2402: "Member name already exists",
	2403: "Failed to create member",
	2404: "Failed to create user member",
	// Get all members
	2411: "Failed to query all members",
	// Get member by page
	2412: "Failed to query member by page",
	// Get member by id
	2413: "Failed to query member by id",
	// Update member
	2421: "Invalid member name",
	2422: "Member is not found",
	2423: "Member name already exists",
	2424: "Failed to update member",
	// Delete member
	2431: "Member is not found",
	2432: "Failed to delete member",
}

// Category
var CategorySuccessMessage = map[int]string {
	// Create category
	3201: "Category created successfully",
	// Get all categories
	3202: "Success to query all categories",
	// Get category by page
	3203: "Success to query category by page",
	// Get category by id
	3204: "Success to query category by id",
	// Update category
	3205: "Category updated successfully",
	// Delete category
	3206: "Category deleted successfully",
}

var CategoryErrorMessage = map[int]string {
	// Create category
	3401: "Invalid category name",
	3402: "Category name already exists",
	3403: "Failed to create category",
	3404: "Failed to create user category",
	// Get all categories
	3411: "Failed to query all categories",
	// Get category by page
	3412: "Failed to query category by page",
	// Get category by id
	3413: "Failed to query category by id",
	// Update category
	3421: "Invalid category name",
	3422: "Category is not found",
	3423: "Category name already exists",
	3424: "Failed to update category",
	// Delete category
	3431: "Category is not found",
	3432: "Failed to delete category",
}

// Subcategory
var SubcategorySuccessMessage = map[int]string {
	// Create subcategory
	4201: "Subcategory created successfully",
	// Get all subcategories
	4202: "Success to query all subcategories",
	// Get subcategory by page
	4203: "Success to query subcategory by page",
	// Get subcategory by id
	4204: "Success to query subcategory by id",
	// Update subcategory
	4205: "Subcategory updated successfully",
	// Delete subcategory
	4206: "Subcategory deleted successfully",
}

var SubcategoryErrorMessage = map[int]string {
	// Create subcategory
	4401: "Invalid subcategory name",
	4402: "Subcategory name already exists",
	4403: "Failed to create subcategory",
	4404: "Failed to create user subcategory",
	// Get all subcategories
	4411: "Failed to query all subcategories",
	// Get subcategory by page
	4412: "Failed to query subcategory by page",
	// Get subcategory by id
	4413: "Failed to query subcategory by id",
	// Update subcategory
	4421: "Invalid subcategory name",
	4422: "Subcategory is not found",
	4423: "Subcategory name already exists",
	4424: "Failed to update subcategory",
	// Delete subcategory
	4431: "Subcategory is not found",
	4432: "Failed to delete subcategory",
}

// Tag
var TagSuccessMessage = map[int]string {
	// Create tag
	5201: "Tag created successfully",
	// Get all tags
	5202: "Success to query all tags",
	// Get tag by page
	5203: "Success to query tag by page",
	// Get tag by id
	5204: "Success to query tag by id",
	// Update tag
	5205: "Tag updated successfully",
	// Delete tag
	5206: "Tag deleted successfully",
}

var TagErrorMessage = map[int]string {
	// Create tag
	5401: "Invalid tag name",
	5402: "Tag name already exists",
	5403: "Failed to create tag",
	5404: "Failed to create user tag",
	// Get all tags
	5411: "Failed to query all tags",
	// Get tag by page
	5412: "Failed to query tag by page",
	// Get tag by id
	5413: "Failed to query tag by id",
	// Update tag
	5421: "Invalid tag name",
	5422: "Tag is not found",
	5423: "Tag name already exists",
	5424: "Failed to update tag",
	// Delete tag
	5431: "Tag is not found",
	5432: "Failed to delete tag",
}

// Color
var ColorSuccessMessage = map[int]string {
	// Get all color categories
	6201: "Success to query all color categories",
	// Get all colors
	6202: "Success to query all colors",
}

var ColorErrorMessage = map[int]string {
	// Get all color categories
	6401: "Failed to query all color categories",
	// Get all colors
	6402: "Failed to query all colors",
}

// Document
var DocumentSuccessMessage = map[int]string {
	// Create document
	7201: "Document created successfully",
	// Get all documents
	7202: "Success to query all documents",
	// Get document by page
	7203: "Success to query document by page",
	// Get document by id
	7204: "Success to query document by id",
	// Update document
	7205: "Document updated successfully",
	// Delete document
	7206: "Document deleted successfully",
}

var DocumentErrorMessage = map[int]string {
	// Create document
	7401: "Invalid document name",
	7402: "Document name already exists",
	7403: "Failed to create document",
	7404: "Failed to create user document",
	// Get all documents
	7411: "Failed to query all documents",
	// Get document by page
	7412: "Failed to query document by page",
	// Get document by id
	7413: "Failed to query document by id",
	// Update document
	7421: "Invalid document name",
	7422: "Document is not found",
	7423: "Document name already exists",
	7424: "Failed to update document",
	// Delete document
	7431: "Document is not found",
	7432: "Failed to delete document",
}

// Document comment
var DocumentCommentSuccessMessage = map[int]string {
	// Create document comment
	8201: "Document comment created successfully",
	// Get all document comments
	8202: "Success to query all document comments",
	// Get document comment by page
	8203: "Success to query document comment by page",
	// Get document comment by id
	8204: "Success to query document comment by id",
	// Update document comment
	8205: "Document comment updated successfully",
	// Delete document comment
	8206: "Document comment deleted successfully",
}

var DocumentCommentErrorMessage = map[int]string {
	// Create document comment
	8401: "Invalid document comment content",
	8402: "Failed to create document comment",
	8403: "Failed to create user document comment",
	// Get all document comments
	8411: "Failed to query all document comments",
	// Get document comment by page
	8412: "Failed to query document comment by page",
	// Get document comment by id
	8413: "Failed to query document comment by id",
	// Update document comment
	8421: "Invalid document comment content",
	8422: "Document comment is not found",
	8423: "Failed to update document comment",
	// Delete document comment
	8431: "Document comment is not found",
	8432: "Failed to delete document comment",
}