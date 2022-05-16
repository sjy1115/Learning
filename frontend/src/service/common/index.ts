import { UserInfo } from '@/pages/register';
import request from '@/utils/request';
export const register = (data: UserInfo) =>
  request('/api/user/register', 'POST', data);

export const getCode = () =>
  request<{ capt_id: string; image: string }>('/api/user/verifycode');

export const login = (data: {
  phone: string;
  password: string;
  capt_id: string;
}) =>
  request<{ token: string; user_id: number; role: number }>(
    '/api/user/login',
    'POST',
    data,
  );
export const logout = () => request('/api/user/logout');

export const getUserInfo = () => request('/api/user/info');
export const changeAvatar = (body: any) =>
  request('/api/user/upload-avatar', 'POST', body);
