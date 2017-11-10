# weatherService #

weatherService is a go version service for weather.

It provice a json responed of weather info.

### Usage


- step 1) Install the weatherService:
      
      $ go get github.com/andy-y-li/weatherService
      
- step 2) Run the weatherService:
      
      $ go run main.go 
      
- step 3) Query weather for example:
      
      $ curl http://localhost:9090/weather/深圳
      
      {"city":"深圳","cityid":"101280601","img1":"n0.gif","img2":"d0.gif","ptime":"18:00","temp1":"14℃","temp2":"23℃","weather":"晴"}

* BTW, I have add the service in my server, you also 
   get the weather iofo via the command as below:
   
   $ curl http://1714w8z674.iok.la:25705/weather/深圳

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.

