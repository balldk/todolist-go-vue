<template>
    <div>
        <h1 id="header">Register</h1>
        <div id="register-container">
            <b-field label="Username">
                <b-input v-model="name" autofocus></b-input>
            </b-field>
            
            <b-field label="Password">
                <b-input v-model="password" type="password"></b-input>
            </b-field>
            <b-field 
                label="Re-enter Password" 
                :type="!isPassMatch && isSubmit?'is-danger':''"
                :message="!isPassMatch && isSubmit?errMsg:''">
                <b-input v-model="rePassword" type="password"></b-input>
            </b-field>
            <div class="buttons">
                <b-button @click="register" type="is-primary" expanded>Register</b-button>
            </div>

            <div style="text-align: center">
                <span>Already have account?</span>
                <router-link to="/login" class="route-link">Login</router-link>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

const host = process.env.VUE_APP_HOST_URL
export default {
    data() {
        return {
            name: "",
            password: "",
            rePassword: "",
            errMsg: "",
            isSubmit: false
        }
    },
    methods: {
        async register() {
            this.isSubmit = true
            if (!this.isPassMatch) return;

            let res
            try {
                res = await axios.post(`${host}/api/register`, {
                    Username: this.name,
                    Password: this.password
                })
                this.$axios.defaults.headers.common['Authorization'] = res.data.token
                localStorage.setItem("token", res.data.token)
                localStorage.setItem("username", this.name)

                this.$router.push('/')
            } catch(e) {
                return console.log(e)
            }
            console.log(res)
        }
    },
    computed: {
        isPassMatch() {
            let isMatch = this.password === this.rePassword
            if (!isMatch) {
                this.errMsg = "Re-enter password is not match with password"
            }
            return isMatch
        }
    }
}
</script>

<style scoped>
    #header {
        font-size: 50px;
        text-align: center;
    }
    #register-container {
        width: 300px;
        height: 300px;
        margin: 0 auto;
        margin-top: 40px;
    }
    .route-link {
        margin-left: 8px;
    }
</style>