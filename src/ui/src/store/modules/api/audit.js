import $http from '@/api'
const actions = {
    getList (context, { params, config }) {
        return $http.post('findmany/audit_list', params, config)
    },
    getDictionary (context, config) {
        return $http.get('find/audit_dict', config)
    },
    getDetails (context, { id, config }) {
        return $http.get(`find/audit/${id}`, config)
    }
}

export default {
    namespaced: true,
    actions
}
