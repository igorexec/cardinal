import React, {Component} from 'react';
import * as d3 from 'd3';

export class LineChart extends Component {
  private el = React.createRef<SVGSVGElement>();

  componentDidMount(): void {
  }

  render() {
    return (
      <div>
        <svg ref={this.el} />
      </div>
    );
  }
}
