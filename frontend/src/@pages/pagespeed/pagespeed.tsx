import React from 'react';
import {PageSpeedChart} from './chart';

export const PageSpeed: React.FC = () => {
  return (
    <div>
      <PageSpeedChart title="Mobile" />
      <PageSpeedChart title="Desktop" />
    </div>
  );
};
