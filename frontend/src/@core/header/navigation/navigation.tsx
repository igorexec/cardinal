import React from 'react';
import {NavLink} from 'react-router-dom';
import {NavigationItem} from '../types';

interface Props {
  items: NavigationItem[];
}

export const Navigation: React.FC<Props> = ({items}) => {
  const navItems = items.map(({id, link, title}) => <NavLink key={id} to={link}>{title}</NavLink>);

  return (
    <nav>
      <ul>{navItems}</ul>
    </nav>
  );
};
