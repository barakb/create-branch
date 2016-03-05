import { store } from "./store";
import { CREATE_BRANCH_REQUEST } from "./actionTypes"


export function createBranchRequest (name)  {
    console.info("createBranchRequest", name)
    store.dispatch({
                type:CREATE_BRANCH_REQUEST,
                name
            });
}

