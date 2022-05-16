import { Button, Form } from 'antd';
import InputItem from './components/InputItem';
import JudgeItem from './components/judgeItem';
import SelectItem from './components/selectItem';
import styles from './index.less';
import { history } from 'umi';
const { Item } = Form;
const Test = () => {
  const [form] = Form.useForm();
  const data = [
    {
      type: 1,
      question: '这是一道测试题目',
      options: ['这是a选项', '这是b选项', '这是c选项'],
    },
    {
      type: 2,
      question: '这是一道测试题目',
      options: ['这是a选项', '这是b选项', '这是c选项'],
    },
    {
      type: 3,
      question: '这是一道测试题目',
      options: ['这是a选项', '这是b选项', '这是c选项'],
    },
  ];
  return (
    <div className={styles.wrapper}>
      <Button type="primary" onClick={() => history.goBack()}>
        返回
      </Button>
      <h2 className={styles.title}>练习题</h2>
      <Form form={form}>
        {data.map((item, index) => {
          if (item.type === 1) {
            return (
              <Item
                key={index}
                name={`question${index + 1}`}
                label={`第${index + 1}题`}
              >
                <JudgeItem question={item.question} />
              </Item>
            );
          } else if (item.type === 2) {
            return (
              <Item
                key={index}
                name={`question${index + 1}`}
                label={`第${index + 1}题`}
              >
                <SelectItem options={item.options} question={item.question} />
              </Item>
            );
          } else {
            return (
              <Item
                key={index}
                name={`question${index + 1}`}
                label={`第${index + 1}题`}
              >
                <InputItem question={item.question} />
              </Item>
            );
          }
        })}
      </Form>
      <div>
        <Button
          onClick={() => {
            form.validateFields().then((res) => {
              console.log(res);
            });
          }}
          type="primary"
        >
          提交
        </Button>
      </div>
    </div>
  );
};
export default Test;
