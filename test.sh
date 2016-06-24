#!/bin/bash
if [ "$1" == "" ]; then
   echo "An argument is required. How many containers do you want to run?"
   exit 1
else
   num=$1
fi

image_name=sorter
echo "-> About to create $num containers from the image $image_name."

# Build all the containers
echo "Building container..."
Tbuild_start=$(date +%s%6N)
sudo docker build -t $image_name .
Tbuild_end=$(date +%s%6N)
echo "Container built"

# Run all containers simultaneuosly
n=1
Trun_start=$(date +%s%6N)
while [  $n -le $num ]; do
   echo "Starting container num: $n."
   sudo docker run -d $image_name
   echo "Started container num: $n."   
   let n+=1 
done
echo "Finished to create containers"

while [ "$(sudo docker ps -q)" != "" ]; do
:
done

Trun_end=$(date +%s%6N)

Tbuild=$(((Tbuild_end-Tbuild_start)/1000))
Trun=$(((Trun_end-Trun_start)/1000))
sudo echo "Time to build: " $Tbuild " ms"
sudo echo "Time to run all containers: " $Trun " ms"
