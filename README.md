# pisense
Store and visualize sensory data from Raspberry Pi via the Azure IoT Hub

## pisense client

The pisense client is written in Go 1.16. In order to use the client on your Rasperry Pi, follow these steps:

1. Download the latest Go compiler for your target architecture (arm6l or arm64) from https://golang.org/dl/
2. Fetch, compile, and install the pisense client 
```
go get github.com/pommestheke/pisense/cmd/pisense
```
3. Run the pisense client
In order to let the pisense client connect to your IoT hub, you need to set the environment variable
`IOTHUB_DEVICE_CONNECTION_STRING` and run the pisense client

```
export IOTHUB_DEVICE_CONNECTION_STRING="<string>";pisense
```

## pisense Azure IoT hub

The pisense Azur IoT hub infrastruture can deployed to Azure using ARM templates located in the azure-templates subdirectory:
https://github.com/pommestheke/pisense/tree/main/azure-templates


