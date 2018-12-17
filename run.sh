#!/bin/bash

docker run -d -v `pwd`/conf:/create-branch-conf -p 4430:4430 --name create-branch-tool barakb/create-branch:2.0

