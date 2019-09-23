para que `require "sinatra/reloader"` funcione primero ejecutra `gem install sinatra-contrib`
`request.body.rewind` es necesario para que no tire error. Wea rara de sinatra


# Dockerizing 

1. `docker build -t ruby-app .`
2. `docker run -it -p 3000:4000 ruby-app`