import { createBranchRequest } from "./actions";

export const CreateBranch = ({ name, onClick, isFetching, fromBranch }) => {
         let input;
         let from = fromBranch ? ("Create From " + fromBranch) : "Create From master"
         let btn = isFetching ? <button type="submit" className="btn btn-default disabled">{from}</button> :
            <button type="submit" className="btn btn-default" onClick={() => onClick(input.value, fromBranch)}>{from}</button>

         return (
             <form className="navbar-form navbar-left" role="search" onSubmit={ev => ev.preventDefault() } >
                 <div className="form-group">
                     <input type="text" className="form-control" placeholder="New Branch" ref={node => {input = node}}  defaultValue={name} />
                 </div>
                 {btn}
                  <div className="form-group">
                    {isFetching ? <img src="/web/images/gears.svg" className="form-control" /> : null}
                  </div>
             </form>
          )
}


const PropTypes = React.PropTypes;
CreateBranch.propTypes = {
  onClick: PropTypes.func.isRequired,
  fromBranch : PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  isFetching: PropTypes.bool.isRequired,
};

const { Component } = React;

const mapStateToProps = (state) => {
    return {
        name : state.createBranch.name,
        fromBranch : state.createBranch.fromBranch,
        isFetching : state.createBranch.isFetching,
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        onClick : (name, fromBranch) =>  { dispatch(createBranchRequest(name, fromBranch)); }
    }
}

const { connect } = ReactRedux;

export const CreateBranchContainer = connect(mapStateToProps, mapDispatchToProps)(CreateBranch);

