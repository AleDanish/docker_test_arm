#!/bin/bash
if [ "$1" == "" ]; then
   echo "An argument is required. How many containers do you want to run?"
   exit 1
else
   num=$1
   echo $num
fi

image_name=sorter
echo "-> About to create $num containers from the image $image_name."

# Build all the containers
echo "Building container..."
sudo docker build -t $image_name .
echo "Container built"   

# Run all containers simultaneuosly
n=1
while [  $n -le $num ]; do
   echo "Starting container num: $n."
   sudo docker run -d $image_name
   echo "Started container num: $n."   
   let n+=1 
done
