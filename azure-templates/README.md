# Create pisense infrastructure: Azure IoT Hub and persistence

Use this ARM template to create the pisense infrastructure:

[![Deploy To Azure](https://raw.githubusercontent.com/Azure/azure-quickstart-templates/master/1-CONTRIBUTION-GUIDE/images/deploytoazure.svg?sanitize=true)](https://portal.azure.com/#create/Microsoft.Template/uri/https%3A%2F%2Fraw.githubusercontent.com%2Fpommestheke%2Fpisense%2Fmain%2Fazure-templates%2Fazuredeploy.json)

Follow these steps:

1. Upload the Telemetry.bacpac database schema from this directory to an Azure storageAccount
2. Hit the above Deploy to Azure button and fill out the template parameters including path and key to the bacpac 
3. After deployment Go tot he IoT Hub and create a new device. Use the resulting connection string with the pisense client

The Azure Time Series Insight environment needs to be commissioned manually due to lack of Gen2 template reference.

