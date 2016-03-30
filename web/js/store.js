import { createBranchRequest, branchAdded } from "./actions"
import { CreateBranch, CreateBranchContainer } from "./createBranch";
import { FilterableBranchesTableComponent } from "./branchesView";
import { rootReducer } from "./reducers";
import { thunkMiddleware as thunk} from './redux-thunk';
const { createStore, applyMiddleware} = Redux;
const { Provider } = ReactRedux;


export const store = createStore(rootReducer, applyMiddleware(thunk));


ReactDOM.render(
   <Provider store={store}>
        <FilterableBranchesTableComponent />
   </Provider>, document.getElementById('listBranches')
);

ReactDOM.render(
   <Provider store={store}>
      <CreateBranchContainer />
   </Provider>, document.getElementById('createBranch')
);



const loadBranches = () => {
    fetch('api/get_branches/', {
    	method: 'get',
    	credentials: 'same-origin',
    	cache: 'no-cache'
    }).then(function(response) {
        return response.json()
    }).then(function(branches){
        console.info("branches ", branches)
        for (var key of Object.keys(branches)) {
           if(key.startsWith(document.currentLoginName + "_")){
                store.dispatch(branchAdded(key,  Object.keys(branches[key])));
           }
        }
    }).catch(function(err) {
        console.info("err is ", err)
    });
}

loadBranches();