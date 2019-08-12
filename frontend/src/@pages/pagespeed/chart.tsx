import React from 'react';
import {LineChart} from '@ui-kit/charts';

interface Props {
  title: string;
  data?: any;
}

export const PageSpeedChart: React.FC<Props> = ({title}) => {
  const lines = [{
    data: [
      {score: 0, date: new Date('2019-08-11T00:00:00.000')},
      {score: 100, date: new Date('2019-08-12T00:00:00.000')},
      {score: 0, date: new Date('2019-08-13T00:00:00.000')},
      {score: 100, date: new Date('2019-08-14T00:00:00.000')},
      {score: 0, date: new Date('2019-08-15T00:00:00.000')},
      {score: 100, date: new Date('2019-08-16T00:00:00.000')},
    ],
    name: 'www.toryburch.com',
  }, {
    data: [
      {score: 33, date: new Date('2019-08-11T00:00:00.000')},
      {score: 71, date: new Date('2019-08-12T00:00:00.000')},
      {score: 54, date: new Date('2019-08-13T00:00:00.000')},
      {score: 81, date: new Date('2019-08-14T00:00:00.000')},
      {score: 13, date: new Date('2019-08-15T00:00:00.000')},
      {score: 77, date: new Date('2019-08-16T00:00:00.000')},
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
