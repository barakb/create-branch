import { createBranchRequest } from "./actions"
import { CreateBranch } from "./createBranch";
import { rootReducer } from "./reducers";

const { createStore} = Redux;

export const store = createStore(rootReducer);


const render = () => {
    ReactDOM.render(
       <CreateBranch
           name={store.getState().createBranch.name}
           onClick={createBranchRequest}
       />, document.getElementById('createBranch')
    );
};

store.subscribe(render);
render();

