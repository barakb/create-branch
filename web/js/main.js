import { store } from "./store";

class MainView extends React.Component {
  render() {
    return (
    	<div>
    		<CreateBranch />
    		<hr/>
          	<FilterableBranchesTable branches={window.branches} />
    	</div>
    	)
  }
}
window.branches  = []

//$(function(){
//        window.branches = [
//          {"name": 'foo', "quantity":3},
//          {"name": 'foo1', "quantity":2},
//          {"name": 'foo2', "quantity":1},
//       ];
//       ReactDOM.render( <MainView  />, document.getElementById('app'))
//});

//var MainView = React.createClass({
//    render:function(){
//        return (
//         <div>
//           <CreateBranch  />
//           <hr/>
//           <FilterableBranchesTable branches={branches} />
//        </div>
//        );
//    }
//});
