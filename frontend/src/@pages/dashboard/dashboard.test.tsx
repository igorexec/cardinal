import React from 'react';
import {shallow, ShallowWrapper} from 'enzyme';
import {Dashboard} from './dashboard';

describe('Dashboard', () => {
  let wrapper: ShallowWrapper;

  beforeAll(() => {
    wrapper = shallow(<Dashboard />);
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
