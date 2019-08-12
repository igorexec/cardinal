import React, {Component} from 'react';
import * as d3 from 'd3';

interface Score {
  score: number;
  date: Date;
}

interface Props {
  data: Array<{
    name: string;
    data: Score[];
  }>;
}

export class LineChart extends Component<Props> {
  private el = React.createRef<SVGSVGElement>();

  componentDidMount(): void {
    const {data} = this.props;

    const scores = data.reduce((acc, l) => {
      return [...acc, ...l.data];
    }, [] as Score[]);
    const [minDate, maxDate] = d3.extent(scores, s => s.date) as [Date, Date];
    const xScale = d3.scaleTime().domain([minDate, maxDate]).range([0, 1000]);
    const yScale = d3.scaleLinear().domain([100, 0]).range([0, 300]);
    const dateFormat = d3.timeFormat('%Y-%m-%d') as () => string;
    const makeXLines = d3.axisBottom(xScale).tickFormat(dateFormat);
    const makeYLines = d3.axisLeft(yScale);

    d3.select(this.el.current).append('g').call(makeXLines);
    d3.select(this.el.current).append('g').call(makeYLines);

    data.forEach(l => {
      const line = d3.line<Score>().x(d => xScale(d.date)).y(d => yScale(d.score));
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
