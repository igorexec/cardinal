import React from 'react';
import {shallow, ShallowWrapper} from 'enzyme';
import {Root} from './root';

describe('Root', () => {
  let wrapper: ShallowWrapper;

  beforeAll(() => {
    wrapper = shallow(<Root />);
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
