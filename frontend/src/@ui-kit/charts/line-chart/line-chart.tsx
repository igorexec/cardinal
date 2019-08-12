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
      {score: 60, date: new Date('2019-08-11T16:10:36.955Z')},
      {score: 30, date: new Date('2019-08-12T16:10:36.955Z')},
      {score: 20, date: new Date('2019-08-13T16:10:36.955Z')},
      {score: 75, date: new Date('2019-08-14T16:10:36.955Z')},
      {score: 92, date: new Date('2019-08-15T16:10:36.955Z')},
    ];
    const minMaxDates = d3.extent(data, d => d.date) as [Date, Date];
    const xScale = d3.scaleLinear().domain(minMaxDates).range([0, 500]);
    const yScale = d3.scaleLinear().domain([100, 0]).range([0, 300]);

    const dateFormat = d3.timeFormat('%Y-%m-%d') as () => string;
    const makeXLines = d3.axisBottom(xScale).tickFormat(dateFormat);
    const makeYLines = d3.axisLeft(yScale);

    d3.select(this.el.current).append('g').attr('transform', `translate(30, 310)`).call(makeXLines);
    d3.select(this.el.current).append('g').attr('transform', 'translate(30, 10)').call(makeYLines);

    const line = d3.line<Score>().x(d => xScale(d.date)).y(d => yScale(d.score));
    d3.select(this.el.current)
      .append('path')
      .datum(data)
      .attr('transform', 'translate(30, 10)')
      .attr('fill', 'none')
      .attr('stroke', 'black')
      .attr('stroke-width', '3')
      .attr('d', line);
  }

  render() {
    return (
      <div>
        <svg ref={this.el} width="100%" height={350} />
      </div>
    );
  }
}
