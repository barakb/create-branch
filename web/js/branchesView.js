const PropTypes = React.PropTypes;
import { updateBranchesFilterRequest, deleteBranchRequest } from "./actions";

export const InternalRepositoryRow = ({ repository }) => {
    return (
            <tr>
                <td>{repository}</td>
            </tr>
    );
}

export const BranchRow = ({ name, repositories, onRemove, isDeleting}) => {
         let s = repositories ? repositories.length : 0;
         let btn= isDeleting ? <button type="button" className="btn btn-default btn-sm disabled" ><span className="glyphicon glyphicon-remove"></span> Remove</button> :
         <button type="button" className="btn btn-default btn-sm" onClick={() => onRemove(name)}><span className="glyphicon glyphicon-remove"></span> Remove</button>;

         let repositoriesRows = repositories.map((repository) => <InternalRepositoryRow repository={repository}></InternalRepositoryRow>);

         let expandbtn= <button type="button" className="btn btn-default btn-sm"><span className="glyphicon glyphicon-eye-open"></span></button>;
         let repositoriesExpandedData =  name + '-repositories-expanded-data';
         let targetToRepositoriesExpandedData =  '#' + repositoriesExpandedData;

         return (
             <tbody>
             <tr key={name} data-toggle="collapse" data-target={targetToRepositoriesExpandedData} className="accordion-toggle">
                 <td>{expandbtn}</td>
                 <td>{name}</td>
                 <td>{s}</td>
                 <td>{btn}</td>
             </tr>
             <tr>
                 <td className="hiddenRow" colSpan="4">
                     <div className="accordian-body collapse" id={repositoriesExpandedData}><table>{repositoriesRows}</table>
                     </div>
                 </td>
             </tr>
             </tbody>
          )
}


export const BranchesTable = ({ filtered, onRemove }) => {
    let rows = filtered.map((branch) => <BranchRow name={branch.name} repositories={branch.repositories} key={branch.name} onRemove={onRemove} isDeleting={branch.isDeleting}></BranchRow>);
    return (
      <table className="table table-hover table-condensed table-responsive">
        <thead>
          <tr>
            <th width="5%">&nbsp;</th>
            <th width="40%">Branch</th>
            <th width="40%">Repositories</th>
            <th width="5%"></th>
          </tr>
        </thead>
          {rows}
      </table>
    );
}


export const SearchBar = ({ filterText, handleUserInput }) => {
    let input;
    return (
      <form className="form-inline">
        <div className="form-group">
           <input type="text" placeholder="Filter..." value={filterText}  ref={node => {input = node}} onChange={() => handleUserInput(input.value)} className="form-control"/>
        </div>
      </form>
    );
}

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
  filtered : PropTypes.arrayOf(
       PropTypes.shape({name : PropTypes.string.isRequired,
                        statuses : PropTypes.arrayOf(PropTypes.shape({name: PropTypes.string.isRequired, success : PropTypes.bool.isRequired}))})),
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