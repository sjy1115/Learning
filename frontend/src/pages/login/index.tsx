import styles from './index.less';
import cat from '@/assets/image/cat.png';
import { Button, Form, Input, Radio } from 'antd';
import { history, useModel } from 'umi';
import { getCode, login } from '@/service/common';
import { useRequest } from 'ahooks';
import { useState } from 'react';
const { Item } = Form;
interface UserInfo {
  phone: string;
  password: string;
  vcode: string;
  role: number;
}
const Login = () => {
  const [role, setRole] = useState(1);
  const { initialState, loading, error, refresh, setInitialState } =
    useModel('@@initialState');
  const { data, run: getVerifycode } = useRequest(getCode);
  const { run } = useRequest(login, {
    manual: true,
    onError: () => {
      form.resetFields(['vcode']);
      getVerifycode();
    },
    onSuccess: (res) => {
      localStorage.setItem('token', res.token);
      console.log(res, '登录成功');

      if (res.role === 0) {
        history.push('/student/home');
        refresh();
      } else {
        history.push('/teacher/home');
        refresh();
      }
    },
  });
  const [form] = Form.useForm<UserInfo>();
  const submit = () => {
    // history.push('/student/home');
    form.validateFields().then((res) => {
      const obj = {
        phone: res.phone,
        password: res.password,
        capt_id: data!.capt_id,
        vcode: res.vcode,
        role: res.role,
      };
      run(obj);
    });
  };

  const register = () => {
    history.push('/register');
  };
  return (
    <div className={styles.wrapper}>
      <div className={styles.login_contain}>
        <div className={styles.row}>
          <img src={cat} className={styles.cat} width={60} height={60} alt="" />
          <h2>用户登录</h2>
        </div>

        <Form
          style={{ width: 350 }}
          form={form}
          labelCol={{ span: 5 }}
          initialValues={{
            role: 1,
          }}
          wrapperCol={{ span: 18 }}
        >
          <Item
            name="phone"
            label="手机号"
            rules={[
              {
                required: true,
                message: '请输入正确格式的手机号',
                pattern: /^1[3456789]\d{9}$/,
              },
            ]}
            required
          >
            <Input autoComplete="new-password" />
          </Item>
          <Item
            name="password"
            label="密码"
            rules={[{ required: true, message: '请输入密码,至少6位', min: 6 }]}
            required
          >
            <Input type="password" autoComplete="new-password" />
          </Item>
          <Item label="验证码" name="vcode" style={{ display: 'flex' }}>
            <div>
              <Input
                className={styles.verifycode}
                style={{ width: 80, display: 'inline', marginRight: 5 }}
              />
              <img
                width={90}
                onClick={() => getVerifycode()}
                src={data?.image}
                height={40}
                alt="1"
              />
            </div>
          </Item>
          <Form.Item wrapperCol={{ span: 18, offset: 6 }}>
            <Button onClick={submit} type="primary">
              登录
            </Button>
            <Button onClick={register} type="ghost" style={{ marginLeft: 8 }}>
              注册
            </Button>
          </Form.Item>
        </Form>
      </div>
    </div>
  );
};
export default Login;
