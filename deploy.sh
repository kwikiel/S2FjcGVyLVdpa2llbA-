#!/usr/bin/env bash

gcloud builds submit --tag gcr.io/airswap/gwp1
gcloud beta run deploy gwp1 --image gcr.io/airswap/hellogo --platform managed
