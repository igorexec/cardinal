import React from 'react';
import {LineChart} from '@ui-kit/charts';

interface Props {
  title: string;
  data?: any;
}

export const PageSpeedChart: React.FC<Props> = ({title}) => {
  return (
    <section>
      <h1>{title}</h1>
      <LineChart />
    </section>
  );
};
