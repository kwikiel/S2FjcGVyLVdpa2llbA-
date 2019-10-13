#!/usr/bin/env bash

gcloud builds submit --tag gcr.io/airswap/gwp
gcloud beta run deploy gwp --image gcr.io/airswap/hellogo --platform managed
