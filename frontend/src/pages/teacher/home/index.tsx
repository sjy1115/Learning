import AddCourse from '@/components/AddCourse';
import { courseDetail, getCourse } from '@/service/teacher';
import { useRequest } from 'ahooks';
import { Button, Form, Input, Select, Table, Typography, Upload } from 'antd';
import moment from 'moment';
import { useState } from 'react';
import { Link } from 'umi';
import styles from './index.less';
const { Paragraph } = Typography;
const Home = () => {
  const { data: list, run } = useRequest(getCourse, {
    defaultParams: [{ page: 1, page_size: 10 }],
  });
  const [type, setType] = useState('add');
  const { data, run: runDetail } = useRequest(courseDetail, {
    manual: true,
    onSuccess: () => {
      setVisible(true);
    },
  });
  console.log(list);
  const onChange = (page: number, pageSize: number) => {
    run({ page: page, page_size: pageSize });
  };
  const [visible, setVisible] = useState(false);
  const column = [
    {
      title: '课程名称',
      dataIndex: 'name',
    },
    {
      title: '学期',
      dataIndex: 'year',
    },
    {
      title: '参加人数',
      dataIndex: 'student_num',
    },
    {
      title: '邀请码',
      width: 50,
      dataIndex: 'invite_code',
      render: (text: string) => (
        <Paragraph style={{ width: 120 }} copyable>
          {text}
        </Paragraph>
      ),
    },
    {
      title: '创建时间',
      dataIndex: 'create_tm',
      render: (time: any) => (
        <span>{moment(time * 1000).format('YYYY-MM-DD HH:mm:ss')}</span>
      ),
    },
    {
      title: '操作',
      //   dataIndex: '',
      render: (_: any, record: any) => (
        <>
          <span
            onClick={() => {
              setType('edit');
              runDetail(record.id);
            }}
            className={styles.text}
          >
            编辑课程
          </span>
          <Link to={`/teacher/course/chapter?course_id=${record.id}`}>
            查看章节
          </Link>
        </>
      ),
    },
  ];

  return (
    <>
      <Form layout="inline" style={{ marginBottom: 20 }}>
        <Form.Item label="学期">
          <Select style={{ width: 140 }} allowClear />
        </Form.Item>
        <Form.Item label="课程名称">
          <Input allowClear />
        </Form.Item>
        <Form.Item>
          <Button type="ghost">搜索</Button>
        </Form.Item>
        <Form.Item>
          <Button
            type="primary"
            onClick={() => {
              setType('add');
              setVisible(true);
            }}
          >
            添加课程
          </Button>
        </Form.Item>
      </Form>

      <Table
        dataSource={list?.items || []}
        pagination={{ total: list?.total || 0, onChange }}
        columns={column}
      />
      <AddCourse
        data={data}
        type={type}
        visible={visible}
        onOk={() => {
          run({ page: 1, page_size: 10 });
          setVisible(false);
        }}
        onCancel={() => setVisible(false)}
      />
    </>
  );
};
export default Home;
