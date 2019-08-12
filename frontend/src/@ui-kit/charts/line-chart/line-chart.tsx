import React, {Component} from 'react';
import * as d3 from 'd3';

interface Score {
  score: number;
  date: Date;
}

export class LineChart extends Component {
  private el = React.createRef<SVGSVGElement>();

  componentDidMount(): void {
    const data: Score[] = [
      {score: 0, date: new Date('2019-08-11T00:00:00.000')},
      {score: 100, date: new Date('2019-08-12T00:00:00.000')},
      {score: 0, date: new Date('2019-08-13T00:00:00.000')},
      {score: 100, date: new Date('2019-08-14T00:00:00.000')},
      {score: 0, date: new Date('2019-08-15T00:00:00.000')},
      {score: 100, date: new Date('2019-08-16T00:00:00.000')},
    ];
    const minMaxDates = d3.extent(data, d => d.date) as [Date, Date];
    const xScale = d3.scaleTime().domain(minMaxDates).range([0, 1000]);
    const yScale = d3.scaleLinear().domain([100, 0]).range([0, 300]);

    const dateFormat = d3.timeFormat('%Y-%m-%d %HH-%MM') as () => string;
    const makeXLines = d3.axisBottom(xScale).tickFormat(dateFormat).ticks(6);
    const makeYLines = d3.axisLeft(yScale);

    d3.select(this.el.current).append('g').call(makeXLines);
    d3.select(this.el.current).append('g').call(makeYLines);

    const line = d3.line<Score>().x(d => xScale(d.date)).y(d => yScale(d.score));
    d3.select(this.el.current)
      .append('path')
      .datum(data)
      .attr('fill', 'none')
      .attr('stroke', 'black')
      .attr('stroke-width', '1')
      .attr('d', line);
  }

  render() {
    return (
      <div>
        <svg ref={this.el} width="100%" height={300} />
      </div>
    );
  }
}
