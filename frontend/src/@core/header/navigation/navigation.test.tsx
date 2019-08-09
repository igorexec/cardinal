import React from 'react';
import {shallow, ShallowWrapper} from 'enzyme';
import {Navigation} from './navigation';

describe('Navigation', () => {
  let wrapper: ShallowWrapper;
  const props = {items: [{id: '1', title: 'title', link: 'link'}]};

  beforeAll(() => {
    wrapper = shallow(<Navigation {...props} />);
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
