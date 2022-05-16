import { IMG_URL } from '@/utils/constant';
import { CaretRightOutlined } from '@ant-design/icons';
import { Button, Collapse, Modal } from 'antd';
import { useState } from 'react';
import { history, Link } from 'umi';
import styles from './index.less';
const { Panel } = Collapse;
const Course = (props) => {
  const [visible, setVisible] = useState(false);
  console.log(props);
  return (
    <>
      <div className={styles.wrapper}>
        <div>
          <img
            className={styles.pic}
            width={160}
            height={160}
            src="https://th.wallhaven.cc/small/qd/qdl8e7.jpg"
            alt=""
          />
        </div>
        <div>
          <p className={styles.title}>java课程设计</p>
          <p>教师：xxxx老师</p>
          <p>课程创建时间：xxxxx</p>
          <p>加入课程时间：xxxxx</p>
        </div>
        <Button
          className={styles.btn}
          onClick={() =>
            history.push(`/student/course/${props.match.params.id}/chat`)
          }
          type="ghost"
        >
          加入课程群聊
        </Button>
      </div>
      <div className={styles.main}>
        {/* {[1, 2, 4].map((item, index) => ( */}
        <Collapse
          bordered={false}
          expandIcon={({ isActive }) => (
            <CaretRightOutlined rotate={isActive ? 90 : 0} />
          )}
          className="site-collapse-custom-collapse"
        >
          {[1, 2, 3, 4].map((item, index) => (
            <Panel
              header={
                <p className={styles.dot}>
                  {index + 1} 、xxxx章节{' '}
                  <span style={{ color: '#666' }}>
                    （已学习/未开始/学习xxx时间）
                  </span>
                </p>
              }
              key={index}
              className="site-collapse-custom-panel"
            >
              <p className={styles.desc}>
                章节介绍:xxxxxxxxxxxxxxxx
                <span style={{ marginLeft: 10 }}>创建时间:xxxxx</span>
              </p>
              <p>章节内容:xxxxxxxxxx</p>
              <p className={styles.other}>
                附件📎:<p className={styles.pdf}>下载链接(pdf)</p>
                <p onClick={() => setVisible(true)} className={styles.watch}>
                  观看视频
                </p>
              </p>
              <p className={styles.watch}>
                <Link to="/student/test/1">练习题</Link>
              </p>
            </Panel>
          ))}
        </Collapse>
        <Modal
          width={1000}
          visible={visible}
          onCancel={() => setVisible(false)}
          footer={null}
          title="课程观看"
        >
          <video
            width={900}
            controls
            src={`${IMG_URL}vEjntxD9SD6F0c8mOdopea7NiVid4N3E-屏幕录制2022-05-07 下午11.56.41.mov`}
          ></video>
        </Modal>
      </div>
    </>
  );
};
export default Course;
