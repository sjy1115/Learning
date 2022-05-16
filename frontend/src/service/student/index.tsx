import request from '@/utils/request';

export const addCourse = (body: Record<string, any>) =>
  request<any>('/api/course/join', 'POST', body);
