var http = require('http');

http.createServer(function (req, res) {
  res.write('Hello, im  bany');
  res.end();

}).listen(3000, function(){
 console.log("server start at port 3000");
});  