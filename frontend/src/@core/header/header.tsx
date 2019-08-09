import React from 'react';
import {NavigationItem} from './types';
import {Navigation} from './navigation';

interface Props {
  navItems: NavigationItem[];
}

export const Header: React.FC<Props> = ({navItems}) => {
  return (
    <header>
      <Navigation items={navItems} />
    </header>
  );
};
