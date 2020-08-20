<template>
    <div>
        <h1 id="header">Login</h1>
        <div id="login-container">
            <b-field label="Username">
                <b-input v-model="name" autofocus></b-input>
            </b-field>
            
            <b-field label="Password">
                <b-input v-model="password" type="password"></b-input>
            </b-field>

            <div v-if="errMsg">
                <p class="has-text-danger">* {{ errMsg }}</p>
                <div style="height: 10px"></div>
            </div>

            <div class="buttons">
                <b-button @click="login" type="is-primary" expanded>Login</b-button>
            </div>

            <div style="text-align: center">
                <span>Haven't got account?</span>
                <router-link to="/register" class="route-link">Register</router-link>
            </div>
        </div>
    </div>
</template>

<script>

const host = process.env.VUE_APP_HOST_URL
export default {
    data() {
        return {
            name: "",
            password: "",
            errMsg: ""
        }
    },
    methods: {
        async login() {
            try {
                let res = await this.$axios.post(`${host}/api/login`, {
                    Username: this.name,
                    Password: this.password,
                })
                this.$axios.defaults.headers.common['Authorization'] = res.data.token
                localStorage.setItem("token", res.data.token)
                localStorage.setItem("username", this.name)

                this.$router.push('/')
            } catch(e) {
                this.errMsg = "Wrong Username or Password"
            }
        }
    }
}
</script>

<style scoped>
    #header {
        font-size: 50px;
        text-align: center;
    }
    #login-container {
        width: 300px;
        height: 300px;
        margin: 0 auto;
        margin-top: 40px;
    }
    .route-link {
        margin-left: 8px;
    }
</style>