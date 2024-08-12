const routes = {
  HOME: '/',
  LOGIN: '/login',
  DASHBOARD: '/dashboard',
  MEMBERS: '/members',
  CATEGORIES: '/categories',
  CATEGORY: '/category/:id',
  TAGS: '/tags',
  TAG: '/tag/:id',
  DOCUMENTS: '/documents',
  ADD_DOCUMENT: '/document/add',
  DOCUMENT: '/document/:id',
};

// Others
routes['RECORDER_UPLOAD_IMAGE'] = '/record/upload-image';

export default {
  routes,
};