import { letters } from '@/utils/constant';
import { Radio } from 'antd';
import React from 'react';

interface IProps {
  onChange?: (e: string) => void;
  question: string;
  options: string[];
}
const SelectItem: React.FC<IProps> = (props) => {
  const { onChange, question, options } = props;
  return (
    <div style={{ marginTop: 5 }}>
      <p>{question}</p>
      <Radio.Group
        onChange={(e) => {
          onChange?.(e.target.value);
        }}
      >
        {options.map((item, index) => (
          <Radio key={index} value={letters[index]}>
            {letters[index]}:{item}
          </Radio>
        ))}
      </Radio.Group>
    </div>
  );
};
export default SelectItem;
