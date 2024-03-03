#include <iostream>
#include <memory>
#include <string>
#include <grpcpp/grpcpp.h>
#include "./control_service.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReaderWriter;
using grpc::Status;
using control::ControlMessage;
using control::ControlService;

class ControlClient {
public:
  ControlClient(std::shared_ptr<Channel> channel)
      : stub_(ControlService::NewStub(channel)) {}

  void SendControlMessage() {
    ClientContext context;
    std::shared_ptr<ClientReaderWriter<ControlMessage, ControlMessage>> stream(
        stub_->StreamControl(&context));

    // Example: send a message
    ControlMessage message;
    message.set_command_name("ExampleCommand");
    message.add_args("arg1");
    message.add_args("arg2");
    stream->Write(message);

    // Read the echo back
    ControlMessage server_message;
    while (stream->Read(&server_message)) {
      std::cout << "Received echo back: " << server_message.command_name() << std::endl;
    }

    Status status = stream->Finish();
    if (!status.ok()) {
      std::cerr << "ControlService stream encountered an error: " << status.error_message() << std::endl;
    }
  }

private:
  std::unique_ptr<ControlService::Stub> stub_;
};

int main(int argc, char** argv) {
  ControlClient client(grpc::CreateChannel("localhost:50051", grpc::InsecureChannelCredentials()));
  client.SendControlMessage();
  return 0;
}
