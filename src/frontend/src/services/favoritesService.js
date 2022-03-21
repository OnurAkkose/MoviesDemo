import axios from 'axios'

const apiAddress = "http://localhost:8000"

export default {
  async addFavorite(data) {
    return axios.post(`${apiAddress}/`, data)
  }, 
  async deleteFavorite(data) {
    return axios.delete(`${apiAddress}/${data.userId}/${data.movieId}`)
  }, 
  async getFavorites(data) {
    return axios.get(`${apiAddress}/${data.userId}`)
  }, 
}
