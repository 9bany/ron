var http = require('http');

http.createServer(function (req, res) {
  res.write('Hello, some time it will be ok change by bany');
  res.end();

}).listen(3000, function(){
 console.log("server start at port 3000");
});  