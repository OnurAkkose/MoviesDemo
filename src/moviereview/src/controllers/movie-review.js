var MovieReview = require("../models/movie-review");

exports.save_review = (req,res,next) => {
    const review = new MovieReview({
        movieId: req.body.movieId,
        userId: req.body.userId,
        title: req.body.title,
        description: req.body.description,
        score: req.body.score 
    })
    return review.save((error) => {
        if (error){
            throw error;
        }      
        res.status(201).json({
            message: "Review saved",
            createdReview: review,
        });
    })
}

exports.getReview = (req, res, next) => {
    const movieId = req.params.movieId;
    MovieReview.find({movieId:movieId}, (error,data) => {
        if (error){
            return res.status(404).json({
                message: "Review not found"
              });
        }        
        res.status(200).json({
            review: data          
          });
    })
}
