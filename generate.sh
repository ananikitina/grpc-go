#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=. --go_opt=paths=source_relative