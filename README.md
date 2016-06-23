# drud-pod-ip-finder
initial pod finder code using the kube go api

##To run:

  Build the image:
  ```make all```
  
  Create the Service:
  ```k create -f service.yaml```
  
  Create the Controller:
  ```k create -f controller.yaml```

