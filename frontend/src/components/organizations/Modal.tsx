import * as React from "react";

export interface Props {
  closeHandler: Function;
  Show: boolean;
  SourceUrl: string;
}

class Modal extends React.Component<Props> {
  closeHandler = () => {
    this.props.closeHandler();
  };
  render() {
    const show = this.props.Show;
    return (
      <div className="modal" style={{ display: show ? "block" : "none" }}>
        <div className="modal-content">
          <div className="close" onClick={this.closeHandler}>
            &times;
          </div>
          <iframe
            src={this.props.SourceUrl}
            className="source-url"
            title="SourceUrl"
          />
        </div>
      </div>
    );
  }
}

export default Modal;
