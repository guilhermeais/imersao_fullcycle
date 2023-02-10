#!/bin/bash

if [ ! -f ".env" ]; then
  echo "No .env file found, copying .env.example to .env"
  cp .env.example .env
fi

npm i --silent

npm run start:dev