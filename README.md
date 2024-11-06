# Protobuf Example
This repo is a quick overview of a client and server being generated on the fly by protobuf. This repo has a structure of protos for all of our proto files and cmds for our applications.

# How to's
### generate protobuf
To generate the protobuf files after modification it is as simple as running:
```
make generate
```
### run server
To run the server, pass the following command from the repo root directory:
```
make server
```

### run client
To run the client, make sure the server is running in another terminal then run the following command:
```
make client
```
This should give your an output of a greeting such as `INFO[0000] Greetings: Hello Andrew`

