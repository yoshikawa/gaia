import * as React from 'react';

// Propsの型定義
interface IProps {
  name: string;
}

interface IState {
  count: number;
}

export class SubComponent extends React.Component<IProps, IState> {
  render() {

    return (
        <div>
          <h2>{this.props.name}</h2>
          <div>{this.state.count}</div>
        </div>
    );
  }
}
