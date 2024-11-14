<script setup>
import { RouterLink } from 'vue-router';
import JobListing from './jobListing.vue';
import { reactive, ref, defineProps, onMounted } from 'vue';
import PulseLoader from 'vue-spinner/src/PulseLoader.vue';
import jobService from '../services/jobService';

// Define component props for limit and showButton
defineProps({
  limit: Number,
  showButton: {
    type: Boolean,
    default: false,
  },
});

// Define reactive states
const jobs = ref([]);
const isLoading = ref(true);

// Fetch jobs from the backend using jobService
async function fetchJobs() {
  console.log('Fetching jobs...');
  try {
    const response = await jobService.getAllJobs();
    jobs.value = response.data;
    console.log('Fetched jobs:', jobs.value);
  } catch (error) {
    console.error('Error fetching jobs:', error);
  } finally {
    isLoading.value = false;
  }
}

// Call fetchJobs on component mount
onMounted(fetchJobs);
</script>

<template>
  <section class="bg-blue-50 px-4 py-10">
    <div class="container-xl lg:container m-auto">
      <h2 class="text-3xl font-bold text-green-500 mb-6 text-center">
        Browse Jobs
      </h2>
      <!-- Show loading spinner while loading is true -->
      <div v-if="isLoading" class="text-center text-gray-500 py-6">
        <PulseLoader />
      </div>

      <!-- Show job listing when done loading -->
      <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <JobListing
          v-for="job in jobs.slice(0, limit || jobs.length)"
          :key="job.id"
          :job="job"
        />
      </div>
    </div>
  </section>

  <section v-if="showButton" class="m-auto max-w-lg my-10 px-6">
    <RouterLink
      to="/jobs"
      class="block bg-black text-white text-center py-4 px-6 rounded-xl hover:bg-gray-700"
    >
      View All Jobs
    </RouterLink>
  </section>
</template>
