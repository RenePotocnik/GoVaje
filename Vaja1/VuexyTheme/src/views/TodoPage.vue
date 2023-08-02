<template>
  <div>
    <b-card title="Add new task">
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
          <b-form-datepicker
              v-model="newTask.predicted_date"
              placeholder="When will the task be completed?"
              format="dd.MM.yyyy"
              value-as-date
          ></b-form-datepicker>
        </b-form-group>

        <b-button type="submit" variant="primary">Add Task</b-button>
      </b-form>
    </b-card>

    <hr>

    <transition-group name="list" tag="div">
      <b-card v-for="(task, index) in tasks" :key="task.id">
        <template #header>
          <h3>{{ task.title }}</h3>
        </template>
        <p>{{ task.description }}</p>
        <p>Date added: <strong>{{ task.date_added }}</strong> | Predicted date of completion: <strong>{{ task.predicted_date }}</strong></p>
        <b-button class="mr-1" @click="completeTask(index)" variant="outline-success">Complete</b-button>
        <b-button class="mr-1" @click="removeTask(index)" variant="outline-danger">Delete</b-button>
      </b-card>
    </transition-group>

  </div>
</template>


<script>
import { BCard, BForm, BFormInput, BButton, BListGroup, BListGroupItem } from 'bootstrap-vue'
import axios from 'axios'
import jwt_decode from 'jwt-decode';
import {logoutUser} from "@/main";
import ToastificationContent from "@core/components/toastification/ToastificationContent.vue";

// Check if token is expired before every request
axios.interceptors.request.use(
    config => {
      let token = localStorage.getItem('token');
      if (isTokenExpired(token)) {
        // Token is expired, remove it from local storage
        logoutUser(this);
      } else {
        config.headers['Authorization'] = token;
      }
      return config;
    },
    error => {
      return Promise.reject(error);
    }
)

export default {
  components: {
    BCard,
    BForm,
    BFormInput,
    BButton,
    BListGroup,
    BListGroupItem,
  },

  data() {
    return {
      newTask: {
        title: '',
        description: '',
        date_added: new Date().toLocaleString().substring(0,10).split('/').join('. '), // Change the date format to "DD. MM. YYYY"
        predicted_date: '',
        user_id: getTokenId(localStorage.getItem('token'))
      },
      tasks: []
    }
  },

  methods: {
    async addTask() {
  // Change the date object to "DD. MM. YYYY" and convert it to string
      this.newTask.predicted_date = this.newTask.predicted_date.toLocaleString().substring(0,10).split('/').join('. ')

      if (this.newTask.title.trim().length === 0 || this.newTask.description.trim().length === 0 || this.newTask.predicted_date.trim().length === 0) {
        this.$toast({
          component: ToastificationContent,
          props: {
            title: 'All fields are required',
            icon: 'MenuIcon',
            variant: 'info',
          },
        })
        return
      }

      let response = await axios.post('http://localhost:80/api/v1/todo/', this.newTask)

      this.tasks.push(response.data)

      // Reset form
      this.newTask = {
        title: '',
        description: '',
        date_added: new Date().toLocaleString().substring(0,10).split('/').join('. '),
        predicted_date: '',
        user_id: getTokenId(localStorage.getItem('token'))
      }

      response = await axios.get('http://localhost:80/api/v1/todo/');
      this.tasks = response.data;

      this.$toast({
        component: ToastificationContent,
        props: {
          title: 'Task Created',
          icon: 'CheckIcon',
          variant: 'success',
        },
      })

    },

    async removeTask(index) {
      await axios.delete(`http://localhost:80/api/v1/todo/${this.tasks[index].id}`);
      this.tasks.splice(index, 1);

      this.$toast({
        component: ToastificationContent,
        props: {
          title: 'Task Removed',
          icon: 'XIcon',
          variant: 'danger',
        },
      })

    },

    async completeTask(index) {
      // await axios.post(`http://localhost:80/api/v1/to do/${this.tasks[index].id}/complete`);
      await axios.delete(`http://localhost:80/api/v1/todo/${this.tasks[index].id}`);

      this.tasks.splice(index, 1);

      this.$toast({
        component: ToastificationContent,
        props: {
          title: 'Task Completed',
          icon: 'CheckIcon',
          variant: 'success',
        },
      })
    },
  },
  async created() {
    const response = await axios.get('http://localhost:80/api/v1/todo/');

    this.tasks = response.data;
  }
}

function isTokenExpired(token) {
  try {
    if (!token) {
      return true;
    }
    const decoded = jwt_decode(token);
    console.log("Remaining token time [s]: ", decoded.exp - Date.now() / 1000)
    return decoded.exp < Date.now() / 1000;
  } catch (err) {
    console.log("Token expired check failed! Error: ", err)
    return false;
  }
}

function getTokenId(token) {
  try {
    if (!token) {
      return null;
    }
    const decoded = jwt_decode(token);
    return decoded.id;
  } catch (err) {
    console.log("Token check failed! Error: ", err)
    return null;
  }
}

</script>

<style>

.list-move {  /* The class of all moving elements */
  transition: all 1s ease;
}

.list-enter-to {  /* The position of the element before the transition */
  transform: translateX(-50%);
  opacity: 0;
}

.list-leave-to {  /* The position of the element after the transition */
  transform: translateX(100%);
  opacity: 0;
}
.list-leave-active {
  position: absolute;
  transition: all 1s ease;
  width: 100%;
}

</style>
