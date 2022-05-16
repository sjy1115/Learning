import { letters } from '@/utils/constant';
import { Button, Col, Input, Radio, RadioChangeEvent, Row, Select } from 'antd';
import React, { useEffect, useState } from 'react';
import styles from './index.less';
interface ValueProps {
  title?: string;
  options?: string[];
  answer?: string;
  type?: string;
}

interface IProps {
  value?: ValueProps;
  onChange?: (value: ValueProps) => void;
}
const AddTopicItem: React.FC<IProps> = (props) => {
  const { value, onChange } = props;
  const [name, setName] = useState('');
  const [options, setOptions] = useState<string[]>(
    value?.options || ['', '', '', ''],
  );
  useEffect(() => {
    console.log('--', value?.options);
    setOptions(value?.options || []);
  }, [value]);
  console.log(value);
  const [answer, setAnswer] = useState(value?.answer || undefined);
  const [type, setType] = useState(value?.type || 'multiple_choice');
  const changeType = (e: RadioChangeEvent) => {
    setType(e.target.value);
    if (e.target.value === 1) {
      onChange?.({
        ...value,
        type: e.target.value,
        options: ['', '', '', ''],
        answer: '',
      });
    } else {
      onChange?.({
        ...value,
        type: e.target.value,
        options: [],
        answer: undefined,
      });
    }
  };
  useEffect(() => {
    onChange?.({
      ...value,
      title: value?.title || name,
      type: value?.type || type,
      options: value?.options || options,
    });
  }, []);
  const changeAnswer = (s: string) => {
    setAnswer(s);
    onChange?.({ ...value, answer: s });
  };
  return (
    <>
      <Row gutter={[0, 20]} className={styles.wrapper}>
        <Col span={4}>题目名称:</Col>
        <Col span={20}>
          <Input.TextArea
            value={value?.title || name}
            onChange={(e) => {
              setName(e.target.value);
              onChange?.({ ...value, title: e.target.value });
            }}
            placeholder="请输入题目名称"
          />
        </Col>
        <Col span={4}>题目类型:</Col>
        <Col span={20}>
          <Radio.Group onChange={changeType} value={value?.type || type}>
            <Radio value="multiple_choice">选择题</Radio>
            <Radio value={2}>判断题</Radio>
          </Radio.Group>
        </Col>
        {(type === 'multiple_choice' || value?.type === 'multiple_choice') && (
          <>
            <Col span={4}>
              <div className={styles.row}>选择项:</div>
            </Col>
            <Col span={20}>
              <Button
                onClick={() => {
                  onChange?.({ ...value, options: [...options, ''] });
                  setOptions([...options, '']);
                }}
                type="primary"
                size="small"
              >
                添加选项
              </Button>
              {options.map((item, index) => (
                <div key={index}>
                  <span style={{ marginRight: 5 }}>{letters[index]}:</span>
                  <Input
                    onChange={(e) => {
                      console.log(e);
                      const temp = [...options];
                      temp[index] = e.target.value;
                      setOptions(temp);
                      onChange?.({ ...value, options: temp });
                    }}
                    defaultValue={options[index]}
                    value={options[index]}
                    style={{ width: 200, marginTop: 10, marginRight: 10 }}
                  />
                  <Button
                    onClick={() => {
                      const temp = [...options];
                      temp.splice(index, 1);
                      setOptions(temp);
                      onChange?.({ ...value, options: temp });
                    }}
                    danger
                  >
                    删除
                  </Button>
                </div>
              ))}
            </Col>
          </>
        )}
        <Col span={4}>答案:</Col>
        {(type === 'multiple_choice' || value?.type === 'multiple_choice') && (
          <Col span={20}>
            <Select
              value={value?.answer || answer}
              onChange={changeAnswer}
              style={{ width: 80 }}
            >
              {options.map((item, index) => (
                <Select.Option key={letters[index]} value={letters[index]}>
                  {letters[index]}
                </Select.Option>
              ))}
            </Select>
          </Col>
        )}

        {(type === 'judge' || value?.type === 'judge') && (
          <Col span={20}>
            <Select
              value={value?.answer || answer}
              onChange={changeAnswer}
              style={{ width: 80 }}
            >
              <Select.Option value="A">正确</Select.Option>
              <Select.Option value="B">错误</Select.Option>
            </Select>
          </Col>
        )}
      </Row>
    </>
  );
};
export default AddTopicItem;
