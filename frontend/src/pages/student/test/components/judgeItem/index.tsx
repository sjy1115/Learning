import { Radio } from 'antd';
import React from 'react';

interface IProps {
  onChange?: (e: string) => void;
  question: string;
}
const JudgeItem: React.FC<IProps> = (props) => {
  const { onChange, question } = props;
  return (
    <div style={{ marginTop: 5 }}>
      <p>{question}</p>
      <Radio.Group
        onChange={(e) => {
          onChange?.(e.target.value);
        }}
      >
        <Radio value="A">A:正确</Radio>
        <Radio value="B">B:错误</Radio>
      </Radio.Group>
    </div>
  );
};
export default JudgeItem;
