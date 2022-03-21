var mongoose = require('mongoose')
var Schema = mongoose.Schema;

var movieReviewSchema = new Schema({
    movieId: String,
    userId: String,
    title: String,
    description: String,
    score: Number    
})
var MovieReview = mongoose.model('MovieReview', movieReviewSchema)

module.exports = MovieReview


