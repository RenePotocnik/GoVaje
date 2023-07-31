<template>
  <div>
    <b-card title="Todo App 📝">
      <b-form @submit.prevent="addTask">
        <b-form-group label="Title">
          <b-form-input
              v-model="newTask.title"
              placeholder="Enter task title"
          ></b-form-input>
        </b-form-group>

        <b-form-group label="Description">
          <b-form-textarea
              v-model="newTask.description"
              placeholder="Enter task description"
          ></b-form-textarea>
        </b-form-group>

        <b-form-group label="Predicted Completion Date">
          <b-form-input
              type="date"
              v-model="newTask.predicted_date"
              placeholder="Enter predicted completion date"
          ></b-form-input>
        </b-form-group>

        <b-button type="submit" variant="primary">Add Task</b-button>
      </b-form>

      <!-- task list goes here -->
    </b-card>
  </div>
</template>


<script>
import { BCard, BForm, BFormInput, BButton, BListGroup, BListGroupItem } from 'bootstrap-vue'
import axios from 'axios'

// Set the Authorization header for all requests
axios.defaults.headers.common['Authorization'] = `${localStorage.getItem('token')}`;
console.log("Token: ", localStorage.getItem('token'));

export default {
  components: {
    BCard,
    BForm,
    BFormInput,
    BButton,
    BListGroup,
    BListGroupItem
  },

  data() {
    return {
      newTask: {
        title: '',
        description: '',
        date_added: new Date().toISOString().substring(0,10), // Assuming the date format is "YYYY-MM-DD"
        predicted_date: '',
        user_id: 0 // TODO: Add dynamically based on the logged in user
      },
      tasks: []
    }
  },


  methods: {
    async addTask() {
      if (this.newTask.title.trim().length === 0 || this.newTask.description.trim().length === 0) {
        return
      }

      const response = await axios.post('http://localhost:80/api/v1/todo/', this.newTask)

      this.tasks.push(response.data)

      // Reset form
      this.newTask = {
        title: '',
        description: '',
        date_added: new Date().toISOString().substring(0,10),
        predicted_date: '',
        user_id: 0 // TODO: Add dynamically based on the logged in user
      }
    },
    async removeTodo(index) {
      await axios.delete(`http://localhost:80/api/v1/todo/${this.todos[index].id}`)

      this.todos.splice(index, 1)
    },
    async completeTodo(index) {
      await axios.delete(`http://localhost:80/api/v1/todo/${this.todos[index].id}`)

      this.todos.splice(index, 1)
    },
  },
  async created() {

    const response = await axios.get('http://localhost:80/api/v1/todo/')

    this.todos = response.data
  }
}
</script>
