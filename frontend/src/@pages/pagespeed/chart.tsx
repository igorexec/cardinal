import React from 'react';
import {LineChart} from '@ui-kit/charts';

interface Props {
  title: string;
  data?: any;
}

export const PageSpeedChart: React.FC<Props> = ({title}) => {
  const lines = [{
    data: [
      {y: 0, x: new Date('2019-08-11T00:00:00.000')},
      {y: 100, x: new Date('2019-08-12T00:00:00.000')},
      {y: 0, x: new Date('2019-08-13T00:00:00.000')},
      {y: 100, x: new Date('2019-08-14T00:00:00.000')},
      {y: 0, x: new Date('2019-08-15T00:00:00.000')},
      {y: 100, x: new Date('2019-08-16T00:00:00.000')},
    ],
    name: 'www.toryburch.com',
  }, {
    data: [
      {y: 61, x: new Date('2019-08-11T00:00:00.000')},
      {y: 71, x: new Date('2019-08-12T00:00:00.000')},
      {y: 54, x: new Date('2019-08-13T00:00:00.000')},
      {y: 81, x: new Date('2019-08-14T00:00:00.000')},
      {y: 13, x: new Date('2019-08-15T00:00:00.000')},
      {y: 77, x: new Date('2019-08-16T00:00:00.000')},
    ],
    name: 'www.toryburch1.com',
  }];
  return (
    <section>
      <h1>{title}</h1>
      <LineChart data={lines} />
    </section>
  );
};
