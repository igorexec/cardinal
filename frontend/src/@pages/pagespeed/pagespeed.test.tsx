import React from 'react';
import {shallow, ShallowWrapper} from 'enzyme';
import {PageSpeed} from './pagespeed';

describe('PageSpeed', () => {
  let wrapper: ShallowWrapper;

  beforeAll(() => {
    wrapper = shallow(<PageSpeed />);
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
