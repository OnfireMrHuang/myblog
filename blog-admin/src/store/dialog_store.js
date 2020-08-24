export default {
    state: {
        formDatas: null,
        token: "1"
    },
    mutations: {
        getFormData(state,data) {
            state.formDatas = data;
        },
        saveToken(state,data) {
            state.token = data;
            window.localStorage.setItem("Token",data); //将token缓存到localstorage中
        }
    }
}