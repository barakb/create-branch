const PropTypes = React.PropTypes;
import { updateBranchesFilterRequest } from "./actions";

export const BranchRow = ({ name, quantity }) => {
         return (
              <tr key={name}>
                <td>{name}</td>
                <td>{quantity}</td>
              </tr>
          )
}


export const BranchesTable = ({ filtered }) => {
    let rows = filtered.map((branch) => <BranchRow name={branch.name}  quantity={branch.quantity} key={branch.name}/>);
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

BranchesTable.propTypes = {
  filtered : PropTypes.arrayOf(PropTypes.shape({name : PropTypes.string.isRequired, quantity : PropTypes.number.isRequired}))
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



export const FilterableBranchesTable = ({ filterText, filtered, handleUserInput }) => {
    return (
      <div>
         <SearchBar filterText={filterText} handleUserInput={handleUserInput}/>
         <BranchesTable filtered={filtered} />
      </div>
    );
}

FilterableBranchesTable.propTypes = {
  handleUserInput: PropTypes.func.isRequired,
  filterText: PropTypes.string.isRequired,
  filtered : PropTypes.arrayOf(PropTypes.shape({name : PropTypes.string.isRequired, quantity : PropTypes.number.isRequired}))
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
        handleUserInput : (text) =>  { dispatch(updateBranchesFilterRequest(text)) }
    }
}

const { connect } = ReactRedux;


export const FilterableBranchesTableComponent = connect(mapStateToProps, mapDispatchToProps)(FilterableBranchesTable);