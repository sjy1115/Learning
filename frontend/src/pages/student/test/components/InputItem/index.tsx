import { Input, Radio } from 'antd';
import React from 'react';

interface IProps {
  onChange?: (e: string) => void;
  question: string;
}
const InputItem: React.FC<IProps> = (props) => {
  const { onChange, question } = props;
  return (
    <div style={{ marginTop: 5 }}>
      <p>{question}</p>
      <Input.TextArea
        rows={4}
        onChange={(e) => {
          onChange?.(e.target.value);
        }}
      />
    </div>
  );
};
export default InputItem;
