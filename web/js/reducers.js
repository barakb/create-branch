const { combineReducers } = Redux;
import { CREATE_BRANCH_REQUEST, CREATE_BRANCH_RESPONSE, UPDATE_BRANCHES_FILTER, DELETE_BRANCH_REQUEST, BRANCH_ADDED, DELETE_BRANCH_RESPONSE, BRANCH_DELETED} from "./actionTypes"

const createBranchInitialState = {name:'', isFetching:false};

function createBranch (state = createBranchInitialState , action) {
   switch (action.type){
       case CREATE_BRANCH_REQUEST:
           return {...state, isFetching:true};
       case CREATE_BRANCH_RESPONSE:
           return {...state, isFetching:false};
       default:
           return state;
   }
}

const branches = [
             ];

const viewBranchesInitialState = {filterText:'', branches, filtered : []};

function viewBranch (state = viewBranchesInitialState , action) {
   switch (action.type){
       case UPDATE_BRANCHES_FILTER:
           let s = { ...state, filterText: action.filterText, filtered : state.branches.filter(b => -1 < b.name.indexOf(action.filterText))};
           return s;
      case DELETE_BRANCH_REQUEST:
           let branches = state.branches.map(b => b.name === action.name ? { ...b, isDeleting:true} : b)
           let filtered = branches.filter(b => -1 < b.name.indexOf(state.filterText))
          return  { ...state, branches, filtered };
      case DELETE_BRANCH_RESPONSE:
           branches = state.branches.map(b => b.name === action.name ? {...b, isDeleting:false} : b)
           filtered = branches.filter(b => -1 < b.name.indexOf(state.filterText))
          return  { ...state, branches, filtered };
      case BRANCH_ADDED:
           branches = [ ...state.branches, {name:action.name, quantity:action.quantity}];
           filtered = branches.filter(b => -1 < b.name.indexOf(state.filterText))
          return  { ...state, branches, filtered };
      case BRANCH_DELETED:
           branches = state.branches.filter(b => b.name !== action.name)
           filtered = branches.filter(b => -1 < b.name.indexOf(state.filterText))
          return  { ...state, branches, filtered };
       default:
           return { ...state, filtered : state.branches.filter(b => -1 < b.name.indexOf(state.filterText))};
   }
}

export const rootReducer = combineReducers({
  createBranch,
  viewBranch
})

