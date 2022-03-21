<template>
  <div>
    <b-card
      :title="data.title"
      :img-src="data.photo.length > 10 ? data.photo : defaultImage"
      img-alt="Image"
      img-top
      tag="article"
      style="
        max-width: 20rem;
        min-height: 290px;
        margin-left: 10px;
        margin-top: 10px;
        border-radius: 5%;
        background-color: #f9f9f9;
      "
    >
      <div style="min-height: 60px">
        <b-card-text>
          {{ data.description }}
        </b-card-text>
      </div>

      <div style="float: right">
        <vue-feather
          type="heart"
          @click="addToFavorites(data.ID)"
          :stroke="favoriteMovies.includes(data.ID) ? 'red' : 'gray'"
          :fill="favoriteMovies.includes(data.ID) ? 'red' : '#ede7d8'"
        ></vue-feather>
        &nbsp;
        <vue-feather
          @click="getMovie(data.ID)"
          type="book-open"
          stroke="gray"
          fill="#ede7d8"
        ></vue-feather>
        &nbsp;
        <vue-feather
          @click="openReviewModal(data.ID)"
          type="edit-2"
          stroke="gray"
          fill="#ede7d8"
        ></vue-feather>
      </div>
    </b-card>
  </div>
  <b-modal v-model="modalShow" id="modal-2" :title="selectedData.name">
    <img :src="selectedData.photo" style="width: 100%" />
    <p class="my-4">{{ selectedData.description }}</p>
    <i>Year: {{ selectedData.year }}</i>
    <br />
    <i>Duration: {{ selectedData.duration }}</i>
    <div>
      <b-table striped hover :items="movieReviews"></b-table>
    </div>
  </b-modal>
  <b-modal v-model="reviewModalShow" id="reviewModal">
    <b-form-textarea
      style="margin-top: 10px"
      id="textarea"
      v-model="review"
      placeholder="Enter something..."
      rows="3"
      max-rows="6"
    ></b-form-textarea>
    <div>
      <label for="range-1">Example range with min and max</label>
      <b-form-input
        id="range-1"
        v-model="score"
        type="range"
        min="1"
        max="10"
      ></b-form-input>
      <div class="mt-2">Value: {{ score }}</div>
    </div>
    <pre
      class="mt-3 mb-0"
    > <b-button @click="addReview(data.ID)">Save Review</b-button></pre>
    <div></div>
  </b-modal>
</template>

<script>
import {
  BCardGroup,
  BCard,
  BCardText,
  BButton,
  BModal,
  BFormTextArea,
  BFormInput,
  BTable,
} from "bootstrap-vue-3";
import moviesService from "../services/moviesService";
import favoritesService from "../services/favoritesService";
import reviewsService from "../services/reviewService";
export default {
  components: {
    BCardGroup,
    BCard,
    BCardText,
    BButton,
    BModal,
    BFormTextArea,
    BFormInput,
    BTable,
  },
  setup() {},
  props: ["data", "favoriteMovies"],
  emits: ["removeFavorite"],
  data() {
    return {
      selectedData: {},
      modalShow: false,
      reviewModalShow: false,
      review: "",
      score: 0,
      movieReviews: [],
      defaultImage: "https://picsum.photos/600/300/?image=25",
    };
  },
  methods: {
    async addReview(id) {
      try {
        await reviewsService.addReview({
          movieId: id,
          userId: localStorage.getItem("user"),
          description: this.review,
          score: this.score,
        });
        this.reviewModalShow = false;
      } catch (error) {
        this.reviewModalShow = false;
      }
    },
    openReviewModal(id) {
      this.reviewModalShow = true;
      this.addReview(id);
    },
    async getReviews(id) {
      try {
        this.movieReviews = [];
        const response = await reviewsService.getReviews(id);
        for (let i = 0; i < response.data.review.length; i++) {
          const element = response.data.review[i];
          this.movieReviews.push({
            score: element.score,
            review: element.description,
          });
        }
      } catch (err) {}
    },
    async addToFavorites(id) {
      try {
        if (this.favoriteMovies.includes(id)) {
          const response = await favoritesService.deleteFavorite({
            movieId: id,
            userId: localStorage.getItem("user"),
          });
          this.$emit("removeFavorite", id);
        } else {
          const response = await favoritesService.addFavorite({
            movieId: id,
            userId: localStorage.getItem("user"),
          });
          this.favoriteMovies.push(id);
        }
      } catch (error) {}
    },
    async getMovie(id) {
      try {
        this.selectedData = {};
        const response = await moviesService.getMovie({
          id: id,
        });
        this.selectedData = response.data;
        this.title = response.data.name;
        this.getReviews(id);
        this.modalShow = true;
      } catch (error) {}
    },
  },
};
</script>

<style scoped>
</style>
