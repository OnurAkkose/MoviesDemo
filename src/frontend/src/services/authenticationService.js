import axios from 'axios'

const apiAddress = "http://localhost:5051"

export default {
  async addUser(data) {
    return axios.post(`${apiAddress}/user`, data)
  }, 
  async getUser(data) {
    return axios.get(`${apiAddress}/user/${data.username}`)
  }, 
  async getUsers() {
    return axios.get(`${apiAddress}/users`)
  }, 
}
