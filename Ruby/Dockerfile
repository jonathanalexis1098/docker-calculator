FROM ruby:2.6-alpine

WORKDIR /app

COPY . /app

RUN gem install sinatra-contrib

CMD [ "ruby","app.rb"]