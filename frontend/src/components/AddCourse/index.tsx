import { addCourse, editCourse } from '@/service/teacher';
import { IMG_URL } from '@/utils/constant';
import {
  LoadingOutlined,
  PlusOutlined,
  QuestionCircleFilled,
} from '@ant-design/icons';
import { useRequest } from 'ahooks';
import {
  Form,
  Modal,
  Input,
  Select,
  Tooltip,
  Upload,
  Button,
  message,
} from 'antd';
import { UploadChangeParam } from 'antd/lib/upload';
import { RcFile, UploadFile } from 'antd/lib/upload/interface';
import { useEffect, useState } from 'react';
const { Item } = Form;
interface IProps {
  visible: boolean;
  onOk: () => void;
  onCancel: () => void;
  data?: any;
  type: string;
}
function getBase64(img: Blob, callback: (i: string) => void) {
  const reader = new FileReader();
  reader.addEventListener('load', () => callback(reader.result as string));
  reader.readAsDataURL(img);
}
const AddCourse: React.FC<IProps> = (props) => {
  const { visible, onOk, onCancel, data, type } = props;
  const [imgUrl, setImgUrl] = useState('');
  const [loading, setLoading] = useState(false);
  const { run: edit } = useRequest(editCourse, {
    manual: true,
    onSuccess: () => {
      onOk?.();
    },
  });
  const [form] = Form.useForm();
  const { run } = useRequest(addCourse, {
    manual: true,
    onSuccess: () => {
      setImgUrl('');

      form.resetFields();
      message.success('添加成功');
      onOk?.();
    },
  });
  useEffect(() => {
    form.resetFields();
    if (type === 'edit') {
      setImgUrl(`${IMG_URL}${data?.avatar}`);
      form.setFieldsValue({
        name: `${data?.name}`,
        desc: data?.introduction,
        seme: data?.semester,
      });
      // setImgUrl(data)
    }
  }, [data, edit]);
  const handleChange = (info: UploadChangeParam<UploadFile<unknown>>) => {
    if (info.file.status === 'uploading') {
      // this.setState({ loading: true });
      setLoading(true);
      return;
    }

    if (info.file.status === 'done') {
      console.log(info.fileList[0].response.data.path);
      // Get this url from response in real world.
      setImgUrl(IMG_URL + info.fileList[0].response.data.path);
      setLoading(false);
    }
  };
  const uploadButton = (
    <div>
      {loading ? <LoadingOutlined /> : <PlusOutlined />}
      <div style={{ marginTop: 8 }}>上传</div>
    </div>
  );
  return (
    <Modal
      maskClosable={false}
      visible={visible}
      onOk={() => {
        form.validateFields().then((res) => {
          console.log(res);
          // return;
          if (type === 'edit') {
            edit(data.id, {
              name: res.name,
              introduction: res.desc,
              semester: res.seme,
              avatar: res?.img?.fileList?.[0].response.data.path || imgUrl,
            });
          } else {
            run({
              name: res.name,
              introduction: res.desc,
              semester: res.seme,
              avatar: res.img.fileList?.[0].response.data.path,
            });
          }
        });
      }}
      onCancel={() => {
        setImgUrl('');
        form.resetFields();
        onCancel?.();
      }}
      title={
        <span>
          添加章节
          <Tooltip overlay="添加成功后复制列表页邀请码，邀请学生加入课程">
            <QuestionCircleFilled style={{ color: '#333' }} />
          </Tooltip>
        </span>
      }
    >
      <Form
        initialValues={{ seme: '1' }}
        form={form}
        labelCol={{ span: 6 }}
        wrapperCol={{ span: 18 }}
      >
        <Item label="课程图片" name="img">
          <Upload
            headers={{
              authorization: localStorage.getItem('token')!,
            }}
            name="file"
            listType="picture-card"
            className="avatar-uploader"
            showUploadList={false}
            onChange={handleChange}
            action="/api/static"
          >
            {imgUrl ? (
              <img src={imgUrl} alt="avatar" style={{ width: '100%' }} />
            ) : (
              uploadButton
            )}
          </Upload>
        </Item>
        <Item label="课程名称" name="name">
          <Input style={{ width: 200 }} />
        </Item>
        <Item label="课程介绍" name="desc">
          <Input.TextArea style={{ width: 500 }} rows={4} />
        </Item>
        <Item label="学期" name="seme">
          <Select style={{ width: 200 }} />
        </Item>
      </Form>
    </Modal>
  );
};
export default AddCourse;
