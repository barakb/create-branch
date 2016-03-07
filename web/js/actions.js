import { CREATE_BRANCH_REQUEST, UPDATE_BRANCHES_FILTER, DELETE_BRANCH_REQUEST, BRANCH_ADDED, CREATE_BRANCH_RESPONSE, DELETE_BRANCH_RESPONSE, BRANCH_DELETED } from "./actionTypes"



export function updateBranchesFilterRequest(filterText) {
    console.info("updateBranchesFilterRequest", filterText)
    return {type:UPDATE_BRANCHES_FILTER, filterText};
}

export function deleteBranchRequest(name) {
    console.info("deleteBranchRequest", name);
    return function(dispatch){
        dispatch({type:DELETE_BRANCH_REQUEST, name});
        setTimeout(() => dispatch(deleteBranchResponse(name, null)), 1000)
    }
}

export function deleteBranchResponse(name, err){
    console.info("deleteBranchResponse", name, err)
    return function(dispatch){
        dispatch({type:DELETE_BRANCH_RESPONSE, name, err:err});
        if(!err){
            dispatch(branchDeleted(name));
        }
    }
}

export function branchAdded(name, quantity) {
    console.info("branchAdded", name, quantity)
    return {type:BRANCH_ADDED, name, quantity};
}

export function branchDeleted(name) {
    console.info("branchAdded", name)
    return {type:BRANCH_DELETED, name};
}

export function createBranchRequest(name) {
    console.info("createBranchRequest", name)
    return function(dispatch){
        dispatch({type:CREATE_BRANCH_REQUEST, name});
        setTimeout(() => dispatch(createBranchResponse(name, 1, null)), 1000)
    }
}

export function createBranchResponse(name, quantity, err) {
    console.info("createBranchResponse", name, quantity)
    return function(dispatch){
        dispatch({type:CREATE_BRANCH_RESPONSE, name, quantity, err:err});
        if(!err){
            dispatch(branchAdded(name, quantity));
        }
    }
}

