import React from 'react';
import {shallow, ShallowWrapper} from 'enzyme';
import {Footer} from './footer';

describe('Footer', () => {
  let wrapper: ShallowWrapper;

  beforeAll(() => {
    wrapper = shallow(<Footer />);
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
