import { IMG_URL } from '@/utils/constant';
import { history } from 'umi';
import styles from './index.less';
interface IProps {
  data: any;
}
const CourseCard: React.FC<IProps> = (props) => {
  const { data } = props;
  const goDetail = (id: number) => {
    history.push(`/student/course/${id}`);
  };
  return (
    <div className={styles.card} onClick={() => goDetail(data.id)}>
      <div
        style={{
          width: '100%',
          height: 280,
          backgroundSize: '100% 100%',
          backgroundImage: `url(${IMG_URL}${data?.avatar})`,
        }}
      />
      <div className={styles.info}>
        <img
          className={styles.avatar}
          src="data:image/jpeg;base64,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"
          alt=""
        />
        <h3 className={styles.teacher_name}>{data?.teacher}老师</h3>
      </div>
      <div className={styles.course_name}>{data?.name || ''}</div>
      <div className={styles.desc}>
        <div className={styles.left}>课程介绍:{data?.introduction}</div>
        <div className={styles.right}>进入课程</div>
      </div>
    </div>
  );
};
export default CourseCard;
