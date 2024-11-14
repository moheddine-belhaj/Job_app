import axios from 'axios';

const API_URL = 'http://localhost:9010/job/';

export default {
  getAllJobs() {
    return axios.get(API_URL + 'all');
  },

  getJobById(jobId) {
    return axios.get(`${API_URL}${jobId}`);
  },

  createJob(job) {
    return axios.post(API_URL + 'create', job);
  },

  updateJob(jobId, job) {
    return axios.put(`${API_URL}update/${jobId}`, job);
  },

  deleteJob(jobId) {
    return axios.delete(`${API_URL}delete/${jobId}`);
  }
};
