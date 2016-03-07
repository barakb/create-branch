import { createBranchRequest } from "./actions"
import { CreateBranch, CreateBranchContainer } from "./createBranch";
import { FilterableBranchesTableComponent } from "./branchesView";
import { rootReducer } from "./reducers";
import { thunkMiddleware as thunk} from './redux-thunk';
const { createStore, applyMiddleware} = Redux;
const { Provider } = ReactRedux;


export const store = createStore(rootReducer, applyMiddleware(thunk));


const render = () => {
    ReactDOM.render(
       <Provider store={store}>
          <div>
            <CreateBranchContainer />
            <hr/>
            <FilterableBranchesTableComponent />
          </div>
       </Provider>, document.getElementById('createBranch')
    );
};

store.subscribe(render);
render();

