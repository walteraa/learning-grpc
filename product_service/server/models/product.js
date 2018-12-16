var dbSettings = require('../database/settings.js')
var db = dbSettings.db;

var productSchema = new db.Schema({
    _id: { type: db.Schema.ObjectId, auto: true },
    price_in_cents: { type: Number },
    title: { type: String, required: true },
    description: { type: String }
},{collection: 'products'});


var Product = db.model('Product', productSchema);
