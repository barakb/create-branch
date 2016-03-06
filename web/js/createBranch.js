import { createBranchRequest } from "./actions";

export const CreateBranch = ({name, onClick}) => {
         let input;
         return (
              <form className="form-inline">
                  <div className="form-group">
                    <input type="text" className="form-control" placeholder="Branch Name" ref={node => {input = node;}}  defaultValue={name} />
                  </div>
                  <div className="form-group">
                    <button type='button' onClick={() => onClick(input.value)} className="btn btn-default form-control">Create</button>
                  </div>
                  <div className="form-group">
                    <img src="/web/images/gears.svg" id="processing" className="form-control" />
                  </div>
              </form>
          )
}


//const PropTypes = React.PropTypes;
//CreateBranch.propTypes = {
//  onClick: PropTypes.func.isRequired,
//  name: PropTypes.string.isRequired
//};

const { Component }  = React;

export class CreateBranchContainer extends Component{

  render() {
    const props = this.props;
    const { store } = this.context;
    return ( <CreateBranch name={ store.getState().createBranch.name } onClick = { branch => store.dispatch(createBranchRequest(branch)) } /> );
  }
};

CreateBranchContainer.contextTypes = {
    store : React.PropTypes.object
};



