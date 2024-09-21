const routes = {
  HOME: '/',
  LOGIN: '/login',
  ACTIVE_ACCOUNT: '/active-account',
  DASHBOARD: '/dashboard',
  MEMBERS: '/members',
  CATEGORIES: '/categories',
  CATEGORY: '/category/:id',
  TAGS: '/tags',
  TAG: '/tag/:id',
  DOCUMENTS: '/documents',
  DOCUMENT: '/document/:id',
  ADD_DOCUMENT: '/document/add',
  EDIT_DOCUMENT: '/document/:id/edit',
  DOCUMENTS_COMMENTS: '/document/:id/comments',
  ADD_DOCUMENT_COMMENT: '/document/:id/comment',
  EDIT_DOCUMENT_COMMENT: '/document/:id/comment/:commentId',
  SETTING: '/setting',
};

// Others
routes['RECORDER_UPLOAD_IMAGE'] = '/record/upload-image';

export default {
  routes,
};