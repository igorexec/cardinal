import React, {Component} from 'react';
import * as d3 from 'd3';
import {LineData, CoordData} from './types';

interface Props<T, U> {
  data: Array<LineData<T, U>>;
}

export class LineChart<T extends d3.Numeric, U extends d3.Numeric> extends Component<Props<T, U>> {
  private el = React.createRef<SVGSVGElement>();

  componentDidMount(): void {
    const {data} = this.props;

    const points = data.reduce((acc, l) => {
      return [...acc, ...l.data];
    }, [] as Array<CoordData<T, U>>);
    const [minX, maxX] = d3.extent(points, p => p.x) as [T, T];
    const [minY, maxY] = d3.extent(points, p => p.y) as [U, U];

    const xScale = d3.scaleTime().domain([minX, maxX]).range([0, 1000]);
    const yScale = d3.scaleLinear().domain([maxY, minY]).range([0, 300]);

    const dateFormat = d3.timeFormat('%Y-%m-%d') as () => string;

    const makeXLines = d3.axisBottom(xScale).tickFormat(dateFormat);
    const makeYLines = d3.axisLeft(yScale);

    d3.select(this.el.current).append('g').call(makeXLines);
    d3.select(this.el.current).append('g').call(makeYLines);

    data.forEach(l => {
      const line = d3.line<CoordData<T, U>>().x(d => xScale(d.x)).y(d => yScale(d.y));
      d3.select(this.el.current)
        .append('path')
        .datum(l.data)
        .attr('fill', 'none')
        .attr('stroke', 'black')
        .attr('stroke-width', '1')
        .attr('d', line);
    });
  }

  render() {
    return (
      <div>
        <svg ref={this.el} width="100%" height={300} />
      </div>
    );
  }
}
