const PropTypes = React.PropTypes;
import { updateBranchesFilterRequest, deleteBranchRequest } from "./actions";

export const BranchRow = ({ name, quantity, onRemove, isDeleting}) => {
         let btn= isDeleting ? <button type="button" className="btn btn-default btn-sm disabled" ><span className="glyphicon glyphicon-remove"></span> Remove</button> :
         <button type="button" className="btn btn-default btn-sm" onClick={() => onRemove(name)}><span className="glyphicon glyphicon-remove"></span> Remove</button>
         return (
              <tr key={name}>
                <td>{name}</td>
                <td>{quantity}</td>
                <td>{btn}</td>
              </tr>
          )
}


export const BranchesTable = ({ filtered, onRemove }) => {
    let rows = filtered.map((branch) => <BranchRow name={branch.name}  quantity={branch.quantity} key={branch.name} onRemove={onRemove} isDeleting={branch.isDeleting}/>);
    return (
      <table className="table table-hover table-responsive">
        <thead>
          <tr>
            <th>Branch</th>
            <th>Quantity</th>
            <th></th>
          </tr>
        </thead>
        <tbody>{rows}</tbody>
      </table>
    );
}

BranchesTable.propTypes = {
  filtered : PropTypes.arrayOf(PropTypes.shape({name : PropTypes.string.isRequired, quantity : PropTypes.number.isRequired})),
  onRemove: PropTypes.func.isRequired
};

export const SearchBar = ({ filterText, handleUserInput }) => {
    let input;
    return (
      <form className="form-inline">
        <div className="form-group">
           <input type="text" placeholder="Search..." value={filterText}  ref={node => {input = node}} onChange={() => handleUserInput(input.value)} className="form-control"/>
        </div>
      </form>
    );
}
SearchBar.propTypes = {
  handleUserInput: PropTypes.func.isRequired,
  filterText: PropTypes.string.isRequired
};



export const FilterableBranchesTable = ({ filterText, filtered, handleUserInput, onRemove }) => {
    return (
      <div>
         <SearchBar filterText={filterText} handleUserInput={handleUserInput}/>
         <BranchesTable filtered={filtered}  onRemove={onRemove}/>
      </div>
    );
}

FilterableBranchesTable.propTypes = {
  handleUserInput: PropTypes.func.isRequired,
  filterText: PropTypes.string.isRequired,
  filtered : PropTypes.arrayOf(PropTypes.shape({name : PropTypes.string.isRequired, quantity : PropTypes.number.isRequired})),
  onRemove: PropTypes.func.isRequired
};


const { Component } = React;

const mapStateToProps = (state) => {
    return {
        filterText : state.viewBranch.filterText,
        filtered : state.viewBranch.filtered
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        handleUserInput : (text) =>  { dispatch(updateBranchesFilterRequest(text)) },
        onRemove : (name) =>  { dispatch(deleteBranchRequest(name)) }
    }
}

const { connect } = ReactRedux;


export const FilterableBranchesTableComponent = connect(mapStateToProps, mapDispatchToProps)(FilterableBranchesTable);