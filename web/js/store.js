import { createBranchRequest } from "./actions"
import { CreateBranch, CreateBranchContainer } from "./createBranch";
import { rootReducer } from "./reducers";
import { thunkMiddleware as thunk} from './redux-thunk';

const { createStore, applyMiddleware} = Redux;
const { Provider } = ReactRedux;


export const store = createStore(rootReducer, applyMiddleware(thunk));


const render = () => {
    ReactDOM.render(
       <Provider store={store}>
          <CreateBranchContainer />
       </Provider>, document.getElementById('createBranch')
    );
};

//       <CreateBranch
//           name={store.getState().createBranch.name}
//           onClick={createBranchRequest}
//       />

store.subscribe(render);
render();

