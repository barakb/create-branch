const { combineReducers } = Redux;
import { CREATE_BRANCH_REQUEST} from "./actionTypes"

const createBranchInitialState = {name:''};

function createBranch (state = createBranchInitialState , action) {
   switch (action.type){
       case CREATE_BRANCH_REQUEST:
           return {...state, name: action.name};
       default:
           return state;
   }
}

export const rootReducer = combineReducers({
  createBranch
})

