var express = require('express');
var cors = require('cors')
var app = express();
app.use(cors())
const bodyParser= require('body-parser')
var mongoose = require('mongoose')
var reviewRoutes = require("./src/routes/movie-review")
mongoose.connect('*',(error) => {
    
})
mongoose.Promise = global.Promise

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());


app.use("/reviews",reviewRoutes)

module.exports = app;