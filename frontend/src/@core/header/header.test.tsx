import React from 'react';
import {shallow, ShallowWrapper} from 'enzyme';
import {Header} from './header';

describe('Header', () => {
  let wrapper: ShallowWrapper;
  const props = {navItems: [{id: '1', title: 'title', link: 'link'}]};

  beforeAll(() => {
    wrapper = shallow(<Header {...props} />);
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
