<script>
export default {
  data () {
    return {
      name:'JOB JOB',
      status: 'active',
      tasks :['task1','task2','task3'],
      link:'https://www.google.com/',

    };
  },

  methods: {

      changeStatus() {
      if (this.status === 'active') {
        this.status = 'pending';
      } else if (this.status === 'pending') {
        this.status = 'inactive';
      } else {
        this.status = 'active';
      }
      }
  }
  };


</script>

<template>
<h1>{{ name}}</h1>
<p v-if="status === 'active'">User is active</p>
<p v-else-if="status === 'pending'">User is pending</p>
<p v-else="status">User is inactive</p>
<h3>Tasks</h3>
<ul>
  <li v-for="task in tasks" :key="task"> {{ task }} </li>
</ul>
<a :href="link"> Go to google</a>
<button @click="changeStatus"  >
change
</button>
</template>

