const { combineReducers } = Redux;
import { CREATE_BRANCH_REQUEST, UPDATE_BRANCHES_FILTER} from "./actionTypes"

const createBranchInitialState = {name:''};

function createBranch (state = createBranchInitialState , action) {
   switch (action.type){
       case CREATE_BRANCH_REQUEST:
           return {...state, name: action.name};
       default:
           return state;
   }
}

const branches = [
                {"name": 'foo', "quantity":3},
                {"name": 'foo1', "quantity":2},
                {"name": 'foo2', "quantity":1},
             ];

const viewBranchesInitialState = {filterText:'', branches, filtered : []};

function viewBranch (state = viewBranchesInitialState , action) {
   switch (action.type){
       case UPDATE_BRANCHES_FILTER:
           console.info("action is ", action);
           let s = {...state, filterText: action.filterText, filtered : branches.filter(b => -1 < b.name.indexOf(action.filterText))};
           console.info("new state is", s);
           return s;
       default:
           return {...state, filtered : state.branches.filter(b => -1 < b.name.indexOf(state.filterText))};
   }
}

export const rootReducer = combineReducers({
  createBranch,
  viewBranch
})

