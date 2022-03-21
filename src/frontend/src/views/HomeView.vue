<template>

  <div class="container">
      <div class="btn-user">   
     <vue-feather
      v-if="logged"
      type="log-out"
      size="36"
      style="color: gray; float: right"
      @click="logout()"
    ></vue-feather>
    <vue-feather
      v-else
      type="log-in"
      size="36"
      style="color: gray; float: right"
      @click="this.loginModalShow = true"
    ></vue-feather>
  </div>
    <br>
    <main class="grid">
      <tr v-for="item in items" :key="item.id">
        <MovieCard
          :data="item"
          :favoriteMovies="favorites"
          @removeFavorite="removeFavorite"
        />
      </tr>
    </main>
  </div>

  <b-button
    class="btn-bot"
    style="background-color: #f9f9f9; border-color: gray; width: 80px"
    v-b-modal.modal-1
  >
    <vue-feather type="plus" size="40" style="color: gray"></vue-feather>
  </b-button>

  <b-modal id="modal-1" title="New Movie" hide-footer>
    <b-form-group id="input-group-1" label="Name:" label-for="input-1">
      <b-form-input
        id="input-1"
        v-model="name"
        type="text"
        required
      ></b-form-input>
    </b-form-group>
    <b-form-group id="input-group-1" label="Description" label-for="input-1">
      <b-form-input
        id="input-1"
        v-model="description"
        type="text"
      ></b-form-input>
    </b-form-group>
    <b-form-group id="input-group-1" label="Image Url" label-for="input-1">
      <b-form-input id="input-1" v-model="imageUrl" type="text"></b-form-input>
    </b-form-group>
    <b-form-group id="input-group-1" label="Year" label-for="input-1">
      <b-form-input id="input-1" v-model="year" type="number"></b-form-input>
    </b-form-group>
    <b-form-group id="input-group-1" label="Duration" label-for="input-1">
      <b-form-input
        id="input-1"
        v-model="duration"
        type="number"
      ></b-form-input>
    </b-form-group>

    <b-button @click="addMovie()">Save</b-button>
  </b-modal>

  <b-modal id="modal-login" v-model="loginModalShow" hide-header hide-footer>
    <b-form-group id="input-group-1" label="username:" label-for="input-1">
      <b-form-input
        id="input-1"
        v-model="username"
        type="text"
        required
      ></b-form-input>
    </b-form-group>
    <b-button @click="login()">Login</b-button>
  </b-modal>
</template>

<script>
import MovieCard from "@/components/MovieCard.vue";
import {
  BCard,
  BCardGroup,
  BButton,
  BModal,
  BFormGroup,
  BFormInput,
  BBadge,
} from "bootstrap-vue-3";
import moviesService from "../services/moviesService";
import favoritesService from "../services/favoritesService";
import authenticationService from "../services/authenticationService";
import { RepeatIcon } from "vue-feather-icons";
export default {
  components: {
    MovieCard,
    BCard,
    BCardGroup,
    BButton,
    BModal,
    BFormGroup,
    BFormInput,
    BBadge,
  },
  data() {
    return {
      defaultImage: "https://picsum.photos/600/300/?image=25",
      defaultTitle: "Denemedfadf",
      defaultDescription: "Description dfdfdf",
      items: [],
      name: "",
      description: "",
      imageUrl: "",
      duration: 0,
      year: 0,
      user: {
        id: "dfdfadf",
        movieId: "dfd",
      },
      favorites: [],
      logged: false,
      username: "",
      loginModalShow: false,

    };
  },

  mounted() {
    this.getMovies();
    this.getFavorites();
    this.userInfo();
  },
  methods: {
    userInfo(){
      if(localStorage.getItem('user') != null){
        this.logged = true;
      }
    },
    removeFavorite(id) {
      this.favorites = this.favorites.filter((f) => f !== id);
    },

    async login() {
      try {
    
        const response = await authenticationService.getUser({
          username: this.username,
        });
        if (response.data != undefined) {
          localStorage.setItem("user", response.data.ID);
          this.logged = true;
          this.loginModalShow = false;
          this.getMovies();
        }
      } catch (err) {
        console.log(err);
      }
    },
    async logout() {
      this.logged = false;
      localStorage.removeItem("user");
    },
    async getMovies() {
      try {
        const response = await moviesService.getMovies();
        this.items = response.data;
      } catch (error) {
        console.log(error);
      }
    },

    async addMovie() {
      try {
        const response = await moviesService.addMovie({
          name: this.name,
          description: this.description,
          photo:
            this.imageUrl.length > 0
              ? this.imageUrl
              : "https://picsum.photos/600/300/?image=25",
          duration: this.duration,
          year: this.year,
        });
        this.items.push(response.data);
      } catch (error) {
        console.log(error);
      }
    },
    async getFavorites() {
      try {
        const response = await favoritesService.getFavorites({
          userId: localStorage.getItem("user"),
        });
        for (let index = 0; index < response.data.length; index++) {
          this.favorites.push(response.data[index].movieId);
        }
      } catch (error) {
        console.log(error);
      }
    },
    async deleteFavorite(id) {
      try {
        const response = await favoritesService.getFavorites({
          favorite: this.favorites.data,
        });
      } catch (error) {
        console.log(error);
      }
    },
  },
};
</script>



<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.container {
  position: relative;
  top: 16px;
}

.btn-bot {
  position: fixed;
  margin-left: -50px;
  left: 50%;
  width: 100px;
  bottom: 5%;
  background-color: #f9f9f9;
}

.btn-user {
  position: absolute;
  margin-left: -50px;
  right: 1%;
  width: 100px;


  width: 80%;
}
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  grid-gap: 20px;
  align-items: stretch;
}
.grid > article {
  border: 1px solid #ccc;
  box-shadow: 2px 2px 6px 0px rgba(0, 0, 0, 0.3);
}
.grid > article img {
  max-width: 100%;
}
.text {
  padding: 0 20px 20px;
}
.text > button {
  background: gray;
  border: 0;
  color: white;
  padding: 10px;
  width: 100%;
}
</style>
