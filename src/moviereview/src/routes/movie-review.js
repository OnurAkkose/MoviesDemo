const express = require("express");
const router = express.Router();

var ReviewControllers = require("../controllers/movie-review");

router.post("/", ReviewControllers.save_review);

router.get("/:movieId", ReviewControllers.getReview);


module.exports = router;