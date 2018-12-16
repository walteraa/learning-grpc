mongoose = require('mongoose');

// TODO: Use environment variable
mongoose.connect('mongodb://localhost:27017/myDB', {useMongoClient: true} );


exports.db = mongoose
