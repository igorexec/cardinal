import React from 'react';
import {shallow, ShallowWrapper} from 'enzyme';
import {App} from './App';

describe('App', () => {
  let wrapper: ShallowWrapper;

  beforeAll(() => {
    wrapper = shallow(<App />);
  });

  it('matches snapshot without props', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
