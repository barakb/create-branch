import { createBranchRequest } from "./actions";

export const CreateBranch = ({ name, onClick, isFetching }) => {
         let input;
         let btn = isFetching ? <button type="submit" className="btn btn-default disabled">Create</button> :
            <button type="submit" className="btn btn-default" onClick={() => onClick(input.value)}>Create</button>

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
  name: PropTypes.string.isRequired,
  isFetching: PropTypes.bool.isRequired
};

const { Component } = React;

const mapStateToProps = (state) => {
    return {
        name : state.createBranch.name,
        isFetching : state.createBranch.isFetching
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        onClick : (name) =>  { dispatch(createBranchRequest(name)); }
    }
}

const { connect } = ReactRedux;

export const CreateBranchContainer = connect(mapStateToProps, mapDispatchToProps)(CreateBranch);



