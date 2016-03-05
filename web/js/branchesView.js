

class BranchRow extends React.Component{
  render() {
    console.info("this.props.branch", this.props.branch)
    return (
      <tr key={this.props.branch.name}>
        <td>{this.props.branch.name}</td>
        <td>{this.props.branch.quantity}</td>
      </tr>
    );
  }
}


class BranchesTable extends React.Component{
  render() {
    var rows = [];
    var props = this.props;
    this.props.branches.forEach(function(branch) {
        console.info("comparing ", branch, " with ", props.filterText)
        if (branch.name.indexOf(props.filterText) === -1){
            return;
        }
        rows.push(<BranchRow branch={branch}  key={branch.name}/>);
    });
    return (
      <table className="table table-hover table-responsive">
        <thead>
          <tr>
            <th>Branch</th>
            <th>Quantity</th>
          </tr>
        </thead>
        <tbody>{rows}</tbody>
      </table>
    );
  }
};


class SearchBar extends React.Component{
  handleChange() {
      this.props.onUserInput(this.refs.filterTextInput.value);
  }
  render() {
    var that = this
    return (
      <form className="form-inline">
        <div className="form-group">
           <input type="text" placeholder="Search..." value={this.props.filterText}  ref="filterTextInput" onChange={this.handleChange.bind(this)} className="form-control"/>
        </div>
      </form>
    );
  }
};


class FilterableBranchesTable extends React.Component{
  constructor() {
    super();
    this.state = {
      filterText: ''
    };
  }
  handleUserInput(filterText) {
    console.info("FilterableBranchesTable:handleUserInput", filterText)
    this.setState({
      filterText: filterText
    });
  }
  render() {
    return (
      <div>
         <SearchBar filterText={this.state.filterText} onUserInput={this.handleUserInput.bind(this)}/>
         <BranchesTable branches={this.props.branches} filterText={this.state.filterText} />
      </div>
    );
  }
};