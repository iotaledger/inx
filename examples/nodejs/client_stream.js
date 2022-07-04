const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const PROTO_PATH = "../../proto/inx.proto";


const INX_ADDRESS = "localhost:9029";


const protoOptions = {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
};


// Load the `inx.proto` file
const packageDefinition = protoLoader.loadSync(PROTO_PATH, protoOptions);
const INX = grpc.loadPackageDefinition(packageDefinition).inx.INX;
// Instantiate an INX client
const client = new INX(
    INX_ADDRESS,
    grpc.credentials.createInsecure()
);

// Listen to the stream of latest milestones
var call = client.ListenToLatestMilestones();
call.on('data', function (answer) {
    console.log(answer);
});
call.on('end', function () {
    // The server has finished sending
});
call.on('error', function (error) {
    // An error has occurred and the stream has been closed.
});
call.on('status', function (status) {
    // process status
});