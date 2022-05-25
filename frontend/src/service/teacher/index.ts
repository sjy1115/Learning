import request from '@/utils/request';

export const getCourse = (body: Record<string, any>) =>
  request<any>('/api/course', 'GET', body);

export const addCourse = (body: Record<string, any>) =>
  request('/api/course', 'POST', body);
export const editCourse = (id: any, body: any) =>
  request(`/api/course/${id}`, 'PUT', body);
export const courseDetail = (id: number) => request(`/api/course/${id}`, 'GET');
export const chapterList = (body: any) =>
  request<any>('/api/chapter', 'GET', body);
export const addChapter = (body: any) =>
  request<any>('/api/chapter', 'POST', body);

export const chapterDetail = (id: number) =>
  request<any>(`/api/chapter/${id}`, 'GET');

export const editChapter = (id: any, body: any) =>
  request(`/api/chapter/${id}`, 'PUT', body);

export const addExercise = (body: any) =>
  request('/api/exercise', 'POST', body);

export const getExerciseDetail = (body: any) =>
  request<any>(`/api/exercise`, 'GET', body);
