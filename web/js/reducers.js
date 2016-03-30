const { combineReducers } = Redux;
import { CREATE_BRANCH_REQUEST,
         CREATE_BRANCH_RESPONSE,
         UPDATE_BRANCHES_FILTER,
         DELETE_BRANCH_REQUEST,
         BRANCH_ADDED,
         DELETE_BRANCH_RESPONSE,
         BRANCH_DELETED,
         FOOBAR_TOGGLE_SHOW_VALUES,
         TOGGLED_BRANCH_ROW
        }
 from "./actionTypes"

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
           branches.sort((b1, b2) => b1.name > b2.name);
           let filtered = branches.filter(b => -1 < b.name.indexOf(state.filterText))
          return  { ...state, branches, filtered };
      case DELETE_BRANCH_RESPONSE:
           branches = state.branches.map(b => b.name === action.name ? {...b, isDeleting:false} : b)
           branches.sort((b1, b2) => b1.name > b2.name);
           filtered = branches.filter(b => -1 < b.name.indexOf(state.filterText))
          return  { ...state, branches, filtered };
      case BRANCH_ADDED:
           branches = [ ...state.branches, {name:action.name, repositories:action.repositories, expanded: false}];
           branches.sort((b1, b2) => b1.name > b2.name);
           filtered = branches.filter(b => -1 < b.name.indexOf(state.filterText))
          return  { ...state, branches, filtered };
      case BRANCH_DELETED:
           branches = state.branches.filter(b => b.name !== action.name)
           branches.sort((b1, b2) => b1.name > b2.name);
           filtered = branches.filter(b => -1 < b.name.indexOf(state.filterText))
          return  { ...state, branches, filtered };
       case TOGGLED_BRANCH_ROW:
           console.log( action );
           branches = state.branches.map( b => b.name === action.name ? {...b, expanded:!b.expanded} : b );
           branches.sort((b1, b2) => b1.name > b2.name);
           filtered = branches.filter(b => -1 < b.name.indexOf(state.filterText))
           return  { ...state, branches, filtered };

       default:
           return { ...state, filtered : state.branches.filter(b => -1 < b.name.indexOf(state.filterText))};
   }
}

const fooBarInitialState = [
    {"name": "name1", "id": "1", "showValues": false, "values": ["a", "b", "c"]},
    {"name": "name2", "id": "2", "showValues": false, "values": ["a", "b", "c"]},
    {"name": "name3", "id": "3", "showValues": false, "values": ["a", "b", "c"]}
]

function fooBar(state = fooBarInitialState, action) {
    switch (action.type){
        case FOOBAR_TOGGLE_SHOW_VALUES:
            return state.map(c => c.id === action.id ? { ...c, showValues:!c.showValues} : c);
       default:
            return state;
    }
}

export const rootReducer = combineReducers({
  createBranch,
  viewBranch,
  fooBar
})

