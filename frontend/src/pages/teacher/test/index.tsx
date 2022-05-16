import { Cascader, Col, Pagination, Row, Table } from 'antd';
import { Link } from 'umi';
import styles from './index.less';
const Test = () => {
  const options = [
    {
      value: 'zhejiang',
      label: 'java程序设计',
      //   disabled: true,
      children: [
        {
          value: 'hangzhou',
          label: '第一章节',
        },
      ],
    },
    {
      value: '22',
      label: 'java程序设计222',
      disabled: true,
    },
  ];
  return (
    <>
      <Cascader
        style={{ width: 400 }}
        placeholder="选择章节"
        options={options}
      />
      <Row gutter={[20, 20]} style={{ marginTop: 20 }}>
        {[1, 2, 3, 4, 6, 7, 8, 9].map((item) => (
          <Col key={item} span={4}>
            <div className={styles.card}>
              <img
                className={styles.avatar}
                src="data:image/jpeg;base64,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"
                alt=""
              />
              <p className={styles.name}>姓名:xxxxx</p>
              <p className={styles.time}>提交时间:2022.10.11 18:00</p>
              <p>
                成绩:<span className={styles.score}>99</span>
              </p>
              <Link to="/777">查看详情</Link>
            </div>
          </Col>
        ))}
      </Row>
      <div className={styles.footer}>
        <Pagination defaultCurrent={6} total={500} />
      </div>
    </>
  );
};
export default Test;
