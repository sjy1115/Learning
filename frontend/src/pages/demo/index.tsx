import { UploadOutlined } from '@ant-design/icons';
import { Button, Input, Upload, message } from 'antd';
import { TextAreaRef } from 'antd/lib/input/TextArea';
import {
  HTMLAttributes,
  HtmlHTMLAttributes,
  useEffect,
  useMemo,
  useRef,
  useState,
} from 'react';

import styles from './index.less';
const Demo = (props) => {
  const { location } = props;
  const ref = useRef<HTMLDivElement>(null);
  const [value, setValue] = useState('');
  const input = useRef<any>();
  const socket = useRef<WebSocket | null>(null);
  const [data, setData] = useState<string[]>([]);
  useEffect(() => {
    socket.current = new WebSocket(
      `ws://121.199.167.227:81/api/course/chat?token=${location.query.token}`,
    );

    socket.current.onmessage = (evt) => {
      console.log(data);
      setData((i) => i.concat(evt.data));
      // setData(data.concat(evt.data));
      // setLatestMessage(evt.data);
    };
    return () => {
      socket.current?.close();
    };
  }, []);
  useEffect(() => {
    ref.current!.scrollTop = ref.current!.scrollHeight;
  }, [data]);

  const onEnter = () => {
    input.current?.focus();
    socket.current?.send(value);
    setValue('');
  };
  const handleKeyDown = (event: {
    keyCode: number;
    metaKey: any;
    preventDefault: () => void;
  }) => {
    // console.log(event);
    if (event.keyCode == 13) {
      if (!event.metaKey) {
        event.preventDefault();
        onEnter();
      } else {
        setValue(value + '\n');
      }
    }
  };
  return (
    <div>
      <h1>demo页面</h1>
      <Upload name="file_name" action="/upload">
        <Button icon={<UploadOutlined />}>上传视频</Button>
      </Upload>
      <div>
        <div className={styles.wrapper}>
          <div ref={ref} className={styles.msg}>
            <pre>
              {data?.map((item, index) => (
                <div key={index}>{item}</div>
              ))}
            </pre>
          </div>
          <Input.TextArea
            ref={input!}
            value={value}
            onKeyDown={handleKeyDown}
            onChange={(e) => {
              setValue(e.target.value);
            }}
            draggable={false}
            style={{ height: 200 - 32, resize: 'none', borderBottom: 'none' }}
          />
          <div className={styles.btn}>
            <Button onClick={onEnter} type="primary">
              发送
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
};
export default Demo;
