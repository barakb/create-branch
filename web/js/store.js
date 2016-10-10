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
    }).then(function(res){
        let branches = res.branches;
        let repos = res.repos;
        for (var key of Object.keys(branches)) {
                if(Object.keys(branches[key]).length ==  repos.length){
                    store.dispatch(branchAdded(key,  Object.keys(branches[key]), !key.startsWith(document.currentLoginName + "_")));
                }
        }
    }).catch(function(err) {
        console.info("err is ", err)
    });
}

loadBranches();