import React from 'react';
import {BrowserRouter, Switch, Route, Redirect} from 'react-router-dom';
import {Header} from '@core/header';
import {Footer} from '@core/footer';
import {Dashboard, PageSpeed} from '@pages';

export const Root: React.FC = () => {
  const navItems = [{
    id: '1', link: '/pagespeed', title: 'PageSpeed',
  }, {
    id: '2', link: '/crawl', title: 'Crawl Issues',
  }];

  return (
    <BrowserRouter>
      <Header navItems={navItems} />
      <Switch>
        <Route exact path="/" component={Dashboard} />
        <Route path="/pagespeed" component={PageSpeed} />
        <Redirect to="/" />
      </Switch>
      <Footer />
    </BrowserRouter>
  );
};
