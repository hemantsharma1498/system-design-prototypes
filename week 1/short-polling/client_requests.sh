#!/bin/bash
#
#

while :
  do 
    response=$(curl -s localhost:3000/short-poll/false)
    if [[ "$response" == "Status has changed" ]]; then
      echo "Response: $response"
      break 
    else echo "Response: $response"
    fi
done
  
