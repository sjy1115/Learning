import AddTopicItem from '@/components/AddTopic';
import { addExercise, getExerciseDetail } from '@/service/teacher';
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Form } from 'antd';
import { useEffect } from 'react';
import { useLocation, useRequest, history } from 'umi';
import styles from './index.less';
const Exercise = () => {
  const [form] = Form.useForm();
  const location: any = useLocation();
  const id = location.query.id;
  useEffect(() => {
    if (id) {
      getExerciseDetail({ chapter_id: +id }).then((res) => {
        console.log(res);
        const a = res.items.map((item) => ({
          name: item,
        }));
        console.log(a);
        form.setFieldsValue({
          test: a,
        });
      });
    }
    form.setFieldsValue({ test: [undefined] });
  }, []);
  const { run: addTest } = useRequest(addExercise, {
    manual: true,
    onSuccess: (res) => {
      console.log(res);
    },
  });
  return (
    <div className={styles.wrapper}>
      <Button
        type="primary"
        onClick={() => {
          history.goBack();
        }}
      >
        返回
      </Button>
      <Form form={form} name="dynamic_form_nest_item" autoComplete="off">
        <Form.List name="test">
          {(fields, { add, remove }) => (
            <>
              {fields.map(({ key, name, ...restField }, index) => (
                <div key={key}>
                  <h1 id={`${index + 1}`} style={{ textAlign: 'center' }}>
                    第{index + 1}题
                    <MinusCircleOutlined onClick={() => remove(name)} />
                  </h1>
                  <Form.Item {...restField} name={[name, 'name']}>
                    <AddTopicItem key={key} />
                  </Form.Item>
                </div>
              ))}
              <Button
                style={{
                  width: 200,
                  margin: 'auto',
                  marginLeft: '50%',
                  transform: 'translateX(-50%)',
                }}
                type="dashed"
                onClick={(e) => {
                  add();
                }}
                block
                icon={<PlusOutlined />}
              >
                添加下一题
              </Button>
            </>
          )}
        </Form.List>
        <Form.Item
          style={{ display: 'flex', justifyContent: 'center' }}
          labelCol={{ offset: 8, span: 3 }}
        >
          <Button
            style={{
              marginTop: 20,
              width: 200,
              margin: 'auto',
              marginLeft: '50%',
              transform: 'translateX(-50%)',
            }}
            type="primary"
            onClick={() => {
              form.validateFields().then((res) => {
                console.log(
                  res.test.map(
                    (item: {
                      options: any;
                      answer: any;
                      type: any;
                      name: any;
                    }) => ({
                      options: item.name.options,
                      answer: item.name.answer,
                      type: item.name.type,
                      title: item.name.name,
                    }),
                  ),
                );
                // return;
                addTest({
                  chapter_id: +location['query'].id,
                  questions: res.test.map(
                    (item: {
                      options: any;
                      answer: any;
                      type: any;
                      name: any;
                    }) => ({
                      options: item.name.options,
                      answer: item.name.answer,
                      type: item.name.type,
                      title: item.name.title,
                    }),
                  ),
                });
              });
            }}
            htmlType="submit"
          >
            提交
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};
export default Exercise;
