#!/bin/bash

clear
(cd backend && go run .) & (cd frontend && npm run dev)