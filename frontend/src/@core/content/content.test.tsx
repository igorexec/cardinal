import React from 'react';
import {shallow, ShallowWrapper} from 'enzyme';
import {Content} from './content';

describe('Content', () => {
  let wrapper: ShallowWrapper;

  beforeAll(() => {
    wrapper = shallow(<Content />);
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
