import AddTopicItem from '@/components/AddTopic';
import {
  addChapter,
  addExercise,
  chapterDetail,
  editChapter,
} from '@/service/teacher';
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import { useRequest } from 'ahooks';
import {
  Anchor,
  Button,
  Form,
  Input,
  message,
  Radio,
  Select,
  Space,
  Switch,
  Upload,
} from 'antd';
import { useEffect, useState } from 'react';
import { history, useLocation } from 'umi';
const { Link } = Anchor;
const { Item } = Form;
const AddCourse = () => {
  const location: any = useLocation();
  const id = location.query.course_id;

  const [form1] = Form.useForm();
  const [form] = Form.useForm();
  const [pdf, setPdf] = useState('');
  const [video, setVideo] = useState('');
  const [checked, setChecked] = useState(false);
  const { run: addTest } = useRequest(addExercise, {
    manual: true,
    onSuccess: (res) => {
      console.log(res);
    },
  });
  const pdfSuccess = (info: any) => {
    console.log(info);
    if (info.file.status === 'done') {
      setPdf(info.fileList[0]?.response.data.path);
    }
  };
  useEffect(() => {
    if (location.query.id) {
      chapterDetail(location.query.id).then((res) => {
        form1.setFieldsValue({
          name: res?.name || '',
          desc: res?.introduction || '',
        });
        setPdf(res?.pdf_url || '');
      });
    }
  }, []);
  const videoSuccess = (info: any) => {
    console.log(info);
    if (info.file.status === 'done') {
      setVideo(info.fileList[0]?.response.data.path);
    }
  };
  return (
    <>
      <Button
        type="primary"
        onClick={() => history.push('/teacher/course/chapter')}
      >
        返回
      </Button>

      <h1 style={{ textAlign: 'center' }}>基本信息</h1>
      <Form form={form1} labelCol={{ span: 8 }} wrapperCol={{ span: 8 }}>
        <Item name="name" label="章节名称">
          <Input />
        </Item>
        <Item name="desc" label="章节介绍">
          <Input.TextArea rows={4} />
        </Item>
        <Item label="上传附件">
          {pdf ? (
            <span style={{ display: 'flex', alignItems: 'center' }}>
              {pdf}
              <Button
                onClick={() => {
                  setPdf('');
                }}
                danger
                size="small"
              >
                删除
              </Button>
            </span>
          ) : (
            <Upload
              headers={{
                authorization: localStorage.getItem('token')!,
              }}
              onChange={pdfSuccess}
              action="/api/static"
              name="file"
            >
              <Button>上传pdf</Button>
            </Upload>
          )}
          {video ? (
            <span style={{ display: 'flex', alignItems: 'center' }}>
              {video}
              <Button
                onClick={() => {
                  setVideo('');
                }}
                danger
                size="small"
              >
                删除
              </Button>
            </span>
          ) : (
            <Upload
              headers={{
                authorization: localStorage.getItem('token')!,
              }}
              onChange={videoSuccess}
              action="/api/static"
              name="file"
            >
              <Button style={{ marginTop: 20 }}>上传视频</Button>
            </Upload>
          )}
        </Item>
      </Form>
      <Button
        style={{ marginLeft: '400px' }}
        type="primary"
        onClick={() => {
          form1.validateFields().then((res) => {
            if (location.query.id) {
              editChapter(location.query.id, {
                id: +location.query.course_id,
                name: res.name,
                introduction: res.desc,
                pdf: pdf,
                video: video,
              }).then((res) => {
                message.success('成功');
              });
            } else {
              addChapter({
                id: +id,
                name: res.name,
                introduction: res.desc,
                pdf: pdf,
                video: video,
              }).then((res) => {
                message.success('添加成功');
                history.goBack();
              });
            }
          });
        }}
      >
        确定
      </Button>
      {/* <h1 style={{ textAlign: 'center' }}>作业</h1>
      <Form labelCol={{ span: 8 }} wrapperCol={{ span: 8 }}>
        <Item label="作业内容">
          <Input.TextArea rows={4} />
        </Item>
      </Form> */}
      {/* <h1
        style={{
          textAlign: 'center',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        随堂测试 <Switch checked={checked} onChange={(e) => setChecked(e)} />
      </h1>
      <Anchor offsetTop={100}>
        <Link href="#1" title="第一题" />
        <Link href="#2" title="第二题" />
        <Link href="#3" title="第三题" />
      </Anchor>
      {checked && (
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
                        title: item.name.name,
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
      )} */}
    </>
  );
};
export default AddCourse;
