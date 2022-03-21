import axios from 'axios'

const apiAddress = "http://localhost:5050"

export default {
  async addMovie(data) {
    return axios.post(`${apiAddress}/movie`, data)
  }, 
  async getMovie(data) {
    return axios.get(`${apiAddress}/movie/${data.id}`)
  }, 
  async getMovies() {
    return axios.get(`${apiAddress}/movies`)
  }, 
}
