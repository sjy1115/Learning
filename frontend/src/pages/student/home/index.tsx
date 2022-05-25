import CourseCard from '@/components/CourseCard';
import { addCourse } from '@/service/student';
import { getCourse } from '@/service/teacher';
import { useRequest } from 'ahooks';
import {
  Button,
  Col,
  Empty,
  Input,
  message,
  Modal,
  Pagination,
  Row,
  Space,
} from 'antd';
import { useState } from 'react';
import styles from './index.less';
const StudentHome = () => {
  const [value, setValue] = useState('');
  const [visible, setVisible] = useState(false);
  const onChange = (page: number, pageSize: number) => {
    run({ page: page, page_size: pageSize });
  };
  const { data, run } = useRequest(getCourse);
  return (
    <div className={styles.wrapper}>
      <Button
        style={{ marginBottom: 20 }}
        type="primary"
        onClick={() => setVisible(true)}
      >
        加入课程
      </Button>
      <Modal
        visible={visible}
        onCancel={() => setVisible(false)}
        onOk={() => {
          addCourse({ invite_code: value }).then((res) => {
            setVisible(false);
            run({ page: 1, page_size: 12 });
            message.success('加入成功');
          });
        }}
        title="加入课程"
      >
        <Space>
          输入邀请码:
          <Input
            style={{ width: 380 }}
            value={value}
            onChange={(e) => {
              setValue(e.target.value);
            }}
          />
        </Space>
      </Modal>
      {/* <Empty description="暂无课程，请先添加课程" /> */}
      {(data?.items || []).length > 0 ? (
        <>
          <Row gutter={[20, 20]}>
            {(data?.items || []).map((item) => (
              <Col span={6} key={item.id}>
                <CourseCard data={item} />
              </Col>
            ))}
          </Row>
          <div className={styles.footer}>
            <Pagination
              onChange={onChange}
              defaultCurrent={1}
              total={data?.total || 0}
            />
          </div>
        </>
      ) : (
        <Empty />
      )}
    </div>
  );
};
export default StudentHome;
