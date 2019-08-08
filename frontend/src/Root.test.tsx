import React from 'react';
import {Root} from './Root';
import {shallow, ShallowWrapper} from 'enzyme';

describe('Root', () => {
  let wrapper: ShallowWrapper;

  beforeAll(() => {
    wrapper = shallow(<Root />);
  });

  it('matches snapshot without props', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
