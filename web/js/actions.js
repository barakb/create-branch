import { CREATE_BRANCH_REQUEST, UPDATE_BRANCHES_FILTER } from "./actionTypes"


export function createBranchRequest(name) {
    console.info("createBranchRequest", name)
    return {type:CREATE_BRANCH_REQUEST, name};
}

export function updateBranchesFilterRequest(filterText) {
    console.info("updateBranchesFilterRequest", filterText)
    return {type:UPDATE_BRANCHES_FILTER, filterText};
}

