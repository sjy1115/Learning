import { message } from 'antd';
import { history } from 'umi';
import { getUserInfo } from './service/common';

export function onRouteChange({ location }: any) {
  console.log(location.pathname);
  if (
    !localStorage.getItem('token') &&
    location.pathname !== '/login' &&
    location.pathname !== '/register'
  ) {
    message.error('请先登录');
    history.push('/login');
  }
}
export async function getInitialState(): Promise<any> {
  if (location.pathname !== '/login' && location.pathname !== '/register') {
    const data = await getUserInfo();
    // if (data.role === 0) {
    //   history.push('/student/home');
    // } else {
    //   history.push('/teacher/home');
    // }
    console.log(data);
    return data;
  }
  return {
    getUserInfo,
  };
}
