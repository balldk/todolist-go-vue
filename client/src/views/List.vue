<template>
    <div>
        <b-button id="logout-btn" size="is-medium" type="is-danger" icon-left="logout-variant" @click="logout">Logout</b-button>
        <div></div>
        <h1 id="header">Todo List</h1>
        <p id="user-text">@{{ username }}</p>
        <div id="list-card-container">
            <b-field class="card-container">
                <b-input
                    placeholder="🤔 What should you do today..."
                    type="search"
                    v-model="newContent"
                    expanded
                ></b-input>
                <p class="control">
                    <b-button class="button is-primary" icon-right="arrow-down-box" @click="newTask">New</b-button>
                </p>
            </b-field>
            <div class="separator" />
            <div class="card-container columns" v-for="task in tasks" :key="task.taskId">
                <b-checkbox type="is-primary" size="is-medium" v-model="task.completed" 
                    @input="val => updateComplete(val, task.taskId)" />
                <p class="card-ctn">{{ task.content }}</p>
                <b-button
                    type="is-danger"
                    size="is-small"
                    icon-right="close"
                    class="delete-btn"
                    @click="() => deleteTask(task.taskId)"
                />
            </div>
        </div>
    </div>
</template>

<script>
const host = process.env.VUE_APP_HOST_URL
export default {
    data() {
        return {
            tasks: [],
            newContent: "",
            username: localStorage.getItem("username")
        }
    },
    methods: {
        logout() {
            localStorage.clear()
            this.$router.push("/login")
        },
        updateComplete(completed, taskId) {
            this.$axios.put(`${host}/api/updateComplete`, {
                taskId,
                completed,
            })
        },
        async deleteTask(taskId) {
            await this.$axios.delete(`${host}/api/deleteTask`, {
                data: {
                    taskId: taskId
                }
            })
            this.tasks = this.tasks.filter(task => task.taskId !== taskId)
        },
        async newTask(content) {
            let res = await this.$axios.post(`${host}/api/newTask`, {
                content: this.newContent,
            })
            this.tasks.unshift({
                content: this.newContent,
                taskId: res.data.taskId,
            })
            this.newContent = ""
        },
    },
    async created() {
        try {
            let res = await this.$axios.get(`${host}/api/getTasks`)
            this.tasks = res.data.sort((a, b) => b.taskId - a.taskId)
        } catch (e) {
            this.$router.replace('/register')
        }
    },
}
</script>

<style scoped>
#header {
    font-size: 50px;
    text-align: center;
    margin-top: 35px;
}
#user-text {
    text-align: center;
    font-weight: bold;
    font-style: italic;
    font-size: 20px;
    margin-bottom: 30px;
}
#list-card-container {
    width: 500px;
    margin: 0 auto;
    margin-bottom: 40px;
}
#logout-btn {
    position: absolute;
    right: 30px;
}
.card-container {
    padding: 20px;
    margin-top: 20px;
    position: relative;

    word-wrap: break-word;
    box-shadow: 0 4px 6px 3px rgba(0, 0, 0, 0.2);
    border-radius: 10px;
    align-items: center;

    margin-left: 0;
    margin-right: 0;
}
.card-ctn {
    font-size: 20px;
}
.separator {
    height: 15px;
}
.delete-btn {
    margin-left: auto;
    margin-right: 0;
}
.add-btn {
    margin-left: auto;
    margin-right: 0;
}
</style>
