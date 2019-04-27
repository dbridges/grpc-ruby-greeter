#!/usr/bin/env ruby

# This is only needed for Go's present tool's code execution
__dir__ = File.expand_path("~/go/src/grpc-ruby-greeter")

# START OMIT

$LOAD_PATH.unshift(File.expand_path(__dir__))

require "grpc"
require "greeter_services_pb"

class Client
  def initialize
    @stub = Greeter::Greeter::Stub.new("localhost:50051", :this_channel_is_insecure)
  end

  def greet(name)
    request = Greeter::GreetRequest.new(name: name)
    puts @stub.greet(request).message
  end
end

client = Client.new
client.greet("Dan")

# END OMIT
