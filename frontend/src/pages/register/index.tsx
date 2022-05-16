import styles from './index.less';
import cat from '@/assets/image/cat.png';
import { Button, Form, Input, message, Radio } from 'antd';
import { useState } from 'react';
import { LeftOutlined } from '@ant-design/icons';
import { history, useRequest } from 'umi';
import { RuleObject } from 'antd/lib/form';
import { StoreValue } from 'antd/lib/form/interface';
import { register } from '@/service/common';
const { Item } = Form;
export interface UserInfo {
  phone: string;
  gender: string;
  role: number;
  username: string;
  password: string;
  college: string;
  number: number;
  re_password: string;
}
const Register = () => {
  const { run: runRegister } = useRequest(register, {
    manual: true,
    onSuccess: (res) => {
      message.success('注册成功');
      history.push('/login');
    },
  });
  const [form] = Form.useForm<UserInfo>();
  const [role, setRole] = useState(0);
  const submit = () => {
    form.validateFields().then((res) => {
      runRegister(res);
    });
  };
  const validatorIsSame = (rule: RuleObject, value: StoreValue) => {
    const values = form.getFieldsValue();
    if (values.password === values.re_password) {
      return Promise.resolve();
    }
    return Promise.reject('两次密码不一致');
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.login_contain}>
        <div
          className={styles.back}
          onClick={() => {
            history.push('/login');
          }}
        >
          <LeftOutlined />
          <span>返回登录</span>
        </div>
        <div className={styles.row}>
          <img src={cat} className={styles.cat} width={60} height={60} alt="" />
          <h2>用户注册</h2>
        </div>

        <Form
          form={form}
          labelCol={{ span: 6 }}
          initialValues={{
            role: 0,
            gender: '1',
          }}
          wrapperCol={{ span: 18 }}
        >
          <Item
            name="username"
            label="用户名"
            rules={[
              {
                required: true,
                message: '请输入用户名',
              },
            ]}
            required
          >
            <Input style={{ width: 300 }} />
          </Item>
          <Item
            name="phone"
            label="手机号"
            rules={[
              {
                required: true,
                message: '请输入手机号',
                pattern: /^1[3456789]\d{9}$/,
              },
            ]}
            required
          >
            <Input style={{ width: 300 }} />
          </Item>
          <Item
            label="学校"
            name="college"
            required
            rules={[
              {
                required: true,
                message: '请输入学校',
              },
            ]}
          >
            <Input style={{ width: 300 }} />
          </Item>
          <Item label="身份" name="role" required>
            <Radio.Group onChange={(e) => setRole(e.target.value)}>
              <Radio value={0}>学生</Radio>
              <Radio value={1}>老师</Radio>
            </Radio.Group>
          </Item>

          <Item label="性别" name="gender" required>
            <Radio.Group onChange={(e) => setRole(e.target.value)}>
              <Radio value="1">男</Radio>
              <Radio value="2">女</Radio>
            </Radio.Group>
          </Item>
          <Item
            label="学号/工号"
            name="number"
            required
            rules={[
              {
                required: true,
                message: '请输入至少6位数工号',
                min: 6,
              },
            ]}
          >
            <Input style={{ width: 300 }} />
          </Item>
          <Item
            name="password"
            label="密码"
            rules={[{ required: true, message: '请输入密码' }]}
            required
          >
            <Input
              autoComplete="new-password"
              style={{ width: 300 }}
              type="password"
            />
          </Item>
          <Item
            name="re_password"
            label="确认密码"
            rules={[
              {
                required: true,
                validator: validatorIsSame,
              },
            ]}
            required
          >
            <Input style={{ width: 300 }} type="password" />
          </Item>
          <Form.Item wrapperCol={{ span: 18, offset: 6 }}>
            <Button onClick={submit} type="primary">
              确认注册
            </Button>
          </Form.Item>
        </Form>
      </div>
    </div>
  );
};
export default Register;
