var model = require('./product.js');
var db = require('../database/settings.js').db;
var Product = db.model('Product')


exports.productDAO = {
    List: function(options, callback){
        Product.find({}, function(error, product_list){
            if(error){
                callback({error: 'Error fetching Products'});
            }else{
                // TODO: Map the mongoose result(prods) to call discount microservice over gRPC client
                // Passing DiscountRequest for each product
                callback(null,{products: product_list});
            }
        
        });
    },
    
    Get: function(options, callback){
        const criteria = {_id: options.product_id};
        Product.findOne(criteria, function(error, product){
            if(error){
                console.log("Product not found!");
            }else{
                callback(null, product); 
            }
        })

    },

    Crate: function(options, callback){
        var product = new Product(options);
        product.save(function(error, product){
            if(error){
                callback(null, {product_id: product._id});
            }else{
                console.log("Error when saving product");
            }
        });   
    }
