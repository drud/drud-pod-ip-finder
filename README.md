# drud-pod-ip-finder
initial pod finder code using the kube go api

##To run:

  Build the image:
  ```make all```
  
  Create the Service:
  ```k create -f service.yaml```
  
  Create the Controller:
  ```k create -f controller.yaml```

##To use:

    Hit port 80 of the service with key=val params denoting labels and the values you want them to match.
    The service will return a list of podd ips that match your query.
