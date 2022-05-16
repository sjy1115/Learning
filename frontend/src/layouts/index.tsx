import BasicLayout from '@/components/StudentBasicLayout';
import { ConfigProvider } from 'antd';
import { IRouteComponentProps } from 'umi';
import zhCN from 'antd/lib/locale/zh_CN';
import moment from 'moment';
import 'moment/locale/zh-cn';
moment.locale('en');
import './index.less';
import TeacherBasicLayout from '@/components/TeacherBasicLayout';
export default function Layout({
  children,
  location,
  route,
  history,
  match,
}: IRouteComponentProps) {
  if (
    location.pathname.includes('login') ||
    location.pathname.includes('register')
  ) {
    return <ConfigProvider locale={zhCN}>{children}</ConfigProvider>;
  }
  if (location.pathname.includes('student')) {
    return (
      <ConfigProvider locale={zhCN}>
        <BasicLayout>{children}</BasicLayout>
      </ConfigProvider>
    );
  } else {
    return (
      <ConfigProvider locale={zhCN}>
        <TeacherBasicLayout>{children}</TeacherBasicLayout>
      </ConfigProvider>
    );
  }
}
