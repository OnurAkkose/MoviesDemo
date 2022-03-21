import axios from 'axios'

const apiAddress = "http://localhost:5052"

export default {
  async getReviews(id) {
    return axios.get(`${apiAddress}/reviews/${id}`)
  }, 
  async addReview(data) {
    return axios.post(`${apiAddress}/reviews/`, data)
  },  
}
