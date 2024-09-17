<script setup>
import { ref, onMounted } from 'vue';
import jobListing from './jobListing.vue';
const jobs = ref([]);


onMounted(async () => {
  try {
    const response = await fetch('/jobs.json');
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    jobs.value = data.jobs;
  } catch (error) {
    console.error('Error fetching jobs:', error);
  }
});
</script>


<template>
<section class="bg-blue-50 px-4 py-10">
        <div class="container-xl lg:container m-auto">
            <h2 class="text-3xl front-bold text-green-500 mb-6 text-center">
                Brwose jobs
            </h2>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">

                <jobListing v-for="job in jobs" :key="job.id" :job="job" />

            </div>
        </div>


</section>

</template>