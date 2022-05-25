import AddCourse from '@/components/AddCourse';
import { chapterList } from '@/service/teacher';
import { useRequest } from 'ahooks';
import { Button, Form, Input, Select, Table, Typography } from 'antd';
import moment from 'moment';
import { useState } from 'react';
import { Link, history, useLocation } from 'umi';

const { Paragraph } = Typography;
const Home = () => {
  const location: any = useLocation();
  const id = location['query'].course_id;
  const { data: list, run } = useRequest(chapterList, {
    defaultParams: [{ id: id, page: 1, page_size: 10 }],
  });
  const onChange = (p: number, z: number) => {
    run({ id: id, page: p, page_size: z });
  };
  const column = [
    {
      title: '章节名称',
      dataIndex: 'name',
    },
    {
      title: '章节介绍',
      dataIndex: 'introduction',
    },
    {
      title: '已学习人数/总人数',
      dataIndex: 'num',
    },

    {
      title: '创建时间',
      dataIndex: 'create_tm',
      render: (time: any) => (
        <span>{moment(time).format('YYYY-MM-DD HH:mm:ss')}</span>
      ),
    },
    {
      title: '附件',
      dataIndex: 'extra',
    },
    {
      title: '操作',
      //   dataIndex: '',
      render: (_: any, record: any) => (
        <>
          <Link
            style={{ marginRight: 10 }}
            to={`/teacher/course/chapter/add?id=${record.id}&course_id=${id}`}
          >
            编辑章节
          </Link>
          <Link to={`/teacher/course/chapter/exercise?id=${record.id}`}>
            添加作业
          </Link>
        </>
      ),
    },
  ];
  const data = [
    {
      id: 1,
      desc: 'xxxxx',
      name: '类',
      year: '2018上学期',
      num: '69/100',
      extra: 'pdf',
      invite_num: '9xjdhihida',
      create_tm: '2021.7.10 12:00:00',
    },
  ];
  return (
    <>
      <Form layout="inline" style={{ marginBottom: 20 }}>
        <Form.Item>
          <Button type="default" onClick={() => history.push('/teacher/home')}>
            返回
          </Button>
        </Form.Item>
        <Form.Item label="章节名称">
          <Input
            allowClear
            style={{ width: 300 }}
            placeholder="输入章节名称按回车搜索"
          />
        </Form.Item>
        {/* <Form.Item>
          <Button type="dashed">搜索</Button>
        </Form.Item> */}
        <Form.Item>
          <Button
            type="primary"
            onClick={() =>
              history.push(`/teacher/course/chapter/add?course_id=${id}`)
            }
          >
            添加章节
          </Button>
        </Form.Item>
      </Form>

      <Table
        rowKey="id"
        pagination={{ total: list?.total || 0, onChange }}
        dataSource={list?.items || []}
        columns={column}
      />
    </>
  );
};
export default Home;
