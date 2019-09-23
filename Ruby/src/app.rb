require 'sinatra'
require 'json'
require "sinatra/reloader"
port = 4000

set :bind, '0.0.0.0'
set :port, port
# set :environment, :development

configure :development do
  enable :reloader
end


get '/api' do
    "HASHER running on #{port}\n"
end

post "/api" do
    request.body.rewind  # en caso de que alguien ya lo haya leído
    data = JSON.parse request.body.read
    num1 = data['num1']
    num2 = data['num2']
    # "Hola #{num1+num2}!"
    content_type :json
    { :result => num1-num2 }.to_json
  end
