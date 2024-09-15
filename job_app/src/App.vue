<script setup>
import { ref } from 'vue';


    const name = ref( 'JOB JOB');
    const status = ref( 'active');
    const tasks = ref(['task1','task2','task3']);
    const newTask = ref('');

    const changeStatus = () => {

      if (status.value === 'active') {
        status.value = 'pending';
      } else if (status.value === 'pending') {
        status.value = 'inactive';
      } else {
        status.value = 'active';
      }
  };

  const addTask = () => {
    tasks.value.push(newTask.value);
    newTask.value = '' ;
  };

  const DeleteTask = (index) => {
    tasks.value.splice(index , 1000000);
  };
</script>

<template>
<h1>{{ name}}</h1>
<p v-if="status === 'active'">User is active</p>
<p v-else-if="status === 'pending'">User is pending</p>
<p v-else="status">User is inactive</p>
<h3>Tasks</h3>
<ul>
  <li v-for="(task , index) in tasks" :key="task">
    <span>
      {{ task }} 
    </span>
    <button @click="DeleteTask()"> X </button>
  </li>
</ul>
<button @click="changeStatus"  >
change
</button>

<form @submit.prevent="addTask">
<label for="newTask"> Add Task</label>
<input type="text" id="newTask" name="newTask" v-model="newTask">
<button type="submit">Submit</button>
</form>
</template>

