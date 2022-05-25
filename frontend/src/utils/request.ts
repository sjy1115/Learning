import { message } from 'antd';
import axios, { Method } from 'axios';
import { history } from 'umi';

const service = axios.create({
  withCredentials: true,
  timeout: 3000,
});
export default async function request<T>(
  url: string,
  method: Method = 'GET',
  data?: any,
): Promise<T> {
  try {
    const res = await service({
      url: url,
      headers: {
        authorization: localStorage.getItem('token') || '',
      },
      data: method === 'POST' || method === 'PUT' ? data : undefined,
      params: method === 'GET' ? data : undefined,
      method: method,
    });
    console.log(res);
    if (res.data.code === 100) {
      message.error(res.data.message);
      return Promise.reject();
    }

    if (res.data.code === 401) {
      history.push('/login');
    }
    return res.data.data;
  } catch (error) {
    message.error(error?.message || '未知错误');
    // history.push('/login');
    return Promise.reject('错误');
  }
}
// export default request;
