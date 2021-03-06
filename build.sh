#! /bin/bash

rm -rf ./build
mkdir ./build
mkdir ./build/assets

# Copy files from assets to build, compile to .js if it's a .coffee file
for f in ./assets/*.*; do
  if echo "$f" | grep -qE ".coffee$"; then
    echo "Compiling $f"
    coffee -o ./build/assets/ -c $f
  else
    echo "Copying $f"
    cp $f ./build/assets/
  fi
done

# Copy assets subdirs (these directories contain nothing that needs to be compiled
for d in $(find ./assets/* -type d); do 
  echo "Copying $d"; 
  cp -R $d ./build/assets/
done

# Compile haml files
for f in ./haml/*; do
  echo "Compiling $f"
  b=`basename $f`
  haml $f > ./build/${b/.haml/}
done

go install
