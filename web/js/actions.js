import { CREATE_BRANCH_REQUEST } from "./actionTypes"


export function createBranchRequest(name) {
    console.info("createBranchRequest", name)
    return {type:CREATE_BRANCH_REQUEST, name};
}

