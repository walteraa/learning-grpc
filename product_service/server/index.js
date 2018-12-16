const PROTO_PATH = '../../proto/product.proto'
const port = process.env.PORT || 50050
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const productDAO = require('./models/product_dao.js').productDAO;


const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        defaults: false,
    });

const server = new grpc.Server();
server.addService(packageDefinition['product.ProductService'], {
    GetProduct: (call, callback) => {
        productDAO.Get(call.request, callback);
    },
    ListProduct: (call, callback) => {
        productDAO.List(call.request, callback);
    },
    CreateProduct: (call, callback) => {
        productDAO.Create(call.request, callback);
    },
});
server.bind(`0.0.0.0:${port}`, grpc.ServerCredentials.createInsecure());
server.start();
