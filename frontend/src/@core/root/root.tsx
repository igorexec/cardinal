import React from 'react';
import {Header} from '@core/header';
import {Content} from '@core/content';
import {Footer} from '@core/footer';

export const Root: React.FC = () => {
  const navItems = [{
    id: '1', link: '/pagespeed', title: 'PageSpeed',
  }, {
    id: '2', link: '/crawl', title: 'Crawl Issues',
  }];

  return (
    <>
      <Header navItems={navItems} />
      <Content />
      <Footer />
    </>
  );
};
