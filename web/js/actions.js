import { CREATE_BRANCH_REQUEST, UPDATE_BRANCHES_FILTER, DELETE_BRANCH_REQUEST, BRANCH_ADDED, CREATE_BRANCH_RESPONSE, DELETE_BRANCH_RESPONSE, BRANCH_DELETED } from "./actionTypes"



export function updateBranchesFilterRequest(filterText) {
//    console.info("updateBranchesFilterRequest", filterText)
    return {type:UPDATE_BRANCHES_FILTER, filterText};
}

export function deleteBranchRequest(name) {
//    console.info("deleteBranchRequest", name);
    return function(dispatch){
        if (!name){
            return
        }
        dispatch({type:DELETE_BRANCH_REQUEST, name});
        fetch('/api/delete_branch/' + name, {
                    method: 'delete',
                    credentials: 'same-origin',
                    cache: 'no-cache'
                }).then(function(response) {
                    return response.json()
                }).then(function(branch){
                    dispatch(deleteBranchResponse(name, null))
                }).catch(function(err) {
                    console.info("err is ", err)
                });
    }
}

export function deleteBranchResponse(name, err){
//    console.info("deleteBranchResponse", name, err)
    return function(dispatch){
        dispatch({type:DELETE_BRANCH_RESPONSE, name, err:err});
        if(!err){
            dispatch(branchDeleted(name));
        }
    }
}

export function branchAdded(name, repositories) {
//    console.info("branchAdded", name, quantity)
    return {type:BRANCH_ADDED, name, repositories};
}

export function branchDeleted(name) {
//    console.info("branchAdded", name)
    return {type:BRANCH_DELETED, name};
}

export function createBranchRequest(name) {
//    console.info("createBranchRequest", name)
    return function(dispatch){
        if (!name){
            return
        }
        const branchName = document.currentLoginName + "_" + name;
        dispatch({type:CREATE_BRANCH_REQUEST, branchName});
        fetch('/api/create_branch/' + branchName, {
            method: 'get',
            credentials: 'same-origin',
            cache: 'no-cache'
        }).then(function(response) {
            return response.json()
        }).then(function(statuses){
            dispatch(createBranchResponse(branchName, statuses, null))
        }).catch(function(err) {
            console.info("err is ", err)
        });
    }
}

export function createBranchResponse(name, statuses, err) {
    return function(dispatch){
        dispatch({type:CREATE_BRANCH_RESPONSE, name, statuses, err:err});
        if(!err){
            dispatch(branchAdded(name, Object.keys(statuses).filter(k => statuses[k])));
        }
    }
}

