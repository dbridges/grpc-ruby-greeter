# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: greeter.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("greeter.proto", :syntax => :proto3) do
    add_message "greeter.GreetRequest" do
      optional :name, :string, 1
    end
    add_message "greeter.MessageReply" do
      optional :message, :string, 1
    end
  end
end

module Greeter
  GreetRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("greeter.GreetRequest").msgclass
  MessageReply = Google::Protobuf::DescriptorPool.generated_pool.lookup("greeter.MessageReply").msgclass
end