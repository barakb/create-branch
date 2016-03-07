import { createBranchRequest } from "./actions";

export const CreateBranch = ({ name, onClick, isFetching }) => {
         let input;
         return (
              <form className="form-inline">
                  <div className="form-group">
                    <input type="text" className="form-control" placeholder="Branch Name" ref={node => {input = node}}  defaultValue={name} />
                  </div>
                  <div className="form-group">
                    <button type='button' onClick={() => onClick(input.value)} className="btn btn-default form-control">Create</button>
                  </div>
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



