#!/bin/bash
echo "Building site code..."
cd ecms-site
hugo -d ../docs/
cd ..